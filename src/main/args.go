/* args.go - this file contains parse cmd args func  */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
2016/08/02,by Han Kai,modify.read version from version file.
*/
/*
DESCRIPTION
This file contains parse cmd args func
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

//cmd params struct
type Args struct {
	Help      bool
	Version   bool
	ConfigDir string
	LogDir    string
}

const (
	VERSIONFILE = "./version"
)

//parse input args
func ParseArgs() *Args {
	var arg Args
	flag.BoolVar(&arg.Help, "h", false, "show help")
	flag.BoolVar(&arg.Version, "v", false, "show version")
	flag.StringVar(&arg.ConfigDir, "c", "conf/spider.conf", "config file position")
	flag.StringVar(&arg.LogDir, "l", "log", "log dir")
	flag.Parse()

	return &arg
}
func (arg *Args) PrintHelp() {
	flag.PrintDefaults()
}

//print version func
func (arg *Args) PrintVersion() {
	version, err := GetVersionContent(VERSIONFILE)
	if err != nil {
		fmt.Println("get version err:%+v", err)
		return
	}
	fmt.Println("version: " + version)
}

//get version from version file
func GetVersionContent(path string) (content string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	content = string(data)
	return
}
