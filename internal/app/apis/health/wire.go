// +build wireinject

package health

import (
	"github.com/blackhorseya/todo-app/internal/app/biz/health"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

// CreateHealthHandler serve user to create health api handler
func CreateHealthHandler(biz health.Biz) (IHandler, error) {
	panic(wire.Build(testProviderSet))
}
