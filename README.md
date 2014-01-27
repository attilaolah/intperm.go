# Permutation

[![Bitdeli](https://d2weczhvl823v0.cloudfront.net/attilaolah/permutation.go/trend.png)](https://bitdeli.com/free "Bitdeli Badge")
[![Build Status](https://travis-ci.org/attilaolah/permutation.go.png?branch=master)](https://travis-ci.org/attilaolah/permutation.go)

This package implements a simple, configurable permutation on the set of 64-bit
integers. The implementation uses `uint64`, but it is also usable for `int64`
types with a simple conversion.

The permutation is based on a bitmask that maps each bit of the input to a bit
of the output. The bitmask is expanded from a random seed using a [PRNG][1], as
described by *George Marsaglia* in his paper called [Xorshift RNGs][2]. The
permutations are thus believed to be unpredictable, provided provided that the
seed is kept secret.

[1]: //en.wikipedia.org/wiki/Pseudorandom_number_generator
[2]: http://www.jstatsoft.org/v08/i14/paper

## Usage

Create a new permutation with `New()`, passing in four parameters.

The first parameter is the seed, which can be any random value.
The next three parameters are used by the XORShift to expand the seed.
Valid values are listed in [George Marsaglia's paper][2], on *page 3*.

```go
p := permutation.New(42, 13, 7, 17)
a := p.Map(42) // 3333656047352411619
b := p.Unmap(3333656047352411619) // 42
```

## Use cases

Use cases may vary, but an example that I find useful is generating
[hard][4]-to-guess, random-looking tokens based on IDs stored in a database.
The IDs can be used together with the seed to decode the original ID, but their
[cardinality][5] is the same as that of the IDs themselves. When used smartly,
this can save you from having to index those tokens in the database.

[4]: //en.wikipedia.org/wiki/NP-hard
[5]: //en.wikipedia.org/wiki/Cardinality

## See also

There are [some][7] [other][8] Go permutation libraries out there, but their
focus is on implementing a more generic mechanism that can be used with
arbitrary structures implementing an interface. This library focuses on very
fast, simple, pseudo-random integer permutations only.

There is currently also a [Python implementation][6] of this library.

[6]: //github.com/attilaolah/permutation.py
[7]: //github.com/alexaandru/permutations
[8]: //github.com/nightlyone/permutation

## License

[Public domain][3].

[3]: //github.com/attilaolah/permutation.go/blob/master/LICENSE
