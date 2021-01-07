package lib

import (
	"fmt"
	"log"
	"os"
)

//Run 执行
func Run(num string) error {
	log.Println("准备中....")
	err, title, page := getTitleAndPage(num)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	log.Println(fmt.Sprintf("文件名: %s", title))
	log.Println(fmt.Sprintf("下载中....总共%s页", page))
	download(num, title, page)
	log.Println("正在打包压缩")
	Zip(title, fmt.Sprintf("%s.zip", title))
	log.Println("移除源文件")
	os.RemoveAll(title)
	log.Println("完成")
	return nil
}
