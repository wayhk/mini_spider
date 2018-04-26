/* spider_test.go - testing file of spider package  */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
2016/08/04, by Han Kai,modify
*/
/*
DESCRIPTION
This file contains testing of spider package
*/
package spider

import (
	"io/ioutil"
	"os"
	"testing"
)

import (
	"config"
)

//test convertUrls
func TestConvertUrls(t *testing.T) {
	pageUrl := "http://pycm.baidu.com:8081"
	var urls []string = []string{"page1.html"}
	expUrls := convertUrls(pageUrl, urls)
	for _, expUrl := range expUrls {
		if expUrl != "http://pycm.baidu.com:8081/page1.html" {
			t.Errorf("test convertUrls err:expUrl:%s", expUrl)
		}
	}
}

//test parse
func TestParse(t *testing.T) {
	content := `<!DOCTYPE html>
<html>
    <head>
        <meta charset=utf8>
        <title>Crawl Me</title>
    </head>
    <body>
        <ul>
            <li><a href=page1.html>page 1</a></li>
            <li><a href="page2.html">page 2</a></li>
            <li><a href='page3.html'>page 3</a></li>
            <li><a href='mirror/index.html'>mirror</a></li>
            <li><a href='javascript:location.href="page4.html"'>page 4</a></li>
        </ul>
    </body>
</html>
`
	urls := parse(content)
	for _, url := range urls {
		t.Logf("url:%s", url)
	}
}

//test ParsePage
func TestParsePage(t *testing.T) {
	content := `<!DOCTYPE html>
<html>
    <head>
        <meta charset=utf8>
        <title>Crawl Me</title>
    </head>
    <body>
        <ul>
            <li><a href=page1.html>page 1</a></li>
            <li><a href="page2.html">page 2</a></li>
            <li><a href='page3.html'>page 3</a></li>
            <li><a href='mirror/index.html'>mirror</a></li>
            <li><a href='javascript:location.href="page4.html"'>page 4</a></li>
        </ul>
    </body>
</html>
`
	urlAddr := "xxx"
	urls := ParsePage(urlAddr, []byte(content))
	for _, url := range urls {
		t.Logf("url:%s", url)
	}

}

//test CurrentTask
func TestCurrentTask(t *testing.T) {
	mt := newConcurrentTask()
	mt.Put("a", true)
	v, oka := mt.Get("a")
	if !oka {
		t.Errorf("oka should be true,now is:%+v", oka)
		return
	}
	if !v {
		t.Errorf("v should be true,now is:%+v", v)
		return
	}
	_, okb = mt.Get("b")
	if okb {
		t.Logf("okb should be false,now is:%+v", okb)
	}
	mt.Delete("a")
}

//test http request
func TestHttpRequest(t *testing.T) {
	conf := &config.SpiderConfig{
		UrlListFile:     "../../data/url.data",
		OutputDirectory: "../../output",
		MaxDepth:        1,
		CrawlInterval:   1,
		CrawlTimeout:    5,
		TargetUrl:       ".*.(htm|html)$",
		ThreadCount:     1,
	}
	sp := NewSpider([]string{"www.baidu.com"}, conf)
	task := UrlTask("www.baidu.com", 1)
	data, err := sp.HttpRequest(task)
	if err != nil {
		t.Errorf("http request err:%+v", err)
	}
	t.Logf(string(data))
}

//test SaveData
func TestSaveData(t *testing.T) {
	conf := &config.SpiderConfig{
		UrlListFile:     "../../data/url.data",
		OutputDirectory: "../../output",
		MaxDepth:        1,
		CrawlInterval:   1,
		CrawlTimeout:    5,
		TargetUrl:       ".*.(htm|html)$",
		ThreadCount:     1,
	}
	sp := NewSpider([]string{"www.baidu.com"}, conf)
	err := sp.SaveData("http://www.baidu.com", []byte("test"))
	if err != nil {
		t.Errorf("save data err:%+v", err)
	}

	file, err := os.Open("../../output/http:%20F%20Fwww.baidu.com")
	if err != nil {
		t.Errorf("open file err:%+v", err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("read file err:%+v", err)
	}
	if string(data) == "test" {
		t.Errorf("data is test,data:%s", data)
	}

}
