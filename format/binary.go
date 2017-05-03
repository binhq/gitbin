package format

import (
	"io"

	binstack "github.com/binhq/gitbin/apis/binstack/v1alpha1"
)

// BinaryUnpacker unpacks from plain binary format
type BinaryUnpacker struct{}

// Unpack implements the Unpacker interface
func (u *BinaryUnpacker) Unpack(r io.Reader, downloadInfo *binstack.DownloadInfo) (io.Reader, error) {
	return r, nil
}
