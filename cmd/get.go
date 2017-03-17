package cmd

import (
	"archive/tar"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"log"

	"strings"

	"io"

	"compress/gzip"

	githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download a binary to a given path (or current directory)",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("Invalid arguments")
		}

		repo := strings.Split(args[0], "/")

		if len(repo) != 2 {
			log.Fatal("Invalid repository")
		}

		search := &githubin.BinarySearch{
			Owner:      repo[0],
			Repository: repo[1],
			Version:    args[1],
			Os:         cmd.Flag("os").Value.String(),
			Arch:       cmd.Flag("arch").Value.String(),
		}

		download := findBinary(search)
		if download == nil {
			log.Fatal("Not found")
		}

		fmt.Printf("Downloading %s\n", download.GetUrl())

		resp, err := http.Get(download.GetUrl())
		if err != nil {
			log.Fatal("Download failed")
		}
		if resp.StatusCode != 200 {
			resp.Body.Close()

			log.Fatal("Download failed")
		}
		defer resp.Body.Close()

		output := cmd.Flag("output").Value.String()

		if _, err := os.Stat(output); os.IsNotExist(err) {
			os.MkdirAll(output, os.FileMode(os.ModeDir+755))
		}

		file, err := os.OpenFile(fmt.Sprintf("%s/%s", output, search.GetRepository()), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(0777))
		if err != nil {
			panic(err)
		}

		switch download.GetType() {
		case githubin.BinaryDownload_BINARY:
			_, err := io.Copy(file, resp.Body)
			if err != nil {
				panic(err)
			}

		case githubin.BinaryDownload_GZIP:
			gz, err := gzip.NewReader(resp.Body)
			if err != nil {
				panic(err)
			}
			defer gz.Close()

			tr := tar.NewReader(gz)

			for {
				header, err := tr.Next()
				if err == io.EOF {
					break
				} else if err != nil {
					panic(err)
				}

				if header.Name == download.GetPath() {
					_, err := io.Copy(file, tr)
					if err != nil {
						panic(err)
					}

					break
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not determine working directory")
	}

	getCmd.Flags().StringP("output", "o", cwd, "Output directory")
	getCmd.Flags().String("os", runtime.GOOS, "Target OS (if matters)")
	getCmd.Flags().String("arch", runtime.GOARCH, "Target Arch (if matters)")
}
