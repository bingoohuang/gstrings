// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Strings is a more capable, UTF-8 aware version of the standard strings utility.
//
// Flags(=default) are:
//
//	-ascii(=false)    restrict strings to ASCII
//	-search=abc       search string abc
//	-min(=6)          minimum length of UTF-8 strings printed, in runes
//	-max(=256)        maximum length of UTF-8 strings printed, in runes
//	-offset(=true)    show file name and offset of start of each string
package main // import "robpike.io/cmd/strings"

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	min    = flag.Int("min", 6, "minimum length of UTF-8 strings printed, in runes")
	max    = flag.Int("max", 256, "maximum length of UTF-8 strings printed, in runes")
	ascii  = flag.Bool("ascii", false, "restrict strings to ASCII")
	search = flag.String("search", "", "search ASCII string")
	n      = flag.Int("n", 0, "print at most n places")
	offset = flag.Bool("offset", true, "show file name and offset of start of each string")
)

var stdout *bufio.Writer

func main() {
	log.SetFlags(0)
	log.SetPrefix("strings: ")
	stdout = bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	flag.Parse()

	if *search != "" {
		*min = len(*search)
	}

	if *max < *min {
		*max = *min
	}

	if flag.NArg() == 0 {
		do(os.Stdin)
		return
	}

	for _, arg := range flag.Args() {
		dealFile(arg)
	}
}

func dealFile(file string) {
	fd, err := os.Open(file)
	if err != nil {
		log.Print(err)
		return
	}

	defer func() {
		fd.Close()
		stdout.Flush()
	}()

	do(fd)
}

func do(file *os.File) {
	str := make([]rune, 0, *max)
	pos := int64(0)
	printTimes := 0
	printer := func() {
		if len(str) >= *min {
			s := string(str)
			if *search == "" || strings.Contains(s, *search) {
				if *offset {
					s = fmt.Sprintf("%s:#%d:\t%s", file.Name(), pos-int64(len(s)), s)
				}

				fmt.Println(s)
				printTimes++

				if *n > 0 && printTimes >= *n {
					os.Exit(0)
				}
			}
		}
		str = str[:0]
	}

	for in := bufio.NewReader(file); ; {
		var r rune
		var wid int
		var err error

		// One string per loop.
		for ; ; pos += int64(wid) {
			if r, wid, err = in.ReadRune(); err != nil {
				if err != io.EOF {
					log.Print(err)
				}
				return
			}
			if !strconv.IsPrint(r) || *ascii && r >= 0xFF {
				printer()
				continue
			}
			// It's printable. Keep it.
			str = append(str, r)
			if len(str) >= cap(str) {
				printer()
			}
		}
	}
}
