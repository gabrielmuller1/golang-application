package provider

import (
	"github.com/aristat/golang-example-app/app/casbin"
	"github.com/aristat/golang-example-app/app/config"
	"github.com/aristat/golang-example-app/app/entrypoint"
	"github.com/aristat/golang-example-app/app/grpc"
	"github.com/aristat/golang-example-app/app/logger"
	"github.com/aristat/golang-example-app/app/tracing"
	"github.com/google/wire"
)

var AwareProductionSet = wire.NewSet(
	entrypoint.ProviderProductionSet,
	logger.ProviderProductionSet,
	config.ProviderSet,
	tracing.ProviderProductionSet,
	grpc.ProviderProductionSet,
	casbin.ProviderProductionSet,
)

var AwareTestSet = wire.NewSet(
	entrypoint.ProviderTestSet,
	logger.ProviderTestSet,
	tracing.ProviderTestSet,
	grpc.ProviderTestSet,
	casbin.ProviderTestSet,
)
