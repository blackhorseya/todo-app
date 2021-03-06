package biz

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health"
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/todo"
	"github.com/google/wire"
)

// ProviderSet is a business provider set
var ProviderSet = wire.NewSet(
	health.ProviderSet,
	todo.ProviderSet,
)
