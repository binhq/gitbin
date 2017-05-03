package format

import (
	"fmt"
	"io"

	binstack "github.com/binhq/gitbin/apis/binstack/v1alpha1"
)

// Unpacker unpacks the binary from a download format
type Unpacker interface {
	Unpack(r io.Reader, download *binstack.DownloadInfo) (io.Reader, error)
}

// AutoUnpacker detects the format and delegates the process to the appropriate unpacker
// This also means that the underlying unpackers MUST be stateless
type AutoUnpacker struct {
	unpackers map[binstack.DownloadInfo_Format]Unpacker
}

// NewAutoUnpacker returns a new AutoUnpacker
func NewAutoUnpacker() Unpacker {
	return &AutoUnpacker{
		unpackers: map[binstack.DownloadInfo_Format]Unpacker{
			binstack.DownloadInfo_BINARY: &BinaryUnpacker{},
			binstack.DownloadInfo_TARGZ:  &TargzUnpacker{},
		},
	}
}

// Unpack implements the Unpacker interface
func (u *AutoUnpacker) Unpack(r io.Reader, downloadInfo *binstack.DownloadInfo) (io.Reader, error) {
	unpacker, ok := u.unpackers[downloadInfo.GetFormat()]
	if !ok {
		return nil, fmt.Errorf("Unknown format: %s", downloadInfo.GetFormat())
	}

	return unpacker.Unpack(r, downloadInfo)
}
