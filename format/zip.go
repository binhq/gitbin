package format

import (
	"io"

	binstack "github.com/binhq/gitbin/apis/binstack/v1alpha1"
	"github.com/goph/stdlib/archive/zip"
)

// ZipUnpacker unpacks from a zip archive.
type ZipUnpacker struct{}

// Unpack implements the Unpacker interface.
func (u *ZipUnpacker) Unpack(r io.Reader, downloadInfo *binstack.DownloadInfo) (io.Reader, error) {
	return zip.NewZipFileReader(r, downloadInfo.GetPath())
}
