package pbstream

import (
	"io"

	pio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"
)

// assert version
const _ = proto.GoGoProtoPackageIsVersion2

// maximum message size
const maxSize = 16 * 1024 * 1024 // 16M

// Decoder can read input streams and decode pb encoded messages
type Decoder struct {
	pio.ReadCloser
}

// NewDecoder creates a new decoder
func NewDecoder(r io.Reader) *Decoder { return &Decoder{ReadCloser: pio.NewDelimitedReader(r, maxSize)} }

// Decode decodes the next item into a message
func (d *Decoder) Decode(msg proto.Message) error { return d.ReadCloser.ReadMsg(msg) }

// Encoder can write a stream of pb messages to a writer
type Encoder struct {
	pio.WriteCloser
}

// NewEncoder creates a new decoder
func NewEncoder(w io.Writer) *Encoder { return &Encoder{WriteCloser: pio.NewDelimitedWriter(w)} }

// Encode encodes a message
func (d *Encoder) Encode(msg proto.Message) error { return d.WriteMsg(msg) }
