package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	f, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	start := time.Now()
	var data []map[string]interface{}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&data); err != nil {
		log.Fatal(err)
	}
	// Extract the Go version.
	fmt.Printf("[%s] %d rows loaded in %s\n", runtime.Version(), len(data), (time.Since(start)))
}
