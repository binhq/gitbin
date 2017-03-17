package api

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"

	githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"
	version "github.com/hashicorp/go-version"
)

// Rule is a template based on which a version of a project can be downloaded and extracted
type Rule struct {
	UrlTemplate  string
	Format       githubin.BinaryDownload_Format
	PathTemplate string
}

// FindBinary finds a binary in the local rule list
func FindBinary(search *githubin.BinarySearch) (*githubin.BinaryDownload, error) {
	repo := fmt.Sprintf("%s/%s", search.GetOwner(), search.GetRepository())

	rules, ok := repositories[repo]
	if !ok {
		return nil, errors.New("Repository not found in the rule list")
	}

	// TODO: fallback to latest if empty
	if search.Version == "" {
		return nil, errors.New("Empty version")
	}

	v, err := version.NewVersion(search.Version)
	if err != nil {
		return nil, errors.New("Version cannot be parsed")
	}

	var currentRule *Rule

	for constraint, rule := range rules {
		if constraint == "*" {
			currentRule = &rule
			break
		} else {
			c, err := version.NewConstraint(constraint)
			if err != nil {
				continue
			}

			if c.Check(v) {
				currentRule = &rule
				break
			}
		}
	}

	if currentRule == nil {
		return nil, errors.New("Rule not found for repository")
	}

	urlTmpl, err := template.New("url").Parse(currentRule.UrlTemplate)
	if err != nil {
		return nil, errors.New("Cannot parse URL template")
	}

	pathTmpl, err := template.New("path").Parse(currentRule.PathTemplate)
	if err != nil {
		return nil, errors.New("Cannot parse Path template")
	}

	urlBuf := &bytes.Buffer{}
	if err := urlTmpl.Execute(urlBuf, search); err != nil {
		return nil, errors.New("Cannot execute URL template")
	}

	pathBuf := &bytes.Buffer{}
	if err := pathTmpl.Execute(pathBuf, search); err != nil {
		return nil, errors.New("Cannot execute Path template")
	}

	return &githubin.BinaryDownload{
		Url:    urlBuf.String(),
		Format: currentRule.Format,
		Path:   pathBuf.String(),
	}, nil
}
