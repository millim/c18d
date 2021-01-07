package lib

import (
	"fmt"
	"log"
)

type DownloadTimer struct {
	MaxPage int
	NowPage int
}

var timer *DownloadTimer

func newTimer(maxPage int) {
	timer = &DownloadTimer{MaxPage: maxPage}
	p()
}

func incTimer() {
	if timer != nil {
		timer.NowPage += 1
		p()
	}
}

func p() {
	log.Print(fmt.Sprintf("进度: %d/%d", timer.NowPage, timer.MaxPage))
}
