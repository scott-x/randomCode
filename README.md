# randomCode
generate random code for validate useage

### useage

```go
package main
import (
    "fmt"
    "github.com/scott-x/randomCode"
)

func main() {
	s, _:= randomCode.Random(8)
	fmt.Println(s)
}
```