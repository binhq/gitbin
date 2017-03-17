package api

import githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"

// rules is a static list of githubin Rules
// TODO: make it a database and create a file format
var repositories = map[string]map[string]Rule{
	"Masterminds/glide": map[string]Rule{
		"*": Rule{
			UrlTemplate:  "https://github.com/{{.Owner}}/{{.Repository}}/releases/download/v{{.Version}}/glide-v{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz",
			Format:       githubin.BinaryDownload_TARGZ,
			PathTemplate: "{{.Os}}-{{.Arch}}/glide",
		},
	},
	"mattes/migrate": map[string]Rule{
		"*": Rule{
			UrlTemplate:  "https://github.com/{{.Owner}}/{{.Repository}}/releases/download/v{{.Version}}/migrate.{{.Os}}-{{.Arch}}.tar.gz",
			Format:       githubin.BinaryDownload_TARGZ,
			PathTemplate: "./migrate.{{.Os}}-{{.Arch}}", // TODO: do not require ./ at the beggining of the path???
		},
	},
}
