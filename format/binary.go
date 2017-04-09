package format

import (
	"io"

	githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"
)

// BinaryUnpacker unpacks from plain binary format
type BinaryUnpacker struct{}

// Unpack implements the Unpacker interface
func (u *BinaryUnpacker) Unpack(r io.Reader, download *githubin.BinaryDownload) (io.Reader, error) {
	return r, nil
}
