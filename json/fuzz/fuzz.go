//go:build ignore
// +build ignore

// Copyright 2015 go-fuzz project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package fuzz

import (
	"github.com/segmentio/encoding/json"
)

func fixS(v any) { _ = "STUB: not implemented"; return }

func Fuzz(data []byte) int { _ = "STUB: not implemented"; return 0 }

// Note: we modified the test to verify that we behavior like the
// standard encoding/json package, whether it's right or wrong.

// both implementations report an error

// both implementations pass

type S struct {
	A int    `json:",omitempty"`
	B string `json:"B1,omitempty"`
	C float64
	D bool
	E uint8
	F []byte
	G any
	H map[string]any
	I map[string]string
	J []any
	K []string
	L S1
	M *S1
	N *int
	O **int
	P json.RawMessage
	Q Marshaller
	R int `json:"-"`
	S int `json:",string"`
}

type S1 struct {
	A int
	B string
}

type Marshaller struct {
	v string
}

func (m *Marshaller) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Marshaller) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
