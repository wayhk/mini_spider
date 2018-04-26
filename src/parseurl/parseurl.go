/* parseurl.go - parse urlfile  func */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
*/
/*
DESCRIPTION
This file contains parse url file func
*/

package parseurl

import (
	"encoding/json"
	"io/ioutil"
)

import (
)

//parse url from urlListFile
func ParseUrl(filename string) (urls []string, err error) {
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	err = json.Unmarshal(text, &urls)
	if err != nil {
		return
	}
	log.Logger.Info("urls:%+v", urls)

	return
}
