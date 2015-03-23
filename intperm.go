// Package intperm implements a simple permutation for 64-bit ints.
// This file also includes a simple XORShift-based PRNG for expanding the seed.
// Example code from http://www.jstatsoft.org/v08/i14/paper (public domain).
package intperm

// 64-bit full-one bitmask
const ones = 0xffffffffffffffff

// A Permutation is a reversible, pseudo-random mapping of 64-bit unsigned ints.
type Permutation []uint64

// New creates a new Permutation.
// The argument can be any random number.
func New(seed uint64) Permutation {
	masks := make([]uint64, 64*2)
	params := triplets[seed%uint64(len(triplets))]
	for i := range masks {
		seed = xorshift(seed, params[0], params[1], params[2])
		masks[i] = seed & ((1 << uint64(i>>1)) ^ uint64(ones))
	}
	return Permutation(masks)
}

// MapTo a number to another random one.
func (p Permutation) MapTo(x uint64) uint64 {
	return p.doMap(x, 0, 64, 1)
}

// MapFrom is the reverse of MapTo.
// In other words, p.MapFrom(p.MapTo(x)) == x.
func (p Permutation) MapFrom(x uint64) uint64 {
	return p.doMap(x, 63, -1, -1)
}

// Used by both MapTo and MapFrom.
// Set `to` to true for MapTo, and `false` for MapFrom behaviour.
func (p Permutation) doMap(x uint64, from, to, step int) uint64 {
	for i := from; i != to; i += step {
		u := uint64(i)
		bit := uint64(1 << u)
		if (bit&x)>>u == 0 {
			x ^= ^(p[(u<<1)+(bit&x)>>u] | (bit ^ bit&x))
		} else {
			x ^= ^(p[(u<<1)+(bit&x)>>u] | (bit & x))
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
