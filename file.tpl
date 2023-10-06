// @Name: {{.fileName}}.go
// @Date: {{.createAt}}
// @Author: {{.author}}

package {{.packageName}}
{{if ne .fileName "serializer" -}}
{{if eq .fileName "router"}}
import "github.com/gin-gonic/gin"
{{end -}}
{{if and .wire .file}}
{{if eq .fileName "provider" -}}
import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	New{{.PackageName}}Handler,
	New{{.PackageName}}Router,
	New{{.PackageName}}Dao,
	New{{.PackageName}}Service,
	//New{{.PackageName}}Cache,
){{- end -}}

{{- if ne .fileName "provider" -}}
type I{{.PackageName}}{{.FileName}} interface {
    {{if eq .fileName "router" -}}
    Router(r *gin.Engine)
    {{- end}}
}

type {{.packageName}}{{.FileName}} struct {
    {{.param}} {{.di}}
}

func New{{.PackageName}}{{.FileName}}({{.param}} {{.di}}) I{{.PackageName}}{{.FileName}} {
    {{if ne .param "" -}}
    return &{{.packageName}}{{.FileName}}{
		{{.param}}: {{.param}},
	}{{else if eq .param "" -}}
    return &{{.packageName}}{{.FileName}}{
    }{{- end}}
}

{{if eq .fileName "router" -}}
func (router {{.packageName}}{{.FileName}}) Router(r *gin.Engine) {

}
{{- end}}
{{- end}}
{{else}}
{{if eq .fileName "router" -}}
func Router(r *gin.Engine) {

}
{{- end}}
{{- end}}
{{- end}}