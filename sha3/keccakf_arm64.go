// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build arm64,!appengine,!gccgo

// This generates the plan9 assembly code from the native ARM assembly code - it assumes that gcc-aarch64-linux-gnu utils is installed on the $PATH
// as well as the tool asm2go : https://github.com/anonymouse64/asm2go
//go:generate asm2go -as aarch64-linux-gnu-as -file src/keccak_arm64_src.s -gofile keccak/keccak_arm64.go -out keccak/keccak_arm64.s

package sha3

//go:noescape
// This function is implemented in keccakf_arm64.s
func KeccakF1600(state *[25]uint64, constants *[24]uint64)

// Uses the assembly implementation
func keccakF1600(a *[25]uint64) {
	KeccakF1600(a, &rc)
}
