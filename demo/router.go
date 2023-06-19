// @Name: router.go
// @Date: 2023-06-19
// @Author: ls

package demo

import "github.com/gin-gonic/gin"

type IDemoRouter interface {
}

type demoRouter struct {
    handler IDemoHandler
}

func NewDemoRouter(handler IDemoHandler) IDemoRouter {
    
    return &demoRouter{
		handler: handler,
	}
}

func (router demoRouter) Router(r *gin.Engine) {
}


