// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gstrings is a more capable, UTF-8 aware version of the standard strings utility.
package gstrings

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type ScanConfig struct {
	Hex     string
	Search  string
	Min     int
	Max     int
	Most    int
	Ascii   bool
	Tab     bool
	Offset  bool
	Verbose bool
}

type Scanner struct {
	*ScanConfig
	file string

	lastPrint string

	printable  []rune
	pos        int64
	printTimes int
}

func (c *ScanConfig) NewScanner(file string) *Scanner {
	return &Scanner{
		file:       file,
		printable:  make([]rune, 0, c.Max),
		ScanConfig: c,
	}
}

func (f *Scanner) Scan(in io.RuneReader) error {
	var r rune
	var wid int
	var err error

	f.Min = lo.Ternary(f.Min <= 0, 6, f.Min)
	f.Max = lo.Ternary(f.Max <= 0, 256, f.Max)

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
		f.printable = append(f.printable, r)
		if len(f.printable) >= cap(f.printable) {
			f.print()
		}
	}
}

func (f *Scanner) print() {
	if len(f.printable) < f.Min {
		f.printable = f.printable[:0]
		return
	}

	s := string(f.printable)
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

		if f.Most > 0 && f.printTimes >= f.Most {
			fmt.Println()
			os.Exit(0)
		}
	}

	f.printable = f.printable[:0]
}
