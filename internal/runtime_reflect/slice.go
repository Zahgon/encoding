package runtime_reflect

import "unsafe"

type Slice struct {
	data unsafe.Pointer
	len  int
	cap  int
}

func (s *Slice) Cap() int { _ = "STUB: not implemented"; return 0 }

func (s *Slice) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *Slice) SetLen(n int) { _ = "STUB: not implemented"; return }

func (s *Slice) Index(i int, elemSize uintptr) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func MakeSlice(elemType unsafe.Pointer, len, cap int) Slice {
	_ = "STUB: not implemented"
	return *new(Slice)
}

func CopySlice(elemType unsafe.Pointer, dst, src Slice) int { _ = "STUB: not implemented"; return 0 }

//go:linkname newarray runtime.newarray
func newarray(t unsafe.Pointer, n int) unsafe.Pointer

//go:linkname typedslicecopy runtime.typedslicecopy
//go:noescape
func typedslicecopy(t unsafe.Pointer, dst, src Slice) int
