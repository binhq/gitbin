package format

import (
	"io"

	githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"
	"github.com/sagikazarmark/utilz/archive/tar"
)

// TargzUnpacker unpacks from a tar.gz archive
type TargzUnpacker struct{}

// Unpack implements the Unpacker interface
func (u *TargzUnpacker) Unpack(r io.Reader, download *githubin.BinaryDownload) (io.Reader, error) {
	return tar.NewTarGzFileReader(r, download.Path)
}
