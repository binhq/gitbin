package format

import (
	"io"

	gitbin "github.com/binhq/gitbin/apis/gitbin/v1alpha1"
	"github.com/sagikazarmark/utilz/archive/tar"
)

// TargzUnpacker unpacks from a tar.gz archive
type TargzUnpacker struct{}

// Unpack implements the Unpacker interface
func (u *TargzUnpacker) Unpack(r io.Reader, download *gitbin.BinaryDownload) (io.Reader, error) {
	return tar.NewTarGzFileReader(r, download.Path)
}
