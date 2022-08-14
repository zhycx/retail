package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"time"
)

func main() {
	t := time.Now()
	number := 1
	c := colly.NewCollector(
		colly.URLFilters(
			regexp.MustCompile("^(https://s\\.weibo\\.com/top/summary)\\?cate=realtimehot"),
		)) // 创建收集器
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0")
		r.Headers.Set("Cookie", "SUB=_2AkMVq_fBf8NxqwFRmP4cy27haYV0yQ3EieKj9wYaJRMxHRl-yT8XqhYytRB6PivZLsfksaG0SmV7y-8z09Hc8E2mU8Ur; SUBP=0033WrSXqPxfM72-Ws9jqgMF55529P9D9W5jhjMTN6uuLkzKJq6cWm6E; _s_tentry=passport.weibo.com; Apache=8908329336088.176.1660385529151; SINAGLOBAL=8908329336088.176.1660385529151; ULV=1660385529165:1:1:1:8908329336088.176.1660385529151:")
		r.Headers.Set("Referer", "https://passport.weibo.com/")
	})
	// 响应的格式为HTML,提取页面中的链接
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if number >= 6 {
			fmt.Print(number-5, ":", e.Text, "\n")
		}
		number += 1

		c.Visit(e.Request.AbsoluteURL(link))
	})
	// 获取电影信息
	//c.OnHTML("div.info", func(e *colly.HTMLElement) {
	//	e.DOM.Each(func(i int, selection *goquery.Selection) {
	//		movies := selection.Find("span.title").First().Text()
	//		director := strings.Join(strings.Fields(selection.Find("div.bd p").First().Text()), " ")
	//		quote := selection.Find("p.quote span.inq").Text()
	//		fmt.Printf("%d --> %s:%s %s\n", number, movies, director, quote)
	//		number += 1
	//	})
	//})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})
	c.Visit("https://s.weibo.com/top/summary?cate=realtimehot")
	c.Wait()
	fmt.Printf("花费时间:%s",time.Since(t))
}