package main

import (
	"fmt"
	"os"

	"github.com/SokolDuck/go-tools/goget"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: goget <url1> <url2> ...")
		os.Exit(1)
	}

	urls := os.Args[1:]

	goget.DownloadFiles(urls)
}
