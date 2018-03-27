package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/bearx3f/a2conf"
)

type options struct {
	tab2space bool
	tabsize   int
	file      string
}

var opt options

func init() {
	flag.BoolVar(&opt.tab2space, "tab2space", false, "Convert to tab to space for indent.")
	flag.IntVar(&opt.tabsize, "tabsize", 4, "indent space size")
	flag.StringVar(&opt.file, "file", "", "Formatting file")
	flag.Parse()
}

func main() {
	var fp *os.File
	var err error

	if _, err := os.Stat(opt.file); os.IsNotExist(err) {
		log.Fatalln(err)
	}

	if fp, err = os.Open(opt.file); err != nil {
		log.Fatalln(err)
	}
	output, _ := a2conf.Format(fp, a2conf.FormatOption{
		Tab2Space: opt.tab2space,
		TabSize:   opt.tabsize,
	})
	if err = fp.Close(); err != nil {
		log.Fatalln(err)
	}

	if fp, err = os.OpenFile(opt.file, os.O_WRONLY|os.O_TRUNC, 0600); err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := fp.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if _, err := io.Copy(fp, strings.NewReader(output)); err != nil {
		log.Println(err)
	}
}
