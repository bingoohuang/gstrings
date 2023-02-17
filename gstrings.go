// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gstrings is a more capable, UTF-8 aware version of the standard strings utility.
package gstrings

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ScanConfig struct {
	Min     int
	Max     int
	Ascii   bool
	Tab     bool
	Search  string
	Most    int
	Offset  bool
	Verbose bool
}

type Scanner struct {
	file string

	str        []rune
	pos        int64
	printTimes int
	lastPrint  string
	*ScanConfig
}

func (c *ScanConfig) NewScanner(file string) *Scanner {
	return &Scanner{
		file:       file,
		str:        make([]rune, 0, c.Max),
		ScanConfig: c,
	}
}

func (f *Scanner) Scan(in *bufio.Reader) error {
	var r rune
	var wid int
	var err error

	if f.Min <= 0 {
		f.Min = 6
	}
	if f.Max <= 0 {
		f.Max = 256
	}

	// One string per loop.
	for ; ; f.pos += int64(wid) {
		if r, wid, err = in.ReadRune(); err != nil {
			return err
		}
		if !strconv.IsPrint(r) || f.Ascii && r >= 0xFF {
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

func (f *Scanner) print() {
	if len(f.str) < f.Min {
		f.str = f.str[:0]
		return
	}

	s := string(f.str)
	if f.Search == "" || strings.Contains(s, f.Search) {
		if !f.Verbose {
			if f.lastPrint == s {
				s = "*"
			} else {
				f.lastPrint = s
			}
		}
		if f.Offset {
			s = fmt.Sprintf("%s:#%d:\t%s", f.file, f.pos-int64(len(s)), s)
		}

		if f.Tab {
			fmt.Print(s)
			fmt.Print("\t")
		} else {
			fmt.Println(s)
		}
		f.printTimes++

		if f.Max > 0 && f.printTimes >= f.Max {
			fmt.Println()
			os.Exit(0)
		}
	}

	f.str = f.str[:0]
}
