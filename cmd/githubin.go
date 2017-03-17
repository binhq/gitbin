package cmd

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/binhq/githubin/apis/githubin/v1alpha1"
)

type downloadTemplate struct {
	Url  string
	Type int32
	Path string
}

var downloads = map[string]downloadTemplate{
	"Masterminds/glide": downloadTemplate{
		Url:  "https://github.com/{{.Owner}}/{{.Repository}}/releases/download/v{{.Version}}/glide-v{{.Version}}-{{.Os}}-{{.Arch}}.tar.gz",
		Type: 1,
		Path: "{{.Os}}-{{.Arch}}/glide",
	},
	"mattes/migrate": downloadTemplate{
		Url:  "https://github.com/{{.Owner}}/{{.Repository}}/releases/download/v{{.Version}}/migrate.{{.Os}}-{{.Arch}}.tar.gz",
		Type: 1,
		Path: "./migrate.{{.Os}}-{{.Arch}}",
	},
}

func findBinary(search *githubin.BinarySearch) *githubin.BinaryDownload {
	repo := fmt.Sprintf("%s/%s", search.GetOwner(), search.GetRepository())

	downloadTemplate, ok := downloads[repo]
	if !ok {
		return nil
	}

	urlTmpl, err := template.New("url").Parse(downloadTemplate.Url)
	if err != nil {
		return nil
	}

	pathTmpl, err := template.New("path").Parse(downloadTemplate.Path)
	if err != nil {
		return nil
	}

	urlBuf := &bytes.Buffer{}
	if err := urlTmpl.Execute(urlBuf, search); err != nil {
		return nil
	}

	pathBuf := &bytes.Buffer{}
	if err := pathTmpl.Execute(pathBuf, search); err != nil {
		return nil
	}

	return &githubin.BinaryDownload{
		Url:  urlBuf.String(),
		Type: githubin.BinaryDownload_DownloadType(downloadTemplate.Type),
		Path: pathBuf.String(),
	}
}
