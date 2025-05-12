// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import (
	"unsafe"
)

// These functions cannot have go:noescape annotations,
// because while ptr does not escape, new does.
// If new is marked as not escaping, the compiler will make incorrect
// escape analysis decisions about the pointer value being stored.

// atomicwb performs a write barrier before an atomic pointer write.
// The caller should guard the call with "if writeBarrier.enabled".
//
// atomicwb should be an internal detail,
// but widely used packages access it using linkname.
// Notable members of the hall of shame include:
//   - github.com/bytedance/gopkg
//   - github.com/songzhibin97/gkit
//
// Do not remove or change the type signature.
// See go.dev/issue/67401.
//
//go:linkname atomicwb runtime.atomicwb
//go:nosplit
func atomicwb(ptr *unsafe.Pointer, new unsafe.Pointer)
