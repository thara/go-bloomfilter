package bloomfilter_test

import (
	"fmt"

	"github.com/thara/go-bloomfilter/bloomfilter"
	"github.com/thara/go-bloomfilter/bloomfilter/testutil"
)

const K = 4

func ExampleBloomFilter() {
	bits := make([]byte, 1024)

	f := bloomfilter.New(K, testutil.MyHash)

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
