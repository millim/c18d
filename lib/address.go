package lib

import "fmt"

func addrView(num string) string {
	return fmt.Sprintf("https://18comic.bet/photo/%s/", num)
}

func addImageAddr(num string) string {
	return fmt.Sprintf("https://dx.18comic.asia/media/photos/%s", num)
}
