// Package runtime_reflect exposes internal APIs of the Go runtime.
//
// This package is internal so it doesn't become part of the exported APIs that
// users of this package can take dependencies on. There is a risk that these
// APIs will be implicitly changed by Go, in which case packages that depend on
// it will break. We use these APIs to have access to optimziations that aren't
// possible today via the reflect package. Ideally, the reflect package evolves
// to expose APIs that are efficient enough that we can drop the need for this
// package, but until then we will be maintaining bridges to these Go runtime
// functions and types.
package runtime_reflect

import "unsafe"

func Assign(typ, dst, src unsafe.Pointer) { _ = "STUB: not implemented"; return }

func MapAssign(t, m, k unsafe.Pointer) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func MakeMap(t unsafe.Pointer, cap int) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

type MapIter struct{ hiter }

func (it *MapIter) Init(t unsafe.Pointer, m unsafe.Pointer) { _ = "STUB: not implemented"; return }

func (it *MapIter) Done() { _ = "STUB: not implemented"; return }

func (it *MapIter) Next() { _ = "STUB: not implemented"; return }

func (it *MapIter) HasNext() bool { _ = "STUB: not implemented"; return false }

func (it *MapIter) Key() unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

func (it *MapIter) Value() unsafe.Pointer {
	_ = "STUB: not implemented"

	// copied from src/runtime/map.go, all pointer types replaced with
	// unsafe.Pointer.
	//
	// Alternatively we could get away with a heap allocation and only
	// defining key and val if we were using reflect.mapiterinit instead,
	// which returns a heap-allocated *hiter.
	return *new(unsafe.Pointer)
}

type hiter struct {
	key         unsafe.Pointer // nil when iteration is done
	value       unsafe.Pointer
	t           unsafe.Pointer
	h           unsafe.Pointer
	buckets     unsafe.Pointer // bucket ptr at hash_iter initialization time
	bptr        unsafe.Pointer // current bucket
	overflow    unsafe.Pointer // keeps overflow buckets of hmap.buckets alive
	oldoverflow unsafe.Pointer // keeps overflow buckets of hmap.oldbuckets alive
	startBucket uintptr        // bucket iteration started at
	offset      uint8          // intra-bucket offset to start from during iteration (should be big enough to hold bucketCnt-1)
	wrapped     bool           // already wrapped around from end of bucket array to beginning
	B           uint8
	i           uint8
	bucket      uintptr
	checkBucket uintptr
}

//go:noescape
//go:linkname makemap reflect.makemap
func makemap(t unsafe.Pointer, cap int) unsafe.Pointer

// m escapes into the return value, but the caller of mapiterinit
// doesn't let the return value escape.
//
//go:noescape
//go:linkname mapiterinit runtime.mapiterinit
func mapiterinit(t unsafe.Pointer, m unsafe.Pointer, it *hiter)

//go:noescape
//go:linkname mapiternext runtime.mapiternext
func mapiternext(it *hiter)

//go:noescape
//go:linkname mapassign runtime.mapassign
func mapassign(t, m, k unsafe.Pointer) unsafe.Pointer

//go:nosplit
//go:noescape
//go:linkname typedmemmove runtime.typedmemmove
func typedmemmove(typ, dst, src unsafe.Pointer)
