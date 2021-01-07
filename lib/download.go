package lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func download(num, t, page string) {
	err := os.MkdirAll(t, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("create dir error-->%s", err.Error()))
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		panic(fmt.Sprintf("page set error-->%s", err.Error()))
	}

	var wg sync.WaitGroup
	imagePath := addImageAddr(num)
	wg.Add(p)
	for i := 1; i <= p; i++ {
		pageName := fmt.Sprintf("%05d.jpg", i)
		pagePath := fmt.Sprintf("%s/%s", imagePath, pageName)
		go func() {
			downloadFile(fmt.Sprintf("%s/%s", t, pageName), pagePath)
			wg.Done()
		}()
	}
	wg.Wait()
}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
