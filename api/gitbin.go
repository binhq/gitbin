package api

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/Masterminds/semver"
	"github.com/Masterminds/sprig"
	gitbin "github.com/binhq/gitbin/apis/gitbin/v1alpha1"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Githubin implements the gRPC client mimicing the future API behaviour
type Githubin struct{}

// Rule is a template based on which a version of a project can be downloaded and extracted
type Rule struct {
	UrlTemplate  string
	Format       gitbin.BinaryDownload_Format
	PathTemplate string
}

// FindBinary finds a binary in the local rule list
func (g *Githubin) FindBinary(ctx context.Context, search *gitbin.BinarySearch, opts ...grpc.CallOption) (*gitbin.BinaryDownload, error) {
	repo := fmt.Sprintf("%s/%s", search.GetOwner(), search.GetRepository())

	rules, ok := repositories[repo]
	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "repository not found in the rule list")
	}

	// TODO: fallback to latest if empty?
	if search.Version == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "empty version")
	}

	v, err := semver.NewVersion(search.Version)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "version cannot be parsed")
	}

	var currentRule *Rule

	for constraint, rule := range rules {
		c, err := semver.NewConstraint(constraint)
		if err != nil {
			continue
		}

		if c.Check(v) {
			currentRule = &rule
			break
		}
	}

	if currentRule == nil {
		return nil, grpc.Errorf(codes.NotFound, "rule not found for repository")
	}

	tplFuncs := sprig.FuncMap()

	// TODO: replace this with more efficient code
	tplFuncs["archReplace"] = func(s string) string {
		switch s {
		case "386":
			return "i386"

		case "amd64":
			return "x86_64"
		}

		return s
	}

	urlTmpl, err := template.New("url").Funcs(tplFuncs).Parse(currentRule.UrlTemplate)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, "cannot parse URL template: %v", err)
	}

	pathTmpl, err := template.New("path").Funcs(tplFuncs).Parse(currentRule.PathTemplate)
	if err != nil {
		return nil, grpc.Errorf(codes.Unknown, "cannot parse Path template: %v", err)
	}

	urlBuf := &bytes.Buffer{}
	if err := urlTmpl.Execute(urlBuf, search); err != nil {
		return nil, grpc.Errorf(codes.Unknown, "cannot execute URL template")
	}

	pathBuf := &bytes.Buffer{}
	if err := pathTmpl.Execute(pathBuf, search); err != nil {
		return nil, grpc.Errorf(codes.Unknown, "cannot execute Path template")
	}

	return &gitbin.BinaryDownload{
		Url:    urlBuf.String(),
		Format: currentRule.Format,
		Path:   pathBuf.String(),
	}, nil
}
