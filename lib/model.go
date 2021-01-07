package lib

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Book struct {
	Number          string
	Title           string
	DownloadDirPath string
	MaxPage         string
}

func GetBook(num string) *Book {
	return &Book{
		Number: num,
	}
}

func (book *Book) SetDownloadDirPath(path string) {
	book.DownloadDirPath = path
}

func (book *Book) GetTitleInfo() error {
	err, title, page := getTitleAndPage(book.Number)
	if err != nil {
		return err
	}
	book.Title = title
	book.MaxPage = page
	return nil
}

func (book *Book) Zip() {
	if book.DownloadDirPath == "" {
		book.DownloadDirPath = book.Title
	}
	Zip(book.DownloadDirPath, fmt.Sprintf("%s.zip", book.DownloadDirPath))
}

func (book *Book) DeleteDir() {
	if book.DownloadDirPath == "" {
		book.DownloadDirPath = book.Title
	}
	os.RemoveAll(book.DownloadDirPath)
}

func (book *Book) Download() {
	if book.DownloadDirPath == "" {
		book.DownloadDirPath = book.Title
	}

	err := os.MkdirAll(book.DownloadDirPath, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("create dir error-->%s", err.Error()))
	}

	p, err := strconv.Atoi(book.MaxPage)
	if err != nil {
		panic(fmt.Sprintf("page set error-->%s", err.Error()))
	}

	var wg sync.WaitGroup
	imageURL := addImageAddr(book.Number)
	wg.Add(p)
	for i := 1; i <= p; i++ {
		pageName := fmt.Sprintf("%05d.jpg", i)
		pageURLPath := fmt.Sprintf("%s/%s", imageURL, pageName)
		go func() {
			downloadFile(fmt.Sprintf("%s/%s", book.DownloadDirPath, pageName), pageURLPath)
			wg.Done()
		}()
	}
	wg.Wait()
}
