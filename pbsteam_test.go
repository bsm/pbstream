package pbstream

import (
	"bytes"
	"io"
	"testing"

	pb "github.com/bsm/pbstream/testdata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encoder", func() {
	var subject *Encoder
	var w *bytes.Buffer

	BeforeEach(func() {
		w = new(bytes.Buffer)
		subject = NewEncoder(w)
	})

	It("should enbcode streams", func() {
		Expect(subject.Encode(&pb.Message{S: "a", N: 1})).NotTo(HaveOccurred())
		Expect(w.Len()).To(Equal(6))
		Expect(subject.Encode(&pb.Message{S: "b", N: 2})).NotTo(HaveOccurred())
		Expect(w.Len()).To(Equal(12))
		Expect(subject.Encode(&pb.Message{S: "boooooooooooooo"})).NotTo(HaveOccurred())
		Expect(w.Len()).To(Equal(30))
	})

})

var _ = Describe("Decoder", func() {
	var subject *Decoder
	var w *bytes.Buffer
	var (
		x1 = &pb.Message{S: "a", N: 1}
		x2 = &pb.Message{S: "b", N: 2}
		x3 = &pb.Message{S: "c"}
	)

	BeforeEach(func() {
		w = new(bytes.Buffer)
		enc := NewEncoder(w)
		Expect(enc.Encode(x1)).NotTo(HaveOccurred())
		Expect(enc.Encode(x2)).NotTo(HaveOccurred())
		Expect(enc.Encode(x3)).NotTo(HaveOccurred())

		subject = NewDecoder(w)
	})

	It("should decode streams", func() {
		var (
			m1 = new(pb.Message)
			m2 = new(pb.Message)
			m3 = new(pb.Message)
			m4 = new(pb.Message)
		)
		Expect(subject.Decode(m1)).NotTo(HaveOccurred())
		Expect(subject.Decode(m2)).NotTo(HaveOccurred())
		Expect(subject.Decode(m3)).NotTo(HaveOccurred())
		Expect(subject.Decode(m4)).To(Equal(io.EOF))

		Expect(m1).To(Equal(x1))
		Expect(m2).To(Equal(x2))
		Expect(m3).To(Equal(x3))
	})

})

// --------------------------------------------------------------------

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "pbstream")
}
