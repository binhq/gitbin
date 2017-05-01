package api

import gitbin "github.com/binhq/gitbin/apis/gitbin/v1alpha1"

// rules is a static list of gitbin Rules
// TODO: make it a database and create a file format
var repositories = map[string]map[string]Rule{
	"Masterminds/glide": map[string]Rule{
		"*": Rule{
			UrlTemplate:  "https://github.com/{{.Owner}}/{{.Repository}}/releases/download/v{{.Version}}/glide-v{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz",
			Format:       gitbin.BinaryDownload_TARGZ,
			PathTemplate: "{{.Os}}-{{.Arch}}/glide",
		},
	},
	"mattes/migrate": map[string]Rule{
		"3.0.0-rc2": Rule{
			UrlTemplate:  "https://github.com/{{.Owner}}/{{.Repository}}/releases/download/v{{.Version}}/migrate.{{.Os}}-{{.Arch}}.tar.gz",
			Format:       gitbin.BinaryDownload_TARGZ,
			PathTemplate: "./migrate.{{.Os}}-{{.Arch}}", // TODO: do not require ./ at the beggining of the path???
		},
		"*": Rule{
			UrlTemplate:  "https://github.com/{{.Owner}}/{{.Repository}}/releases/download/v{{.Version}}/migrate.{{.Os}}-{{.Arch}}.tar.gz",
			Format:       gitbin.BinaryDownload_TARGZ,
			PathTemplate: "./migrate.{{.Os}}-{{.Arch}}", // TODO: do not require ./ at the beggining of the path???
		},
	},
	"goreleaser/goreleaser": map[string]Rule{
		"*": Rule{
			UrlTemplate:  "https://github.com/{{.Owner}}/{{.Repository}}/releases/download/v{{.Version}}/goreleaser_{{.Os | title}}_{{.Arch | archReplace}}.tar.gz",
			Format:       gitbin.BinaryDownload_TARGZ,
			PathTemplate: "goreleaser", // TODO: do not require ./ at the beggining of the path???
		},
	},
}
