package pbstream

import (
	"bufio"
	"encoding/binary"
	"io"

	"github.com/gogo/protobuf/proto"
)

// Decoder can read input streams and decode pb encoded messages
type Decoder struct {
	r *bufio.Reader
}

// NewDecoder creates a new decoder
func NewDecoder(r io.Reader) *Decoder { return &Decoder{r: bufio.NewReader(r)} }

// Decode decodes the next item into a message
func (d *Decoder) Decode(pb proto.Message) error {
	u, err := binary.ReadUvarint(d.r)
	if err != nil {
		return err
	}

	msg := make([]byte, int(u))
	if _, err := io.ReadFull(d.r, msg); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return err
	}
	return proto.Unmarshal(msg, pb)
}

// Encoder can write a stream of pb messages to a writer
type Encoder struct {
	w io.Writer
}

// NewEncoder creates a new decoder
func NewEncoder(w io.Writer) *Encoder { return &Encoder{w: w} }

// Decode decodes the next item into a message
func (d *Encoder) Encode(pb proto.Message) error {

	msg, err := proto.Marshal(pb)
	if err != nil {
		return err
	}

	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, uint64(len(msg)))

	_, err = d.w.Write(append(buf[:n], msg...))
	return err
}
