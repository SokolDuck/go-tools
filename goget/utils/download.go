package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type DownloadStatus struct {
	Total   int64
	Current int64
}

type StatusWriter struct {
	status *DownloadStatus
}

func (sw *StatusWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	sw.status.Current += int64(n)

	return n, nil
}

func Download(url string, statuses map[string]*DownloadStatus) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	outFile, err := os.Create(path.Base(url))
	if err != nil {
		fmt.Printf("Error creating file for %s: %v\n", url, err)
		return
	}
	defer outFile.Close()

	status := &DownloadStatus{
		Total: resp.ContentLength,
	}
	statuses[url] = status

	_, err = io.Copy(outFile, io.TeeReader(resp.Body, &StatusWriter{status: status}))
	if err != nil {
		fmt.Printf("Error saving %s: %v\n", url, err)
	}
}

func PrintStatuses(statuses map[string]*DownloadStatus, urls []string) {
	fmt.Printf("\r")
	for _, url := range urls {
		s, ok := statuses[url]
		if !ok {
			fmt.Print("[-----] ")
			continue
		}

		if s.Total != -1 {
			percentage := float64(s.Current) / float64(s.Total) * 100
			fmt.Printf("[%4.0f%%] ", percentage)
		} else {
			// The case when the get request does not return the correct ContentLength
			fmt.Print("[-xx%-] ")
		}
	}
	fmt.Print("\n") // Comment if you want a nice one-line status update.
}
