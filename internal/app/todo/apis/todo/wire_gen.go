// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package todo

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// Injectors from wire.go:

// CreateIHandler serve user to create health api handler
func CreateIHandler(logger *zap.Logger, biz todo.IBiz) (IHandler, error) {
	iHandler := NewImpl(logger, biz)
	return iHandler, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
