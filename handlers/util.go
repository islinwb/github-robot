package handlers

import "regexp"

var (
	labelReg  = regexp.MustCompile("^/[Ll][Aa][Bb][Ee][Ll]")
	testReg   = regexp.MustCompile("^/[Tt][Ee][Ss][Tt]")
	retestReg = regexp.MustCompile("^/[Rr][Ee][Tt][Ee][Ss][Tt]")
)
