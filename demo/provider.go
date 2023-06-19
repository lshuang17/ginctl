// @Name: provider.go
// @Date: 2023-06-19
// @Author: ls

package demo

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewDemoHandler,
	NewDemoRouter,
	NewDemoDao,
	NewDemoService,
	//NewDemoCache,
)
