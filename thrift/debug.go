package thrift

import (
	"io"
	"log"
)

func NewDebugReader(r Reader, l *log.Logger) Reader { _ = "STUB: not implemented"; return *new(Reader) }

func NewDebugWriter(w Writer, l *log.Logger) Writer { _ = "STUB: not implemented"; return *new(Writer) }

type debugReader struct {
	r Reader
	l *log.Logger
}

func (d *debugReader) log(method string, res any, err error) { _ = "STUB: not implemented"; return }

func (d *debugReader) Protocol() Protocol { _ = "STUB: not implemented"; return *new(Protocol) }

func (d *debugReader) Reader() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (d *debugReader) ReadBool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (d *debugReader) ReadInt8() (int8, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *debugReader) ReadInt16() (int16, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *debugReader) ReadInt32() (int32, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *debugReader) ReadInt64() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *debugReader) ReadFloat64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *debugReader) ReadBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *debugReader) ReadString() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *debugReader) ReadLength() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *debugReader) ReadMessage() (Message, error) {
	_ = "STUB: not implemented"
	return *new(Message), nil
}

func (d *debugReader) ReadField() (Field, error) {
	_ = "STUB: not implemented"
	return *new(Field), nil
}

func (d *debugReader) ReadList() (List, error) { _ = "STUB: not implemented"; return *new(List), nil }

func (d *debugReader) ReadSet() (Set, error) { _ = "STUB: not implemented"; return *new(Set), nil }

func (d *debugReader) ReadMap() (Map, error) { _ = "STUB: not implemented"; return *new(Map), nil }

type debugWriter struct {
	w Writer
	l *log.Logger
}

func (d *debugWriter) log(method string, arg any, err error) { _ = "STUB: not implemented"; return }

func (d *debugWriter) Protocol() Protocol { _ = "STUB: not implemented"; return *new(Protocol) }

func (d *debugWriter) Writer() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (d *debugWriter) WriteBool(v bool) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteInt8(v int8) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteInt16(v int16) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteInt32(v int32) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteInt64(v int64) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteFloat64(v float64) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteBytes(v []byte) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteString(v string) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteLength(n int) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteMessage(m Message) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteField(f Field) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteList(l List) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteSet(s Set) error { _ = "STUB: not implemented"; return nil }

func (d *debugWriter) WriteMap(m Map) error { _ = "STUB: not implemented"; return nil }
