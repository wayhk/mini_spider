/* config_test.go - test func of analyse config  */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
*/
/*
DESCRIPTION
This file contains test func of analyse config
*/

package config

import (
	"testing"
)

//test LoadConfig
func TestLoadConfig(t *testing.T) {
	fileName := "a.txt"
	_, err := LoadConfig(fileName)
	if err != nil {
		t.Logf("test LoadConfig err:%#v", err)
	}

	fileName = "../../conf/spider.conf"
	sc, err := LoadConfig(fileName)
	if err != nil {
		t.Errorf("test LoadConfig err :%#v", err)
	}
	if sc.CrawlInterval != 1 && sc.UrlListFile != "data/url.data" {
		t.Errorf("get conf data failed,crawlInterval:%d,urlListFile:%s", sc.CrawlInterval, sc.UrlListFile)
	}

}

//test IsValid()
func TestIsValid(t *testing.T) {
	sp := &SpiderConfig{
		UrlListFile:     "a.txt",
		OutputDirectory: "out",
		MaxDepth:        1,
		CrawlInterval:   1,
		CrawlTimeout:    1,
		TargetUrl:       "\b \t \n \f \r \"\\",
		ThreadCount:     1,
	}
	err := sp.IsValid()
	if err != nil {
		t.Errorf("IsValid test err:%#v", err)
	}
}
