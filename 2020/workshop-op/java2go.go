package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/hdlopez/go-talks/2019/gopherconuk/examples/context"
)

// START 1 OMIT

// Private, available only from within the package
type header struct {
}

// Public, available from other packages
func Export() {
}

// Private, available only from within the package
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

// START 5 OMIT
type geometry interface {
	area() float64
	perim() float64
}

// rect implements geometry interface // HL
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// END 5 OMIT

func constructors() {
	var reader io.Reader

	// START 6 OMIT
	context.New() // means new instance of "Context" type

	http.NewRequest("GET", "/users/1", reader) // another standard way to declare a constructor
	// END 6 OMIT
}

// START 7.1 OMIT
type Event struct{}

func (e *Event) Log(msg string) {
	if e == nil {
		return
	}
	// Log the msg on the event...
}

// END 7.1 OMIT

func useEvent() {
	// START 7.2 OMIT
	var e *Event
	e.Log("this is a message")
	// END 7.2 OMIT
}

// START 8.1 OMIT
type MyFakeClass struct {
	attribute1 string
}

// END 8.1 OMIT

// START 8.2 OMIT
func (mc MyFakeClass) printMyAttribute() {
	fmt.Println(mc.attribute1)
}

// END 8.2 OMIT

// START 8.3 OMIT
func printMyAttribute(mc MyFakeClass) {
	fmt.Println(mc.attribute1)
}

// END 8.3 OMIT
