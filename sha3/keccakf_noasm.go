//  +build !amd64,!arm,!arm64 appengine gccgo

package sha3

// Use generic implementation
func keccakF1600(a *[25]uint64) {
	keccakF1600Generic(a)
}
