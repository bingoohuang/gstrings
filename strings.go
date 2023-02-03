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
	tab     = pflag.BoolP("tab", "t", false, "print strings separated by tabs other than new lines")
	search  = pflag.StringP("search", "s", "", "search ASCII sub-string)")
	files   = pflag.StringArrayP("files", "f", nil, "target file names")
	n       = pflag.IntP("most", "n", 0, "print at most n places")
	offset  = pflag.Bool("offset", false, "show file name and offset of start of each string")
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
	}

	for _, f := range *files {
		dealFile(f)
		fmt.Println()
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
	f := newFile(file)

	for in := bufio.NewReader(file); ; {
		if err := f.read(in); err != nil {
			if err != io.EOF {
				log.Print(err)
			}

			return
		}
	}
}

type File struct {
	file *os.File

	str        []rune
	pos        int64
	printTimes int
	lastPrint  string
}

func newFile(file *os.File) *File {
	return &File{
		file: file,
		str:  make([]rune, 0, *max),
	}
}

func (f *File) print() {
	if len(f.str) < *min {
		f.str = f.str[:0]
		return
	}

	s := string(f.str)
	if *search == "" || strings.Contains(s, *search) {
		if !*verbose {
			if f.lastPrint == s {
				s = "*"
			} else {
				f.lastPrint = s
			}
		}
		if *offset {
			s = fmt.Sprintf("%s:#%d:\t%s", f.file.Name(), f.pos-int64(len(s)), s)
		}

		if *tab {
			fmt.Print(s)
			fmt.Print("\t")
		} else {
			fmt.Println(s)
		}
		f.printTimes++

		if *n > 0 && f.printTimes >= *n {
			fmt.Println()
			os.Exit(0)
		}
	}

	f.str = f.str[:0]
}

func (f *File) read(in *bufio.Reader) error {
	var r rune
	var wid int
	var err error

	// One string per loop.
	for ; ; f.pos += int64(wid) {
		if r, wid, err = in.ReadRune(); err != nil {
			return err
		}
		if !strconv.IsPrint(r) || *ascii && r >= 0xFF {
			f.print()
			continue
		}
		// It's printable. Keep it.
		f.str = append(f.str, r)
		if len(f.str) >= cap(f.str) {
			f.print()
		}
	}
}
