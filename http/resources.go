// Code generated by go-bindata.
// sources:
// http/resources/combo.css
// http/resources/combo.html
// http/resources/combo.js
// DO NOT EDIT!

package http

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

var _comboCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x52\xc1\x6e\xdb\x30\x0c\xbd\xfb\x2b\x08\x14\xbb\x04\x50\x9a\x6c\xc8\x61\xde\x69\x28\x7a\x28\xb0\xec\xb2\xed\x03\x64\x89\xb1\x89\xc8\x52\x20\xc9\x4e\xd2\xa2\xff\x3e\xd2\x76\xba\x74\xab\x37\xec\x64\x99\x8f\x4f\x7c\xef\x89\xb7\x8b\x02\x16\x70\x17\x0e\xe7\x48\x75\x93\xe1\xfd\x6a\xbd\x81\x6d\x47\x11\xb6\xda\x5b\x8c\x69\x09\xf0\xd9\x39\x18\xd0\x04\x11\x13\xc6\x1e\xed\x52\x58\x3f\x12\x42\xd8\x41\x6e\x28\x41\x0a\x5d\x34\x08\x26\x58\x04\xfe\xad\x43\x8f\xd1\xa3\x85\xea\x0c\x1a\xb6\x0f\xdf\x55\xca\x67\x87\xc2\x72\x64\xd0\x33\x33\x37\x3a\x83\xd1\x1e\x2a\x84\x5d\xe8\xbc\x05\xf2\x5c\x44\xf8\xf2\x70\x77\xff\xf5\xdb\x3d\xec\xc8\xa1\xcc\xb9\x2d\x16\xf0\x54\x00\x54\xe1\xa4\x12\x3d\x92\xaf\x4b\x3e\x47\x16\xa7\xb8\xf4\x89\x11\x75\xc4\x6a\x4f\x59\xe5\xd0\x99\x46\x19\xed\x5c\xe8\x72\x09\x3e\x78\xbc\x86\x3b\xd6\xae\x12\x3a\x34\xd7\xe0\xbe\xc9\xad\x9b\xc1\xda\xf0\x38\x87\xa4\xb7\x81\x37\x8a\xcf\x45\x71\x53\x05\x1d\xed\xe0\xe2\x48\x36\x37\x25\x7c\xdc\xf4\x2d\x79\x61\x34\x28\xd1\x5e\x57\x46\x73\x25\xac\x0f\x27\xce\xd5\x11\xa7\xe8\xb4\xd9\x0b\x64\x29\x1d\x9c\x3e\x97\xb0\x73\x78\x7a\x55\xb8\x98\xbc\x00\xf2\x55\x96\x22\x0b\xa1\xe0\x4b\x7e\x19\xd7\xb5\xfe\x3a\x8e\xb9\x0e\xd6\xbb\x8c\xe1\x38\xa8\x95\x1e\x16\xf2\x3b\x6d\x2a\x4d\x5e\xd6\xab\xd5\xbb\xff\x11\x27\x03\x0c\xf2\x52\xfd\x7b\xc2\x5f\xa2\xa8\xf8\x5b\x47\x59\x1c\xc5\xca\x03\x37\x99\x26\xf0\x49\xe7\xe1\x21\x76\xc1\x67\x59\x17\x14\x7d\x97\x64\x33\x9e\xb2\xd2\x8e\x6a\xb1\x8b\x3e\x63\x1c\xd5\x1c\x08\x79\x79\x9f\x5e\x06\xaa\xa8\x2d\x75\xa9\x84\x8d\x38\x93\x8e\x63\x43\x79\xea\xf8\x63\xee\x80\xc9\xf5\xd3\xff\x24\x51\x68\xc3\x71\x86\xf6\xe2\xe4\xf5\x35\x42\xeb\x59\xa3\xe5\xed\xeb\xe7\x46\xd6\x11\x71\x7c\xaa\x9b\x16\x53\xd2\x35\xa6\x31\xcd\x5f\xae\x3f\x8c\xa6\x9f\x8b\x9f\x01\x00\x00\xff\xff\xeb\x56\xad\x73\xe3\x03\x00\x00")

func comboCssBytes() ([]byte, error) {
	return bindataRead(
		_comboCss,
		"combo.css",
	)
}

func comboCss() (*asset, error) {
	bytes, err := comboCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "combo.css", size: 995, mode: os.FileMode(420), modTime: time.Unix(1434758159, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _comboHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x91\xbd\x8e\xdb\x30\x10\x84\x7b\x3d\xc5\x46\xbd\xc8\xc8\x48\x9a\x80\x32\x10\x18\x2e\x0c\xc4\x69\x92\x3c\x00\x45\xae\xa4\x75\x28\x52\xe1\x52\x86\xf5\xf6\xa1\xa4\xc3\x1d\x0e\x57\x5c\x71\x15\x31\x33\x3b\x1f\xff\xd4\xa7\xaa\x2a\x4e\x61\x5a\x22\xf5\x43\x82\xc3\xe7\xfa\x2b\x5c\x67\x8a\x70\xd5\xde\x62\x64\x01\xf0\xdd\x39\xd8\x52\x86\x88\x8c\xf1\x8e\x56\x14\x7f\x18\x21\x74\x90\x06\x62\xe0\x30\x47\x83\x60\x82\x45\xc8\xb2\x0f\x77\x8c\x1e\x2d\xb4\x0b\x68\xb8\x5e\x7e\x57\x9c\x16\x87\x85\x23\x83\x3e\xd7\xd2\xa0\x13\x18\xed\xa1\x45\xe8\xc2\xec\x2d\x90\xcf\x26\xc2\x8f\xcb\xe9\xfc\xf3\xd7\x19\x3a\x72\x28\x8a\xaa\x3a\x16\x6a\x48\xa3\x5b\x17\xd4\xf6\x58\x00\x28\x47\xfe\x6f\x3e\x85\x6b\xca\x8d\xc9\x03\x62\x2a\x21\x2d\x13\x36\x65\xc2\x47\x92\x86\xb9\x84\x21\x62\xd7\x94\x26\x8c\x6d\x10\xab\xb1\x55\xd9\x44\x9a\x12\x70\x34\x4d\x39\xa4\x34\xf1\x37\x29\xf5\x4d\x3f\x44\x1f\x42\xef\x50\x4f\xc4\x22\x57\x36\x4f\x3a\x6a\x59\xde\xfe\xcd\x18\x17\x79\x10\xb5\xf8\xf2\x24\xc4\x48\x5e\xdc\x32\x51\xc9\x9d\xf7\x21\xf4\x4c\xb2\x16\xf5\x0b\xbd\x9a\xe9\xfd\x0d\xe4\x7e\xaf\x37\x33\x89\x92\xc3\xe3\x69\x0d\x95\xdc\x45\xa1\xe4\xfe\x72\xaa\x0d\x76\xd9\xa6\x2c\xdd\x81\x6c\x53\xb6\x41\x47\xbb\x12\xb2\xf1\x2a\x18\x91\x59\xf7\xc8\xcf\x99\x92\x7b\x39\xb3\xb6\xcf\xf8\x1f\x00\x00\xff\xff\x85\x2a\xb4\x5f\x32\x02\x00\x00")

func comboHtmlBytes() ([]byte, error) {
	return bindataRead(
		_comboHtml,
		"combo.html",
	)
}

func comboHtml() (*asset, error) {
	bytes, err := comboHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "combo.html", size: 562, mode: os.FileMode(420), modTime: time.Unix(1434758159, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _comboJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x57\x4d\x8f\xdb\x36\x13\xbe\xfb\x57\x10\xca\x1e\x64\xac\x97\x9b\xbc\x40\xde\x16\x76\xd2\xa2\x58\xe4\x90\xa2\x49\x0f\x49\xd1\x83\xb1\x30\x68\x89\xb6\x89\xc8\xa2\x4a\x52\xb6\x85\xc4\xff\xbd\x33\x43\x49\xa4\xbc\x72\x9a\x14\xbd\xac\xc5\xf9\xe2\x70\xe6\x99\x8f\xbd\xbf\x67\x0f\xba\x6a\x8c\xda\xee\x1c\xfb\xdf\xf3\x17\x2f\xd9\xbb\x5a\x19\xf6\x4e\x94\xb9\x34\x96\x33\xf6\x4b\x51\x30\xe2\x5a\x66\xa4\x95\xe6\x20\x73\x3e\x01\xad\x3f\xac\x64\x7a\xc3\xdc\x4e\x59\x66\x75\x6d\x32\xc9\x32\x9d\x4b\x06\xc7\xad\x3e\x48\x53\xca\x9c\xad\x1b\x26\xd8\xbb\xb7\x1f\xef\xac\x6b\x0a\x89\x5a\x85\xca\x64\x09\x9a\x6e\x27\x1c\xcb\x44\xc9\xd6\x92\x6d\x74\x5d\xe6\x4c\x95\x40\x94\xec\xb7\xb7\x0f\x6f\xde\x7f\x78\xc3\x36\xaa\x90\x7c\x32\x39\x08\x03\x66\xf7\x6b\xcd\x5e\xb3\xcf\x13\xc6\xf6\x60\xda\xce\xd9\xe7\xf3\x0c\x0e\xb6\x2a\x94\xa3\xd3\xe4\xbc\x98\x4c\x48\x8e\xab\x52\xb9\xd5\x5a\x0b\x93\x83\xca\xa6\x2e\x33\xa7\x74\x99\x1e\x55\xee\x76\x33\xb6\x93\xf8\x92\x29\x99\x22\xd3\xb2\x28\x2c\xc8\x2d\x1f\x41\x9f\xd1\x63\x38\x89\x02\x8d\x7e\x17\x1d\xd5\x6b\x02\xd9\x7f\x90\xf8\x46\x1b\x96\xa2\x99\x06\xe8\xcf\x17\xf0\xf3\xaa\x63\xb3\xe6\xf6\xd6\x5f\xc3\xfc\x25\xcb\xe6\xb1\xbd\x07\x49\xa8\x64\xf4\x11\x28\x37\x69\xf2\x2a\x57\x07\x96\x15\xc2\xda\xd7\x40\xfb\x29\x99\x7a\x99\xde\xfa\xc9\x5b\x3f\x81\x75\xef\x13\x3b\x05\xe3\xc1\xfc\xf2\xf4\xf8\xc4\x1e\xf2\xc0\x20\x17\xce\x99\x34\xc9\x85\x13\x77\xa7\x64\xc6\x4e\x03\x4a\x03\x94\x06\x28\x55\x25\xcb\xfc\xa3\x4e\xc1\x89\x29\xcf\x8d\xae\x2a\xb1\x2e\x64\xda\xdd\x03\x81\xd0\x85\x34\xa2\xcc\xe4\x9c\x25\xaa\x74\x80\x0f\x99\xb9\x64\xd6\xf3\x45\x96\xc9\xca\x01\x93\x3b\x7d\x97\xdc\x9e\x6e\x13\xf8\xdb\xc4\x7c\xa7\x0e\xf2\x01\x3d\x03\xa1\x83\x28\x54\x7e\x87\x09\x8d\x4c\xe0\xb5\xf3\x90\x36\x79\x90\xa5\x9b\xb1\x5a\x85\xe7\xd2\x93\x29\xd3\x47\xcb\x2d\x78\x9c\xfe\xfa\xe1\xf7\xf7\xdc\x3a\xa3\xca\xad\xda\x34\x69\x2c\x48\xa2\x7b\xc0\x32\xdc\x77\x71\x13\x39\x64\xb6\x08\x9f\x01\x0d\x02\x6f\xf4\x1e\xa8\xa7\x39\xbb\xad\x15\x04\x42\x6c\xb7\x18\x88\x61\x10\xa7\x10\xb3\xaf\x08\x34\xc9\xf4\x3c\xbb\xb0\xeb\x74\x6b\xf5\x26\x45\x48\x4d\x47\x0d\x8e\xf0\xc6\x6c\x11\xf4\xe7\x6d\x20\xe8\xe0\xe0\xf9\x03\xa1\x73\x74\x3a\x4f\x5b\x50\xc5\x8c\x73\x4b\xf3\x67\xc8\x7a\x80\x40\xf2\x8c\x2a\xc8\x43\xf1\xdc\x97\x46\x57\x2e\xf4\xbb\x88\x8a\x2e\x57\xe0\x83\x68\x56\x7b\x69\xad\xd8\xca\xb8\xf2\xf6\x76\xeb\x93\x07\xc8\x7c\xd6\xf2\x2d\x40\xd2\xc9\x93\x23\x66\x6c\x47\xc3\xfd\xab\xa3\x8d\xf5\x43\xb1\x76\x85\xf9\x63\x57\xc7\xf8\x4d\x85\x48\x5c\xd4\xf2\x05\x8c\x5f\xa5\x3c\xb2\x3f\xe5\xfa\x83\xce\x3e\x49\x97\x26\x47\x3b\xbf\xbf\x4f\xd8\x2d\x2b\x74\x26\xd0\x2e\xdf\x69\xeb\xe0\x9c\xdc\x67\xba\x2c\x11\xc7\x53\x32\x05\xa0\xd2\xe5\xc8\x33\x08\x89\x1d\x0a\xa9\x75\xec\xb1\xc1\x10\xf6\x2a\x01\x95\xe0\x25\x38\x66\xac\x8d\xab\x3d\x2a\x97\xed\x58\x0a\x92\xbc\x85\x61\xdf\x12\x04\xf4\x40\x8f\xc8\xf9\x24\xc6\x34\x92\x48\x01\xa1\xd9\xe7\x6c\x6d\xa4\xf8\xb4\x88\x34\xb7\x62\x2f\x57\xd8\x65\x2f\xd4\x2f\xf2\xd0\x5b\xe2\x2d\xa1\xb7\x08\xc9\xe0\x95\x92\x99\x4c\xa6\x01\xc0\xe9\x67\xd0\xc7\x0f\x28\x17\x67\x6a\x79\x1e\x73\x00\xd1\x72\x0e\xa1\xc2\x8c\x3d\x4d\x17\x31\xbf\x56\x9c\xa1\x2c\x21\x51\x2b\x7c\x4e\x5f\x9a\x97\x45\x49\x69\x9f\xfb\x9f\x50\x06\x1e\x01\xf3\xf6\x77\x12\x23\x9b\xc0\x3e\x09\x41\x09\x13\xe1\x62\x0c\x2c\x86\x4f\x91\xc6\x40\xcb\x8d\x73\xde\xa7\x4b\x97\x16\x7a\x1f\x2f\xf4\x16\xa0\x24\xd7\x96\x50\xc5\x48\x01\x9e\x00\x30\x92\x9d\xb1\x80\x65\x4c\x65\x6c\x8d\x32\x4a\x06\x09\xa5\x34\xc9\x70\xaa\x9d\xfb\x19\xe3\xe7\x59\x4b\x8b\xe7\x8b\xf2\x13\x40\xc1\x04\xf0\xd9\x44\x5d\x5e\xc8\x72\x8b\xd3\x40\x85\x69\x80\xc2\x7b\x10\x0e\x52\x4b\xf5\xd8\x06\x03\x79\xd8\xda\x80\xbd\xe7\xf8\xc1\x4f\x88\xff\x3b\x74\xbf\x25\x34\x61\x3a\x39\x4d\x72\x4e\x0f\xa4\xe0\xd8\x44\xe6\x5c\x53\x49\x12\x23\xcf\xd9\xcf\x83\x67\xcc\xa3\x77\x7a\xbb\x28\xbe\xc4\x8b\x70\x50\x45\x87\x2f\x5f\xfa\xc1\x18\xa8\xbc\xaa\xed\x2e\x75\x3a\xb4\xa0\xb1\x81\x18\x26\xf6\x60\x2a\x8e\x4d\xe6\x68\x8c\x0f\xc6\xb3\x7f\x8a\xfd\xab\x16\x46\x76\xa1\x23\xb4\xc0\x40\x85\xb1\xda\x3e\x37\xec\x0b\xd8\xff\x28\xc1\xd1\xe0\xed\x85\x90\xc6\x37\x0a\x80\x1f\xea\xcb\x48\xaa\xea\x69\x2f\xa4\x36\x2c\xf5\xf7\x79\x99\x55\x06\x1b\x10\x74\x33\x70\x35\x9e\x73\x80\x3a\x68\xea\xb5\x5c\xf4\xe8\x8e\x3c\x21\xc5\x27\xf3\x9e\xa8\x34\xf0\xbb\x4e\x8e\x0e\xf5\x65\x4c\xec\xa7\xcb\xc0\x55\xb6\xdf\x0c\x7a\xbf\x3d\x9f\x1a\xf7\x53\xf7\x2f\xad\xe4\x39\x0d\xfa\x34\x69\xc3\x70\x85\x7d\x61\xa8\xd0\x26\x9a\x4a\x61\xf2\x58\xe9\x56\x18\xc5\x15\x02\x24\x1d\xcc\x8c\x01\x6b\x7c\x72\xe8\x22\x27\xee\x0c\xc7\x02\x7d\xe1\x1d\x98\x86\x00\x58\x9c\x9f\x5d\xf4\x3b\xf9\x6e\xa2\x44\x18\xee\x0c\x74\x2c\x0f\x76\xc2\x28\x93\x05\xf4\xe6\x71\x0b\x41\xec\xa9\x89\xde\xfa\x55\x94\xb7\xfb\xce\x37\xc1\xdc\xcb\x5e\xc5\x79\x84\x1b\x1f\xf4\xa5\x07\xc2\x6b\xda\xd7\x1e\xfd\xa9\x81\x53\x03\xa7\x90\x35\x8c\x95\xcf\x9c\xef\x3a\xdf\x07\x56\xa7\xb1\xab\x75\x31\x59\xb6\x8b\xe1\x63\x6c\x1c\x44\x62\x73\x63\xad\x0f\x44\xc6\x7a\x5e\x0c\x2b\x5f\x6a\x2d\xf0\x68\x09\x05\x25\x68\x81\x63\xdb\x4f\xe7\xa0\x77\xae\x4b\xca\xb7\x38\xe7\xef\x8a\xc6\x66\xe4\x87\x81\x35\xc0\x38\x5a\x90\x69\xc1\x1d\x6c\x9c\xba\x12\x99\x72\xb0\xe6\x3d\xe7\x3f\xbc\x8c\x19\x61\xee\x6e\x04\x80\x28\x38\x1b\x39\xfe\x2f\x43\x12\x0a\xf1\x1f\xe2\x11\xc3\x77\xf4\x95\xd7\x96\x83\x73\x54\xaf\x17\xff\x87\x7d\x92\xcd\x6a\x07\x73\x1e\xfe\x6f\x18\x59\xea\x6e\xd2\x5c\x67\xf5\x1e\x17\x2b\xbe\x81\x2f\xdb\xf6\xc8\x98\x0e\x16\x72\x7d\x2c\xd3\xf1\x45\x0c\x93\xe3\xf7\x2e\x90\x7b\xc0\x7f\x3f\x01\x97\x2f\xfe\x1f\xfd\x73\x34\x5c\x90\xb1\xe4\x4c\x00\xe9\x48\x03\x49\xe3\xf6\x33\xea\x4e\x5d\xfd\x77\xce\x50\xb6\xbf\xc3\x1b\x0c\xef\x4d\x7a\x25\xc0\x7c\x8d\x73\x87\x98\xb8\x01\x75\x72\xed\x4e\x7d\xc1\xfd\x3b\x00\x00\xff\xff\x5a\x95\xd2\x7b\xf8\x0f\x00\x00")

func comboJsBytes() ([]byte, error) {
	return bindataRead(
		_comboJs,
		"combo.js",
	)
}

func comboJs() (*asset, error) {
	bytes, err := comboJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "combo.js", size: 4088, mode: os.FileMode(420), modTime: time.Unix(1434758159, 0)}
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
	"combo.css": comboCss,
	"combo.html": comboHtml,
	"combo.js": comboJs,
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
	"combo.css": &bintree{comboCss, map[string]*bintree{
	}},
	"combo.html": &bintree{comboHtml, map[string]*bintree{
	}},
	"combo.js": &bintree{comboJs, map[string]*bintree{
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

