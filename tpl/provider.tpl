// @Name: {{.fileName}}.go
// @Date: {{.createAt}}
// @Author: {{.author}}

package {{.packageName}}

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	New{{.PackageName}}Handler,
	New{{.PackageName}}Router,
	New{{.PackageName}}Dao,
	New{{.PackageName}}Service,
	//New{{.PackageName}}Cache,
)
