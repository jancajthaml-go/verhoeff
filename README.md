## verhoeff checksum algorithm

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/verhoeff)](https://goreportcard.com/report/jancajthaml-go/verhoeff)

### Usage ###

```
import "github.com/jancajthaml-go/verhoeff"

ok := verhoeff.Validate("00123014764700968325")

digit, error := verhoeff.Digit("x")

signed := verhoeff.Generate("1")
```

### Performance ###

```
BenchmarkVerhoeffSmall-4          100000000  15.6 ns/op  0 B/op  0 allocs/op
BenchmarkVerhoeffLarge-4          20000000   83.5 ns/op  0 B/op  0 allocs/op
BenchmarkVerhoeffSmallParallel-4  200000000  6.43 ns/op  0 B/op  0 allocs/op
BenchmarkVerhoeffLargeParallel-4  50000000   38.2 ns/op  0 B/op  0 allocs/op
```

verify your performance by running `make benchmark`

### Resources ###

* [Wikipedia - Verhoeff algorithm](https://en.wikipedia.org/wiki/Verhoeff_algorithm)
