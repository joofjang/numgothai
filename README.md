# numgothai
Go module for generate Thai text and Baht currency text from number.

# Usage
```go
package main

import (
	"fmt"

	"github.com/joofjang/numgothai"
)

func main() {
	fmt.Println(numgothai.IntText(1500)) //หนึ่งพันห้าร้อย
	fmt.Println(numgothai.IntBaht(150050)) //หนึ่งพันห้าร้อยบาทห้าสิบสตางค์
}
```

![Go Test and Build](https://github.com/joofjang/numgothai/workflows/Go%20Test/badge.svg)