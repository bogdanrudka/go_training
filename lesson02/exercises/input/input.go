package main

import (
	"bufio"
	"io"
	"log"
	"strings"
)

// ReadAll reads all the lines of text from r and returns
// all the data read as a string
func ReadAll(r io.Reader) string {
	fr := bufio.NewScanner(r)
	var lns []string
	for fr.Scan(){
		lns = append(lns, fr.Text())
	}
	//Why we should check this error after reading file?
	//if fr.Err() != nil {
	//	log.Fatal(fr.Err())
	//}
	return strings.Join(lns, "\n")
}
