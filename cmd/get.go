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
	"github.com/binhq/githubin/api"
	githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"
	"github.com/binhq/githubin/format"
	"github.com/spf13/cobra"
	context "golang.org/x/net/context"
)

type getOpts struct {
	binaryOutput string
	binaryMode   int
	binaryOs     string
	binaryArch   string
}

type getCommand struct {
	opts     *getOpts
	githubin githubin.GithubinClient
}

func (g *getCommand) Run(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("Not enough arguments")
	}

	repo := strings.Split(args[0], "/")

	if len(repo) != 2 {
		return fmt.Errorf("Invalid repository: \"%s\"", args[0])
	}

	search := &githubin.BinarySearch{
		Owner:      repo[0],
		Repository: repo[1],
		Version:    args[1],
		Os:         cmd.Flag("os").Value.String(),
		Arch:       cmd.Flag("arch").Value.String(),
	}

	logger.WithFields(logrus.Fields{
		"repository": args[0],
		"version":    args[1],
	}).Info("Searching binary")

	ctx := context.Background()
	download, err := g.githubin.FindBinary(ctx, search)
	if err != nil {
		logger.Fatal(err)
	}

	logger.WithField("format", download.Format).Info("Binary found")

	unpacker, err := format.FindUnpacker(download.Format)
	if err != nil {
		logger.Fatal(err)
	}

	logger.WithField("url", download.Url).Info("Downloading binary")

	resp, err := http.Get(download.Url)
	if err != nil {
		logger.Fatalf("Download failed: %v", err)
	}
	if resp.StatusCode != 200 {
		resp.Body.Close()

		if resp.StatusCode == 404 {
			log.Fatal("File not found")
		} else {
			log.Fatal("Download failed")
		}
	}
	defer resp.Body.Close()

	binary, err := unpacker(resp.Body, download)
	if err != nil {
		logger.Panic(err)
	}

	if r, ok := binary.(io.ReadCloser); ok {
		defer r.Close()
	}

	// TODO: Create the directory if does not exists
	if _, err := os.Stat(g.opts.binaryOutput); os.IsNotExist(err) {
		log.Fatal(err)
	}

	target := fmt.Sprintf(
		"%s/%s",
		g.opts.binaryOutput,
		search.Repository,
	)

	logger.WithField("target", target).Info("Unpacking")

	file, err := os.OpenFile(
		target,
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
		os.FileMode(g.opts.binaryMode),
	)
	if err != nil {
		logger.Panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, binary)
	if err != nil {
		logger.Panic(err)
	}

	return nil
}

func init() {
	var opts getOpts

	getCmd := &getCommand{
		opts:     &opts,
		githubin: &api.Githubin{},
	}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Download a binary to a given path (or current directory)",
		RunE:  getCmd.Run,
	}

	rootCmd.AddCommand(cmd)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not determine working directory")
	}

	cmd.Flags().StringVarP(&opts.binaryOutput, "output", "o", cwd, "Output directory")
	cmd.Flags().IntVarP(&opts.binaryMode, "mode", "m", int(os.ModePerm), "File mode")
	cmd.Flags().StringVar(&opts.binaryOs, "os", runtime.GOOS, "Target OS (if matters)")
	cmd.Flags().StringVar(&opts.binaryArch, "arch", runtime.GOARCH, "Target Arch (if matters)")
}
