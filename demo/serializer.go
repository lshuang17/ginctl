// @Name: serializer.go
// @Date: 2023-06-19
// @Author: ls

package demo

type IDemoSerializer interface {
}

type demoSerializer struct {
    handler IDemoHandler
}

func NewDemoSerializer(handler IDemoHandler) IDemoSerializer {
    return &demoSerializer{
		handler: handler,
	}
}
