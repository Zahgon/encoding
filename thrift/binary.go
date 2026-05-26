package thrift

import (
	"io"
)

// BinaryProtocol is a Protocol implementation for the binary thrift protocol.
//
// https://github.com/apache/thrift/blob/master/doc/specs/thrift-binary-protocol.md
type BinaryProtocol struct {
	NonStrict bool
}

func (p *BinaryProtocol) NewReader(r io.Reader) Reader {
	_ = "STUB: not implemented"
	return *new(Reader)
}

func (p *BinaryProtocol) NewWriter(w io.Writer) Writer {
	_ = "STUB: not implemented"
	return *new(Writer)
}

func (p *BinaryProtocol) Features() Features { _ = "STUB: not implemented"; return *new(Features) }

type binaryReader struct {
	p *BinaryProtocol
	r io.Reader
	b [8]byte
}

func (r *binaryReader) Protocol() Protocol { _ = "STUB: not implemented"; return *new(Protocol) }

func (r *binaryReader) Reader() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (r *binaryReader) ReadBool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r *binaryReader) ReadInt8() (int8, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *binaryReader) ReadInt16() (int16, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *binaryReader) ReadInt32() (int32, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *binaryReader) ReadInt64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *binaryReader) ReadFloat64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *binaryReader) ReadBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *binaryReader) ReadString() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *binaryReader) ReadLength() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *binaryReader) ReadMessage() (Message, error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

// non-strict

func (r *binaryReader) ReadField() (Field, error) {
	_ = "STUB: not implemented"
	return *new(Field), nil
}

func (r *binaryReader) ReadList() (List, error) { _ = "STUB: not implemented"; return *new(List), nil }

func (r *binaryReader) ReadSet() (Set, error) { _ = "STUB: not implemented"; return *new(Set), nil }

func (r *binaryReader) ReadMap() (Map, error) { _ = "STUB: not implemented"; return *new(Map), nil }

func (r *binaryReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *binaryReader) read(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type binaryWriter struct {
	p *BinaryProtocol
	b [8]byte
	w io.Writer
}

func (w *binaryWriter) Protocol() Protocol { _ = "STUB: not implemented"; return *new(Protocol) }

func (w *binaryWriter) Writer() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (w *binaryWriter) WriteBool(v bool) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteInt8(v int8) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteInt16(v int16) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteInt32(v int32) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteInt64(v int64) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteFloat64(v float64) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteBytes(v []byte) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteString(v string) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteLength(n int) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteMessage(m Message) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteField(f Field) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteList(l List) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteSet(s Set) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) WriteMap(m Map) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) write(b []byte) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) writeString(s string) error { _ = "STUB: not implemented"; return nil }

func (w *binaryWriter) writeByte(b byte) error {
	_ = "STUB: not implemented"
	// The special cases are intended to reduce the runtime overheadof testing
	// for the io.ByteWriter interface for common types. Type assertions on a
	// concrete type is just a pointer comparison, instead of requiring a
	// complex lookup in the type metadata.
	return nil
}
