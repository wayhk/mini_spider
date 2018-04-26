/* config.go - the main structure of config  */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
2016/07/25,by Han Kai,modify,add check config in LoadConfig func
*/
/*
DESCRIPTION
This file contains spider config structure and analyze func
*/

package config

import (
	"errors"
	"regexp"
)

import (
	"code.google.com/p/gcfg"
)

//config struct
type Config struct {
	Spider SpiderConfig
}

//spider config struct
type SpiderConfig struct {
	UrlListFile     string
	OutputDirectory string
	MaxDepth        int64
	CrawlInterval   int64
	CrawlTimeout    int64
	TargetUrl       string
	ThreadCount     int64
}

//load config from file
func LoadConfig(filename string) (spiderConfig *SpiderConfig, err error) {
	var config Config
	err = gcfg.ReadFileInto(&config, filename)
	if err != nil {
		return
	}

	spiderConfig = config.GetSpiderConfig()

	//check config is or not valid
	err = spiderConfig.IsValid()
	return
}

//get spider config
func (config *Config) GetSpiderConfig() *SpiderConfig {
	return &config.Spider
}

//check config is or not valid.
//urlListFile,outputDirectory, has default set,so not check
func (sc *SpiderConfig) IsValid() error {
	if sc.MaxDepth < 0 {
		return errors.New("maxDepth invalid,should >=0")
	}

	if sc.CrawlInterval < 1 {
		return errors.New("crawlInterval invalid,should >=1")
	}

	if sc.CrawlTimeout < 0 {
		return errors.New("crawlTimeout invalid,should >=0")
	}

	if _, err := regexp.Compile(sc.TargetUrl); err != nil {
		return errors.New("targetUrl invaild")
	}

	if sc.ThreadCount < 1 {
		return errors.New("threadCount invalid,should >=1")
	}

	return nil
}
