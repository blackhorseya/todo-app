package health

import (
	"time"

	"github.com/blackhorseya/todo-app/internal/app/todo/biz/health/repo"
	"github.com/blackhorseya/todo-app/internal/pkg/base/contextx"
	"github.com/blackhorseya/todo-app/internal/pkg/entity/er"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IRepo
}

// NewImpl is a constructor of implement business with parameters
func NewImpl(logger *zap.Logger, healthRepo repo.IRepo) IBiz {
	return &impl{
		logger: logger.With(zap.String("type", "biz")),
		repo:   healthRepo,
	}
}

// Readiness to handle application has been ready
func (i *impl) Readiness(ctx contextx.Contextx) (ok bool, err error) {
	err = i.repo.Ping(ctx, 2*time.Second)
	if err != nil {
		i.logger.Error(er.ErrPing.Error(), zap.Error(err))
		return false, er.ErrPing
	}

	return true, nil
}

// Liveness to handle application was alive
func (i *impl) Liveness(ctx contextx.Contextx) (ok bool, err error) {
	err = i.repo.Ping(ctx, 5*time.Second)
	if err != nil {
		i.logger.Error(er.ErrPing.Error(), zap.Error(err))
		return false, er.ErrPing
	}

	return true, nil
}
