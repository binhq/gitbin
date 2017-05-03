package api

// repositories is a static list of binstack Rules
// TODO: make it a database and create a file format
var repositories = map[string]map[string]Rule{
	"Masterminds/glide": map[string]Rule{
		"*": Rule{
			Homepage:     "https://glide.sh",
			Description:  "Package Management for Golang",
			UrlTemplate:  "v{{.Version}}/glide-v{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz",
			Format:       2, // binstack.DownloadInfo_TARGZ
			PathTemplate: "{{.Os}}-{{.Arch}}/glide",
		},
	},
	"mattes/migrate": map[string]Rule{
		"*": Rule{
			Description:  " Database migrations. CLI and Golang library.",
			UrlTemplate:  "v{{.Version}}/migrate.{{.Os}}-{{.Arch}}.tar.gz",
			Format:       2,                             // binstack.DownloadInfo_TARGZ
			PathTemplate: "./migrate.{{.Os}}-{{.Arch}}", // TODO: do not require ./ at the beggining of the path???
		},
	},
	"goreleaser/goreleaser": map[string]Rule{
		"*": Rule{
			Homepage:     "https://goreleaser.github.io/",
			Description:  "Deliver Go binaries as fast and easily as possible",
			UrlTemplate:  "v{{.Version}}/goreleaser_{{.Os | title}}_{{.Arch | archReplace}}.tar.gz",
			Format:       2, // binstack.DownloadInfo_TARGZ
			PathTemplate: "goreleaser",
		},
	},
}
