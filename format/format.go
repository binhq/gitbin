package format

import (
	"fmt"
	"io"

	gitbin "github.com/binhq/gitbin/apis/gitbin/v1alpha1"
)

// Unpacker unpacks the binary from a download format
type Unpacker interface {
	Unpack(r io.Reader, download *gitbin.BinaryDownload) (io.Reader, error)
}

// AutoUnpacker detects the format and delegates the process to the appropriate unpacker
// This also means that the underlying unpackers MUST be stateless
type AutoUnpacker struct {
	unpackers map[gitbin.BinaryDownload_Format]Unpacker
}

// NewAutoUnpacker returns a new AutoUnpacker
func NewAutoUnpacker() Unpacker {
	return &AutoUnpacker{
		unpackers: map[gitbin.BinaryDownload_Format]Unpacker{
			gitbin.BinaryDownload_BINARY: &BinaryUnpacker{},
			gitbin.BinaryDownload_TARGZ:  &TargzUnpacker{},
		},
	}
}

// Unpack implements the Unpacker interface
func (u *AutoUnpacker) Unpack(r io.Reader, download *gitbin.BinaryDownload) (io.Reader, error) {
	unpacker, ok := u.unpackers[download.GetFormat()]
	if !ok {
		return nil, fmt.Errorf("Unknown format: %s", download.GetFormat())
	}

	return unpacker.Unpack(r, download)
}
