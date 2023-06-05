{{define "provider"}}
import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	New{{.PackageName}}Handler,
	New{{.PackageName}}Router,
	New{{.PackageName}}Repo,
	New{{.PackageName}}Service,
	//New{{.PackageName}}Cache,
)
{{end}}