package wire

import (
	"bytes"

	"github.com/nxenon/h3spacex/internal/protocol"
	"github.com/nxenon/h3spacex/internal/qerr"
	"github.com/nxenon/h3spacex/quicvarint"
)

// A ResetStreamFrame is a RESET_STREAM frame in QUIC
type ResetStreamFrame struct {
	StreamID  protocol.StreamID
	ErrorCode qerr.StreamErrorCode
	FinalSize protocol.ByteCount
}

func parseResetStreamFrame(r *bytes.Reader, _ protocol.Version) (*ResetStreamFrame, error) {
	var streamID protocol.StreamID
	var byteOffset protocol.ByteCount
	sid, err := quicvarint.Read(r)
	if err != nil {
		return nil, err
	}
	streamID = protocol.StreamID(sid)
	errorCode, err := quicvarint.Read(r)
	if err != nil {
		return nil, err
	}
	bo, err := quicvarint.Read(r)
	if err != nil {
		return nil, err
	}
	byteOffset = protocol.ByteCount(bo)

	return &ResetStreamFrame{
		StreamID:  streamID,
		ErrorCode: qerr.StreamErrorCode(errorCode),
		FinalSize: byteOffset,
	}, nil
}

func (f *ResetStreamFrame) Append(b []byte, _ protocol.Version) ([]byte, error) {
	b = append(b, resetStreamFrameType)
	b = quicvarint.Append(b, uint64(f.StreamID))
	b = quicvarint.Append(b, uint64(f.ErrorCode))
	b = quicvarint.Append(b, uint64(f.FinalSize))
	return b, nil
}

// Length of a written frame
func (f *ResetStreamFrame) Length(version protocol.Version) protocol.ByteCount {
	return 1 + quicvarint.Len(uint64(f.StreamID)) + quicvarint.Len(uint64(f.ErrorCode)) + quicvarint.Len(uint64(f.FinalSize))
}
