package lib

import (
	"fmt"
	"log"
)

//Run 执行
func Run(num string) error {
	log.Println("准备初始化中....")
	book := GetBook(num)
	err := book.GetTitleInfo()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	log.Println(fmt.Sprintf("文件名: %s", book.Title))
	log.Println(fmt.Sprintf("下载中....总共%s页", book.MaxPage))
	book.Download()
	log.Println("正在打包压缩")
	book.Zip()
	log.Println("移除源文件")
	book.DeleteDir()
	log.Println("完成")
	return nil
}
