// @Name: handler.go
// @Date: 2023-06-19
// @Author: ls

package demo

type IDemoHandler interface {
}

type demoHandler struct {
    svc IDemoService
}

func NewDemoHandler(svc IDemoService) IDemoHandler {
    return &demoHandler{
		svc: svc,
	}
}
