// Code generated by go-bindata.
// sources:
// model/model.swagger.json
// DO NOT EDIT!

package model

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

var _modelModelSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\x5f\x6f\xdb\x38\x12\x7f\xcf\xa7\x18\xe8\xee\x61\x17\x68\x9b\xde\xe2\x70\x0f\x7d\xcb\x5e\xd3\x6d\x1f\x16\xdb\x6b\x72\xf7\x72\x57\x18\xb4\x38\xb6\xb8\xa1\x48\x85\xa4\xec\x0d\x8a\x7c\xf7\x03\x49\xfd\x37\x25\x4b\xb6\x63\x3b\xdb\x2c\xb0\xa8\x4d\x71\xa8\xe1\xcc\x8f\xf3\x97\xce\xb7\x0b\x80\x48\xaf\xc9\x72\x89\x2a\x7a\x07\xd1\x4f\x6f\xde\x46\xaf\xec\x18\x13\x0b\x19\xbd\x03\xfb\x1c\x20\x32\xcc\x70\xb4\xcf\x53\x49\x91\xbf\xc9\x94\x34\xd2\xcd\x03\x88\x56\xa8\x34\x93\xc2\x3e\x2d\x3e\x82\x90\x06\x34\x9a\xe8\x02\xe0\xd1\xad\xa6\xe3\x04\x53\xd4\xd1\x3b\xf8\xaf\x27\x4a\x8c\xc9\xca\x05\xec\x67\x6d\xe7\x7e\x75\x73\x63\x29\x74\xde\x9a\x4c\xb2\x8c\xb3\x98\x18\x26\xc5\xe5\xef\x5a\x8a\x7a\x6e\xa6\x24\xcd\xe3\x91\x73\x89\x49\x74\xbd\xa5\xcb\xd5\xdf\x2e\x17\x5c\xae\x2f\x63\x85\xc4\x60\xf5\xc0\xce\x94\xda\x34\xbe\x03\x44\x32\x43\xe5\xd6\xfc\x44\xed\x3e\xff\xe9\x48\x7e\x51\x24\x4b\x8a\x4d\xb8\x59\x0a\x75\x26\x85\x46\xdd\x22\x06\x88\x7e\x7a\xfb\xb6\x33\x04\x10\x51\xd4\xb1\x62\x99\x29\x64\xd7\x58\xc8\x3d\x76\x22\x23\x1b\x64\x00\xd1\x5f\x15\x2e\x2c\xc5\x5f\x2e\x29\x2e\x98\x60\x76\x05\x7d\xe9\x34\xd3\x60\xec\x4b\xc1\x4c\xd4\x22\x7f\xbc\x08\x7d\x7e\x6c\x6c\x22\x23\x8a\xa4\x68\x50\xd5\x42\xf5\xff\x75\xd8\x17\x24\x75\x88\x98\x4b\xfa\xd0\xe5\x9d\x89\xbe\x27\x0a\xef\x73\xa6\xd0\x4a\xd1\xa8\x1c\x0f\xbe\xe7\xfb\x1c\xb5\x19\xb3\xe5\xaf\x8d\x2d\x1b\xb2\xec\x6e\x36\xfa\xc0\xe5\xfa\x06\xd5\x8a\xc5\x0d\x09\x7e\xbd\x68\x2e\x53\x48\xad\x46\x92\x36\x0a\x49\xda\x44\xd2\x12\x87\x81\x74\xe3\x28\xae\x57\x28\x8c\x3e\x18\x92\x7e\xf0\x7c\x30\xb1\x84\x6a\x9d\x1f\x0f\x02\x2f\x27\x64\xc7\xed\x53\xa3\x6a\x69\xdf\xf4\xc6\x4a\x75\xc6\x68\x18\x5e\xf7\x39\xaa\x21\x7c\x2d\x08\xd7\x5d\x80\x99\x87\xcc\x2d\xaf\x8d\x62\x62\xd9\xdc\xc4\xe3\xab\xd1\x4c\x29\x99\xce\x34\xde\x3f\x19\x57\x9d\xa7\x0b\xa9\x52\x62\x61\x14\xe5\x4c\x98\x7f\xfc\x3d\x3a\x1a\x9e\xbf\x15\xf2\x7f\x9c\x02\xe9\x5f\xd0\x38\x98\xdc\x18\x6b\x55\xcf\xca\x3a\xb6\x58\x3b\x96\x7d\x1c\xc4\xb0\xf5\x49\xd3\x4c\xe4\x00\x82\x8f\x83\x85\xcb\x58\xa6\x29\x33\xd3\xfc\xa5\x27\x39\x2f\x30\x34\x1c\xc6\x67\x25\x63\xd4\x1a\xe9\x0b\x28\x76\x04\x05\x45\x4e\x1e\x26\x61\xe2\x8a\xd2\xf7\x8e\xe8\xac\x50\x71\x45\xe9\x8d\x21\xcb\x3f\xa1\x75\x18\xe1\xdf\x4e\x1e\xca\x95\x98\x28\x54\x70\xd2\x60\xae\xc6\x36\x13\x2b\x79\x37\x2d\x41\xb8\xa2\xf4\x93\xa3\xfa\x90\x8b\xd8\x01\xf4\x05\xe5\x2f\x28\xaf\xa5\xdf\x06\xc7\x79\xc1\x5d\x5b\x6e\xa6\xa2\xdd\x6d\xe1\x05\xe4\x2f\x20\xdf\x94\xfe\x39\xc1\x5a\x5f\x7e\x73\xff\xba\x11\xb2\x26\xed\x38\x76\x5b\x6a\x73\x65\x09\x4a\x4c\xe5\xfc\xcc\x02\xda\x2e\x77\xdf\x25\xee\x4b\xed\x9e\x07\x37\x86\xa5\x28\x73\x33\x4b\xf5\x77\x54\x2c\x08\x9c\xb4\x58\xa6\x19\xc7\xa9\x45\xd6\x82\xc8\x41\xfa\xfa\x0f\x83\x4a\x10\xce\xcf\x2c\x5f\xe8\x61\xf2\xe5\xec\x9d\x9c\x9b\x93\x7b\xc0\x5e\x68\x9c\x89\x43\x5c\x11\x9e\x4f\x8e\xf3\xfe\x63\x89\x5e\x82\xbd\xf3\x3a\x78\x27\x87\xfa\x15\xa5\x25\xda\x1b\x08\x39\x32\xd2\xab\x66\x63\x83\xbd\xba\xdd\x17\xac\xfd\x16\x6c\x66\x0a\x35\x0a\x43\x0a\x44\x56\x07\xa2\xd4\x8a\x9c\xff\x8e\x71\x1d\xed\x45\x99\xb2\x07\xc3\xb0\x0e\xde\xcb\xf9\xad\x13\xd0\xa7\xd9\x26\xf8\xb4\x21\x26\xdf\x38\x3b\x63\x28\x29\x66\x28\x28\x8a\xb8\xcb\x4b\x83\x9e\x28\x45\xda\x20\x88\x98\xc1\xb4\x3b\x7f\x64\x19\xf1\x31\x68\x5e\x46\xa0\x60\x0f\xb9\x96\x87\x6d\x07\x01\x75\x8d\x1c\x8c\x33\xda\x4c\x8a\x22\xc4\x0f\xae\x1a\x4b\x8a\x33\x2e\xe3\x2e\x62\x46\xb3\x15\x13\xce\x51\x4d\xd9\xd3\x45\x67\x91\xba\x25\x3f\x28\x76\xf0\xdd\x6d\x0d\x04\x04\xae\xc1\xb9\x68\x58\x33\x93\x00\x01\x9d\x61\xcc\x16\x2c\x06\x2f\xa4\xa0\x42\x37\xcb\x71\xa7\xd1\xa3\xab\x2f\xcf\x36\x40\x3b\x14\x0d\x37\x62\xe1\x4e\x28\x7c\xce\xaa\xdc\x10\x78\x43\x83\x4e\x08\x5e\x87\x61\x6d\x0d\x94\x95\x4e\xa3\xb6\x45\xc1\xca\x8e\xe4\x44\x2d\x27\x9c\xdd\x8f\xb7\xb7\x9f\xbf\xe0\xfd\x7b\x62\xf2\xf4\x79\x28\xbb\x5f\x5f\x40\x28\xb5\x2a\x2f\x05\x08\x4c\xac\x4a\xbe\x83\xaa\x3f\x03\x65\x57\xd1\xe2\x4e\xf6\xf6\xb7\x8a\x3a\x2c\x79\x2e\x75\xae\xa6\x98\xf2\x9f\xb9\x9c\x0f\x60\x81\x62\x76\x2c\x97\x79\xc6\x08\x0c\x61\xae\xe3\x28\xca\x18\xc3\x47\x19\x60\x24\x98\x04\xc1\xdd\x85\x18\x84\x62\x11\x88\x9f\x06\x8b\x55\x2e\x7a\x48\x29\xf9\x1d\x81\x42\x93\x2b\xa1\x9d\x18\xbc\x9c\x3e\xbd\x07\xb9\x70\xdf\xbd\xb1\xa6\xfd\x46\xba\xaf\x48\xf7\x4c\xc4\xd4\x4e\xfa\x6c\x80\x74\x90\xd8\xaa\x3f\xa8\xac\x4f\xf1\x1e\x12\x9a\x73\x39\xdf\x71\x97\xb1\x14\x06\x85\x99\xed\x1a\xdb\x73\x14\x4b\x93\xec\x16\xb6\x74\x4b\x78\x03\x20\xad\xc4\x04\x89\xe4\xee\x14\x2b\x5c\xa0\x42\x11\xa3\x3d\xb1\x04\xac\x08\x1c\x42\x89\xd6\x32\x66\x0e\xa4\x6e\x4c\x1b\xa9\x42\x48\xdd\x52\xbe\x78\x7e\x78\xfd\x0e\x52\x81\x61\x9d\x41\x4a\xd4\x9d\x06\x22\x00\xff\x60\xda\x30\xb1\x2c\xcc\x17\xd1\x50\x56\x68\x47\xe7\x06\xdb\x0a\x9f\xcf\x0f\x1f\x3a\x8f\x63\xd4\x7a\x91\xf3\x3e\xea\xb9\x94\x1c\x89\xe8\x3b\xae\xe5\xe3\x31\x66\x2d\x14\xf7\x04\x44\xd6\xb1\x0f\x11\x0a\x67\x08\xeb\xda\x48\x94\x8b\x3b\x21\xd7\x62\x56\x07\x5f\xcd\xf8\x39\x8e\x31\x33\xd7\xcc\x24\xa8\x5a\xe3\x59\xc6\x1f\x6e\xe5\xe6\x03\x93\xa0\xb8\x72\x44\x3f\xcb\x56\x0d\xcb\x3f\xb1\x64\xdd\xc1\x2f\xb9\x08\xaf\xd0\x1d\xb5\x5b\x96\x1a\x03\xc3\x73\x26\x5a\xc3\xeb\x62\xb6\x6b\x19\x34\xc6\x13\x22\x28\x6f\x8d\xe8\xbc\xcb\x12\xeb\xbd\x76\x11\xb7\x92\xe4\xa8\x9b\x5c\xb6\xe4\xc3\xf9\x6f\x8b\xd6\x80\x78\x68\x0f\x60\x81\xf8\x5a\x8f\xed\xa7\x56\x00\x4c\x6e\x34\x2c\x0c\xaa\x94\x09\xa7\xa7\x8f\x52\xde\xf5\xd2\x94\xb2\x2a\x1e\x57\xe5\xb1\x88\xe2\x82\x78\xaf\x3b\xa0\xf9\x0d\x83\xd0\x84\x19\xf8\xf2\xeb\x1c\x7d\x00\x63\x91\x06\x44\x50\x98\x63\x42\x56\x4c\xe6\xca\x86\x32\xa4\xb0\x0c\x45\x54\xd3\x17\xec\x6d\x98\xc8\x3d\x4e\xfd\x13\x1c\xbf\xa6\x8a\x3b\x21\x04\x6c\x31\xfe\xdd\xc4\x61\xab\xcd\x65\xa2\x14\x42\xe5\x81\x9d\xe9\x04\xef\x95\xad\x84\xad\x4f\x76\x13\x7c\xbd\x0f\x7e\xa8\xb7\x7c\xb9\x20\x8c\x23\xfd\x31\x24\xe3\xcd\xeb\xed\xfb\xd8\xd6\xfd\x12\xf2\xa9\xa6\x79\xc0\x00\x06\x7e\xa9\x70\x44\x9f\xd1\xcf\xd8\xde\xc1\x26\xa6\x99\x79\x98\x00\xb5\x6b\x3b\x7f\x20\x51\xb5\x51\xda\xe1\xd2\x5e\x54\x4a\xaa\x29\xdc\xd9\xf9\x03\xeb\x79\x8f\xeb\xe9\xc7\xae\x59\xe4\x3f\x8b\x81\x65\x13\x63\xb2\x99\xc2\xfb\x83\x96\x7f\x8a\x45\x75\x36\x79\x55\x9d\x0d\x8b\xc0\x4c\x09\x2c\x5d\x0b\x62\xbc\x81\xf1\x06\x04\x45\x4c\xb2\x5c\xe7\xdc\xd7\x02\x39\x87\x4c\x6a\xcd\xe6\x1c\xbd\x9d\xb1\xd6\x9c\xd8\x10\xef\x01\xe6\xad\x20\xbf\x0c\xe8\x9c\x3d\x67\xc2\xd9\x73\x7b\x5a\x7a\x8d\xba\xaf\x40\xc6\x09\xd2\x9c\x23\xf5\xbf\xf6\xd8\xc7\xae\xef\x11\x92\x19\x96\xe2\x53\x14\x7d\xcd\x8e\x4b\x52\x62\xf0\xb5\x65\x6a\x5c\xe1\xa0\x28\xd8\x26\x44\x5b\xf1\x2b\xab\x8d\xd7\x60\x12\xa6\x8b\x48\xdc\xaa\x42\x21\x27\x86\xad\xd0\xcf\x70\xbe\x57\x00\x5a\xa1\x83\x0d\x83\xaa\xaa\xaf\x90\x14\x81\x69\x50\x18\xcb\x15\x2a\xa4\x01\xcd\x35\x2c\xc9\x08\x85\x55\x7c\xd6\x64\xa0\xca\x5e\x98\xab\x03\xe5\x9c\x83\x54\xe0\x0c\x5a\x6f\x22\xd0\x30\x10\x07\x6e\x9d\x8d\xb2\x4a\xb7\x96\x30\xa8\xe4\x14\xb5\x6e\xdf\x64\x1c\xc6\x5d\xbf\x22\xeb\xb7\xb5\x04\x24\xc0\x19\x53\x7f\xf0\xac\x92\x65\x1c\xe7\xaa\x38\x72\xcc\xaa\xae\x3e\x67\xaf\xea\x60\x8b\x09\xca\x62\x77\x8a\x5d\xc1\x88\xe4\xba\x8a\xb3\xfc\x7a\x36\x56\xb0\xdf\x8a\x0d\x58\xad\x33\xe1\x41\xe8\xc2\xc3\x41\x25\xdc\xb6\x05\xb9\x43\x26\xe1\x3d\xc4\x86\x89\x2f\xee\x32\x6d\x3e\xf0\xa1\x4b\x14\xaa\xfa\x07\x68\xea\x00\xc4\x05\xea\x01\x6a\xbf\x2a\x97\xda\x74\xe2\x7a\xc2\x19\x9d\x95\xfe\xa6\xdd\xe6\x1f\x0c\x90\xdb\x1b\x0a\xa9\xd5\x0a\x0d\x1c\xc8\xd0\xe9\x95\x09\x1f\xdd\x17\x0a\xb1\x39\x3e\x2c\x51\xd8\x08\xba\xab\xde\x80\x32\x3e\x10\xa2\x3f\x55\xe5\xfa\xaa\x37\x77\x52\x4b\x7a\xc8\x62\xdd\x13\x19\xd2\x4e\x4d\xe4\x40\xc5\x5b\xb0\x6b\x95\x45\xeb\x0f\x84\xdc\x54\xa5\x8e\x90\x05\x6d\xab\xee\xc6\x1b\xed\xd3\xba\xc0\xa7\x91\xf0\x2e\x79\xc0\x24\x29\x37\x9c\x5e\x40\xce\xe1\x5f\xe2\xed\x2b\xe2\x5e\x51\x75\x56\x71\xcf\x08\xa5\xcc\xdb\xd3\xcf\xe1\x35\xa1\xff\x84\x8c\xbe\x4c\xd2\x58\x2c\xdc\x0f\x3a\x9b\x7c\xcc\xed\xe7\x70\xb6\xea\x74\x3d\xe3\xe3\xc4\x76\x4e\x5c\xd6\x33\x57\x85\x53\xe7\xb6\xd7\x8c\x73\x10\x12\xb8\x14\x4b\x54\x10\x27\x44\x04\xdb\x40\x5e\xda\xbe\x57\xf4\x22\xeb\x51\xb2\x2e\x3a\x6b\x7d\xc2\xdc\xdf\x48\x6f\x26\x9b\x3b\x76\x69\xf6\x38\x9f\x0d\x52\x17\x36\xce\xca\x5d\x8f\x77\xdc\x9b\xd0\x1a\x58\xde\xd5\x23\x89\x30\x96\x8b\x89\xaf\xb8\x2d\x4b\x99\x62\xb9\xf5\x35\xb5\xcb\x9d\xba\x8f\xb6\x41\x0a\xbe\xc2\x5f\x0d\xd2\x65\xc2\x3a\xa5\xc8\x17\xc8\x74\x83\xaf\xf0\x0e\x9c\x50\x3a\x69\x79\xe7\x10\xae\x2c\xd1\xd6\xa5\x77\x11\x90\x5b\x7e\x8c\x80\xea\x57\x48\xbd\xdb\x1b\x2c\xdd\xc0\x0b\x16\x84\xe8\x59\x7d\x41\x65\x56\xfa\xfd\xf1\x6f\x1a\x88\xbb\x46\xbd\x71\x17\xf9\x0d\x86\xe9\xa3\x5d\x66\xef\x2f\xc8\xcf\xa2\x98\x19\x3e\xa9\xcf\xd2\xd7\xec\x5e\x65\xeb\x2f\x11\x1c\xd3\x7f\x35\x5a\x3f\x9d\x8a\x90\x90\x90\x4a\x85\x45\xb4\xa0\x21\x26\x02\xe6\x08\x29\xa1\xe8\x43\x6a\xa6\x7d\x11\xe1\x7f\xc2\x7d\x76\x41\xc6\x1c\x61\xc1\x38\x97\x6b\xa4\x30\x7f\x00\x52\x06\x22\x76\xf9\x66\xf9\x48\x3c\xb4\xdf\x2c\xe5\x9d\x86\x84\xac\x10\x54\x1e\xba\x4a\xf6\xf1\xf6\xf6\xf3\x47\x24\x14\xd5\x3e\x30\xb9\xc3\x8d\x0a\xf8\xce\x2d\xfa\x1d\xf2\x91\x7a\x13\xb0\x56\x24\xd3\x40\x40\x33\xb1\xe4\x08\x89\x1f\xbd\xc3\x87\xcb\xbe\x82\x96\x25\xfe\x15\x4d\x22\xe9\x7e\x75\x94\xd4\xaf\xd1\x74\x86\xd8\x2a\x69\x58\x5e\x9a\xdf\xb3\x4e\xc9\x23\x6b\x57\x4e\x28\x76\xfb\xa1\x32\xf3\x97\xdc\x9b\x34\xc4\xc4\xc9\xa8\xaa\x48\x87\xbd\xb6\xf0\xfc\xfe\xeb\x92\x48\x7d\x1f\xc0\x3e\x85\x82\x36\x2c\xbb\xaa\x06\xbf\xcf\x85\x1d\x49\xa7\xb4\x50\xb6\xf4\x3c\xbc\xd2\x0f\x70\xdb\x6f\x4b\x9f\xa0\x38\x37\x0d\x9a\x70\xd2\x97\x76\xd1\x35\x66\xf1\x5f\x1b\x32\x87\xed\xf0\x2f\xb5\xd0\x68\xf7\x12\x01\x32\x37\x4b\xc9\xc4\x12\xa4\x02\x26\x62\xe9\xfe\x48\x94\x53\x69\xe1\xc7\x5e\x01\x33\xc5\x99\xb1\xa9\x7c\x21\xb9\x57\x85\xc6\x5d\x8e\x23\x8b\x2e\x35\x38\x25\xf5\x61\xa0\xec\x98\xbc\x80\xa0\x07\x04\xbe\x09\x3c\x8b\x25\xed\xb5\x78\x4c\x18\x5c\xb6\x2e\x68\x0c\x36\x38\xb6\x01\xa2\x50\x49\x69\x12\x07\xd1\x70\xfd\xaf\x7f\x5f\xdf\xdc\xf6\xa1\xa1\x68\x60\xbb\xba\x68\x09\x09\x21\xfb\x21\xd1\x8d\x84\x4f\x54\x4b\x93\x53\x1a\x7f\x27\xbd\x99\x7c\xec\x1f\xf5\x1c\xa1\xa4\x7b\x4e\xf7\xe4\xae\xca\xfb\xcd\x44\x83\xcb\xea\xb6\xdd\x6a\x0e\xe5\x5a\x2f\x95\xfc\x5d\x5a\xa2\x5e\xf0\x55\xbe\x06\xf0\x1a\xa8\x5c\x0b\xff\x27\x0b\xfd\x53\x5d\x76\xb1\x8d\x62\xcb\x65\x4f\xaf\x33\x90\x9b\x9e\x48\x21\x65\x62\x3d\x3b\x97\x32\xff\x38\xdc\x97\x6c\x03\x13\x46\x02\x94\xcc\x5b\x85\x54\x9f\x5d\xa6\x61\x27\x62\x7d\x37\xd4\xdd\x7a\x27\x98\x96\x97\x99\x1a\x4b\xd5\x12\xe8\xd1\x57\x75\xe7\xe3\x98\xaa\xea\x17\x47\x8b\xa7\x76\xef\xbb\x7d\x7b\x5a\x48\x93\xa0\x6a\xdf\xa1\x18\x30\x15\xe5\xed\x8e\xe3\xf5\xc3\x7b\x93\xdd\xc1\xcd\x17\x24\x55\x8e\xd4\xf8\x61\x76\xd9\x91\x76\x59\x77\xeb\x1a\x20\xb8\xd6\x27\xad\xdd\x57\xcf\xee\xf7\x6f\x44\xeb\xee\x5f\x90\x74\xf7\xe5\x90\x76\x9a\xcd\x1b\x0d\xe4\x98\x88\x18\x79\x67\xf0\x8e\xb9\x91\x62\x60\x30\x33\x6a\xbf\x37\x24\xb1\xa2\x5f\xdc\xbc\x47\x59\x09\xaa\xba\x3a\xa9\x5a\x10\xb9\xb0\xff\x3f\x5e\xfc\x3f\x00\x00\xff\xff\x23\x38\xdd\x70\x70\x59\x00\x00")

func modelModelSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_modelModelSwaggerJson,
		"model/model.swagger.json",
	)
}

func modelModelSwaggerJson() (*asset, error) {
	bytes, err := modelModelSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/model.swagger.json", size: 22896, mode: os.FileMode(420), modTime: time.Unix(1511348799, 0)}
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
	"model/model.swagger.json": modelModelSwaggerJson,
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
	"model": &bintree{nil, map[string]*bintree{
		"model.swagger.json": &bintree{modelModelSwaggerJson, map[string]*bintree{}},
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

