// Code generated by go-bindata.
// sources:
// ../../../../images/base/Dockerfile
// ../../../../images/base/clean-install
// ../../../../images/base/entrypoint
// DO NOT EDIT!

package sources

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

var _imagesBaseDockerfile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\xfb\x6f\xdb\x46\xf2\xff\x9d\x7f\xc5\x40\x0a\xf2\xb5\x53\x91\xf4\x23\x6e\xdd\x7c\xe1\x1e\x54\xc5\x49\x85\x3a\x52\x20\x29\x2d\x82\xbb\x43\xb0\x5c\x8e\xc8\xad\x96\xbb\xbc\x7d\x88\x56\x72\xfe\xdf\x0f\xb3\xa4\x5e\x4e\xd2\x2b\xae\x77\x41\x00\xdb\xdc\x99\xd9\xcf\xcc\x7c\xe6\xb1\x7d\x18\xe9\x7a\x63\x44\x51\x3a\xb8\x38\x3b\xbf\x86\x45\x89\xf0\xb3\xcf\xd0\x28\x74\x68\x61\xe8\x5d\xa9\x8d\x4d\xa2\x7e\xd4\x87\x3b\xc1\x51\x59\xcc\xc1\xab\x1c\x0d\xb8\x12\x61\x58\x33\x5e\xe2\xf6\x64\x00\xbf\xa0\xb1\x42\x2b\xb8\x48\xce\xe0\x84\x04\x7a\xdd\x51\xef\xf4\xff\xa3\x3e\x6c\xb4\x87\x8a\x6d\x40\x69\x07\xde\x22\xb8\x52\x58\x58\x0a\x89\x80\xf7\x1c\x6b\x07\x42\x01\xd7\x55\x2d\x05\x53\x1c\xa1\x11\xae\x0c\xd7\x74\x46\x92\xa8\x0f\xef\x3b\x13\x3a\x73\x4c\x28\x60\xc0\x75\xbd\x01\xbd\x3c\x94\x03\xe6\x02\x60\xfa\x57\x3a\x57\xbf\x48\xd3\xa6\x69\x12\x16\xc0\x26\xda\x14\xa9\x6c\x05\x6d\x7a\x37\x1e\xdd\x4e\xe6\xb7\xf1\x45\x72\x16\x54\xde\x29\x89\xd6\x82\xc1\x7f\x78\x61\x30\x87\x6c\x03\xac\xae\xa5\xe0\x2c\x93\x08\x92\x35\xa0\x0d\xb0\xc2\x20\xe6\xe0\x34\xe1\x6d\x8c\x70\x42\x15\x03\xb0\x7a\xe9\x1a\x66\x30\xea\x43\x2e\xac\x33\x22\xf3\xee\x28\x58\x5b\x74\xc2\x1e\x09\x68\x05\x4c\x41\x6f\x38\x87\xf1\xbc\x07\x3f\x0e\xe7\xe3\xf9\x20\xea\xc3\xaf\xe3\xc5\x4f\xd3\x77\x0b\xf8\x75\x38\x9b\x0d\x27\x8b\xf1\xed\x1c\xa6\x33\x18\x4d\x27\x2f\xc7\x8b\xf1\x74\x32\x87\xe9\x2b\x18\x4e\xde\xc3\xcf\xe3\xc9\xcb\x01\xa0\x70\x25\x1a\xc0\xfb\xda\x10\x7e\x6d\x40\x50\x18\x31\xa7\x98\xcd\x11\x8f\x00\x2c\x75\x0b\xc8\xd6\xc8\xc5\x52\x70\x90\x4c\x15\x9e\x15\x08\x85\x5e\xa3\x51\x42\x15\x50\xa3\xa9\x84\xa5\x64\x5a\x60\x2a\x8f\xfa\x20\x45\x25\x1c\x73\xe1\xcb\x67\x4e\x25\x51\xd4\x87\x95\x50\x39\x70\xe9\xad\x43\x03\x19\x23\x57\x2b\x56\x60\x08\xec\x2b\x6d\xc0\x6e\xac\xc3\x2a\x87\x6f\x20\xd7\x7c\x85\x06\xb8\x56\x4b\x51\x78\x13\xac\x12\x23\x72\xc8\x50\xea\x66\x00\xb6\x83\xbc\xd4\x52\xea\x86\x00\x19\x5c\xa2\x41\xc5\xd1\xbe\x88\xfa\x21\xad\xb6\xcb\xeb\x92\xb2\x81\x76\xe5\x74\x1d\x92\xdb\x88\x95\x48\xe7\x5d\x36\xd2\xee\xd2\x74\xa4\x15\x31\x06\xcd\x58\x39\x34\x4b\xc6\x31\x3d\xb0\x93\xe3\x1a\xa5\xae\xd1\xd8\xc4\x60\x5e\x32\x97\x70\x5d\xa5\x99\xd4\x45\x7a\x71\x76\xfe\x3c\x3d\xbb\xa2\xff\xc6\x2b\x0a\x4e\xdc\xd9\x8c\x89\xa0\x42\xc5\xad\x37\x31\xdf\xde\xf0\xc7\x0d\x7f\x9b\x9e\x7d\x9f\x9e\x5f\x7e\x66\x58\xa8\x98\xc5\x4a\xab\xb8\x36\x62\x2d\x24\x16\x98\x1f\x9a\x8f\x86\xb3\xd7\xc4\x94\xdb\x0f\xe3\x37\xc3\xd7\xb7\x37\x3d\x9f\x79\xe5\xfc\x8b\xf3\xef\x93\xb3\xe7\xbd\xe8\xd5\x6c\xfa\x06\x9e\x7c\xda\x0b\x3c\x50\x76\x2c\x3a\x22\x2a\xbc\xbc\xfd\x71\x3c\x9c\x7c\x78\x35\x9b\x4e\x16\xb7\x93\x97\x37\x4a\x2b\x41\x21\x61\xdc\x89\x35\x82\x75\xba\xb6\x60\x75\x85\xc0\x6a\x07\x0d\x0b\x7c\xb0\x83\xb6\x58\x85\x0d\xc5\x1b\xf5\x81\x81\x41\x26\x81\x99\xc2\x57\xa8\xdc\x00\x1a\xfc\x3f\x83\x70\xc2\xb2\x53\x6f\xe9\x1e\xc2\xe8\x34\x14\xe8\x80\x81\xc3\xaa\xd6\x86\x99\x0d\xdc\x4e\x7e\x01\x56\x30\xa1\x92\xe0\xc5\xef\xa3\x89\xa2\xd1\xf4\xed\x7b\xe0\x12\x99\x8a\x85\xb2\x8e\x49\x09\xa9\xb7\x26\x95\x9a\x33\x99\x66\x42\xa5\x47\x87\xd1\xec\xdd\x04\x78\x59\xe9\x1c\xbe\xb9\xff\x7d\xc9\xa8\x0f\xaf\xd1\x41\x8e\x35\xaa\x1c\x15\x17\x68\xa3\x7e\x68\x80\x7b\xea\x02\x93\x06\x59\xbe\x81\x92\xd9\x17\x60\x6d\x39\xa0\xa0\x0c\xc0\x2a\x56\xe7\x41\xba\x8d\x49\x66\xf4\x0a\x15\xe4\xba\x51\x20\x94\xd3\x70\x82\x8c\x97\xa1\xb4\x41\x0a\x85\xa7\xc4\xd9\x18\x6a\xc6\x57\xac\x40\x0b\x8a\x18\x1b\x3a\x88\xf1\x0a\x2c\x9a\xb5\xe0\x68\xe1\xa4\xcb\xff\x69\x90\x1e\xcd\xc6\x70\xb2\xcb\x7a\xf7\xf1\xb1\x09\xaa\xe5\xd5\xbe\x61\x53\xf7\xd4\x0a\x95\xb3\x41\xba\x12\x96\xef\x55\x42\x7d\x7a\x8b\x16\x84\xb3\x28\x97\xad\xbb\x0a\x1a\x6c\x23\xec\x6b\x38\x31\x58\xe9\x35\xa5\xcf\xab\x86\x29\xea\x4f\xdb\xaa\xdd\xa2\x24\x1c\xaf\x84\x62\x52\x6e\x48\x33\x17\x36\xf4\xc6\x55\x65\x0b\xea\x87\xbf\x69\x6f\x14\x93\xf9\xff\xae\xb8\x42\x8a\x8f\x18\xf1\xb7\x28\xf4\xfa\x3d\xd4\xce\x84\xdd\xd8\x35\x48\x91\x75\x7f\x9f\xed\x04\xf7\x51\x3d\xfc\xa4\x9c\x61\x7c\x05\xa2\x76\xe4\x91\x05\x51\x1b\xed\x1d\x5e\x00\xba\xd2\x69\x2d\xc1\x6a\xce\x1c\x78\x27\x64\x2c\x85\xf2\xf7\x50\x69\xaf\x1c\x60\xd6\x29\xf8\x1c\xd7\xb0\x22\xf2\x6d\xad\x66\xcc\x96\xc0\x59\xcc\xd1\x38\x6a\xb5\x2c\x24\xc9\x1b\x09\xc6\x6e\x14\xef\xe4\x9e\x3e\x85\x25\xe5\x26\x95\x22\xdb\x35\xac\xf6\x27\xfd\x10\x4a\xb8\xc4\x31\x53\xa0\x4b\x28\x29\x36\x85\x58\xb1\x0a\xa1\xb7\x75\xd4\x55\x35\x8d\x51\x1b\x5b\x74\xbe\x4e\xba\x54\xf5\x20\xce\x51\xa2\xc3\xfd\x35\xa6\x82\x78\xf9\xc5\x7b\x2a\x2f\x9d\x88\xbd\x45\x73\x7c\xd5\xb3\xcf\x94\xd1\xf1\xc7\xca\xcf\xbe\x2a\xfc\x85\x9b\x42\x41\xc6\x4b\xfb\xef\xee\xf9\x52\x30\x88\x0c\xee\xb1\x26\x85\xfd\x4f\xa8\x53\x78\xb9\x93\x7f\xc8\x42\xc6\xac\xe0\x5f\x05\x8e\xbc\xd4\xd0\x9b\x21\xcb\x7f\x7e\x63\x8b\x1b\xa5\x7b\xf0\xc3\x0f\xc7\x01\xdb\x96\x47\x42\xb3\x6f\xaf\xd9\x1e\x73\x27\x01\x55\x28\xa7\xcf\x18\xba\x33\x9f\x6b\x85\xd0\x31\x3f\x4c\xea\xae\xbc\x7b\xd4\xd1\x72\xcc\x7c\x71\xa8\xbc\xee\x36\x32\x46\xa3\xd9\x20\x73\x08\x39\x2e\x99\x97\xae\x9b\xbe\x6d\x39\xed\x15\xe2\x78\xab\xb2\xbb\xb7\x5a\xe5\xc2\x40\x5c\xb7\x9e\x1c\xc8\x3e\x7d\x7a\xa8\xd9\xda\xdb\x99\xff\xe1\xb1\x78\xda\x0a\x24\x4e\x57\xa1\xfb\x8e\xbb\xea\x1d\x4d\xc6\x90\x09\xc5\x8c\x40\x4b\x2d\x31\xd5\xb5\x4b\xb9\x12\xd4\xb3\xa9\x45\x4d\x5f\x4e\x4f\x32\x54\xae\x44\x94\x39\x9a\xd3\x17\xb4\x3e\x40\x53\x6e\x20\x85\xa6\x64\x0e\x4a\x34\x18\x86\xc9\x68\x32\xfe\xf0\xcb\xed\x6c\x3e\x9e\x4e\x6e\x7a\x67\xc9\x77\xc9\x55\x6f\xf7\x3d\x4c\xc3\x77\xb3\xbb\x9b\xde\xb6\x2d\x59\xa7\x0d\x2b\x30\x29\xb4\x2e\x24\xb2\x5a\xd8\xd0\x97\xf6\xdd\x34\x36\x28\x91\x59\x4c\x15\xba\x46\x9b\x55\x5c\x4b\x5f\x08\x65\xd3\x5e\x08\x1a\xde\xd7\xda\x38\x18\xce\x46\x3f\xdd\x3c\x39\xc9\xeb\x55\x01\x31\x0d\x6c\xe5\x62\x66\x78\x29\x1c\x72\xe7\x0d\x9e\x1e\x24\xb0\xd5\x20\x38\x8b\xe1\xec\xc7\xe1\xdd\xdd\x4d\x8f\x2b\xb1\xb5\x1b\x3f\xf9\x44\xc6\x1e\xe2\xf5\x93\x4f\x07\xae\x3c\x24\xae\xf8\xd8\xfb\xa2\x95\xe0\x4f\x2b\xbc\xf5\xef\xa1\xfd\xb3\xb3\xff\x70\xa0\x17\xda\x4e\x6c\xe7\x77\x10\xc7\x06\x9d\xd9\xc0\x15\xc4\xb1\xf6\xae\xf6\x0e\x52\x57\xd5\x14\x74\xba\x0b\x3a\x93\x64\xed\x40\xdf\x96\xec\xe2\xea\x5b\xeb\xab\x63\xe1\x2f\xd0\xe4\x20\x81\xfb\x63\xc7\x0c\xc4\xa3\xe3\xc3\xf8\xfe\xe3\xf2\x2b\xd6\xa8\x06\xcd\xf1\xe1\x21\x69\xb8\x11\x54\x2d\x44\x97\xa3\x21\xdf\xe6\x7b\x36\x1e\x2d\xee\xf6\x54\x58\x9f\x27\xe7\xcf\x93\xb3\xff\x34\x6d\x6d\xe0\x96\xf3\x3b\xd8\x71\xa7\x10\xae\xf4\xd9\x63\xbe\x58\x51\xd8\x94\x1b\x11\xd3\xb4\xb0\x69\x47\x1f\x9b\xd2\x5e\x20\x35\xcb\xd3\x27\x9f\x8e\xa1\x3d\xa4\xad\x1f\xf1\x67\x07\xed\x80\xd9\x32\x82\x3a\x4e\x42\x24\xf8\x67\x08\xe3\xfd\xc7\xd1\x63\xb7\xa3\x3e\x38\x94\x72\x37\x02\x1d\x15\x86\x70\xb4\x9c\x08\xb5\xdd\xb7\x4f\x84\x83\x46\x50\xf4\x4a\xe4\xab\xdd\x3b\x60\x57\xa2\x80\x6a\x7d\xfa\x5f\xdb\xae\x69\xd3\xdb\x9b\x6e\x21\xd0\x1e\xda\x21\xc4\x7b\xe1\x2c\xad\x48\xf3\xf1\xeb\xd9\xe2\xcd\x78\xf2\xcd\xe5\x20\x2c\x97\xf3\xf1\xeb\xc5\xed\xec\x0d\x9c\x34\xa5\xe0\x25\x18\x8c\xf1\x1e\xb9\x77\x61\x6f\x39\x84\x97\xf9\xe2\xa3\x90\x92\x1d\x6e\x15\xb6\xd4\xcd\x87\xcc\x17\x09\x2f\xc4\x5f\x44\x7e\x73\x4e\x5b\xf6\xd5\x77\xd1\x7c\x31\x7d\x3b\x1f\xbf\x9e\x0c\xef\x0e\xee\xa3\xa8\x35\x86\xd5\x3b\x4c\xe1\xc5\xa9\xbd\x69\x9f\x46\x8c\xba\xb1\x33\x9b\x5a\x0b\x5a\x70\xe9\x4d\x52\xaf\x8a\x34\xf3\x42\xb6\x8b\x57\xa9\x9b\xdd\x5a\x4c\x5f\x5d\xd4\x87\x30\x21\xc2\x72\x14\x4e\x7e\xf3\xd6\x81\x44\x67\xc1\x5b\x5a\xc1\xc1\xd7\xed\x72\x4d\xcb\x4d\x61\x21\xc3\xa5\x36\x6d\x0a\x84\xf2\xd4\xcf\xb5\x22\x56\x77\x88\x08\x60\x49\x0f\x64\x7a\xd7\xd1\x54\x57\x45\x9b\xda\x2d\x62\x61\xe1\xed\xf8\xe5\x79\xd4\xdf\x01\x6a\x10\x24\xae\x91\x5a\x5b\x80\xb0\x07\xde\x3d\xcd\xda\x9d\xfa\xaf\xd0\xdb\x3b\xd7\x1b\x40\xef\xd1\xb2\xdc\x83\xbf\xd3\x2b\x14\xc3\xa2\xb9\xa7\x95\x86\x0c\xc3\x8d\xdb\xe5\x95\x08\xb4\x66\x46\xe8\xe0\xdf\x76\x91\x6d\xb3\x3d\x08\xab\xa9\x44\x37\x00\x74\x3c\xa1\xdc\x4d\xa6\x8b\xdb\x17\xbb\xa0\x3d\xd3\x4a\x6e\x9e\x05\xe8\xb9\xe6\xe1\x21\x11\xde\x82\x83\x60\x76\x8f\x8f\x64\xe9\x69\x6a\x44\x9e\xa3\x02\xe6\xe8\x6a\x27\x2a\x8c\x6e\x27\x8b\xd9\xfb\xb7\xd3\xf1\x64\x41\x1e\x3d\xf2\xe1\x91\x83\x96\xbe\xd1\xb4\x27\xdf\xbe\x32\x59\xe8\x4d\x13\x58\xb0\xf5\xf8\xcd\x70\x04\x2c\xcf\xc3\xb3\x9a\x59\x2b\x0a\x45\x28\x0f\x58\x78\xd0\x0a\x8e\x97\x85\x3c\x15\xd6\x7a\xb4\xe9\xe5\xe5\x77\xcf\xfb\xe1\x77\xae\x2b\xd2\x8e\x2f\xae\xaf\xaf\xaf\x2f\x2e\xaf\xae\xfe\xa4\x9d\xcb\xcb\xef\x2f\xae\xae\x9f\x5f\x5f\x46\xff\x0a\x00\x00\xff\xff\x43\x57\x8c\x3b\xca\x11\x00\x00")

func imagesBaseDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_imagesBaseDockerfile,
		"images/base/Dockerfile",
	)
}

func imagesBaseDockerfile() (*asset, error) {
	bytes, err := imagesBaseDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "images/base/Dockerfile", size: 0, mode: os.FileMode(438), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _imagesBaseCleanInstall = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x51\x6f\xdb\x36\x14\x85\xdf\xf9\x2b\xce\xac\x62\xd8\x8a\xc8\x4a\xf3\x32\x60\xc5\x86\x79\x49\x86\x19\x2d\x6c\x20\x76\x57\x14\xdb\x1e\xae\xa8\x2b\xe9\x22\x12\xc9\x91\x57\x75\xfc\xef\x07\x2a\x71\x10\xa3\xd1\x93\xc8\x73\xc8\xf3\xf1\x90\xc5\x77\x55\x2d\xae\x4a\xbd\x31\x05\xae\x7d\x38\x46\xe9\x7a\xc5\xd5\xe5\xbb\x9f\xb0\xef\x19\x1f\xa6\x9a\xa3\x63\xe5\x84\xd5\xa4\xbd\x8f\x69\x69\x0a\x53\xe0\xa3\x58\x76\x89\x1b\x4c\xae\xe1\x08\xed\x19\xab\x40\xb6\xe7\x93\x72\x81\xbf\x38\x26\xf1\x0e\x57\xcb\x4b\xfc\x90\x0d\x8b\x27\x69\xf1\xe3\x7b\x53\xe0\xe8\x27\x8c\x74\x84\xf3\x8a\x29\x31\xb4\x97\x84\x56\x06\x06\x3f\x58\x0e\x0a\x71\xb0\x7e\x0c\x83\x90\xb3\x8c\x83\x68\x3f\xc7\x3c\x6d\xb2\x34\x05\xbe\x3c\x6d\xe1\x6b\x25\x71\x20\x58\x1f\x8e\xf0\xed\x4b\x1f\x48\x67\xe0\xfc\xf5\xaa\xe1\xe7\xaa\x3a\x1c\x0e\x4b\x9a\x61\x97\x3e\x76\xd5\xf0\x68\x4c\xd5\xc7\xf5\xf5\xed\x66\x77\x5b\x5e\x2d\x2f\xe7\x25\x9f\xdc\xc0\x29\x21\xf2\x7f\x93\x44\x6e\x50\x1f\x41\x21\x0c\x62\xa9\x1e\x18\x03\x1d\xe0\x23\xa8\x8b\xcc\x0d\xd4\x67\xde\x43\x14\x15\xd7\x5d\x20\xf9\x56\x0f\x14\xd9\x14\x68\x24\x69\x94\x7a\xd2\xb3\xb2\x4e\x74\x92\xce\x0c\xde\x81\x1c\x16\xab\x1d\xd6\xbb\x05\x7e\x5f\xed\xd6\xbb\x0b\x53\xe0\xf3\x7a\xff\xe7\xf6\xd3\x1e\x9f\x57\x77\x77\xab\xcd\x7e\x7d\xbb\xc3\xf6\x0e\xd7\xdb\xcd\xcd\x7a\xbf\xde\x6e\x76\xd8\xfe\x81\xd5\xe6\x0b\x3e\xac\x37\x37\x17\x60\xd1\x9e\x23\xf8\x21\xc4\xcc\xef\x23\x24\xd7\xc8\x4d\xee\x6c\xc7\x7c\x06\xd0\xfa\x47\xa0\x14\xd8\x4a\x2b\x16\x03\xb9\x6e\xa2\x8e\xd1\xf9\xaf\x1c\x9d\xb8\x0e\x81\xe3\x28\x29\x5f\x66\x02\xb9\xc6\x14\x18\x64\x14\x25\x9d\x67\xbe\x39\xd4\x32\xbf\xa5\x15\x92\x8d\x12\x14\xec\x2c\x85\x34\x0d\x94\x9b\x99\xef\x68\x1c\xbd\xc3\x8d\xb7\xf7\x1c\x65\xcc\x49\x81\x54\x39\xba\x99\x45\x5c\x52\x1a\x86\x39\x96\xec\x3d\x75\x9c\x4c\x91\x53\x73\x82\x83\x1d\x98\x66\xa6\x29\xcc\x91\x93\x73\x6c\x39\x25\x8a\xc7\xd3\x52\x50\x54\x69\xc9\x6a\x7e\xaa\xe0\x65\xb7\x7c\x5c\x55\x9e\x74\x09\x9a\x6f\x30\x81\xeb\xa7\x1f\xeb\x9d\xd3\x48\xf6\xde\x98\xc4\x8a\xd2\x83\x63\xe4\x07\x51\x63\xa4\xc5\xdf\x78\x53\xe0\x17\x5c\xe2\xdf\xf7\x33\x84\x01\xd8\xf6\x1e\xbf\x7e\x7f\x85\xc5\xc6\x3f\x73\x9e\x3a\xe4\x66\x91\x2d\x0f\xa2\x78\x67\x5a\x31\x86\x82\x96\x1d\x2b\xa6\xd0\x90\xf2\xf3\xf0\xc4\x53\x1e\x51\x96\xce\x9f\xf8\xca\xc8\xb9\x23\x76\x4d\xc2\xe2\xcd\x6f\x8b\x67\xff\x7c\x0a\x94\x47\x13\x47\x94\xb1\xc5\x3f\x06\x40\xf5\x95\x62\x65\xf3\x63\xae\x1a\xae\xad\x77\x6d\xf5\xf6\x85\x32\x48\x5d\x51\xd0\x6a\x90\xa4\xe9\x5c\xf1\xdd\xf3\x58\xc7\x70\xa6\xbd\x1c\x4f\x29\x56\xa9\xa7\xc8\x55\xe3\xed\x2b\xb3\x23\xb9\x57\x66\x07\x6f\x69\xa8\xde\x9a\xff\x03\x00\x00\xff\xff\x02\x8e\x04\x89\x60\x04\x00\x00")

func imagesBaseCleanInstallBytes() ([]byte, error) {
	return bindataRead(
		_imagesBaseCleanInstall,
		"images/base/clean-install",
	)
}

func imagesBaseCleanInstall() (*asset, error) {
	bytes, err := imagesBaseCleanInstallBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "images/base/clean-install", size: 0, mode: os.FileMode(438), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _imagesBaseEntrypoint = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\x61\x8f\xd3\x46\x10\xfd\x8c\x7f\xc5\x6b\xc2\x07\x90\x62\x07\xf8\x56\x0a\x55\xd3\xbb\x43\x44\xa0\x04\x9d\x43\x29\xaa\xaa\x63\x6d\x8f\xed\x51\xec\x59\x77\x77\x7d\x4e\x84\xf8\xef\xd5\xae\x9d\xcb\x1d\xa8\x52\xef\xcb\x79\x77\xdf\xcc\xbe\x9d\xf7\x66\x32\xff\x69\x99\xb1\x2c\x33\x65\xeb\x28\x9a\xe3\x42\x77\x47\xc3\x55\xed\xf0\xe2\xd9\xf3\x9f\xb1\xab\x09\xef\xfa\x8c\x8c\x90\x23\x8b\x55\xef\x6a\x6d\x6c\x12\xcd\xa3\x39\xde\x73\x4e\x62\xa9\x40\x2f\x05\x19\xb8\x9a\xb0\xea\x54\x5e\xd3\xe9\x64\x81\x3f\xc8\x58\xd6\x82\x17\xc9\x33\x3c\xf1\x80\xd9\x74\x34\x7b\xfa\x4b\x34\xc7\x51\xf7\x68\xd5\x11\xa2\x1d\x7a\x4b\x70\x35\x5b\x94\xdc\x10\xe8\x90\x53\xe7\xc0\x82\x5c\xb7\x5d\xc3\x4a\x72\xc2\xc0\xae\x0e\xd7\x4c\x49\x92\x68\x8e\xcf\x53\x0a\x9d\x39\xc5\x02\x85\x5c\x77\x47\xe8\xf2\x3e\x0e\xca\x05\xc2\xfe\xaf\x76\xae\x7b\xb9\x5c\x0e\xc3\x90\xa8\x40\x36\xd1\xa6\x5a\x36\x23\xd0\x2e\xdf\xaf\x2f\xae\x36\xe9\x55\xfc\x22\x79\x16\x42\x3e\x4a\x43\xd6\xc2\xd0\x3f\x3d\x1b\x2a\x90\x1d\xa1\xba\xae\xe1\x5c\x65\x0d\xa1\x51\x03\xb4\x81\xaa\x0c\x51\x01\xa7\x3d\xdf\xc1\xb0\x63\xa9\x16\xb0\xba\x74\x83\x32\x14\xcd\x51\xb0\x75\x86\xb3\xde\x3d\x28\xd6\x89\x1d\xdb\x07\x00\x2d\x50\x82\xd9\x2a\xc5\x3a\x9d\xe1\xf7\x55\xba\x4e\x17\xd1\x1c\x9f\xd6\xbb\xb7\xdb\x8f\x3b\x7c\x5a\x5d\x5f\xaf\x36\xbb\xf5\x55\x8a\xed\x35\x2e\xb6\x9b\xcb\xf5\x6e\xbd\xdd\xa4\xd8\xbe\xc1\x6a\xf3\x19\xef\xd6\x9b\xcb\x05\x88\x5d\x4d\x06\x74\xe8\x8c\xe7\xaf\x0d\xd8\x97\x91\x0a\x5f\xb3\x94\xe8\x01\x81\x52\x8f\x84\x6c\x47\x39\x97\x9c\xa3\x51\x52\xf5\xaa\x22\x54\xfa\x96\x8c\xb0\x54\xe8\xc8\xb4\x6c\xbd\x98\x16\x4a\x8a\x68\x8e\x86\x5b\x76\xca\x85\x9d\x1f\x1e\x95\x44\x91\x25\x87\x58\x83\x8c\xa1\x03\xbb\xd3\x52\x74\x2f\x96\xee\x96\x1d\x77\x54\x2a\x6e\xa2\x28\xd7\x52\x72\xd5\x1b\xba\xe9\x8c\x3e\x1c\x9f\x3c\xc5\xd7\xc8\x0b\xd6\xee\x0b\x36\x88\x3b\x2c\xc9\xe5\x4b\x7b\xb4\x8e\xda\x62\xfa\xbf\xcc\xb5\x78\xd9\xc9\x14\x89\x25\x73\xcb\x39\x25\xc5\x32\x84\xe5\xca\xe1\xd5\xab\xab\xed\x1b\xfc\xfa\xbf\x03\xbd\x37\xe2\x70\x7b\xe2\xd9\x44\x7f\xa5\xe3\xd1\xdf\xd1\x95\xdc\xb2\xd1\xd2\x92\xb8\xd7\xb3\xb7\xbb\xdd\x87\x9b\x0f\xd7\xdb\x3f\x3f\xbf\x7e\xfc\xf5\xbc\x78\x19\x7f\x9b\xfd\x08\x4c\x1f\x20\xd3\xff\x82\x6e\xb6\x77\xb8\xd3\xe7\x08\xda\xbe\x89\xbe\x45\x51\xc9\x87\x9b\x56\xf7\xe2\xee\xca\x32\x87\x50\x4e\xd6\x2a\x73\x84\x96\xe6\x88\xa1\x26\xf1\x4d\x64\xc4\xc6\x86\x5a\xd5\x79\x5f\x91\x78\x9f\x06\x4f\x79\x71\x6a\x6d\xdd\x02\x59\xef\x50\x2b\xd3\x7a\x67\x4f\xb9\x7c\x9b\x67\x2c\x3e\x59\x98\x06\xe1\x2e\xd8\x5a\xf7\x4d\x81\x8c\xa0\x07\x19\xcd\x6f\xb4\x76\x5e\x7f\xd4\xea\x76\x34\x91\x25\xd7\x73\x81\x8c\xdd\x58\xf7\x5a\x0f\x12\x60\x2f\x03\xf6\x9c\x6e\x3a\x6e\x75\x81\xd8\xde\xdf\x9f\x38\x4c\x0a\xc5\x2c\xb1\x8a\xef\xf4\x39\x91\x08\xf7\x19\x52\xc5\xf8\x5a\xaf\xe3\x14\xe7\x55\xb3\x53\x4b\x97\xbe\x11\xc9\xee\x9d\xee\x42\x5f\x0f\xbc\xe7\x65\x3a\x35\xe2\x9d\x07\x2e\x4e\xc9\xd7\xe2\xc8\x94\x2a\xa7\xe5\x29\x97\x1e\xe8\x96\xcc\x02\x03\x41\x7c\x4f\xeb\xd0\x46\xae\x66\xa9\x2c\x4a\xa3\x5b\x7c\x29\x74\xbe\x27\x03\xd3\x0b\xe2\xb8\x33\x7c\xcb\x0d\x55\x54\x7c\x41\x92\x24\x53\x1a\x5f\xa0\x71\x96\x35\xaa\x82\x6a\xac\x46\xad\xba\x8e\xc4\xfa\x21\xd1\xaa\x3d\x85\x07\xc0\x0c\x0b\xa8\x56\x4b\x65\xdd\x83\x9b\x46\xe3\x07\x0d\x62\x0d\x43\xe1\x73\x61\xf4\xf8\xec\x93\x1f\x54\x5e\xb3\xd0\x0d\x17\xf7\x4c\x71\x49\x4d\x18\xd4\x5e\x99\x09\x10\x73\x01\x6a\x33\x2a\x0a\x2a\xfc\x80\xf2\x47\xa2\x0b\x02\xb7\xbe\xc5\x3d\xd7\x8a\x84\x8c\xf2\x71\x0a\x42\x03\xb4\x50\x12\x3d\xf2\xb6\x60\xeb\x6d\x74\xb6\x5a\x46\xb9\xf2\x93\x3a\xd3\xae\xc6\xbe\xcf\xfc\x75\x21\xc5\x48\xdf\x0f\x6b\x2d\x24\xce\xa2\xe1\x3d\x61\x20\x2f\x9b\x90\xf3\xd9\x7c\xdc\x3d\x4e\xec\x8b\x2f\xaa\x69\x8e\xbe\x28\x7e\x02\xb2\x54\x3d\xdb\x3a\xb0\xb3\x49\xf4\xc8\xb4\x88\xcb\xb1\xef\xcf\x71\xd1\xa3\x93\x51\xce\x7b\xb1\x37\x61\xe7\x0b\x33\x0f\xba\x74\x86\x62\x16\x76\x28\xf9\xd0\x77\xf6\xdc\x3e\xdf\x15\xee\xfb\xa1\xe3\xe3\x07\xc2\xa0\xc4\x85\x32\xe5\xba\x6d\xfd\xe3\x9e\xd0\xa1\xa3\xdc\x8d\x33\x3e\xa3\x93\x55\x9f\x4e\xcb\x0f\xeb\xcb\xe7\x7e\xde\x83\x0e\x94\x87\x9f\x01\x17\x85\xcf\xd9\xe3\xdf\x66\xd1\xbf\x01\x00\x00\xff\xff\x70\x78\x32\xcc\x61\x07\x00\x00")

func imagesBaseEntrypointBytes() ([]byte, error) {
	return bindataRead(
		_imagesBaseEntrypoint,
		"images/base/entrypoint",
	)
}

func imagesBaseEntrypoint() (*asset, error) {
	bytes, err := imagesBaseEntrypointBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "images/base/entrypoint", size: 0, mode: os.FileMode(438), modTime: time.Unix(0, 0)}
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
	"images/base/Dockerfile":    imagesBaseDockerfile,
	"images/base/clean-install": imagesBaseCleanInstall,
	"images/base/entrypoint":    imagesBaseEntrypoint,
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
	"images": {nil, map[string]*bintree{
		"base": {nil, map[string]*bintree{
			"Dockerfile":    {imagesBaseDockerfile, map[string]*bintree{}},
			"clean-install": {imagesBaseCleanInstall, map[string]*bintree{}},
			"entrypoint":    {imagesBaseEntrypoint, map[string]*bintree{}},
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
