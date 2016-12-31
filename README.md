# primes
Prime check in 4 nanosecond
##Usage
```go
import (
  "fmt"
  "github.com/aydos/primes"
)

func main() {
  // calculate all primes up to some max number
  // you can put this line in init()
  primes.InitPrimes(300)
  
  // do the prime check
  // CAUTION: do not check i > 300
  for i := 0; i <= 300; i++ {
    if primes.Primes(i) {
      fmt.Println(i, "is prime")
    }
  }
}
```
## CAUTION
Primes() do not return error. It returns only true or false. If you check bigger number than the initializing one, it will say "its not prime". So, choose it wisely and do not check biggers than it.
