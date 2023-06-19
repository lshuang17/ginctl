// @Name: {{.fileName}}.go
// @Date: {{.createAt}}
// @Author: {{.author}}

package {{.packageName}}

import "github.com/gin-gonic/gin"
{{if and .wire}}
type I{{.PackageName}}{{.FileName}} interface {
}

type {{.packageName}}{{.FileName}} struct {
    {{.param}} {{.di}}
}

func New{{.PackageName}}{{.FileName}}({{.param}} {{.di}}) I{{.PackageName}}{{.FileName}} {
    {{if ne .param ""}}
    return &{{.packageName}}{{.FileName}}{
		{{.param}}: {{.param}},
	}{{else if eq .param ""}}
    return &{{.packageName}}{{.FileName}}{
    }{{end}}
}

func (router {{.packageName}}{{.FileName}}) Router(r *gin.Engine) {
}
{{else}}
func Router(r *gin.Engine) {
}
{{end}}

