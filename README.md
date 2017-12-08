## Performant and straight-forward implementation of verhoeff checksum algorithm

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
BenchmarkVerhoeffSmall-4            200000000     9.54 ns/op
BenchmarkVerhoeffLarge-4            20000000      61.7 ns/op
BenchmarkVerhoeffSmallParallel-4    300000000     4.60 ns/op
BenchmarkVerhoeffLargeParallel-4    50000000      31.9 ns/op
```

test on your own by running `make benchmark`

### Resources ###

* [Wikipedia - Verhoeff algorithm](https://en.wikipedia.org/wiki/Verhoeff_algorithm)
