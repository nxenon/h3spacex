package wire

import (
	"bytes"
	"fmt"

	"github.com/nxenon/h3spacex/internal/protocol"
	"github.com/nxenon/h3spacex/quicvarint"
)

// A StreamsBlockedFrame is a STREAMS_BLOCKED frame
type StreamsBlockedFrame struct {
	Type        protocol.StreamType
	StreamLimit protocol.StreamNum
}

func parseStreamsBlockedFrame(r *bytes.Reader, typ uint64, _ protocol.Version) (*StreamsBlockedFrame, error) {
	f := &StreamsBlockedFrame{}
	switch typ {
	case bidiStreamBlockedFrameType:
		f.Type = protocol.StreamTypeBidi
	case uniStreamBlockedFrameType:
		f.Type = protocol.StreamTypeUni
	}
	streamLimit, err := quicvarint.Read(r)
	if err != nil {
		return nil, err
	}
	f.StreamLimit = protocol.StreamNum(streamLimit)
	if f.StreamLimit > protocol.MaxStreamCount {
		return nil, fmt.Errorf("%d exceeds the maximum stream count", f.StreamLimit)
	}
	return f, nil
}

func (f *StreamsBlockedFrame) Append(b []byte, _ protocol.Version) ([]byte, error) {
	switch f.Type {
	case protocol.StreamTypeBidi:
		b = append(b, bidiStreamBlockedFrameType)
	case protocol.StreamTypeUni:
		b = append(b, uniStreamBlockedFrameType)
	}
	b = quicvarint.Append(b, uint64(f.StreamLimit))
	return b, nil
}

// Length of a written frame
func (f *StreamsBlockedFrame) Length(_ protocol.Version) protocol.ByteCount {
	return 1 + quicvarint.Len(uint64(f.StreamLimit))
}
