package main

import (
	"fmt"
	"os"
)

type KVPair struct {
	key string
	value string
}

func main() {
	path := "data/Moscow.txt"
    input, err := os.ReadFile(path)
    if err != nil {
        panic(err)
    }
    map_out := Map(path, string(input))
    fmt.Println(map_out)
}
