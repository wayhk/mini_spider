/* parseurl_test.go - testing file of parseurl package  */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
*/
/*
DESCRIPTION
This file contains testing of parseurl package
*/

package parseurl

import (
	"testing"
)

//test ParseUrl()
func TestParseUrl(t *testing.T) {
	fileName := "a.txt"
	urls, err := ParseUrl(fileName)
	if err != nil {
		t.Logf("parse url err:%#v,return urls:%+v", err, urls)
	}

	fileName = "../../data/url.data"
	urls, err = ParseUrl(fileName)
	if err != nil {
		t.Errorf("parse url err:%#v,return urls:%+v", err, urls)
	}
}
