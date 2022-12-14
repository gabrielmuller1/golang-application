// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package products_router

import (
	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/aristat/golang-example-app/app/grpc"
	"github.com/aristat/golang-example-app/app/logger"
	"github.com/aristat/golang-example-app/app/tracing"
)

// Injectors from injector.go:

// Build
func Build() (*Manager, func(), error) {
	context, cleanup, err := entrypoint.ContextProvider()
	if err != nil {
		return nil, nil, err
	}
	viper, cleanup2, err := config.Provider()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	loggerConfig, cleanup3, err := logger.ProviderCfg(viper)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	zap, cleanup4, err := logger.Provider(context, loggerConfig)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	configuration, cleanup5, err := tracing.Cfg(viper)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracerProvider, cleanup6, err := tracing.Provider(context, configuration, zap)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	grpcConfig, cleanup7, err := grpc.Cfg(viper)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	poolManager, cleanup8, err := grpc.Provider(context, tracerProvider, zap, grpcConfig)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceManagers := ServiceManagers{
		PoolManager: poolManager,
	}
	products_routerConfig, cleanup9, err := Cfg(viper)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	manager, cleanup10, err := Provider(context, zap, serviceManagers, products_routerConfig)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return manager, func() {
		cleanup10()
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func BuildTest() (*Manager, func(), error) {
	context, cleanup, err := entrypoint.ContextProviderTest()
	if err != nil {
		return nil, nil, err
	}
	loggerConfig, cleanup2, err := logger.ProviderCfgTest()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mock, cleanup3, err := logger.ProviderTest(context, loggerConfig)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracerProvider, cleanup4, err := tracing.ProviderTest()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	grpcConfig, cleanup5, err := grpc.CfgTest()
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	poolManager, cleanup6, err := grpc.Provider(context, tracerProvider, mock, grpcConfig)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serviceManagers := ServiceManagers{
		PoolManager: poolManager,
	}
	products_routerConfig, cleanup7, err := CfgTest()
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	manager, cleanup8, err := Provider(context, mock, serviceManagers, products_routerConfig)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return manager, func() {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
