// Code generated by go-bindata.
// sources:
// seed/META-INF/manifest.xml
// seed/content.xml
// seed/manifest.rdf
// seed/meta.xml
// seed/mimetype
// seed/styles.xml
// DO NOT EDIT!

package ods

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _metaInfManifestXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x92\xc1\x6a\x03\x21\x10\x86\xef\x79\x8a\xc5\x6b\x59\x6d\x7b\x2a\xb2\x6e\x6e\x7d\x82\xf6\x01\x44\x67\x1b\x41\x47\x71\x66\x43\xf6\xed\x6b\x42\x9a\x4d\x29\x81\x06\x72\x73\xc6\x7f\xfe\xff\x13\x67\xd8\x1e\x52\xec\xf6\x50\x29\x64\x34\xe2\x45\x3e\x8b\x0e\xd0\x65\x1f\xf0\xcb\x88\xcf\x8f\xf7\xfe\x4d\x6c\xc7\xcd\x90\x2c\x86\x09\x88\xf5\xcf\xa1\x6b\x73\x48\x97\xd2\x88\xb9\xa2\xce\x96\x02\x69\xb4\x09\x48\xb3\xd3\xb9\x00\xfa\xec\xe6\x04\xc8\xfa\xb7\x5e\x9f\x92\x2e\xd5\x15\xc0\xab\x18\x37\xdd\x9a\x37\x85\x08\x7d\x9b\xaf\xcb\xaa\x9e\xe6\x18\xfb\x62\x79\x67\x84\xba\x65\xb2\xb6\x13\xf8\x60\x7b\x5e\x0a\x18\x61\x4b\x89\xc1\x59\x6e\x32\xb5\x47\x2f\x4f\xc0\xf2\x9a\x53\x52\xa9\x60\x3d\xed\x00\x58\xa8\x7b\x50\x5c\x46\x3e\x1a\xb4\x87\xde\x48\x67\x38\xb0\x3a\x5e\xdf\xe5\x9b\x80\xed\xc3\x4d\x89\x97\x08\xf4\x78\xd6\x73\x4f\x56\x3f\xfd\xe3\x0b\x9a\xea\xe9\x9c\x31\xa8\x3f\x2b\x36\x6e\xbe\x03\x00\x00\xff\xff\x8e\x73\x64\x7e\x9e\x02\x00\x00")

func metaInfManifestXmlBytes() ([]byte, error) {
	return bindataRead(
		_metaInfManifestXml,
		"META-INF/manifest.xml",
	)
}

func metaInfManifestXml() (*asset, error) {
	bytes, err := metaInfManifestXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "META-INF/manifest.xml", size: 670, mode: os.FileMode(436), modTime: time.Unix(1434533756, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _contentXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x58\xcd\x6e\x1b\x37\x10\xbe\xf7\x29\x84\x2d\xd0\x1b\x45\x4b\x8e\x01\x7b\x2b\x29\x28\x50\xf8\x64\x5f\x9a\xb4\xe8\x95\xe2\xce\x4a\xac\xc9\xe5\x82\xe4\xea\xe7\x96\x17\x69\x1f\x2e\x4f\xd2\x21\xf7\x47\x5c\x59\xab\x6c\x9a\xa0\xc9\x45\x89\x66\xbe\x19\x7e\x1c\x7e\x1c\x8e\xbc\x78\x7b\x50\x72\xb2\x03\x63\x85\x2e\x96\xc9\x6c\x7a\x93\x4c\xa0\xe0\x3a\x13\xc5\x66\x99\xfc\xfe\xfe\x91\xdc\x27\x6f\x57\x3f\x2c\x74\x9e\x0b\x0e\x69\xa6\x79\xa5\xa0\x70\x84\xeb\xc2\xe1\xbf\x13\x8c\x2e\x6c\x5a\x7b\x97\x49\x65\x8a\x54\x33\x2b\x6c\x5a\x30\x05\x36\x75\x3c\xd5\x25\x14\x6d\x54\x1a\xa3\xd3\xb0\x56\x6d\xb1\xee\x28\x47\x87\x07\x70\x1c\xed\xe0\xe0\xc6\x06\x7b\x6c\x2f\x96\xad\xc7\xaf\x1c\xc0\x71\x74\x66\xd8\x7e\x6c\xb0\xc7\x62\x51\xe3\xf0\x5c\x8f\x0d\x3e\x58\x49\x72\x8d\x55\x57\x25\x73\xe2\x8c\xc5\x41\x8a\xe2\x65\x99\x6c\x9d\x2b\x53\x4a\xf7\xfb\xfd\x74\x7f\x3b\xd5\x66\x43\x67\x0f\x0f\x0f\x34\x78\x3b\xc2\xbc\xc3\x95\x95\x91\x01\x95\x71\x0a\x12\xfc\x62\x96\xce\xa6\x33\xda\x62\x15\x38\x36\x96\x9f\xc7\xc6\x94\x8a\x4a\xad\xc1\x8c\x2e\x0d\x73\xec\xd5\xa9\x96\x06\x2c\x42\x70\xbb\x5e\x98\xe3\x12\xc5\x31\x3d\x7d\xed\x36\xa3\xd5\xb5\xdb\x0c\x94\x99\x6f\x99\x19\xad\xb3\x00\xee\x4b\xe5\x36\x1b\x2f\x95\xdb\x2c\x8e\x55\xcc\x6d\x07\xce\xf7\x9e\x3e\xa3\x33\x7c\x3c\x3f\x9d\x74\x65\xd4\xd8\xb5\x3c\xb6\x57\x2a\x6e\x44\x39\x7a\x9b\x35\x3a\x8e\xd7\x5a\x77\x54\x7d\x40\x7d\xd9\x03\xdd\xf9\xcd\xcd\x1b\x5a\x7f\x8f\xd0\xfb\xab\xf0\xbd\x11\x0e\x4c\x04\xe7\x57\xe1\x9c\x49\xde\x55\x5c\xab\x4b\x45\x43\xdc\x8c\x22\x82\xc0\xce\x4b\xbe\xbb\x44\xbe\x10\x76\x20\x60\x4e\x6b\x77\x07\xb6\xd9\x60\xea\x3f\x9f\x9f\xde\xf1\x2d\x28\x76\x02\x8b\x4f\x83\x89\x28\xac\x63\xc5\xa9\x32\xc6\x1f\xc2\xe0\x4e\xef\xa8\x81\x52\x1b\xd7\x15\x26\x1f\xdf\x7c\xf1\xb4\xe6\x1d\xb7\xad\x53\x72\xb8\x75\x78\x6f\x0b\xdd\x98\x2c\xbb\x08\x45\x3a\xb7\x14\xdb\x08\x5e\x62\xb2\x13\xb0\xff\xb1\xd7\x5b\xaf\xeb\xe1\x81\x06\x50\xdc\x4f\xaf\x06\xcc\x6e\xa8\xc7\x74\x57\x12\x8f\xfb\xd4\xfc\xcd\xa6\x7b\x9f\x72\x5d\x15\x59\xdd\x06\xea\x62\xc0\xa1\x04\x23\xbc\x8b\xc9\x10\x96\xf6\x32\xc4\x0a\x96\xfa\x3f\xa4\x6c\xde\xb4\x28\x43\xaf\xd1\x0b\x90\xed\xed\xef\xb6\x74\x31\x8d\xd6\x44\x59\x14\x03\x8a\x5e\x97\x69\x14\xdd\x7f\x37\x8c\x3a\x8c\x4b\xe7\x85\xae\xb3\xfc\x3c\xe3\xd9\xa5\xe7\xd6\xde\xba\x4b\x67\xfb\xfe\x37\xea\x7d\xc4\x3f\x9b\xf8\x30\x34\x2b\x45\xe3\xc2\x3c\x59\xb5\xb3\x41\xdd\x0b\x2c\xed\x0c\x39\xce\x08\x24\x67\x1c\x48\x06\x5c\xda\xd5\xa2\xee\xf1\x9d\x79\x52\x7f\xf7\xbc\x97\xc9\x93\xc0\x07\x23\x94\x77\xf2\x8e\x15\x78\xd3\xb0\x11\xb7\x50\x25\xe4\x71\x99\xfc\xc4\x4a\x6d\x7f\x3e\xc3\xd5\xc6\x64\xd2\x4b\xed\xf1\x64\x03\x05\x16\x02\xdb\x85\xdd\x0b\x6b\x7b\x88\x52\x38\x8e\x0d\x75\xc7\x8c\x08\xea\xa3\xd7\xa9\xfd\x0a\x7f\xb1\x3f\xaa\xeb\xb4\x22\xcc\x18\x4a\x47\xeb\x40\x7d\x11\x27\xa3\x45\x16\x96\x9b\x3c\x32\x29\xd7\x8c\xbf\x0c\x73\x7b\x8d\xfd\x5f\x38\x3e\x1a\x80\x81\xa2\x45\xae\x2f\x65\x40\x87\xd4\xd6\xd8\x59\xe5\x34\x3e\xa0\x82\x93\x90\xa7\x93\x61\xf8\xec\xf1\xe5\x7a\xd6\x2d\xd6\xf0\x0c\xdd\x09\xe7\x01\x59\xa9\x22\x69\x23\x63\x23\x29\xf1\x56\x81\x71\x02\xec\x24\xd7\xe9\xda\x00\x7b\x21\x6b\xc0\x0b\x86\x09\xfd\xd2\x6d\xc6\x06\xbe\x17\x99\x7f\xcb\xe7\xf3\xe9\xdd\xbd\x52\x81\x7f\x44\x67\x98\x9b\x19\xe2\x66\xf4\xfe\x8c\x18\x5a\x62\x56\xb5\xcb\x1b\xb7\x20\x36\x5b\xbc\xe6\x6f\xa6\x77\x73\x5c\xfb\x3a\xdf\xca\x02\xd1\xa5\x13\x8a\x49\x12\x07\x3b\x53\xc1\x78\xda\x8e\x5d\xa6\xdd\x1a\x15\xce\x7d\x60\x48\xc9\x36\x40\xda\xcb\x96\xb3\x4a\xba\xb3\x3d\x45\xfb\xa9\x87\xf0\x4c\xd8\x52\xb2\x63\xc3\xa7\xc9\xe6\x87\x05\x1c\xb1\x89\xd2\x19\x66\x92\x86\xb8\xf5\x6b\xaa\x74\x50\x17\x8d\x63\xad\xb3\xe3\xa9\xa9\xe1\x48\xc9\x32\xbb\x05\x70\xab\x45\xbd\x74\xf8\x6c\x68\xd4\x9c\x9f\xab\xe2\x85\x49\x56\xe2\x5e\x6b\x73\xc8\x48\x4e\x25\xe8\x85\x36\xca\xb9\x00\x0d\x02\x6c\xf6\x57\x97\x81\x70\x90\x92\xc4\x98\xb6\x3e\xb4\x9f\x13\x8f\xe8\x42\x42\xaf\x9a\xb3\xb5\x31\x5f\xd7\xc6\x99\xac\x80\xb8\x63\x89\x48\xeb\x0c\x56\x2e\x99\xb4\xef\xe1\x05\x1f\x26\xf2\x9e\x72\xf5\xcb\x82\x36\xff\x5b\xd0\xf3\xdc\x67\x26\xa4\xf5\x75\x79\xe6\x52\x33\x97\xf4\x3c\xf8\x0c\x5d\xe6\x5d\x63\x3b\xda\xb3\x6f\x48\xfb\x73\xca\xfb\xf1\xc3\x3f\xf3\x8f\x1f\xfe\xfe\x4c\xb6\xb1\xe5\xd3\x4a\x9d\x7f\x8f\x4a\x6d\x99\x86\xdf\x8e\x1e\x67\x09\x4e\xb9\xc0\x1c\xe0\xec\x34\xbf\x50\x69\xfa\x3d\x1d\x9b\xb5\xf6\x2b\x1c\x99\x67\x98\x11\x1c\xe2\xf0\x97\xac\x9f\xb1\x6c\xf4\xc2\xf5\x7a\x11\xed\x75\x2b\x3a\xf0\xf7\x99\xd5\xbf\x01\x00\x00\xff\xff\x1f\x8d\x19\xe0\xe0\x11\x00\x00")

func contentXmlBytes() ([]byte, error) {
	return bindataRead(
		_contentXml,
		"content.xml",
	)
}

func contentXml() (*asset, error) {
	bytes, err := contentXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "content.xml", size: 4576, mode: os.FileMode(436), modTime: time.Unix(1434526146, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _manifestRdf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x93\xcd\x6e\xc2\x30\x10\x84\xef\x3c\x85\x65\xce\xb1\x21\xbd\x94\x88\xc0\xa1\x88\x73\xd5\x3e\x81\xeb\x6c\xc0\x6a\xe2\x8d\xbc\x9b\x86\xbc\x7d\x8d\xa1\x15\x42\xaa\xaa\xfe\x49\x3d\x7a\x35\x9a\xf9\x66\x24\x2f\xd7\x87\xb6\x11\x2f\x10\xc8\xa1\x2f\xe5\x5c\xcd\xa4\x00\x6f\xb1\x72\x7e\x57\xca\x9e\xeb\xec\x56\xae\x57\x93\x65\xa8\xea\xe2\x61\xb3\x15\x51\xed\xa9\x88\xaf\x52\xee\x99\xbb\x42\xeb\x61\x18\xd4\x70\xa3\x30\xec\xf4\x7c\xb1\x58\xe8\x59\xae\xf3\x3c\x8b\x8a\x8c\x46\xcf\xe6\x90\x79\x9a\xca\xd5\x44\x88\xe4\xb1\x01\xb2\xc1\x75\x1c\xd3\xc4\xf1\x6d\x9e\xb0\xe7\x52\x12\x8f\x0d\x90\x8a\xee\x49\x7a\x16\xf3\xd8\x41\x52\x05\x20\xec\x83\x85\xf7\xd0\x0a\x2d\x29\x34\xe4\x28\xc3\x0e\x7c\x4a\xf7\xa4\xb1\xae\x9d\x05\x3d\x57\xb9\x6e\x81\x8d\xc6\xaa\x9e\x3e\x26\xeb\xad\x6b\x40\xea\x84\xa1\xaf\x38\x3e\x63\x7b\x23\xf2\x34\x2b\xf6\x86\xee\x4d\xe0\xf3\x0c\xf1\xf2\x55\xa2\xee\x79\x37\x95\x57\x9d\x2e\xca\x7f\x8f\xd0\xa2\x67\xf0\xfc\x27\xf3\xdd\x9d\xbc\xff\xf3\x7e\x97\xf5\x7f\x86\xf8\x2b\xab\x1d\x11\x37\x68\xfb\x36\x32\x7d\xc8\x73\x3a\xc5\x1f\xb5\x9a\xbc\x06\x00\x00\xff\xff\xb4\xf7\x68\xd2\x83\x03\x00\x00")

func manifestRdfBytes() ([]byte, error) {
	return bindataRead(
		_manifestRdf,
		"manifest.rdf",
	)
}

func manifestRdf() (*asset, error) {
	bytes, err := manifestRdfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifest.rdf", size: 899, mode: os.FileMode(436), modTime: time.Unix(1434526146, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _metaXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x93\x4f\x8f\xdb\x20\x14\xc4\xef\xfd\x14\x16\xdd\xab\x8d\x21\x94\x8d\x91\xed\x95\x7a\xe8\x69\xa5\x56\x6a\x2a\xf5\x16\x39\xf0\x36\xa5\xc5\x10\x01\x5e\xa7\xdf\xbe\xfe\x9b\x26\x6d\x0e\x7b\xcc\xf0\x9b\x37\xc3\x23\x2e\x9f\xce\xad\x49\x5e\xc1\x07\xed\x6c\x85\x48\x96\xa3\x04\xac\x74\x4a\xdb\x63\x85\xbe\xed\x3e\xa5\x5b\xf4\x54\xbf\x2b\xdd\xcb\x8b\x96\x20\x94\x93\x5d\x0b\x36\xa6\x2d\xc4\x26\x19\xac\x36\x88\xf9\xa8\x42\x9d\xb7\xc2\x35\x41\x07\x61\x9b\x16\x82\x88\x52\xb8\x13\xd8\xd5\x22\xae\x69\x31\x05\xcd\xca\xd9\x68\xfb\xab\x42\x3f\x62\x3c\x09\x8c\xfb\xbe\xcf\xfa\x4d\xe6\xfc\x11\x93\xa2\x28\xf0\x74\xba\xa2\x4a\x5e\xb8\x53\xe7\xcd\x44\x29\x89\xc1\xc0\x98\x10\x30\xc9\x08\x5e\xd9\xb1\xe1\x5b\x4b\x8d\xec\x75\x25\xe7\xdc\x25\x68\xc4\xe7\xd2\x53\x1c\xcd\x73\x86\xe7\xdf\x2b\x7d\xf4\x4a\x99\x7b\x17\x18\xd8\x0d\x1e\x1a\x36\xb1\x49\x5f\x35\xf4\xef\x51\xb2\x5c\xff\x6a\xe1\x14\xd5\xeb\x76\xc7\x1a\x75\x39\x95\x91\x1e\x9a\x38\x10\xe9\x60\x86\x9a\xe6\xe4\x43\x9a\xf3\x94\x3c\xee\x08\x11\x74\x2b\x28\xcf\x0a\xca\x39\x21\x9c\xb1\x12\xdf\x71\x94\x4a\x8a\x7b\xd6\x42\xe4\x3c\xa3\x8f\x74\xc3\xe9\x96\xd1\x12\xaf\xd8\x9c\x0a\x4a\xc7\xe1\xe1\x53\xd5\xf9\x69\x56\xfd\x65\xb7\x29\xbe\x2e\x01\xff\x1d\xde\x7a\xe4\x6f\x69\x20\xd4\xe4\x1f\x7a\x91\x67\xf6\xf2\xf7\x09\x71\x18\x11\xa2\x96\xc9\xa4\xc7\xe6\x60\x20\x95\xae\xb3\xb1\x42\x14\xcd\xa2\x04\x63\x56\x8d\x2d\x9a\x3b\xfc\x04\x19\x57\x35\x47\x78\x19\x7c\x04\x0b\x43\x2b\xe7\xeb\x67\x7d\xf0\xf0\x79\xda\x27\x66\x19\xcb\x68\x46\x1f\x9e\xb5\xed\xce\xfb\xef\x5b\xbe\xe7\x2c\xb9\x02\xf6\x27\xef\xc6\x79\x98\xe5\x6d\xfe\xf0\xb1\xd3\x46\xa5\x74\xe9\xff\x77\x62\x89\x6f\x9e\x07\xdf\xfb\x14\xea\x3f\x01\x00\x00\xff\xff\x76\xae\x40\x34\x48\x03\x00\x00")

func metaXmlBytes() ([]byte, error) {
	return bindataRead(
		_metaXml,
		"meta.xml",
	)
}

func metaXml() (*asset, error) {
	bytes, err := metaXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "meta.xml", size: 840, mode: os.FileMode(436), modTime: time.Unix(1434526146, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _mimetype = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x04\xc0\x81\x09\x80\x00\x08\x04\xc0\x8d\x6c\x26\xd1\x87\x84\xd2\xa7\xb7\xe6\xef\x9c\xbc\x2a\x7c\x6b\xfa\xf8\x3a\x6d\x5c\x25\x1b\xa2\x73\xe2\xbd\xd1\x6b\xe2\x03\x4f\x9d\xc0\xfe\x01\x00\x00\xff\xff\x85\x6c\x39\x8a\x2e\x00\x00\x00")

func mimetypeBytes() ([]byte, error) {
	return bindataRead(
		_mimetype,
		"mimetype",
	)
}

func mimetype() (*asset, error) {
	bytes, err := mimetypeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mimetype", size: 46, mode: os.FileMode(436), modTime: time.Unix(1434526146, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _stylesXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x59\x5d\x8f\xd4\x36\x17\xbe\x7f\x7f\x45\x94\x95\x10\x48\x6f\x26\x99\x59\xbe\x66\xd8\x99\xbd\x28\x5a\x51\x09\x68\x45\xa1\xf7\xde\xc4\xc9\xb8\x38\x71\x64\x3b\x33\x3b\xfc\xfa\x1e\xdb\xb1\x27\xdf\x1b\x96\x42\x2b\xb6\xab\x42\xec\xe7\x1c\x3f\x3e\xdf\x09\x57\xd7\x77\x39\xf5\x0e\x98\x0b\xc2\x8a\xad\xbf\x5c\x44\xbe\x87\x8b\x98\x25\xa4\xc8\xb6\xfe\xa7\x8f\x37\xc1\x4b\xff\x7a\xf7\xbf\x2b\x96\xa6\x24\xc6\x9b\x84\xc5\x55\x8e\x0b\x19\x08\x79\xa2\x58\x78\x20\x5c\x88\x8d\xd9\xdc\xfa\x15\x2f\x36\x0c\x09\x22\x36\x05\xca\xb1\xd8\xc8\x78\xc3\x4a\x5c\x58\xa1\x4d\x13\xbd\xd1\x47\x99\x15\xad\x6c\xae\xb8\x06\x37\xa5\x25\xbe\x93\x73\x85\x15\xb6\x25\x8b\x6e\xe7\x9f\xac\xc1\x4d\xe9\x84\xa3\xe3\x5c\x61\x85\x05\x9b\x36\xc5\x53\x36\x57\xf8\x4e\xd0\x20\x65\x41\xcc\xf2\x12\x49\xd2\x61\x71\x47\x49\xf1\x79\xeb\xef\xa5\x2c\x37\x61\x78\x3c\x1e\x17\xc7\xcb\x05\xe3\x59\xb8\x5c\xaf\xd7\xa1\xde\x75\x84\x63\x87\x2b\x2b\x4e\x35\x2a\x89\x43\x4c\xb1\x3a\x4c\x84\xcb\xc5\x32\xb4\xd8\x1c\x4b\x34\x97\x9f\xc2\x36\x29\x15\x55\x7e\x8b\xf9\x6c\xd3\x20\x89\x7a\x5e\x2d\x39\x16\x00\x81\xeb\xaa\xb8\x9c\xa7\xa8\x29\xd3\x8a\xaf\x43\x36\x3b\xba\x0e\xd9\x88\x99\xe3\x3d\xe2\xb3\xe3\x4c\x83\xdb\xa1\x72\x99\xcc\x0f\x95\xcb\xa4\x29\x9b\x23\xb9\x1f\xf1\xef\xcb\xf0\x1d\x6c\xea\xff\xbd\x7b\x7b\x8e\x2b\x9e\xcf\x3d\x4b\x61\x5b\xa6\x8a\x39\x29\x67\x5f\xd3\xa0\x9b\xf2\x8c\x31\x47\x55\x09\x98\x64\xd7\x74\x57\x51\xf4\x34\x34\xcf\x0d\xf4\x71\x12\x7e\xe4\x44\x62\xde\x80\xc7\x93\xf0\x18\xd1\xd8\x59\x9c\xe5\x43\x46\x03\xdc\x32\x04\x44\x80\x0f\x2a\xe4\x2d\x9a\xab\x4b\x8f\x6a\x7e\x16\x72\x5c\x32\x2e\x1d\x91\x74\x7e\xb1\x03\xeb\xac\x5c\xaa\xee\x65\x4e\xc7\x53\x55\xed\x5a\x68\xc6\x93\x64\x10\x0a\x74\x2e\x43\x48\x5b\x48\x9a\xe0\x40\xf0\xf1\xa2\x55\xcb\xa6\xed\xbf\x0e\x35\xa8\x59\xbf\x26\x05\x96\x51\xa8\x30\x2e\x05\xc0\xbc\xe7\x62\xcb\x33\xd7\x0e\x52\x56\x15\x89\x49\x3b\x63\x0c\x7c\x57\x62\x4e\xd4\x16\xa2\x5a\x6c\xd3\xd2\xd0\x8c\x18\xca\x1e\xa0\xb2\xee\x21\x0d\x0d\xad\x54\x15\xe2\x52\x0e\xd9\xee\xe3\x87\x50\xed\x05\xaa\x0d\x40\xa1\xab\xb5\x34\xba\xdf\xca\xdf\xd9\x56\x97\x32\x68\x73\x29\x8a\x71\x90\xe0\x98\x8a\xdd\x95\x29\x51\x6e\xd9\x33\xcf\x8a\xdc\xd6\x7f\x4b\xa0\xde\x69\xb6\xde\x1f\xa8\x80\xa8\x82\x3a\x62\xa1\x39\xa1\xa7\xad\xff\x08\x95\x4c\xbc\xea\xe0\xcc\xa2\xef\xb5\x54\x2b\x7c\x90\xe1\x02\x6e\x0b\xd1\x2e\x8e\x44\x88\x16\xa2\x24\x32\x86\x7a\x70\x40\x9c\x68\x67\x86\xd3\xd4\x5e\xe3\xbf\xd0\x9f\xd5\x34\xad\x06\x66\x0e\xa5\x93\x90\x38\xff\x26\x4e\x9c\x91\x44\x1f\xe7\xdd\x20\x4a\x6f\x51\xfc\x79\x9c\x5b\x1f\xfb\x43\x38\xde\x70\x8c\x47\x8c\xd6\xd8\xfa\x56\x06\xe1\x58\xb4\xd5\xeb\x66\xce\xb2\x4c\x13\x9c\xa2\x8a\xd6\xd3\x97\xd5\x5c\x93\xd2\x99\x1d\xc4\x98\x52\xdf\xc2\x4b\xc4\x51\xc6\x51\xb9\x0f\x4a\x0e\xb9\xcd\x25\x81\x91\xcd\x6c\x01\x1a\xb4\xb0\x32\x48\x88\x90\xa8\x50\xd3\xdb\x72\xb5\x78\x96\xe7\x67\xb3\xa8\x24\xe9\x0b\x6a\xa2\x23\x41\x9f\xb2\x0d\x45\x45\x56\xa1\x0c\x76\xf7\x95\x5e\x88\x21\x8f\x25\x07\x7e\x6f\x3e\xf9\x5d\x15\x01\x94\x4f\x54\x74\x43\x54\x63\xac\x1e\x0b\xf9\xb2\xb7\x3b\xb5\x42\xbb\xf1\xcb\xfb\xbe\x5a\xd5\xc0\x29\xbe\x9b\x56\xec\x40\x7b\xd2\x55\xed\xb6\x7e\x7d\xaf\x7d\x34\x60\xfc\xdd\x95\x99\x71\xea\x51\xa7\xe5\x11\x63\x9d\xf7\x91\xdf\x01\x79\xf5\x53\x4e\x8a\x80\x14\x12\x67\x20\x97\x90\x8c\x48\x01\xc6\xd7\x07\x0d\xe8\x74\x3a\xe2\x8a\x73\x98\xcc\x4f\x43\x47\x2d\xa3\xa7\xbf\x47\xf6\x16\x07\x46\xc1\x29\x6a\xac\x95\xbc\xc2\x63\x24\x20\xcc\x48\x8e\x68\x50\x52\x08\x3a\x20\x00\x3d\x6a\x92\x9e\xdd\xcd\x38\xab\x4a\xfd\x72\xa0\xb5\x87\x4e\xbd\x0a\x97\x9d\xe7\xee\xa0\x1f\xfb\xdc\x4f\xf9\x2d\xa3\x56\x57\x3b\x58\x2c\xb6\x11\x30\xbb\x1b\xe9\x14\x76\x74\xec\x06\x36\xbe\xc6\x5e\xfe\x58\x98\xeb\x98\xa5\x0c\xa6\xd7\x8b\x34\x8d\xe0\xa7\x7b\xc7\x60\xf8\x8e\x3f\x9b\x79\x8d\x75\x72\x54\xba\xe4\x28\x12\x62\xa6\xf1\x03\xa2\x15\x7e\xfc\xe4\x51\x26\x5f\x6d\x5d\xd8\xa1\xb2\xa4\xb5\xb1\x83\x56\x58\x86\xe3\xae\x32\x92\x7d\x07\xbd\x36\xa9\xe6\xdf\x5f\xe2\xee\x29\x52\xae\xc2\x0c\x36\x9c\x5e\xf1\xae\xd1\x0f\x6f\x3c\x56\xc3\x58\xf1\xb7\xfb\xae\x05\x8c\xd7\xae\xa9\x26\x33\x0b\x64\x29\x39\xf0\x28\x29\x87\x68\x75\xa6\x86\x73\xc6\x3d\xf5\x01\x8b\x69\x47\x79\xae\x15\xb9\xcf\x06\x41\xdb\xc9\x13\x89\xa8\x39\xd6\x5f\x07\x08\xcc\x7d\x24\xb6\xfa\x34\x18\x26\x44\xcc\xe1\xfd\x16\x5b\x8c\x60\x94\x24\x23\x90\x23\x49\xd4\x6b\x14\xaa\x24\x1b\x41\xd4\x69\xaf\x0f\xd5\x7f\xf7\x1d\x87\x23\x26\xd9\x1e\x46\x4a\xc8\x8c\xe4\x6b\x8d\xb3\x7a\x90\x75\xda\x86\xd5\xd3\x7e\x37\xb7\xce\xad\xba\x7f\xf4\x1b\x8c\xd4\x27\x9c\x7f\xc6\x31\x4e\x68\x60\x8e\x50\x26\x04\xc7\x64\x45\x20\x58\xc5\xd5\x1c\x91\x92\x3b\xab\x1c\xde\x9a\x30\x52\xc6\x84\x7a\x57\x80\xf9\x52\x44\x45\x63\xf0\x1a\x9c\x4f\xc0\xe4\x67\xa5\x5b\x3f\x06\x41\x78\x03\x1c\x9d\x4a\x5c\x94\x90\x2f\x6a\x86\x79\x5e\x4a\x7f\x2c\x72\xbe\xcd\x99\xb5\x45\x97\x0f\x32\xa9\x75\xc7\x1c\x93\x72\x66\xbe\x60\x04\x50\xbe\x15\xfd\x75\xd4\x27\x19\x76\xc6\xc3\xfa\x51\x05\x77\x0e\xc2\x71\xd0\x9e\x1b\x4b\x35\xf0\x50\x74\x62\x95\x6c\xdd\xe9\x5d\x99\x2f\xfd\x01\x50\x9f\x94\x7a\x11\x87\x0b\x04\x39\x4b\x40\x8e\xf2\x40\xde\x9e\x7d\xb2\x87\xdb\x9d\xe7\x95\xd6\x5a\xca\x18\xb8\xaf\xe3\x30\xd5\x02\xf7\xb5\x13\x5e\xe8\xa1\x53\xaf\x22\x9e\xc1\x06\xc5\x29\x2c\x47\xed\x45\x6e\xc0\x9d\xd5\x5b\x26\xa5\x7a\xcb\x77\x83\x6b\x38\xce\xa7\x26\xf2\xc3\x39\xc2\x8c\xdd\x27\xd8\x26\x13\xf6\xec\x3f\xc3\x6f\xab\x9f\xc6\x6f\x6a\xe3\x96\xf1\x44\x7d\x30\x5c\x2d\x9e\xae\x4b\xb8\xab\xaa\xe3\xde\x45\xa4\x7f\x34\xa0\x44\x89\xf9\x24\x1d\x2d\x96\x2f\xad\x10\x74\x64\x35\x2f\x15\x89\x2d\xdd\x17\x71\xa4\xfe\xfb\xef\x87\xc2\x77\xbd\xf3\xfd\xd1\x15\x8e\xd6\x8b\x7a\x23\x47\xc2\xa9\x10\xe7\x59\x50\x2f\x2a\x4d\x53\xd3\x5a\x33\x1e\x07\xca\x8c\xb1\xf4\xee\x4a\x7f\x8a\x2f\xeb\x3f\xc5\x1e\x63\x83\xde\x5d\x5f\x5f\x5f\x85\xdd\xc5\x7a\xa5\xec\xf8\xb5\xe3\x3d\xe5\x0b\xdb\x30\x89\x80\xd1\xfb\xd4\x6b\x39\xc6\x36\xee\xf4\xdf\x68\x82\xa8\x57\x3f\x29\xde\x66\x50\xdd\x2d\xed\x81\x8d\xb5\x1e\x07\xab\xab\x65\xf5\x49\x0e\x61\xcf\x90\xf7\xd9\xf6\x43\xfd\xdd\x71\xc2\xb4\xab\x9e\x69\x6d\xf3\xcd\x54\x17\x51\x84\xbe\xd2\xda\xde\x63\x83\x93\x44\xd2\x26\xc4\x3c\x3f\xe9\x19\xa2\x75\x52\x6b\x49\x27\x42\xe7\x74\x98\x65\xf0\xe8\x58\x03\xc3\x92\x03\x05\xfa\x3d\x03\x12\x24\x5a\x3e\x0b\xa2\xe7\xc1\xf2\x85\xbf\x53\xc9\xb1\xd0\xbf\x35\x0b\x05\xdc\xfd\xdf\xb3\x84\x81\x7d\x14\x6d\xf4\xaf\x23\x3d\x14\x3f\x6d\x7e\xff\x5e\x50\x79\x61\x13\xa8\x5f\xce\x76\xeb\x75\x13\x68\xd6\xbe\x53\xf0\x85\xc3\xf9\x1e\x0e\xff\x73\xdf\xee\xef\x00\x00\x00\xff\xff\xc8\x80\x14\x7c\x2e\x1c\x00\x00")

func stylesXmlBytes() ([]byte, error) {
	return bindataRead(
		_stylesXml,
		"styles.xml",
	)
}

func stylesXml() (*asset, error) {
	bytes, err := stylesXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "styles.xml", size: 7214, mode: os.FileMode(436), modTime: time.Unix(1434526146, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if (err != nil) {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"META-INF/manifest.xml": metaInfManifestXml,
	"content.xml": contentXml,
	"manifest.rdf": manifestRdf,
	"meta.xml": metaXml,
	"mimetype": mimetype,
	"styles.xml": stylesXml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"META-INF": &bintree{nil, map[string]*bintree{
		"manifest.xml": &bintree{metaInfManifestXml, map[string]*bintree{
		}},
	}},
	"content.xml": &bintree{contentXml, map[string]*bintree{
	}},
	"manifest.rdf": &bintree{manifestRdf, map[string]*bintree{
	}},
	"meta.xml": &bintree{metaXml, map[string]*bintree{
	}},
	"mimetype": &bintree{mimetype, map[string]*bintree{
	}},
	"styles.xml": &bintree{stylesXml, map[string]*bintree{
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
        data, err := Asset(name)
        if err != nil {
                return err
        }
        info, err := AssetInfo(name)
        if err != nil {
                return err
        }
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
        if err != nil {
                return err
        }
        err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
        if err != nil {
                return err
        }
        err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
        if err != nil {
                return err
        }
        return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        // File
        if err != nil {
                return RestoreAsset(dir, name)
        }
        // Dir
        for _, child := range children {
                err = RestoreAssets(dir, path.Join(name, child))
                if err != nil {
                        return err
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

