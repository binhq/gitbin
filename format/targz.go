package format

import (
	"archive/tar"
	"compress/gzip"
	"io"

	"errors"

	"bytes"

	githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"
)

// TargzUnpacker unpacks from a tar.gz archive
func TargzUnpacker(r io.Reader, download *githubin.BinaryDownload) (io.Reader, error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	tr := tar.NewReader(gz)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if header.Name == download.Path {
			buf := new(bytes.Buffer)
			_, err := io.Copy(buf, tr)
			if err != nil {
				return nil, err
			}

			return buf, nil
		}
	}

	return nil, errors.New("Binary not found in archive")
}
