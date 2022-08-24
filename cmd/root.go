package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/aristat/golang-example-app/cmd/jwt"

	"github.com/aristat/golang-example-app/cmd/migrate"

	health_check_service "github.com/aristat/golang-example-app/cmd/health-check-service"
	product_service "github.com/aristat/golang-example-app/cmd/product-service"

	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/aristat/golang-example-app/app/logger"

	"go.uber.org/automaxprocs/maxprocs"

	"github.com/aristat/golang-example-app/cmd/daemon"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configPath string
	debug      bool
	v          *viper.Viper
	log        logger.Logger
)

const prefix = "cmd.root"

// Root command
var rootCmd = &cobra.Command{
	Use:           "bin [command]",
	Long:          "",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		l, c, e := logger.Build()
		defer c()
		if e != nil {
			panic(e)
		}

		log = l.WithFields(logger.Fields{"service": prefix})

		v.SetConfigFile(configPath)

		if configPath != "" {
			e := v.ReadInConfig()
			if e != nil {
				log.Error("can't read config, %v", logger.Args(errors.WithMessage(e, prefix)))
				os.Exit(1)
			}
		}

		if debug {
			b, _ := json.Marshal(v.AllSettings())
			var out bytes.Buffer
			e := json.Indent(&out, b, "", "  ")
			if e != nil {
				log.Error("can't prettify config")
				os.Exit(1)
			}
			fmt.Println(string(out.Bytes()))
		}

		_, _ = maxprocs.Set(maxprocs.Logger(log.Printf))
	},
}

func init() {
	v = viper.New()
	v.SetConfigType("yaml")
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	v.AutomaticEnv()

	// pflags
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug mode")

	// initializing
	wd := os.Getenv("APP_WD")
	if len(wd) == 0 {
		wd, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}
	wd, _ = filepath.Abs(wd)
	ep, _ := entrypoint.Initialize(wd, v)

	// bin pflags to viper
	_ = v.BindPFlags(rootCmd.PersistentFlags())

	go func() {
		reloadSignal := make(chan os.Signal)
		signal.Notify(reloadSignal, syscall.SIGHUP)
		for {
			sig := <-reloadSignal
			ep.Reload()
			fmt.Printf("OS signaled `%v`, reload", sig.String())
		}
	}()
}

func Execute() {
	rootCmd.AddCommand(daemon.Cmd, product_service.Cmd, health_check_service.Cmd, migrate.Cmd, jwt.Cmd)
	if e := rootCmd.Execute(); e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", e.Error())
		os.Exit(1)
	}
}
