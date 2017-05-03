package format

import (
	"io"

	binstack "github.com/binhq/gitbin/apis/binstack/v1alpha1"
	"github.com/sagikazarmark/utilz/archive/tar"
)

// TargzUnpacker unpacks from a tar.gz archive
type TargzUnpacker struct{}

// Unpack implements the Unpacker interface
func (u *TargzUnpacker) Unpack(r io.Reader, downloadInfo *binstack.DownloadInfo) (io.Reader, error) {
	return tar.NewTarGzFileReader(r, downloadInfo.GetPath())
}
