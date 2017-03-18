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

	"github.com/binhq/githubin/api"
	githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"
	"github.com/binhq/githubin/format"
	"github.com/spf13/cobra"
)

const fileFlags = os.O_CREATE | os.O_TRUNC | os.O_WRONLY

var (
	binaryOutput string
	binaryMode   int
	binaryOs     string
	binaryArch   string
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download a binary to a given path (or current directory)",
	RunE: func(cmd *cobra.Command, args []string) error {
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

		download, err := api.FindBinary(search)
		if err != nil {
			log.Fatal(err)
		}

		unpacker, err := format.FindUnpacker(download.Format)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Downloading %s\n", download.Url)

		resp, err := http.Get(download.Url)
		if err != nil {
			log.Fatalf("Download failed: %v", err)
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
			panic(err)
		}

		// TODO: Create the directory if does not exists
		if _, err := os.Stat(binaryOutput); os.IsNotExist(err) {
			log.Fatal(err)
		}

		file, err := os.OpenFile(fmt.Sprintf("%s/%s", binaryOutput, search.Repository), fileFlags, os.FileMode(binaryMode))
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = io.Copy(file, binary)
		if err != nil {
			panic(err)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(getCmd)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not determine working directory")
	}

	getCmd.Flags().StringVarP(&binaryOutput, "output", "o", cwd, "Output directory")
	getCmd.Flags().IntVarP(&binaryMode, "mode", "m", int(os.ModePerm), "File mode")
	getCmd.Flags().StringVar(&binaryOs, "os", runtime.GOOS, "Target OS (if matters)")
	getCmd.Flags().StringVar(&binaryArch, "arch", runtime.GOARCH, "Target Arch (if matters)")
}
