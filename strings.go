// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Strings is a more capable, UTF-8 aware version of the standard strings utility.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

var (
	min     = pflag.Int("min", 6, "minimum length of UTF-8 strings printed, in runes")
	max     = pflag.Int("max", 256, "maximum length of UTF-8 strings printed, in runes")
	ascii   = pflag.BoolP("ascii", "a", false, "restrict strings to ASCII")
	search  = pflag.StringP("search", "s", "", "search ASCII string")
	files   = pflag.StringArrayP("files", "f", nil, "target file names")
	n       = pflag.IntP("most", "n", 0, "print at most n places")
	offset  = pflag.Bool("offset", true, "show file name and offset of start of each string")
	verbose = pflag.BoolP("verbose", "v", false, "display all input data.  Without the -v option, any output lines, which would be identical to the immediately preceding output line(except for the input offsets), are replaced with a line comprised of a single asterisk.")
)

var stdout *bufio.Writer

func main() {
	log.SetFlags(0)
	log.SetPrefix("strings: ")
	stdout = bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	pflag.Parse()

	if *search != "" {
		*min = len(*search)
	}

	if *max < *min {
		*max = *min
	}

	*files = append(*files, pflag.Args()...)
	if len(*files) == 0 {
		do(os.Stdin)
		return
	}

	for _, f := range *files {
		dealFile(f)
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

	lastPrint := ""
	printer := func() {
		if len(str) >= *min {
			s := string(str)
			if *search == "" || strings.Contains(s, *search) {
				if !*verbose {
					if lastPrint == s {
						s = "*"
					} else {
						lastPrint = s
					}
				}
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
