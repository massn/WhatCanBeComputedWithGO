![Go](https://github.com/massn/WhatCanBeComputedWithGo/workflows/Go/badge.svg?branch=main)
# WhatCanBeComputedWithGo
Go programs of ["What Can Be Computed?: A Practical Guide to the Theory of Computation(John Maccormick)"](https://press.princeton.edu/books/hardcover/9780691170664/what-can-be-computeda).　　

I am reading the Japanese version([計算できるもの、計算できないもの
――実践的アプローチによる計算理論入門](https://www.oreilly.co.jp/books/9784873119335/))

## Tests
Computings in the textbook are implemented as tests.
```
$ go test -v -race -cover ./...
```

## References
`chap2/geneticString.txt` is from https://whatcanbecomputed.com/
