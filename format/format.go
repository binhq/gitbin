package format

import (
	"io"

	"fmt"

	githubin "github.com/binhq/githubin/apis/githubin/v1alpha1"
)

// Unpacker unpacks the binary from a download format
type Unpacker func(r io.Reader, download *githubin.BinaryDownload) (io.Reader, error)

// FindUnpacker returns an unpacker or an error
func FindUnpacker(format githubin.BinaryDownload_Format) (Unpacker, error) {
	switch format {
	case githubin.BinaryDownload_BINARY:
		return BinaryUnpacker, nil
	case githubin.BinaryDownload_TARGZ:
		return TargzUnpacker, nil
	}

	return nil, fmt.Errorf("Error not found for format: %d", format)
}
