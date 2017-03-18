package format

import (
	"archive/tar"
	"compress/gzip"
	"io"

	"errors"

	githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"
)

type targzBinaryReader struct {
	gzipReader io.ReadCloser
	tarReader  *tar.Reader
}

func newTargzBinaryReader(r io.Reader, path string) (io.ReadCloser, error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	tr := tar.NewReader(gz)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			gz.Close()

			return nil, errors.New("Binary not found in archive")
		} else if err != nil {
			gz.Close()

			return nil, err
		}

		if header.Name == path {
			break
		}
	}

	return &targzBinaryReader{
		gzipReader: gz,
		tarReader:  tr,
	}, nil
}

func (r *targzBinaryReader) Read(p []byte) (int, error) {
	return r.tarReader.Read(p)
}

// Close implements io.ReadCloser
func (r *targzBinaryReader) Close() error {
	return r.gzipReader.Close()
}

// TargzUnpacker unpacks from a tar.gz archive
func TargzUnpacker(r io.Reader, download *githubin.BinaryDownload) (io.Reader, error) {
	return newTargzBinaryReader(r, download.Path)
}
