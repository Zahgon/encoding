package proto

import (
	"unsafe"
)

type sizeFunc = func(unsafe.Pointer, flags) int

func sizeOfVarint(v uint64) int { _ = "STUB: not implemented"; return 0 }

func sizeOfVarintZigZag(v int64) int { _ = "STUB: not implemented"; return 0 }

func sizeOfVarlen(n int) int { _ = "STUB: not implemented"; return 0 }

func sizeOfTag(f fieldNumber, t wireType) int { _ = "STUB: not implemented"; return 0 }
