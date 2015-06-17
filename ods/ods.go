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

package ods

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/xml"
	"io"

	"gopkg.in/errgo.v1"
)

// After modifying anything under ods/seed/, run "go generate github.com/tgulacsi/csv2ods/ods"
//
// This may need "go get github.com/jteeuwen/go-bindata/go-bindata"
//
//go:generate go-bindata -o seed.go -prefix seed/ -pkg ods seed/...

var BufferSize = 1 << 20 // 1MiB

func NewWriter(w io.Writer) *writer {
	bw := bufio.NewWriterSize(w, BufferSize)
	zw := zip.NewWriter(bw)
	return &writer{zw: zw, bzw: bw}
}

type writer struct {
	zw  *zip.Writer
	bzw *bufio.Writer
	w   *errWriter
	bw  *bufio.Writer
}

// Closes the last written sheet (if any) and opens a new for writing.
func (w *writer) BeginSheet(name string, header []string) error {
	if w.w == nil {
		for _, fn := range []string{"manifest.rdf", "META-INF/manifest.xml", "meta.xml", "mimetype", "styles.xml"} {
			cw, err := w.zw.Create(fn)
			if err != nil {
				return errgo.Notef(err, "create %q", fn)
			}
			if _, err = cw.Write(MustAsset(fn)); err != nil {
				return errgo.Notef(err, "write %q", fn)
			}
		}
		cw, err := w.zw.Create("content.xml")
		if err != nil {
			return err
		}
		w.bw = bufio.NewWriterSize(cw, BufferSize)
		w.w = &errWriter{w: w.bw}
		content := MustAsset("content.xml")
		w.w.Write(content[:bytes.Index(content, []byte(`<office:spreadsheet>`))])
		io.WriteString(w.w, "<office:spreadsheet>\n")
	} else { // close last sheet
		_, _ = io.WriteString(w.w, "</table:table>\n")
	}
	_, _ = io.WriteString(w.w, `<table:table table:name="`)
	if err := xml.EscapeText(w.w, []byte(name)); err != nil {
		return errgo.Notef(err, "escape sheet name %q", name)
	}
	_, _ = io.WriteString(w.w, `" table:style-name="ta1"><table:table-column table:style-name="co1" table:default-cell-style-name="Default"/>`+"\n")
	return w.w.Err()
}

// AddRow adds a new row to the end of the actual sheet.
func (w *writer) AddRow(fields []string) error {
	_, _ = io.WriteString(w.w, `  <table:table-row table:style-name="ro1">`)
	for _, cell := range fields {
		_, _ = io.WriteString(w.w, `    <table:table-cell office:value-type="string" calcext:value-type="string"><text:p>`)
		if err := xml.EscapeText(w.w, []byte(cell)); err != nil {
			return errgo.Notef(err, "escape cell value %q", cell)
		}
		_, _ = io.WriteString(w.w, `</text:p></table:table-cell>`+"\n")
	}
	_, err := io.WriteString(w.w, "  </table:table-row>\n")
	return err
}

// Close the last sheet and the full document, too. Flushes all buffers, but does not close the underlying io.Writer.
func (w *writer) Close() error {
	if _, err := io.WriteString(w.w, "</table:table>\n</office:spreadsheet></office:body></office:document-content>"); err != nil {
		return err
	}
	if err := w.bw.Flush(); err != nil {
		return err
	}
	if err := w.zw.Close(); err != nil {
		return err
	}
	if err := w.bzw.Flush(); err != nil {
		return err
	}
	return nil
}

type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) Write(p []byte) (int, error) {
	if ew.err != nil {
		return 0, ew.err
	}
	var n int
	n, ew.err = ew.w.Write(p)
	return n, ew.err
}

func (ew errWriter) Err() error {
	return ew.err
}
