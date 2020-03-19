// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/tpl/curd.tpl
// assets/tpl/entity.tpl
// assets/tpl/example.tpl
// assets/tpl/init.tpl
// assets/tpl/markdown.tpl
// assets/tpl/tables.tpl
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsTplCurdTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\x5f\x6b\xdb\x56\x14\x7f\xb6\x40\xdf\xe1\x54\x94\x20\xb5\x46\xce\x60\xec\x21\x43\x1d\xb1\xad\xac\x66\x89\xd3\xd9\xce\xc6\x28\x65\xc8\xd6\x75\x26\xaa\x3f\xce\xd5\x55\x9c\x20\x2e\x74\xb0\x87\x15\xd6\xae\x2f\xc9\x4a\x17\x36\x52\xba\x51\x06\x69\xf2\xb0\xb1\xbe\xed\xcb\xc4\x5a\x3e\xc6\xb8\x57\xb2\x2d\xff\x8b\xe5\x90\x41\xc7\x0a\x25\x44\xd2\xfd\x9d\x73\x7e\xe7\x77\xfe\xdc\x90\xfd\x0e\x82\x30\x54\xeb\x04\x07\x2d\xd2\x30\x9a\x36\xaa\x1a\x0e\xa2\x74\xc3\x33\x91\x0d\x3e\x7f\x0c\xa1\x28\x94\x8b\x70\xcb\xdf\xb1\xd5\x72\x51\x14\xa8\x28\x88\x42\x3b\x70\x5b\x50\x45\xdd\x69\x87\x65\xb3\x09\xaa\xaa\x26\x07\x14\xb8\x35\xdb\x42\x28\x0a\x56\x1b\x6c\xe4\xca\x66\x53\x81\x3b\xb0\xcc\x9e\x00\x00\x60\x44\x02\xec\xc2\xd2\xcc\xa3\xc9\x77\xec\x5f\xb9\xb8\x02\x66\xf3\xfe\xf2\x83\x7c\xfc\x8c\x72\x1f\xb3\x22\xb0\xd3\x8e\xe1\x13\x84\xcb\xc5\x3c\x3f\xc9\x03\x2c\x14\xe0\xe2\xe9\x9f\xbd\xef\x0f\xa3\xc7\x8f\xa2\xa3\xc7\x7f\xbf\xf8\xe6\xe2\xf8\x75\xef\xe4\x87\xe8\xcd\x1f\x49\xf4\xb2\x73\x49\x64\x0a\x6c\x23\x52\xf2\xec\xc0\x71\x7d\x59\x61\x4c\x5a\xee\x36\x0b\x2e\x71\x4b\x62\xb4\xaf\xda\xf6\x9a\x85\x6c\x73\xdd\xf2\x09\xa5\x20\x8d\x9b\xee\xbd\x7a\x71\x71\xfc\x5d\x74\x70\x16\x3d\x79\xa3\x66\xb6\x5a\xf3\xba\xbe\xec\xef\xd8\x8d\x3d\x92\xd8\xcd\x43\xc7\xc0\x86\xe3\xb3\xac\x58\x2e\x41\xb8\x6d\xb4\x50\x48\x15\x90\xb1\xd7\xf5\x6b\xc8\x0f\x6c\x02\xf7\x1f\x30\xdc\x7b\x0f\xb7\x75\x97\x58\x64\x9f\xd2\x69\x56\xf2\x80\x30\x66\xff\x3d\xac\xb0\x68\x76\x02\x84\xf7\xe3\x87\x2b\x1a\x38\x6a\xb9\xa8\x7e\xca\x1e\x25\xf6\xfb\x86\x55\x55\x55\x78\xa6\xd9\x87\x37\x34\x70\x2d\x7b\xc8\x05\x8f\xda\x44\x6d\x84\x81\xc3\xa9\x25\xdb\xf3\x91\xac\x88\x42\xdb\xeb\x3f\xaa\xa2\x3d\x22\x73\x8b\xd8\xeb\x32\x53\x13\xae\x56\x03\xdb\x9e\x70\x37\xa4\xa2\xc0\x4c\x6a\x09\x4c\xbd\x65\xb8\xb2\x28\x84\x21\x36\xdc\x6d\x04\xfc\x10\xcf\x80\x5f\x71\xdb\x1e\xa5\x4b\xd8\xeb\xaa\x61\xa8\xde\x0d\x9c\x4e\x12\x70\xa1\xc0\x6c\x95\x3c\xc7\x41\x2e\xa1\x94\x1d\x46\xae\x49\x69\x1c\x0f\x0b\xe4\x86\xc6\xc3\x0a\x45\xa1\xe5\xb9\xc4\x72\x03\x14\xeb\x6f\x48\xad\x06\x46\xa7\x83\x5c\x33\x45\x77\x9e\x0b\x73\x0e\xdb\xe1\x25\xae\x86\x21\xa3\x73\x07\xd4\x8f\xbd\x06\xab\x62\xa9\x6d\x7b\x06\xf9\xe0\x7d\x89\x43\x0d\x03\x58\x99\x08\x49\x5d\x8b\xbf\xcc\x17\x0a\x13\x91\xd9\x3e\x82\x31\x5c\xcb\xcd\x86\x5a\x71\xb3\x63\x12\xcb\x41\x6a\xc3\x72\x50\x06\x5c\xf6\xd9\x0c\xd8\xf9\x87\xeb\xb1\xfe\x67\x25\x31\xf9\xc1\x92\x49\x47\xe4\x98\x2e\xc2\x27\x07\x83\x22\x5c\xa4\x06\x79\x09\xa0\x6c\x35\x98\xe8\xe4\xaa\x05\x38\x5a\x7b\x29\xdb\xa3\xf5\xf7\x6e\x95\x4e\xd2\x0a\xd8\x90\xd0\x31\xae\x7a\xac\x6d\x8d\x35\x85\x21\x33\xda\x3b\x50\x2e\xf0\x7f\xaa\x97\x4b\xcb\x85\xd5\xc7\x78\xb1\x7c\x19\xfd\xf8\x7b\x74\x78\xb6\x48\x9d\xd4\x8d\x5d\x34\x3e\xa8\x76\x0d\x3b\x40\x93\x35\xd2\x84\xa6\xe7\xd9\xe3\xe2\xf7\x89\x43\x46\x87\xcf\x3d\x8c\x3a\x06\xee\xa3\x66\x9a\x39\x0c\x64\x38\x72\x70\xd2\x9d\x13\x50\xfe\x52\xdf\x43\x2d\x99\x3b\x36\x7f\x8e\xed\x1a\x18\x8c\x76\x1b\xb5\x48\xc9\x0b\x5c\x02\x5c\x0d\xa2\x90\x7a\x14\x63\x6b\x10\x5b\x52\x99\xec\x57\xf9\x5b\x64\xca\x73\xc0\x9b\x6c\x92\xa4\xc0\xef\xc0\xf2\x64\x22\xa2\xc3\xb3\xde\xcb\x9f\xce\xff\x3a\x8e\xbe\x3e\xcd\x96\x87\x12\x46\x06\x41\x71\x84\x59\x7a\x90\x02\xb2\x6d\xf8\xa4\x62\xc6\xd1\x8d\x67\xa5\xe5\xb9\x3e\x81\x7e\xf7\xd3\x40\xaa\x54\xeb\x7a\xad\x01\x95\x6a\x63\x13\x24\xb8\x9d\x34\x20\x0e\xc8\x0d\x6c\x75\x3a\x08\xa7\xf0\xe1\x36\x48\x20\x87\xa1\x5a\x71\x7d\x84\x49\x6a\x45\x52\xe0\xb3\xd5\xf5\x2d\xbd\x9e\x7a\xbb\x61\xe0\x87\x94\x2a\xd2\x3c\x35\xa0\xeb\x95\x43\xaa\xd5\xc4\x7e\xc4\x6d\x26\x56\xc9\x15\xda\xe0\xa4\x3b\x31\xc5\x63\x6a\x59\x67\x0f\x63\x7b\x73\xd5\x32\x21\x8c\x85\x0b\x74\xab\x63\x2e\x2a\x8c\xe9\x85\x3a\x50\xc3\x8a\x06\xd2\xd6\xbd\xf2\x6a\x43\xcf\xae\x84\xba\xde\x00\xfe\x8e\x39\x93\xde\x97\x3f\xbf\xab\xd7\x74\x8e\x81\x2d\xc7\xc0\xfb\x9f\xa0\x7d\x4a\x41\x83\x8f\x24\x51\x48\x86\x2d\x53\x82\xf1\x10\xc9\xf7\x1f\xa4\xfa\x49\x1e\x96\x95\x61\xfe\x6e\x5a\x79\xb8\xb9\x6b\xd8\xec\xdb\xc4\x08\xc3\xe7\x86\x28\x4d\x70\x06\x1b\x5c\xfc\x7b\x1e\xc2\x90\x9d\xe1\x49\x4c\xd2\x39\x58\xed\x1d\x75\xd0\xd8\x26\x46\xf0\x20\x17\x3f\xff\x72\x71\xfa\x32\xbd\xdf\x67\xcb\xc8\x9a\xe5\x9a\x0b\xe5\x03\x23\x9f\x45\x73\xb5\xed\x7e\x24\x6d\x75\x7d\x5d\x2f\x35\x58\xd6\x1c\x75\xe4\x6e\x73\x1b\xa4\xb5\xda\xe6\x46\xa6\x84\xf2\x8a\x62\x0e\xf5\x75\xcd\xb1\x06\x37\x96\xb8\x48\xaf\x6b\x0f\x5b\xb3\xb0\x4f\x16\x65\xeb\x8a\x7b\xd8\xf5\x53\xc5\xb5\xbf\x5e\xd9\xa8\x34\xe0\x3d\x69\xb4\x13\x0d\x69\xcb\xd8\xda\xa6\x33\x1a\x1d\x3d\xea\x3d\x7b\x7a\xfe\xf6\xd1\x82\xbc\xb2\x26\xf4\xdf\xa6\x75\xb3\x56\xd6\x6b\x50\xfc\x02\x2a\x65\x28\xeb\xf5\xd2\xbf\xcb\x73\xef\xc9\xc1\xf9\xdb\xdf\x16\x61\x78\xd3\x45\x72\xe0\x23\xdc\x1f\xb1\x43\x0e\xa7\x4e\xdc\xdc\x24\x4b\x60\x99\xb0\x18\x27\xdd\xaf\x10\x46\x60\x99\x1a\x6b\x9f\x39\x76\x53\xbd\xe4\x5a\x11\x3b\xa7\x88\x42\x2e\xa1\x43\x03\x76\x22\xbe\x25\x2c\xc5\xbe\x2a\x1f\x8e\x11\x95\xcb\xf5\xf9\xc9\x51\x66\x62\x2a\x59\xb1\x16\x33\x2e\x2e\x6c\x0b\x92\x15\x90\x5b\xc3\x5d\x2b\x83\x82\xa0\xb4\xb9\x55\x6d\xc8\xb7\x94\xec\x0c\xcd\xbb\x68\x29\x53\x2e\x4b\x4b\xdc\xab\x05\x15\xd3\xfb\xf6\x55\x74\x78\x12\x3d\x3f\xed\x3d\xfb\xb5\x77\xf2\xbc\x77\xf4\x3a\x1b\x15\xfa\x9e\xe5\x13\x5f\xb6\x86\x82\x99\x3e\x89\xa7\x49\x65\x71\x3a\x46\x04\x93\x4c\xdc\xdc\xbc\xab\xa8\xc5\xf5\xc2\x16\xe4\x56\x7a\x35\xce\xcd\x26\x2e\x37\xc1\xdc\x98\x84\xac\x76\x82\x95\xfc\xe5\x30\x97\x63\x2b\x32\xc1\x01\x9a\xad\xb6\x7f\x02\x00\x00\xff\xff\x0c\x0b\x48\x40\xf4\x14\x00\x00")

func assetsTplCurdTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplCurdTpl,
		"assets/tpl/curd.tpl",
	)
}

func assetsTplCurdTpl() (*asset, error) {
	bytes, err := assetsTplCurdTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/curd.tpl", size: 5364, mode: os.FileMode(438), modTime: time.Unix(1584586545, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplEntityTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\x41\x0a\xc2\x40\x0c\x45\xf7\x85\xde\x21\x8b\x2e\xa5\x07\x10\x5c\x09\x5d\x76\x21\xbd\xc0\x68\x83\x54\x67\x3a\x25\x4d\x29\x25\xe4\xee\x12\x75\xa4\x54\xc1\x59\x4d\x5e\xfe\x0f\x2f\xcf\x44\xca\xc6\x9d\x3d\x1e\x63\x08\xd8\xb3\x6a\x9e\xf1\x32\x20\x24\xae\x0a\x23\xd3\x74\x61\x10\x0b\x93\xeb\xaf\x08\xc5\x6d\x07\x45\xc7\x18\x60\x7f\x80\xb2\xea\xd0\xb7\xa3\xaa\xc8\x93\x95\xb5\x0b\x56\x7b\xbd\x04\x9b\x65\x78\xc3\x44\xaa\x48\xc1\x71\x2a\xc3\x26\x7f\xc2\xe0\xe8\x6e\x3a\x22\xd8\xb7\xf6\xd1\x2f\xdb\x7a\xf2\x7e\x6d\x6c\xf3\x1f\x6b\x8a\xf3\x56\x9a\xe2\xbc\x72\x4e\xc0\x4e\x7d\x9c\x61\xb5\xf9\x65\xf6\x08\x00\x00\xff\xff\x6e\x2e\x32\xff\x48\x01\x00\x00")

func assetsTplEntityTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplEntityTpl,
		"assets/tpl/entity.tpl",
	)
}

func assetsTplEntityTpl() (*asset, error) {
	bytes, err := assetsTplEntityTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/entity.tpl", size: 328, mode: os.FileMode(438), modTime: time.Unix(1584357534, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplExampleTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x6d\x6b\xdc\x46\x10\xfe\x2e\xd0\x7f\xd8\x0a\xee\x22\x99\x8b\xe4\xd4\x49\x09\x82\xa3\xed\xbd\x84\x06\x92\xd4\xc5\x2e\xfd\x50\x4a\xd8\x93\xf6\xee\x16\x4b\xbb\xba\xdd\x55\x1a\xd7\x18\x2e\x50\x9a\x84\x24\x76\x20\x4d\x02\xad\xd3\x77\xf7\x85\x42\xed\x0f\xa5\x31\x97\x94\xfe\x19\x4b\x67\xff\x8b\x32\x92\xee\xcd\x75\x8a\x69\x09\xe4\xdb\xec\xec\xcc\xce\xcc\xf3\x3c\xb3\x11\xf6\x56\x70\x87\xa0\x70\x55\xf6\x02\x5d\xd3\x35\x1a\x46\x5c\x28\x64\xea\x9a\xe1\x63\x85\x5b\x58\x12\x47\xf6\x02\x43\xd7\x8c\x76\xa8\x0c\x5d\xbb\x8a\x8c\x0e\x55\xdd\xb8\x65\x7b\x3c\x74\x3a\xfc\xb4\xec\x05\xa7\x7d\x41\xaf\x11\xe1\x64\xaf\x40\xa8\x22\x52\x51\xd6\x01\x53\x2a\x41\x59\x47\x82\xc9\x88\x72\x62\x01\x01\x16\x94\x52\xab\x11\x41\x8d\x5a\x9d\xb3\x36\xed\x34\x99\xa2\x6a\x15\x49\x25\x62\x4f\xa1\x35\x5d\x7b\x87\x4b\x85\x10\x42\x79\x3e\x72\x9c\x64\x6b\x37\x79\xd2\xd7\xb5\x45\xe8\x0f\x21\x44\x99\x42\x8e\x33\xfc\x75\x27\xd9\xfc\x5e\xd7\xae\xe0\x90\xcc\x84\x0f\x3f\xff\x39\xbd\xf5\x54\xd7\x16\xb1\x94\xb3\xef\xec\x7c\x36\xfc\xe6\x86\xae\x35\x6a\x79\xce\xe4\x62\xf0\x20\xb9\x7f\x4f\xd7\xea\x5d\x2c\x24\x51\x53\x4f\x3d\x7f\x94\x65\x2c\xd3\x90\x7c\xc2\x19\x99\xdc\xa4\x8f\xff\x48\xee\x0e\x74\xed\x32\xbe\x7e\xd1\x0f\x48\xd1\x53\xba\xd5\x4f\x7e\xf8\x69\xf8\xcb\xe0\xf0\xf1\xef\x07\x7f\x7d\x95\x6e\x6c\x67\x11\xef\x46\x84\xcd\x44\xe4\x77\xe9\xc3\x5d\x5d\x5b\x07\x40\x1c\x07\x8d\x5d\xe9\xbd\xdf\x92\xc1\x03\x5d\x6b\xc7\xcc\x43\x17\x19\x55\x8d\x9a\xe9\xb5\x3b\x47\xf0\xb2\xd0\x9c\xec\x05\x76\xa3\x06\x88\xd1\x76\xd1\x98\xb4\x9b\xbd\x18\x07\x17\x78\xe0\x43\x8e\x3d\xea\xbb\x82\x0c\xc3\x82\xc8\x69\x27\xaa\x22\xe3\xd4\xdb\x92\x62\x67\xa9\x8b\x59\xa7\x8b\xe9\x29\x23\xeb\xc7\x97\x0c\xb9\x55\xd4\x0e\x95\xbd\x14\x09\xca\x54\xdb\x34\x4a\xd2\x2d\xc9\xb7\x94\x17\x99\x60\xf9\x96\x53\x92\x6f\x7a\x39\x5c\xd5\x92\x2c\x47\x60\xc1\xc3\x55\x25\x62\x52\x0e\xb8\x57\xbd\xc4\x3d\x1c\x94\x15\x0d\xc9\x55\xa8\x56\x2d\x49\xa3\x92\x37\x00\xe8\x17\x26\x70\x54\x98\x40\xfb\xc8\xcb\xc5\xc8\xcc\xb9\x2a\x0e\x05\x3f\x15\x5d\x8b\x45\x60\xbf\x17\x13\xb1\xda\x94\x1e\x8e\xc8\xcc\xb0\x56\x25\x93\x99\xc7\x19\x23\x9e\xa2\x9c\x55\x10\x11\x02\x26\x02\xc4\x80\x0b\xd3\xc8\xf5\x5a\x41\xbe\x64\x56\x86\x1f\x44\xbc\x56\x45\x8c\x06\x00\x53\x84\x19\xf5\x4c\x22\x84\x95\x01\x32\x79\xca\x5e\x22\xaa\x20\xb4\xce\x19\x93\x59\xe1\xc2\x61\x1d\x13\x08\xda\x98\x09\x04\x87\xa5\x6b\x82\xa8\x58\x30\x34\x89\xcf\xea\x5c\xc3\x02\x35\x6a\x23\x66\xf3\xf3\xda\x5a\x86\xd7\xfa\x7a\xa3\x85\xe6\xc6\x87\xcb\xdc\x27\xd9\xd6\x66\x2a\xa1\x8c\xaa\xfa\xf8\x29\x73\xc4\x34\x8c\x3c\xab\x9a\x62\xbb\x5c\x23\x00\x72\xba\x5c\x2a\xe0\x04\xe0\x76\x17\x16\xe6\xcf\x57\xf2\x6d\x72\x0d\xc1\x79\x7e\x83\xa5\x74\x8d\x33\xaf\x2f\x9c\x3d\xf7\x06\x9c\x73\x36\x5c\x63\x85\x32\x3f\x20\x3e\xb8\x0a\x4e\x5c\x23\x56\xed\xf3\x61\xeb\x2c\xf8\x0a\x40\xdc\x33\xf3\x95\xf1\x86\xb8\xe8\x5c\x25\x1b\xb2\x51\x43\xd5\x29\x5d\x5b\xba\x36\x3d\x62\x15\x5d\x21\x1f\x8f\x1d\x66\xa3\x66\x4d\x56\x24\xfd\x7a\xfb\x60\xe7\xbb\xf4\x76\x3f\xdd\xba\x3d\xfc\xe2\xd3\x7c\x57\x0a\x08\x96\x89\x54\xe3\xb4\x0b\x94\xf9\x1f\x74\x89\x20\xa6\x42\x73\xc5\x8f\x64\x2f\x67\xb0\x1c\x85\x4a\xd7\x5a\x9c\xaf\x00\x52\xe3\xec\x35\x28\xa6\xb8\xcf\xb3\xca\x82\xc8\x38\x50\x63\x09\x4d\xf5\x6a\x43\x19\xb3\x0c\xf9\x16\x7c\x22\x37\xff\x4c\x1f\xee\x1e\xde\xbc\x7b\xb8\xd5\x3f\xf8\xf1\x46\x7a\xe7\x79\xba\xb1\xbd\xbf\xd7\xdf\xdf\xbb\x93\xfd\xb0\xff\x10\x99\xb2\x9b\x42\x70\x31\x91\x19\x2c\xdc\x22\xec\x5b\xc0\xcc\xbc\xec\x87\xf3\x1f\xe5\x57\xf0\x41\x6c\x3c\x4d\x36\x1f\xc1\xff\x71\x7f\x63\x7f\xaf\x9f\x3e\xf9\xf6\xc5\x00\x5c\xc2\x52\x9d\x64\xf6\x7f\x1d\x4e\x48\x65\x32\x1a\x1c\xb7\x1f\x27\x68\xfd\x68\xdf\xfd\x67\x19\x3a\x9b\xc7\xb5\x5b\xe7\x31\x7b\x65\xfa\x4d\x6e\x7d\x99\x3c\x1b\xbc\x18\xdb\xba\x20\x58\x9d\x48\x59\xbe\xff\xdf\x74\x55\x54\x28\xfb\xfe\x4b\x9a\xa5\x79\x1d\x87\x51\x40\xde\x8f\xfc\x97\x3c\x49\x51\xe1\xff\x4d\xf2\x77\x00\x00\x00\xff\xff\x0f\x0d\x50\xf5\xaa\x08\x00\x00")

func assetsTplExampleTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplExampleTpl,
		"assets/tpl/example.tpl",
	)
}

func assetsTplExampleTpl() (*asset, error) {
	bytes, err := assetsTplExampleTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/example.tpl", size: 2218, mode: os.FileMode(438), modTime: time.Unix(1584585779, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplInitTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\xdf\x8b\xeb\x44\x14\x7e\x6e\x20\xff\xc3\x18\xe8\xde\x44\x4a\xb3\xed\xde\x7b\x91\x4a\x50\xdb\x5e\x51\x70\xfd\x41\x7d\x97\xd9\x64\xd2\x0e\x26\x33\xd9\x99\xc9\xe2\x2a\x85\xf5\x45\x7c\x70\x55\x58\x75\x5f\x56\x11\x54\x54\x04\x77\x1f\x44\xa1\x08\xfe\x35\x6d\xf5\xbf\x90\x33\x33\x49\x93\xba\xde\xa7\x64\xbe\xf3\x9d\x33\xe7\x7c\xf3\xcd\xd0\xbc\xe0\x42\x21\xdf\x75\x3a\x5e\x82\x15\x3e\xc1\x92\x84\xf2\x34\xf3\x00\x48\x73\xa5\xbf\x8c\xa8\xb0\x14\x06\x93\x4a\x50\x36\x97\x9e\xeb\x04\xae\xe3\x3a\x67\x58\xe8\xe4\x1c\x4b\x45\xc4\x74\x8c\x9e\x95\xa7\x59\x7f\x3a\x76\x9d\x8e\xcc\xf0\x19\x69\x22\x3a\x21\x2d\x59\x8c\x28\xa3\xca\x0f\xd0\x07\xae\xd3\x89\xd3\x39\x1a\x45\x28\x39\x99\x70\x96\xd2\x39\x40\x9d\x57\xb8\x54\x23\x84\x10\xf2\x32\x1e\xe3\x6c\xc1\xa5\xf2\x7a\x10\x78\x93\x0b\x13\x38\x3a\x3a\x7c\xac\x91\xd7\x71\x4e\x0c\x55\x70\x5e\xb1\xb0\x94\x06\x1b\x0c\x8f\x1e\x3e\x7a\x6c\xd0\xe9\xd8\x72\xbd\x77\x29\x4b\x32\x92\x18\x78\xb2\xc0\x42\x12\x35\x42\x5e\xa9\xd2\xe7\xf2\x93\x87\x06\x3e\xc6\xef\xbd\x51\x10\x36\x42\x83\xc3\xc3\x0a\x78\x35\xc9\xc8\x08\x3d\xd2\xeb\x65\x73\xe6\x48\x0f\x34\x1d\xfb\x71\x3a\x0f\x5c\xa7\x13\x86\xd5\xec\x7b\x91\x25\x28\xa0\xce\x0b\x52\xcf\x8b\xa4\x12\x65\xac\xb4\x14\x30\x36\xb4\x8d\x8c\xc6\x28\x0c\xd7\x37\x77\xeb\xaf\x2f\x5c\x47\x0f\xae\x43\x94\xe9\x6f\x18\x6e\x7f\xb9\x5d\x7f\xf6\x9d\xeb\x68\x05\xda\x59\xdb\x2f\x7e\xda\x7c\xfc\x07\x64\x61\x29\xf7\x0a\xde\x7e\xb4\xfd\xf6\x43\xd7\xb1\x6a\xb4\x42\xab\xab\xf5\xe7\x97\xae\x53\x29\xd2\x2c\xf8\xe7\x57\x26\xeb\x6d\x9a\x93\xf7\x39\x23\xbb\xd0\xe6\xfa\xf7\xf5\x27\x2b\xd7\xa9\xf4\x69\x74\xb8\xb9\xb9\x58\x7f\xff\xe3\xf6\xe7\xd5\x3f\xd7\xbf\xfd\xfd\xd7\x37\x9b\x4f\x7f\x30\x34\xd0\xf5\x3f\x34\x43\xd8\x7c\x79\x67\x45\x0a\x43\x54\x43\x9b\xcb\x5f\xd7\xab\xab\x86\x75\x8c\x9e\xb5\x86\x41\x65\x31\xad\x22\x4d\x6d\x77\xb2\xff\xe4\xb4\xc4\xd9\xcb\x3c\x4b\x80\xde\xaf\x9a\xef\x21\xcf\x33\xde\xeb\x34\x61\x14\x21\xef\xc1\x4b\x92\xe2\x70\xb6\xc0\x6c\xbe\xc0\xf4\x81\x67\x0f\x3a\x91\x0c\x3c\x9a\xe6\xaa\x3f\x2b\x04\x65\x2a\xf5\xbd\xae\x1c\x75\xe5\x8b\x2a\x2e\x7c\xf8\x4b\x82\xb0\x2b\x5f\x88\x8d\x72\x51\x57\x1e\x14\xf0\x07\xa5\x23\x25\x4a\x72\x90\xf1\x38\x7a\x0d\xbc\x7c\xa0\x68\x4e\xde\x81\xfd\xa2\xae\x34\x56\x83\x26\xe0\x30\xea\x05\x1c\x5b\xbd\x00\x53\xec\x22\x5c\xec\x16\xe6\x08\xeb\xa5\x3d\x36\xbd\x2e\x45\xd6\x7f\xab\x24\xe2\xfc\x89\x8c\x71\x41\x5a\xe3\x07\xc0\x00\x97\x26\x24\x25\x02\x81\xa8\xf6\x2a\x82\x74\x44\x08\x18\x55\x90\x98\x9f\x11\xe1\x07\xcf\x6b\xe4\x99\x08\x31\x9a\x19\x92\x65\x0d\x80\x76\x6c\x6f\xc0\x24\xe3\x92\x58\xf2\xa0\xcd\xee\x14\x98\xd1\xd8\x87\x40\xa0\x81\x65\xa3\xc8\x10\x8a\xcc\xcc\x5d\x69\xd6\x18\xfe\x5f\x8d\x61\xb3\x46\x8d\x6a\x10\xb0\xa5\x0f\xbf\x31\x67\x8c\xc4\x8a\x72\xd6\xab\xe6\x01\x7f\x80\xef\x7c\x2f\x3f\x87\xd7\xad\x87\x12\xc9\x02\xe3\x96\xfd\x01\x5b\x65\x97\xad\x7a\xfd\x19\x51\xd6\xc1\x13\xce\x98\xd4\xc2\x5a\x20\xb8\x8f\x09\x57\xa2\xc5\x04\x00\x98\x82\xa8\x52\x30\xb4\x4b\xb0\xbe\xd7\x1e\xdf\x93\x15\x3a\xe4\xa2\x32\x77\xfd\xea\x34\x7b\xb6\xe5\xaa\x58\xdf\x66\xda\x01\x6c\x94\xd1\x4c\xef\xa2\xf7\x68\xab\xde\xde\xa2\x7a\xbd\xee\xd9\xc1\x86\x9e\xb6\xc1\xbf\x01\x00\x00\xff\xff\x8c\x02\xa6\x98\x53\x06\x00\x00")

func assetsTplInitTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplInitTpl,
		"assets/tpl/init.tpl",
	)
}

func assetsTplInitTpl() (*asset, error) {
	bytes, err := assetsTplInitTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/init.tpl", size: 1619, mode: os.FileMode(438), modTime: time.Unix(1584438333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplMarkdownTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\xc1\x6a\xea\x40\x14\x86\xf7\x81\xbc\xc3\x81\xb8\x50\xb8\xc9\x03\x08\xf7\x6e\xae\x1b\xb9\x17\xeb\x42\xba\x11\x17\xb1\x9e\x16\x69\x92\x96\x68\xa0\x61\x66\xc0\x45\xa1\x85\xd2\xd6\x45\xc1\xac\x0a\x05\x11\xbb\xa8\x52\xda\x82\x98\x3e\x8e\x19\xf5\x2d\x4a\x66\x9c\xd4\x56\x9a\xc5\x84\xf9\xff\xc3\xf9\x4f\xbe\x13\xc3\x80\xd5\xc3\x98\xf7\x62\x3e\x3a\xd7\x35\x5d\xa3\xc9\xfc\x36\xb9\x1c\xd0\x54\x4d\xfa\xd7\x40\x61\x35\x7d\xe5\xd1\x0d\x50\x48\x86\x17\xfc\x65\x0c\x54\xd7\x68\xd1\x34\x4d\x71\xa8\x13\xe4\x4b\xd7\x08\xf1\x6d\xef\x08\xc1\xaa\xd9\x4d\x07\xff\xb7\x3b\x5d\xc6\x28\x10\x62\x95\xbd\x16\x9e\x31\x46\xeb\xe9\x45\x98\x15\xdb\x45\xc6\x1a\x79\xe3\xd3\xfd\x6a\x15\x40\x3e\x94\x10\xeb\xef\x89\xeb\xa2\x27\x9a\x89\x14\xf4\x5a\x8c\xa5\x03\xab\xc0\xdc\x31\x86\xbf\x20\xd7\xee\xa2\x0b\xc5\xdf\x60\x95\xb0\x73\x20\xe3\x75\xcd\x30\x0c\x20\x44\x78\x2a\xc9\x52\xf7\xad\x3c\x5d\xfb\x93\x95\x65\x79\x69\x86\x29\x68\x3c\x0d\xf8\xe4\x4d\xd2\x90\xa4\x84\x00\x14\x96\xcf\x71\x72\x7f\x05\x14\x78\x34\x4d\xfa\xa3\xc5\x6c\xbe\x7c\x9c\x53\x58\xc7\xd1\x6a\x32\x4c\x7a\xef\x99\xc3\xa3\xe9\x62\x16\xaf\xef\x26\xbb\x58\x15\x55\x73\x97\xed\x37\x61\x0b\xb2\x1c\x75\x03\x59\x40\x72\x02\xd7\x93\x1f\x03\x20\xb0\x4b\xa9\x16\x9e\x62\x25\x70\x9b\xe8\x33\x26\xf5\x72\xa7\x12\x38\xce\x66\x37\x25\x3c\xb4\x03\xa7\xbb\x6f\x3b\x01\xaa\x82\xaa\xdf\x76\x6d\x3f\xfc\x87\xa1\x52\x64\xab\x1f\x16\x51\xaf\xed\x55\x1b\x79\x23\xfb\x97\x0a\x99\xf7\x11\x00\x00\xff\xff\x54\x04\xaf\xad\x62\x02\x00\x00")

func assetsTplMarkdownTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplMarkdownTpl,
		"assets/tpl/markdown.tpl",
	)
}

func assetsTplMarkdownTpl() (*asset, error) {
	bytes, err := assetsTplMarkdownTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/markdown.tpl", size: 610, mode: os.FileMode(438), modTime: time.Unix(1584357534, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTplTablesTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe2\xe5\x4a\xce\xcf\x2b\x2e\x51\xd0\xe0\xe5\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\x50\xc9\xd4\x51\x50\xc9\x2c\x49\xcd\x55\xb0\xb2\x55\xd0\xab\xad\x05\xc9\x80\xf9\x7a\xa1\x05\x05\xa9\x45\x21\x89\x49\x39\xa9\x7e\x89\xb9\xa9\xb5\xb5\x0a\x0a\x0a\x0a\xb6\x0a\x4a\x30\x69\x24\x19\x25\x05\x7d\x7d\x05\x98\xb8\x73\x7e\x6e\x6e\x6a\x5e\x49\x6d\x6d\x75\x75\x6a\x5e\x0a\xc8\x40\x4d\x40\x00\x00\x00\xff\xff\xf6\x4d\x87\xcf\x77\x00\x00\x00")

func assetsTplTablesTplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTplTablesTpl,
		"assets/tpl/tables.tpl",
	)
}

func assetsTplTablesTpl() (*asset, error) {
	bytes, err := assetsTplTablesTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/tpl/tables.tpl", size: 119, mode: os.FileMode(438), modTime: time.Unix(1584357858, 0)}
	a := &asset{bytes: bytes, info: info}
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
	if err != nil {
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
	"assets/tpl/curd.tpl":     assetsTplCurdTpl,
	"assets/tpl/entity.tpl":   assetsTplEntityTpl,
	"assets/tpl/example.tpl":  assetsTplExampleTpl,
	"assets/tpl/init.tpl":     assetsTplInitTpl,
	"assets/tpl/markdown.tpl": assetsTplMarkdownTpl,
	"assets/tpl/tables.tpl":   assetsTplTablesTpl,
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
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"assets": &bintree{nil, map[string]*bintree{
		"tpl": &bintree{nil, map[string]*bintree{
			"curd.tpl":     &bintree{assetsTplCurdTpl, map[string]*bintree{}},
			"entity.tpl":   &bintree{assetsTplEntityTpl, map[string]*bintree{}},
			"example.tpl":  &bintree{assetsTplExampleTpl, map[string]*bintree{}},
			"init.tpl":     &bintree{assetsTplInitTpl, map[string]*bintree{}},
			"markdown.tpl": &bintree{assetsTplMarkdownTpl, map[string]*bintree{}},
			"tables.tpl":   &bintree{assetsTplTablesTpl, map[string]*bintree{}},
		}},
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
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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
		err = RestoreAssets(dir, filepath.Join(name, child))
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
