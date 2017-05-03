package api

import (
	"bytes"
	"html/template"
	"strings"

	"fmt"

	"github.com/Masterminds/semver"
	"github.com/Masterminds/sprig"
	binstack "github.com/binhq/gitbin/apis/binstack/v1alpha1"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Githubin implements the gRPC client mimicing the future API behaviour
type Githubin struct{}

// Rule is a template based on which a version of a project can be downloaded and extracted
type Rule struct {
	Homepage     string
	Description  string
	UrlTemplate  string
	Format       binstack.DownloadInfo_Format
	PathTemplate string
}

// FindBinary finds a binary in the local rule list
func (g *Githubin) FindBinary(ctx context.Context, search *binstack.BinarySearch, opts ...grpc.CallOption) (*binstack.Binary, error) {
	repo := search.GetName()

	rules, ok := repositories[repo]
	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "repository not found in the rule list")
	}

	// TODO: fallback to latest if empty?
	if search.Version == nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "empty version")
	}

	var v *semver.Version
	var err error

	// We do not allow constraints for now
	if search.GetExactVersion() != "" {
		if v, err = semver.NewVersion(search.GetExactVersion()); err != nil {
			return nil, grpc.Errorf(codes.InvalidArgument, "version cannot be parsed")
		}
	} else {
		// TODO: Find the latest matching version here
		return nil, grpc.Errorf(codes.InvalidArgument, "empty version")
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

	args := map[string]string{
		"Version": v.String(),
		"Os":      search.GetOs(),
		"Arch":    search.GetArch(),
	}

	urlBuf := &bytes.Buffer{}
	if err := urlTmpl.Execute(urlBuf, args); err != nil {
		return nil, grpc.Errorf(codes.Unknown, "cannot execute URL template")
	}

	pathBuf := &bytes.Buffer{}
	if err := pathTmpl.Execute(pathBuf, args); err != nil {
		return nil, grpc.Errorf(codes.Unknown, "cannot execute Path template")
	}

	homepage := currentRule.Homepage
	if homepage == "" {
		homepage = fmt.Sprintf("https://github.com/%s", repo)
	}

	url := urlBuf.String()
	if strings.HasPrefix(url, "http") == false {
		url = fmt.Sprintf("https://github.com/%s/releases/download/%s", repo, url)
	}

	return &binstack.Binary{
		Homepage:    homepage,
		Description: currentRule.Description,
		Version:     v.String(),
		DownloadInfo: &binstack.DownloadInfo{
			Url:    url,
			Format: binstack.DownloadInfo_Format(currentRule.Format),
			Path:   pathBuf.String(),
		},
	}, nil
}
