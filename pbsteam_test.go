package pbstream

import (
	"bytes"
	"io"
	"testing"

	pb "github.com/bsm/pbstream/testdata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("pbstream", func() {
	var buf *bytes.Buffer
	var (
		x1 = &pb.Message{S: "a", N: 1}
		x2 = &pb.Message{S: "b", N: 2}
		x3 = &pb.Message{S: "c"}
		x4 = &pb.Message{S: "d", N: 4, V: &pb.Message_B{B: true}}
		x5 = &pb.Message{S: "boooooooooooooo"}
	)

	BeforeEach(func() {
		buf = new(bytes.Buffer)
	})

	Describe("Encoder", func() {
		var subject *Encoder

		BeforeEach(func() {
			subject = NewEncoder(buf)
		})

		It("should enbcode streams", func() {
			Expect(subject.Encode(x1)).NotTo(HaveOccurred())
			Expect(buf.Len()).To(Equal(6))
			Expect(subject.Encode(x2)).NotTo(HaveOccurred())
			Expect(buf.Len()).To(Equal(12))
			Expect(subject.Encode(x4)).NotTo(HaveOccurred())
			Expect(buf.Len()).To(Equal(20))
			Expect(subject.Encode(x5)).NotTo(HaveOccurred())
			Expect(buf.Len()).To(Equal(38))
		})

	})

	Describe("Decoder", func() {
		var subject *Decoder

		BeforeEach(func() {
			buf = new(bytes.Buffer)
			enc := NewEncoder(buf)
			Expect(enc.Encode(x1)).NotTo(HaveOccurred())
			Expect(enc.Encode(x2)).NotTo(HaveOccurred())
			Expect(enc.Encode(x3)).NotTo(HaveOccurred())
			Expect(enc.Encode(x4)).NotTo(HaveOccurred())

			subject = NewDecoder(buf)
		})

		It("should decode streams", func() {
			var (
				m1 = new(pb.Message)
				m2 = new(pb.Message)
				m3 = new(pb.Message)
				m4 = new(pb.Message)
				m5 = new(pb.Message)
			)

			Expect(subject.Decode(m1)).NotTo(HaveOccurred())
			Expect(subject.Decode(m2)).NotTo(HaveOccurred())
			Expect(subject.Decode(m3)).NotTo(HaveOccurred())
			Expect(subject.Decode(m4)).NotTo(HaveOccurred())
			Expect(subject.Decode(m5)).To(Equal(io.EOF))

			Expect(m1).To(Equal(x1))
			Expect(m2).To(Equal(x2))
			Expect(m3).To(Equal(x3))
			Expect(m4).To(Equal(x4))
		})

	})
})

// --------------------------------------------------------------------

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "pbstream")
}
