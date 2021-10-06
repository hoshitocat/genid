package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

var (
	numOfIDs = flag.Uint("n", 1, "number generate IDs")
)

func init() {
	flag.Parse()
}

func main() {
	num := *numOfIDs
	if num == 0 {
		num = 1
	}

	for i := 0; i < int(num); i++ {
		fmt.Println(gen())
	}
}

func gen() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
