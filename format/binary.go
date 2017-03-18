package format

import (
	"io"

	githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"
)

// BinaryUnpacker unpacks from plain binary format
func BinaryUnpacker(r io.Reader, download *githubin.BinaryDownload) (io.Reader, error) {
	return r, nil
}
