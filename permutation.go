// Package permutation implements a simple permutation for 64-bit ints.
// This file also includes a simple XORShift-based PRNG for expanding the seed.
// Example code from http://www.jstatsoft.org/v08/i14/paper (public domain).
package permutation

// 64-bit full-one bitmask
const ones = 0xffffffffffffffff

type permutation []uint64

// New creates a new permutation.
// The first argument, `seed`, can be any random number.
// The other three should be one of the 275 available triplets from the paper (page 3).
// For unpredictable permutations, choose different values from http://www.jstatsoft.org/v08/i14/paper.
func New(seed, a, b, c uint64) permutation {
	masks := make([]uint64, 64*2)
	for i := range masks {
		seed = xorshift(seed, a, b, c)
		masks[i] = seed & ((1 << uint64(i>>1)) ^ uint64(ones))
	}
	return permutation(masks)
}

// Map a number to another random one.
func (b permutation) Map(x uint64) uint64 {
	for i := 0; i < 64; i++ {
		u := uint64(i)
		bit := uint64(1 << (u))
		if (bit&x)>>u == 0 {
			x ^= ^(b[(u<<1)+(bit&x)>>u] | (bit ^ bit&x))
		} else {
			x ^= ^(b[(u<<1)+(bit&x)>>u] | (bit & x))
		}
	}
	return x
}

// Unmap is the reverse of Map.
// In other words, b.Unmap(b.Map(x)) == x.
func (b permutation) Unmap(x uint64) uint64 {
	for i := 63; i >= 0; i-- {
		u := uint64(i)
		bit := uint64(1 << (u))
		if (bit&x)>>u == 0 {
			x ^= ^(b[(u<<1)+(bit&x)>>u] | (bit ^ bit&x))
		} else {
			x ^= ^(b[(u<<1)+(bit&x)>>u] | (bit & x))
		}
	}
	return x
}

// PRNG for expandirg the seed.
// This is required because we need 128 (64*2) random values,
// but to initialise the permutation, it is sufficient to give one seed.
func xorshift(seed, a, b, c uint64) uint64 {
	seed ^= seed << a
	seed ^= seed >> b
	seed ^= seed << c
	return seed
}
