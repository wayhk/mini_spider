/* args_test.go - testing file of args package  */
/*
modification history
--------------------
2016/06/24, by Han Kai, created
*/
/*
DESCRIPTION
This file contains testing of args package
*/

package main

import (
	"testing"
)

//test args
func TestArgs(t *testing.T) {
	arg := ParseArgs()
	t.Logf("c:%s", arg.ConfigDir)
	t.Logf("d:%s", arg.LogDir)
	arg.PrintHelp()
	arg.PrintVersion()
}
