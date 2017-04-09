package format

import (
	"fmt"
	"io"

	githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"
)

// Unpacker unpacks the binary from a download format
type Unpacker interface {
	Unpack(r io.Reader, download *githubin.BinaryDownload) (io.Reader, error)
}

// AutoUnpacker detects the format and delegates the process to the appropriate unpacker
// This also means that the underlying unpackers MUST be stateless
type AutoUnpacker struct {
	unpackers map[githubin.BinaryDownload_Format]Unpacker
}

// NewAutoUnpacker returns a new AutoUnpacker
func NewAutoUnpacker() Unpacker {
	return &AutoUnpacker{
		unpackers: map[githubin.BinaryDownload_Format]Unpacker{
			githubin.BinaryDownload_BINARY: &BinaryUnpacker{},
			githubin.BinaryDownload_TARGZ:  &TargzUnpacker{},
		},
	}
}

// Unpack implements the Unpacker interface
func (u *AutoUnpacker) Unpack(r io.Reader, download *githubin.BinaryDownload) (io.Reader, error) {
	unpacker, ok := u.unpackers[download.GetFormat()]
	if !ok {
		return nil, fmt.Errorf("Unknown format: %s", download.GetFormat())
	}

	return unpacker.Unpack(r, download)
}
