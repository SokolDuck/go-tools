package goget

import (
	"fmt"
	"path"
	"sync"
	"time"

	"github.com/SokolDuck/go-tools/goget/utils"
)

func DownloadFiles(urls []string) {
	statuses := make(map[string]*utils.DownloadStatus)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			utils.Download(url, statuses)
		}(url)
	}

	done := make(chan bool)
	go func() {
		for index, url := range urls {
			fmt.Printf("%v. %s ", index, path.Base(url))
		}
		fmt.Println("")

		for {
			select {
			case <-done:
				return
			default:
				utils.PrintStatuses(statuses, urls)
				time.Sleep(1 * time.Second)
			}
		}

	}()

	wg.Wait()
	close(done)
}
