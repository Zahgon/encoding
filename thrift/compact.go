package thrift

import (
	"encoding/binary"
	"io"
)

// CompactProtocol is a Protocol implementation for the compact thrift protocol.
//
// https://github.com/apache/thrift/blob/master/doc/specs/thrift-compact-protocol.md#integer-encoding
type CompactProtocol struct{}

func (p *CompactProtocol) NewReader(r io.Reader) Reader {
	_ = "STUB: not implemented"
	return *new(Reader)
}

func (p *CompactProtocol) NewWriter(w io.Writer) Writer {
	_ = "STUB: not implemented"
	return *new(Writer)
}

func (p *CompactProtocol) Features() Features { _ = "STUB: not implemented"; return *new(Features) }

type compactReader struct {
	protocol *CompactProtocol
	binary   binaryReader
}

func (r *compactReader) Protocol() Protocol { _ = "STUB: not implemented"; return *new(Protocol) }

func (r *compactReader) Reader() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (r *compactReader) ReadBool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r *compactReader) ReadInt8() (int8, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *compactReader) ReadInt16() (int16, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *compactReader) ReadInt32() (int32, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *compactReader) ReadInt64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *compactReader) ReadFloat64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *compactReader) ReadBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *compactReader) ReadString() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *compactReader) ReadLength() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *compactReader) ReadMessage() (Message, error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

func (r *compactReader) ReadField() (Field, error) {
	_ = "STUB: not implemented"
	return *new(Field), nil
}

func (r *compactReader) ReadList() (List, error) { _ = "STUB: not implemented"; return *new(List), nil }

func (r *compactReader) ReadSet() (Set, error) { _ = "STUB: not implemented"; return *new(Set), nil }

func (r *compactReader) ReadMap() (Map, error) { _ = "STUB: not implemented"; return *new(Map), nil }

// empty map

func (r *compactReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *compactReader) readUvarint(typ string, max uint64) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *compactReader) readVarint(typ string, min, max int64) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type compactWriter struct {
	protocol *CompactProtocol
	binary   binaryWriter
	varint   [binary.MaxVarintLen64]byte
}

func (w *compactWriter) Protocol() Protocol { _ = "STUB: not implemented"; return *new(Protocol) }

func (w *compactWriter) Writer() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (w *compactWriter) WriteBool(v bool) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteInt8(v int8) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteInt16(v int16) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteInt32(v int32) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteInt64(v int64) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteFloat64(v float64) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteBytes(v []byte) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteString(v string) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteLength(n int) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteMessage(m Message) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteField(f Field) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteList(l List) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteSet(s Set) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) WriteMap(m Map) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) writeUvarint(v uint64) error { _ = "STUB: not implemented"; return nil }

func (w *compactWriter) writeVarint(v int64) error { _ = "STUB: not implemented"; return nil }
