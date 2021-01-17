![Go](https://github.com/massn/WhatCanBeComputedWithGo/workflows/Go/badge.svg?branch=main)
# WhatCanBeComputedWithGo
Go programs of ["What Can Be Computed?: A Practical Guide to the Theory of Computation(John Maccormick)"](https://press.princeton.edu/books/hardcover/9780691170664/what-can-be-computeda).　　

I am reading the Japanese translated version([計算できるもの、計算できないもの
――実践的アプローチによる計算理論入門](https://www.oreilly.co.jp/books/9784873119335/))

## Tests
Computings in the textbook are implemented as tests.
```
$ go test -v -race -cover ./...
```

## Eratta of the Japanese translated version

|場所　 　　 |誤            | 正              |
|----      | ---          | ---            |
|p.37 2行目 | 24ページを参照 | 27ページを参照   |
|p.37 4行目 | 1000字以上なら | 1000字より大きい |
|p.143 7行目| ![\text{DecidesOneOf}_{S'}](https://render.githubusercontent.com/render/math?math=%5Ctextstyle+%5Ctext%7BDecidesOneOf%7D_%7BS%27%7D) |![\large{\text{D}}\small{\text{ECIDES}}\large{\text{O}}\small{\text{NE}}\large{\text{O}}\small{\text{F}}_{S}](https://render.githubusercontent.com/render/math?math=%5Ctextstyle+%5Clarge%7B%5Ctext%7BD%7D%7D%5Csmall%7B%5Ctext%7BECIDES%7D%7D%5Clarge%7B%5Ctext%7BO%7D%7D%5Csmall%7B%5Ctext%7BNE%7D%7D%5Clarge%7B%5Ctext%7BO%7D%7D%5Csmall%7B%5Ctext%7BF%7D%7D_%7BS%7D)| 
|p.143 7行目| ![\large{\text{R}}\small{\text{ECOGNIZES}}\large{\text{O}}\small{\text{NE}}\large{\text{O}}\small{\text{F}}_{S'}](https://render.githubusercontent.com/render/math?math=%5Ctextstyle+%5Clarge%7B%5Ctext%7BR%7D%7D%5Csmall%7B%5Ctext%7BECOGNIZES%7D%7D%5Clarge%7B%5Ctext%7BO%7D%7D%5Csmall%7B%5Ctext%7BNE%7D%7D%5Clarge%7B%5Ctext%7BO%7D%7D%5Csmall%7B%5Ctext%7BF%7D%7D_%7BS%27%7D)| ![\large{\text{R}}\small{\text{ECOGNIZES}}\large{\text{O}}\small{\text{NE}}\large{\text{O}}\small{\text{F}}_{S}](https://render.githubusercontent.com/render/math?math=%5Ctextstyle+%5Clarge%7B%5Ctext%7BR%7D%7D%5Csmall%7B%5Ctext%7BECOGNIZES%7D%7D%5Clarge%7B%5Ctext%7BO%7D%7D%5Csmall%7B%5Ctext%7BNE%7D%7D%5Clarge%7B%5Ctext%7BO%7D%7D%5Csmall%7B%5Ctext%7BF%7D%7D_%7BS%7D)|


## References
`chap2/geneticString.txt` is from https://whatcanbecomputed.com/
