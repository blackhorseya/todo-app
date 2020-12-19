package task

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// IHandler declare task api
type IHandler interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	Remove(c *gin.Context)
	ModifyInfo(c *gin.Context)
}

// ProviderSet is a task of api provider set
var ProviderSet = wire.NewSet(NewImpl)
