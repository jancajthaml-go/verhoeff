## zero-alloc verhoeff checksum algorithm

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/verhoeff)](https://goreportcard.com/report/jancajthaml-go/verhoeff)

### Usage ###

```
import "github.com/jancajthaml-go/verhoeff"

ok := verhoeff.Validate("00123014764700968325")

digit, error := verhoeff.Digit("x")

signed := verhoeff.Generate("1")
```

### Resources ###

* [Wikipedia - Verhoeff algorithm](https://en.wikipedia.org/wiki/Verhoeff_algorithm)
