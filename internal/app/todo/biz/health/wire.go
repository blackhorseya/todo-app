// +build wireinject

package health

import (
	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repo"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(
	NewImpl,
)

// CreateIBiz serve user to create health biz
func CreateIBiz(logger *zap.Logger, repo repo.IRepo) (IBiz, error) {
	panic(wire.Build(testProviderSet))
}
