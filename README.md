# primes
Primes for projects which extremely do prime check.
## Usage
```go
package main

import (
  "fmt"
  "github.com/aydos/primes"
)

func init() {
  // calculate all primes upto some max number
  // call it only once
  primes.InitPrimes(300)
}

func main() {
  // do the prime check
  // CAUTION: do not check i > 300
  for i := 0; i <= 300; i++ {
    if primes.IsPrime(i) {
      fmt.Println(i, "is prime")
    }
  }
  // you can use the the primes.Primes slice
  fmt.Println("There are", len(primes.Primes), "primes upto 300")
}
```
## CAUTION
```primes.IsPrime()``` do not return error. It returns only true or false. If you check bigger number than the initializing one, it will say "its not prime". So, choose it wisely and do not check biggers.
## Benchmarks
On my virtual machine, primes.IsPrime is 10-20 times better then big.ProbablyPrime
```
[root@test primes]# go test -bench=.
BenchmarkIsPrime                10000000               188 ns/op
BenchmarkProbablyPrime3          1000000              2033 ns/op
BenchmarkProbablyPrime10          500000              2780 ns/op
BenchmarkProbablyPrime20          500000              3865 ns/op
PASS
ok      primes  9.123s
```
