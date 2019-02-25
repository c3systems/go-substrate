package rpc

import (
	"bufio"

	inet "github.com/libp2p/go-libp2p-net"
	"github.com/ugorji/go/codec"
)

// streamWrap wraps a libp2p stream. We encode/decode whenever we
// write/read from a stream, so we can just carry the encoders
// and bufios with us
type streamWrap struct {
	stream inet.Stream
	enc    *codec.Encoder
	dec    *codec.Decoder
	w      *bufio.Writer
	r      *bufio.Reader
}

// wrapStream takes a stream and complements it with r/w bufios and
// decoder/encoder. In order to write to the stream we can use
// wrap.w.Write(). To encode something into it we can wrap.enc.Encode().
// Finally, we should wrap.w.Flush() to actually send the data. Similar
// for receiving.
func wrapStream(s inet.Stream) *streamWrap {
	reader := bufio.NewReader(s)
	writer := bufio.NewWriter(s)
	h := &codec.MsgpackHandle{}
	dec := codec.NewDecoder(reader, h)
	enc := codec.NewEncoder(writer, h)
	return &streamWrap{
		stream: s,
		r:      reader,
		w:      writer,
		enc:    enc,
		dec:    dec,
	}

}
