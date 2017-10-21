package api

// repositories is a static list of binstack Rules
// TODO: make it a database and create a file format
var repositories = map[string]map[string]Rule{
	"Masterminds/glide": {
		"*": {
			Homepage:     "https://glide.sh",
			Description:  "Package Management for Golang",
			UrlTemplate:  "v{{.Version}}/glide-v{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz",
			Format:       2, // binstack.DownloadInfo_TARGZ
			PathTemplate: "{{.Os}}-{{.Arch}}/glide",
		},
	},
	"mattes/migrate": {
		"*": {
			Description:  " Database migrations. CLI and Golang library.",
			UrlTemplate:  "v{{.Version}}/migrate.{{.Os}}-{{.Arch}}.tar.gz",
			Format:       2,                             // binstack.DownloadInfo_TARGZ
			PathTemplate: "./migrate.{{.Os}}-{{.Arch}}", // TODO: do not require ./ at the beggining of the path???
		},
	},
	"goreleaser/goreleaser": {
		"*": {
			Homepage:     "https://goreleaser.github.io/",
			Description:  "Deliver Go binaries as fast and easily as possible",
			UrlTemplate:  "v{{.Version}}/goreleaser_{{.Os | title}}_{{.Arch | archReplace}}.tar.gz",
			Format:       2, // binstack.DownloadInfo_TARGZ
			PathTemplate: "goreleaser",
		},
	},
	"golang/dep": {
		">0.3.0": {
			Homepage:    "https://github.com/golang/dep",
			Description: "Go dependency management tool",
			UrlTemplate: "v{{.Version}}/dep-{{.Os}}-{{.Arch}}",
			Format:      1, // binstack.DownloadInfo_BINARY
		},
		"<=0.3.0": {
			Homepage:     "https://github.com/golang/dep",
			Description:  "Go dependency management tool",
			UrlTemplate:  "v{{.Version}}/dep-{{.Os}}-{{.Arch}}.zip",
			Format:       3, // binstack.DownloadInfo_ZIP
			PathTemplate: "dep",
		},
	},
}
