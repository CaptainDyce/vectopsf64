# vectopsf64
Vector ([]float64) library - immediate operations on float64 arrays

### Quick Example

```go
package main

import (
	"fmt"
	v "github.com/CaptainDyce/vectopsf64"
	"math"
)

func main() {
	// defining an identity vector 0..100, dividing each entry by 2*pi, apply math.Sin and reverse the vector
	v1 := v.OnIdent(100).Divl(2 * math.Pi).ApplyOp(math.Sin).Rev()
	fmt.Println(v1)
	// print the sum of the section [50..70[
	fmt.Println(v1.Slice(50, 70).Sum())

	a := make([]float64, 100)
	v2 := v.On(a).Ident().Divl(2 * math.Pi).ApplyOp(math.Sin).Rev()
	fmt.Println(v2)
	// print the sum of the section [50..70[
	fmt.Println(v2.Slice(50, 70).Sum())
}
```
