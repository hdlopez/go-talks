// START 1 OMIT
package gopherconuk

import (
	"bytes"
	"io"
	"log"
	"sync"
)

// Private, available only from within the gopherconuk package
type header struct {
}

// Public, available from other packages
func Export() {
}

// Private, available only from within the gopherconuk package
func doExport(h header) {
}

// END 1 OMIT

// START 2 OMIT

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
	Reader
	Writer
}

// END 2 OMIT

// START 3.1 OMIT

type File struct {
	sync.Mutex
	rw io.ReadWriter
}

// END 3.1 OMIT

func file() {
	// START 3.2 OMIT

	f := File{}
	f.Lock()

	// END 3.2 OMIT

	// START 4 OMIT

	var r io.Reader

	r = bytes.NewBufferString("hello")

	buf := make([]byte, 2048)
	if _, err := r.Read(buf); err != nil {
		log.Fatal(err)
	}

	// END 4 OMIT
}
