# go-bloomfilter

generic &amp; simple bloom filter implementation in Go

## Features

- No hash function implementation
  - Implement your own.
- No 3rd party dependencies
- No complex implementation.

## Example

```go
import (
	"fmt"

	"github.com/thara/go-bloomfilter/bloomfilter"
)

const K = 4

func ExampleBloomFilter() {
	bits := make([]byte, 1024)

	f := bloomfilter.New(K, myHash)

	f.Set(bits, "smith")
	f.Set(bits, "james")
	f.Set(bits, "mary")
	f.Set(bits, "johnson")
	f.Set(bits, "john")
	f.Set(bits, "patricia")

	result := f.Test(bits, "mary")
	fmt.Println(result)
	// Output:
	// true
}
```

## License

MIT

## Author

Tomochika Hara (a.k.a [thara](https://thara.dev))
