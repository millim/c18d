package lib

import (
	"errors"
	"github.com/gocolly/colly"
	"strings"
)

func collyConfig() *colly.Collector {
	return colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"),
	)
}

func getTitleAndPage(num string) (error, string, string) {
	c := collyConfig()
	title, page := "", ""
	c.OnHTML("#phpage > a > span", func(e *colly.HTMLElement) {

		page = e.Text

	})
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	err := c.Visit(addrView(num))
	if err != nil {
		return err, "", ""
	}

	c.Wait()
	if title == "" || page == "" {
		return errors.New("not title or page data"), "", ""
	}
	t := strings.Split(title, "|")[0]
	p := strings.Split(page, "/")[1]
	return nil, t, p
}
