// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package health

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repo"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// Injectors from wire.go:

// CreateIBiz serve user to create health biz
func CreateIBiz(logger *zap.Logger, repo2 repo.IRepo) (IBiz, error) {
	iBiz := NewImpl(logger, repo2)
	return iBiz, nil
}

// wire.go:

var testProviderSet = wire.NewSet(
	NewImpl,
)
