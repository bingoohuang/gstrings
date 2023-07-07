package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bingoohuang/gstrings"
	"github.com/spf13/pflag"
)

var (
	conf  gstrings.ScanConfig
	files []string
)

func init() {
	pflag.IntVar(&conf.Min, "min", 6, "minimum length of UTF-8 strings printed, in runes")
	pflag.IntVar(&conf.Max, "max", 256, "maximum length of UTF-8 strings printed, in runes")
	pflag.BoolVarP(&conf.Ascii, "ascii", "a", false, "restrict strings to ASCII")
	pflag.BoolVarP(&conf.Tab, "tab", "t", false, "print strings separated by tabs other than new lines")
	pflag.StringVarP(&conf.Search, "search", "s", "", "search ASCII sub-string)")
	pflag.StringArrayVarP(&files, "files", "f", nil, "target file names")
	pflag.IntVarP(&conf.Most, "most", "n", 0, "print at most n places")
	pflag.BoolVar(&conf.Offset, "offset", false, "show file name and offset of start of each string")
	pflag.BoolVarP(&conf.Verbose, "verbose", "v", false, "display all input data.  Without the -v option, any output lines, which would be identical to the immediately preceding output line(except for the input offsets), are replaced with a line comprised of a single asterisk.")
	pflag.Parse()
}

var stdout *bufio.Writer

func main() {
	log.SetFlags(0)
	log.SetPrefix("strings: ")
	stdout = bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	if conf.Search != "" {
		conf.Min = len(conf.Search)
	}

	if conf.Max < conf.Min {
		conf.Max = conf.Min
	}

	files = append(files, pflag.Args()...)
	if len(files) == 0 {
		do(os.Stdin)
	}

	for _, f := range files {
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
	f := conf.NewScanner(file.Name())

	for in := bufio.NewReader(file); ; {
		if err := f.Scan(in); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Print(err)
			}

			return
		}
	}
}
