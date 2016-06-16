package main

import (
	"fmt"
	"os"
	// "bufio"
	"log"
	// "io"
	// "time"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, err := os.Open("./Video/cat.mp4")
	check(err)

	data := make([]byte, 100)
	count, err1 := file.Read(data)
	if err1 != nil {
		log.Fatal(err1)
	}
	// var z int64
	// z = int64(count)
	// o2, err := file.Seek(z, 1)
	fmt.Println(count)
	fmt.Printf("\n\n\n\n\n\n")
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	// tobestreamed := make([]byte, 100)
}
