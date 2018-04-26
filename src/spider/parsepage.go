/* parsepage.go - the main file of parse page  */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
2016/08/15,by Han Kai,modify,add err log when parse url failed
*/
/*
DESCRIPTION
This file contains parse page func and convert url func
*/

package spider

import (
	"net/url"
	"strings"
)

import (
	"golang.org/x/net/html"
)

//parse page func.param:page url,page content.return sub urls
func ParsePage(urlAddr string, content []byte) []string {
	// parse
	urls := parse(string(content))

	// convert relative urls to absolute urls
	taskUrls := convertUrls(urlAddr, urls)
	return taskUrls
}

//parse page content,return urls
func parse(con string) (urls []string) {
	r := strings.NewReader(con)
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						urls = append(urls, a.Val)
						break
					}
				}
			}
		}
	}
	return
}

//convert relative urls to abs urls
func convertUrls(pageUrl string, urls []string) []string {
	u, _ := url.Parse(pageUrl)

	var absUrls []string
	for _, urlItem := range urls {
		urlObj, err := u.Parse(urlItem)
		if err != nil {
			log.Logger.Error("parse url:%s failed,err:%+v", urlItem, err)
			continue
		}
		absUrls = append(absUrls, urlObj.String())
	}
	return absUrls
}
