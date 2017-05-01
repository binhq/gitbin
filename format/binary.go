package format

import (
	"io"

	gitbin "github.com/binhq/gitbin/apis/gitbin/v1alpha1"
)

// BinaryUnpacker unpacks from plain binary format
type BinaryUnpacker struct{}

// Unpack implements the Unpacker interface
func (u *BinaryUnpacker) Unpack(r io.Reader, download *gitbin.BinaryDownload) (io.Reader, error) {
	return r, nil
}
