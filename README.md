# Permutation

This package implements a simple, configurable permutation on the set of 64-bit
integers. The implementation uses `uint64`, but it is also usable for `int64
types with a simple conversion.

The permutation uses a simple bit mask to map each bit of the input to a bit of
the output. The bit mask is expanded from a seed using a simple [PRNG][1]
described by *George Marsaglia* [in this paper][2].

[1]: //en.wikipedia.org/wiki/Pseudorandom_number_generator
[2]: http://www.jstatsoft.org/v08/i14/paper

## Usage

```go
p := permutation.New(42, 13, 7, 17)
a := p.Map(42) // 3333656047352411619
b := p.Unmap(3333656047352411619) // 42
```

## License

[Public domain][3].

[3]: //github.com/attilaolah/permutation.go/blob/master/LICENSE
