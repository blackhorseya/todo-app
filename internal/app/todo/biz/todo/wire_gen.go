// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package todo

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo/repo"
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// Injectors from wire.go:

// CreateIBiz serve user to create health biz
func CreateIBiz(logger *zap.Logger, repo2 repo.IRepo, node *snowflake.Node) (IBiz, error) {
	iBiz := NewImpl(logger, repo2, node)
	return iBiz, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
