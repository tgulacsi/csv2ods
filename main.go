// Copyright 2015 Tamás Gulácsi
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tgulacsi/csv2ods/ods"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/htmlindex"
	"golang.org/x/text/transform"
)

//go:generate go generate ./ods

func main() {
	var (
		OutFileName string
		SheetTitles []string
		titles      string
		noHeading   bool
		ShowHelp    bool
		err         error
	)
	charset := os.Getenv("LANG")
	if i := strings.IndexByte(charset, '.'); i < 0 {
		charset = ""
	} else {
		charset = charset[i+1:]
	}

	flag.StringVar(&OutFileName, "o", "", "output file name or stdout if empty")
	flag.BoolVar(&noHeading, "noheading", false, "first row is heading, by default")
	flag.StringVar(&titles, "titles", "", "sheet titles, comma separated - the default is the names of input files")
	flag.StringVar(&charset, "charset", charset, "input CSV's charset")
	flag.BoolVar(&ShowHelp, "h", false, "show help")
	flag.Parse()

	if ShowHelp {
		flag.Usage()
		os.Exit(0)
	}

	var enc encoding.Encoding
	if charset != "" {
		if enc, err = htmlindex.Get(charset); err != nil {
			log.Fatalf("cannot find encoding for charset %q: %v", charset, err)
		}
	}

	if strings.TrimSpace(titles) != "" {
		SheetTitles = strings.Split(titles, ",")
	}

	inStreams, outStream := make([]io.Reader, 0, flag.NArg()+1), os.Stdout
	for i, fn := range flag.Args() {
		if fn == "" {
			continue
		}
		fh, err := os.Open(fn)
		if err != nil {
			log.Fatalf("open input %q: %v", fn, err)
		}
		defer fh.Close()
		inStreams = append(inStreams, fh)
		if len(SheetTitles) <= i {
			SheetTitles = append(SheetTitles, strings.TrimSuffix(filepath.Base(fh.Name()), ".csv"))
		}
	}
	log.Printf("sheet titles: %q", SheetTitles)
	if len(inStreams) == 0 {
		inStreams = append(inStreams, os.Stdin)
		defer os.Stdin.Close()
	}

	if OutFileName != "" {
		if outStream, err = os.Create(OutFileName); err != nil {
			log.Fatalf("create output file %q: %v", OutFileName, err)
		}
		defer func() {
			if err := outStream.Close(); err != nil {
				log.Printf("Close output: %v", err)
			}
		}()
	}

	out := ods.NewWriter(outStream)

	for n, in := range inStreams {
		if enc != nil {
			in = transform.NewReader(in, enc.NewDecoder())
		}
		bin := bufio.NewReaderSize(in, 16<<20)
		line, _ := bin.Peek(1024)
		inCSV := csv.NewReader(bin)
		inCSV.LazyQuotes = true
		inCSV.FieldsPerRecord = -1
		if bytes.Contains(line, []byte(`";"`)) {
			inCSV.Comma = ';'
		} else if bytes.Contains(line, []byte("\"\t\"")) {
			inCSV.Comma = '\t'
		}
		//log.Printf("inCSV=%#v comma=%q", inCSV, string([]rune{inCSV.Comma}))

		if err := out.BeginSheet(SheetTitles[n], nil); err != nil {
			log.Fatalf("begin %d. sheet: %v", n+1, err)
		}
		for {
			row, err := inCSV.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatalf("read csv %q: %v", in, err)
			}

			if err = out.AddRow(row); err != nil {
				log.Fatalf("add row %q: %v", row, err)
			}
		}
	}

	if err := out.Close(); err != nil {
		log.Fatalf("Close output: %v", err)
	}
}
