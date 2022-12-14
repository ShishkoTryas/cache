# Go in-memory-cache

### Example

```go
package main

import (
	"fmt"

	"github.com/ShishkoTryas/cache"
)

func main() {
	cache := cache.New()

	if err := cache.Set("", 42); err != nil {
		fmt.Println(err)
	}

	if err := cache.Set("userId", 42); err != nil {
		fmt.Println(err)
	}

	if value, err := cache.Get("userId"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	if err := cache.Delete("xxx-userId-xxx"); err != nil {
		fmt.Println(err)
	}

	if err := cache.Delete("userId"); err != nil {
		fmt.Println(err)
	}

	if userId, err := cache.Get("userId"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userId)
	}

}
```
