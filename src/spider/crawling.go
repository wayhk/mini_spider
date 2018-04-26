/* crawling.go - the main file of execute crawling task */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
2016/07/25,by Han Kai,modify,split Crawl func
2016/08/02,by Han Kai,modify,async push taskurl
*/
/*
DESCRIPTION
This file contains crawling func
*/

package spider

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

import (
)

//crawling controller
func (sp *Spider) Crawling() {
	for {
		select {
		case task := <-sp.Tasks:
			sp.Crawl(task)
			time.Sleep(time.Duration(sp.Conf.CrawlInterval) * time.Second)
		}
	}
}

//crawl func,contains get page,parse page,append task
func (sp *Spider) Crawl(urlTask UrlTask) {
	defer sp.Wg.Done()

	u, err := url.Parse(urlTask.Url)
	if err != nil {
		log.Logger.Error("url is invaild:%+v", err)
		return
	}

	//same host,set Interval
	for {
		timestamp, ok := sp.HostCache.Get(u.Host)
		if !ok || (time.Now().Unix()-timestamp > sp.Conf.CrawlInterval) {
			sp.HostCache.Put(u.Host, time.Now().Unix())
			break
		} else {
			time.Sleep(time.Duration(sp.Conf.CrawlInterval) * time.Second)
		}
	}
	//get page content
	data, err := sp.HttpRequest(urlTask)
	if err != nil {
		log.Logger.Error("get page data err:%+v", err)
		return
	}

	//save page data to file
	err = sp.SaveData(urlTask.Url, data)
	if err != nil {
		log.Logger.Error("save data to file err:%+v", err)
		return
	}

	//parse page,get valid url
	urls := ParsePage(urlTask.Url, data)
	var taskUrl []string
	for _, url := range urls {
		match := regexp.MustCompile(sp.Conf.TargetUrl).MatchString(url)
		if !match {
			log.Logger.Warn("url:%s is not match given targetUrl:%s", urlTask.Url, sp.Conf.TargetUrl)
		}
		taskUrl = append(taskUrl, url)
	}

	//append to tasks chan
	for _, url := range taskUrl {
		_, ok := sp.TaskCache.Get(url)
		if ok {
			log.Logger.Info("url has exist:%s", url)
			continue
		}
		var task UrlTask
		task.Url = url
		task.Depth = urlTask.Depth + 1
		if task.Depth > sp.Conf.MaxDepth {
			log.Logger.Info("url:%s has reach max depth", url)
			break
		}
		//add first,async push
		sp.AddTask(task)
	}
	log.Logger.Info("url:%s parse finished", urlTask.Url)
	return
}

//get page data
func (sp *Spider) HttpRequest(urlTask UrlTask) (data []byte, err error) {

	client := http.Client{
		//set timeout
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Duration(sp.Conf.CrawlTimeout)*time.Second)
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		},
	}
	resp, err := client.Get(urlTask.Url)
	if err != nil {
		log.Logger.Error("get url:%s response failed:%+v", urlTask.Url, err)
		return
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		log.Logger.Error("read response failed:%#v", err)
		return
	}
	data = body.Bytes()
	return
}

//save page content to file
func (sp *Spider) SaveData(url string, data []byte) (err error) {

	//deal with spec word
	newname := strings.Replace(url, "/", "%20F", -1)

	//use md5 as filename
	h := md5.New()
	h.Write([]byte(newname))
	fileName := hex.EncodeToString(h.Sum(nil))

	//write content to file
	filePath := filepath.Join(sp.Conf.OutputDirectory, fileName)
	fileAbsPath, err := filepath.Abs(filePath)
	if err != nil {
		log.Logger.Error("get file abs path failed:%#v", err)
		return
	}
	err = ioutil.WriteFile(fileAbsPath, data, 0644)
	if err != nil {
		log.Logger.Error("write content to file:%s failed:%#v", fileAbsPath, err)
		return
	}
	return
}
