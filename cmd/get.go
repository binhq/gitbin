package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/binhq/gitbin/api"
	gitbin "github.com/binhq/gitbin/apis/gitbin/v1alpha1"
	"github.com/binhq/gitbin/format"
	"github.com/spf13/cobra"
	context "golang.org/x/net/context"
)

type getOpts struct {
	output   string
	fileMode int
	os       string
	arch     string
}

type getCommand struct {
	opts     *getOpts
	gitbin   gitbin.GitbinClient
	unpacker format.Unpacker
}

func (g *getCommand) Run(cmd *cobra.Command, args []string) error {
	repo := strings.Split(args[0], "/")

	if len(repo) != 2 {
		return fmt.Errorf("Invalid repository: \"%s\"", args[0])
	}

	search := &gitbin.BinarySearch{
		Owner:      repo[0],
		Repository: repo[1],
		Version:    args[1],
		Os:         g.opts.os,
		Arch:       g.opts.arch,
	}

	logger.WithFields(logrus.Fields{
		"repository": args[0],
		"version":    args[1],
	}).Info("Searching binary")

	ctx := context.Background()
	download, err := g.gitbin.FindBinary(ctx, search)
	if err != nil {
		logger.Fatalf("Cannot find binary: %v", err)
	}

	logger.WithField("format", download.Format).Info("Binary found")
	logger.WithField("url", download.Url).Info("Downloading binary")

	resp, err := http.Get(download.Url)
	if err != nil {
		logger.Fatalf("Download failed: %v", err)
	}
	if resp.StatusCode != 200 {
		resp.Body.Close()

		if resp.StatusCode == 404 {
			logger.Fatal("Download failed: file not found")
		} else {
			logger.Fatal("Download failed: unknown reason")
		}
	}
	defer resp.Body.Close()

	// TODO: is this a good idea?
	logrus.RegisterExitHandler(func() {
		resp.Body.Close()
	})

	binary, err := g.unpacker.Unpack(resp.Body, download)
	if err != nil {
		logger.Fatalf("Unpacking failed: %v", err)
	}

	if r, ok := binary.(io.Closer); ok {
		defer r.Close()
	}

	// TODO: Create the directory if does not exists
	if _, err := os.Stat(g.opts.output); os.IsNotExist(err) {
		log.Fatal(err)
	}

	target := fmt.Sprintf(
		"%s/%s",
		g.opts.output,
		search.Repository,
	)

	logger.WithField("target", target).Info("Saving binary")

	file, err := os.OpenFile(
		target,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		os.FileMode(g.opts.fileMode),
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, binary)
	if err != nil {
		logger.Fatal(err)
	}

	return nil
}

func init() {
	var opts getOpts

	getCmd := &getCommand{
		opts:     &opts,
		gitbin:   &api.Githubin{},
		unpacker: format.NewAutoUnpacker(),
	}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Download a binary to a given path (or current directory)",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("Not enough arguments")
			}

			return nil
		},
		RunE: getCmd.Run,
	}

	rootCmd.AddCommand(cmd)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not determine working directory")
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", cwd, "Output directory")
	cmd.Flags().IntVarP(&opts.fileMode, "mode", "m", int(os.ModePerm), "File mode")
	cmd.Flags().StringVar(&opts.os, "os", runtime.GOOS, "Target OS (if matters)")
	cmd.Flags().StringVar(&opts.arch, "arch", runtime.GOARCH, "Target Arch (if matters)")
}
