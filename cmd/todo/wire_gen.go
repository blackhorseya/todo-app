// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/blackhorseya/todo-app/internal/app/todo"
	"github.com/blackhorseya/todo-app/internal/app/todo/apis"
	health2 "github.com/blackhorseya/todo-app/internal/app/todo/apis/health"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repository"
	"github.com/blackhorseya/todo-app/internal/pkg/app"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/config"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/database"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/log"
	"github.com/blackhorseya/todo-app/internal/pkg/infra/transports/http"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateApp(path string) (*app.Application, error) {
	viper, err := config.New(path)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	todoOptions, err := todo.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	client, err := database.NewMongo(databaseOptions)
	if err != nil {
		return nil, err
	}
	healthRepo := repository.NewImpl(client)
	biz := health.NewImpl(healthRepo)
	iHandler := health2.NewImpl(biz)
	initHandlers := apis.CreateInitHandlerFn(iHandler)
	engine := http.NewRouter(httpOptions, logger, initHandlers)
	server, err := http.New(httpOptions, logger, engine)
	if err != nil {
		return nil, err
	}
	application, err := todo.New(todoOptions, logger, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, http.ProviderSet, database.ProviderSet, todo.ProviderSet, apis.ProviderSet, biz.ProviderSet)
