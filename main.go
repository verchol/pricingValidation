package main

import (
	"bytes"
	"flag"
	"fmt"
)

func main() {
	bytesLimit := flag.Int("bytesLimit", 100000, "put number of flags")
	flag.Parse()
	var b bytes.Buffer
	longString := "very long string !!!!!"
	for i := 0; i < *bytesLimit; i++ {
		b.Write([]byte(longString))
		fmt.Printf("total: %d\n",
			b.Len(),
		)
	}
}
