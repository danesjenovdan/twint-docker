// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// .docker/templates/docker-compose.tmpl
// .docker/templates/docker-entrypoint.tmpl
// .docker/templates/docker-sync.tmpl
// .docker/templates/dockerfile_alpine.tmpl
// .docker/templates/dockerfile_debian-slim.tmpl
// .docker/templates/dockerfile_ubuntu.tmpl
// .docker/templates/dockerignore.tmpl
// .docker/templates/env.tmpl
// .docker/templates/makefile.tmpl
// .docker/templates/readme.tmpl
// .docker/templates/readme_alpine.tmpl
// .docker/templates/travis.tmpl
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _DockerTemplatesDockerComposeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\x41\x4f\xe3\x3c\x10\xbd\xe7\x57\x8c\xf8\x90\x38\x39\x4d\x84\x3e\x2d\x6b\xa9\x87\xb2\xf4\xc0\xee\xaa\x45\xb4\xe2\x1a\x99\x64\x0a\x56\x1d\xbb\x1a\x3b\xa1\x6c\x94\xff\xbe\x72\x9c\x84\x96\x85\xee\x5e\x9c\xcc\xcc\x9b\xe7\x37\xf3\xcc\x18\x8b\x6a\x24\x2b\x8d\xe6\x70\x71\x19\x7f\xb9\x88\x2c\x52\x2d\x73\xb4\x3c\x8a\x00\xdc\x8b\xd4\x8e\x47\x00\x00\xb2\x14\x4f\xc8\x61\x9f\xd0\xaf\xad\xa9\x27\xa1\xd2\x34\x10\xdf\x48\x82\xb6\x6d\x1a\xb9\x81\xf8\x5a\x58\x6c\x5b\xd6\x34\xa8\x0b\x9f\x0b\x19\x68\xdb\x8e\xe2\xb1\x92\xaa\x08\x6c\x00\xb9\xd1\x0e\xf7\x8e\x43\xdc\x27\x0a\x93\x6f\x91\x36\x52\x21\x87\x9b\xf1\x3f\x1a\xb0\x42\x6a\xa4\x4c\x8b\x12\x79\x90\xc5\x72\x25\x47\x8d\xcc\xa2\xa0\xfc\xf9\x73\xa9\x03\x40\x09\x87\xd6\x9d\x60\x0d\xb8\x0e\xb0\x33\xe4\x6c\xa0\x64\x70\x99\x24\x09\xf7\x47\x17\xa3\xae\x25\x19\x5d\xe2\xb0\x1e\x06\x8b\xe5\xcd\x3c\x9b\x2f\x1e\xa6\x3b\x32\x45\x95\x3b\x69\x74\x27\xcf\x10\xdb\x91\xac\xcd\xfe\x95\x9f\xb8\xd6\x19\xea\x51\x87\x23\xf8\x9d\x88\x27\x9c\x1c\x92\x1c\x4c\x70\x24\xf0\xec\x6b\xf2\x7f\xc2\xfd\x71\x06\xff\xc1\xda\x10\xec\x68\xa0\x0b\xd5\xd4\x57\xd3\xa1\xea\x75\x90\x51\x1d\xc9\x00\xba\x4a\xd3\x2b\xee\x0f\x0f\xba\x0b\x17\x02\xf8\x31\x50\x09\xeb\x64\xfe\xc1\x9a\x83\x6f\x71\x0f\x88\x73\x33\x39\xc2\x1e\x47\xfc\xbc\x99\xff\x9c\xad\xd6\xb7\xdf\xb2\x87\xf9\xfd\xea\x76\xb9\x68\x4f\x2c\xa5\x6f\xfd\x6c\xe3\xda\x14\x18\x7b\xfc\xf4\xe8\x8e\xbe\x9a\xab\xca\x3a\xa4\x58\x6a\xe9\xa4\x50\x59\x29\x7c\x98\xf9\x26\x7b\xb2\xa1\x63\x0c\x53\xb1\x3e\xd7\x23\x1e\x8d\x71\xd6\x91\xd8\xc5\x25\x96\x86\x5e\x33\x65\xf2\xed\xd4\x51\x85\xc3\x02\xe7\xab\xec\xfb\xec\x61\x96\x2d\xef\xd6\xab\xe9\xdb\xac\x63\xae\x3d\xeb\x90\x95\x92\xa5\x1c\xac\x03\x28\xb1\xf4\x4c\x43\x08\x60\xcd\xc6\x71\x60\xe9\x98\x78\x16\x54\x8c\x89\xda\xa8\xaa\xc4\xd1\x78\xb4\x85\x70\x22\x49\xf9\xa4\xb2\x34\xb1\xcf\x82\xf0\x9d\x05\xbe\xfe\xe7\x83\x79\x93\x77\xb7\xbc\x5f\xb7\xfc\x5d\xec\x5d\xdf\xca\x47\xa1\xc5\x5f\xec\x0e\xa0\xfe\xc3\xcf\x9b\x1f\xb7\xd7\xb3\xc5\xec\x5f\xfc\x0d\x2d\x1f\x29\xeb\x39\x06\x61\x87\x61\x14\x1d\xcc\x3f\xce\xde\x75\x16\x24\x6b\x24\x0e\xca\xe4\x42\x45\x91\x46\xf7\x62\x68\xdb\x01\x0b\xdc\x88\x4a\xf5\x4f\x07\xf7\x0e\x49\x0b\x35\x2c\x3c\x68\xd2\x2f\xac\x93\xf5\x3b\x00\x00\xff\xff\x93\x51\xb3\x35\x15\x05\x00\x00")

func DockerTemplatesDockerComposeTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesDockerComposeTmpl,
		".docker/templates/docker-compose.tmpl",
	)
}

func DockerTemplatesDockerComposeTmpl() (*asset, error) {
	bytes, err := DockerTemplatesDockerComposeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/docker-compose.tmpl", size: 1301, mode: os.FileMode(420), modTime: time.Unix(1576920052, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesDockerEntrypointTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xd4\x4f\xca\xcc\xd3\x4f\x4a\x2c\xce\xe0\x52\x71\xe0\x02\x04\x00\x00\xff\xff\xc2\x78\x36\x2c\x0f\x00\x00\x00")

func DockerTemplatesDockerEntrypointTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesDockerEntrypointTmpl,
		".docker/templates/docker-entrypoint.tmpl",
	)
}

func DockerTemplatesDockerEntrypointTmpl() (*asset, error) {
	bytes, err := DockerTemplatesDockerEntrypointTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/docker-entrypoint.tmpl", size: 15, mode: os.FileMode(420), modTime: time.Unix(1576745907, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesDockerSyncTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcd\xb1\x6a\xc3\x40\x0c\xc6\xf1\xdd\x4f\x21\x8e\xc2\x4d\x6e\xdd\x8e\x07\x5d\x4a\x1f\xc0\x5b\x47\x73\xb5\x95\xcb\x41\x90\x0e\x49\xb6\x63\x42\xde\x3d\x24\x86\x38\xe0\x4d\xfc\xfe\x82\x6f\x42\xd1\xcc\x14\xc0\x7d\xb9\x8a\x8b\x65\x26\x0d\x15\xc0\x84\xf2\xcf\x8a\x01\x4c\x46\xac\x74\xa1\xfe\xc1\xb1\x94\x6e\x88\x16\xef\x37\x80\x4a\x1f\xe0\xed\xd2\xfe\xfd\x5e\x3f\x6c\xce\x64\xab\x2e\xd4\x77\x6a\x12\x0d\xd3\x12\xc0\x53\xb4\x3c\x61\xc7\x7a\xf6\x5b\x1f\x15\x25\x0f\x01\x3e\x9b\xa6\xd9\x34\x09\x8f\x65\xcf\x51\x92\xae\x8b\x35\xb8\xba\x08\x1e\x50\x80\x70\x46\x71\x4f\xcd\x89\x58\xf0\xdb\xb7\xd1\x8e\xf0\x9e\xb2\xf9\x7d\xfb\xc1\x13\xcf\x2f\x0f\xb7\x00\x00\x00\xff\xff\x22\x21\x4b\x7c\xfd\x00\x00\x00")

func DockerTemplatesDockerSyncTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesDockerSyncTmpl,
		".docker/templates/docker-sync.tmpl",
	)
}

func DockerTemplatesDockerSyncTmpl() (*asset, error) {
	bytes, err := DockerTemplatesDockerSyncTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/docker-sync.tmpl", size: 253, mode: os.FileMode(420), modTime: time.Unix(1576854421, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesDockerfile_alpineTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\xd1\x4e\x1b\x3b\x10\x7d\xbe\xfe\x8a\x51\xb8\x82\x07\xae\x13\x50\xde\xb8\x97\xab\x46\x21\x6d\xa3\x96\x24\x5a\x42\x2a\x54\xaa\xca\xb1\x27\xd9\x29\xbb\xb6\x65\xcf\x2e\x89\x10\xff\x5e\xed\x6e\x02\x01\x05\xf1\xd0\xa7\xb5\x67\x8e\xec\x73\x8e\xe7\xec\xc7\x64\x7c\x09\x2a\xf3\x64\xf1\xac\xdb\x3e\x3d\x81\xde\x15\xcc\x0b\xca\x8c\x10\xdf\xc6\xc9\x97\x8b\x61\x02\x1d\xe7\xb9\xa3\xbc\x17\xe2\x00\x86\x36\xb2\xca\x32\x98\xac\x39\x75\x16\x94\x35\x80\x2b\xc6\x60\x55\x06\x06\x3d\x5a\x83\x56\x13\xc6\x7f\x80\xac\xce\x0a\x43\x76\x09\x29\x2a\x83\x21\xd6\xe0\x4f\xfd\xbe\x48\xae\x47\xa0\xfc\x1d\x28\x63\x40\x4a\xeb\xa4\x56\x3a\x45\xf0\xf5\x91\xdd\xed\x57\x1a\x2c\xc1\xaf\xbb\xd2\x93\x87\x8c\xe6\x8b\x05\x6d\x3e\x75\x27\x2f\x62\x56\x2f\x96\x5a\xc3\x92\x18\xb4\x92\x1a\x03\xd3\x82\xb4\x62\x8c\xe0\x3c\xda\x79\xa6\xe2\x2b\xf4\xf1\xf1\x0b\x19\xe4\xd1\x96\x35\x23\x4f\xbe\x0b\xb4\xa9\xfb\xa6\x2e\x0e\xa0\x1f\x50\x31\x82\x82\x92\x02\x17\x2a\x03\xb4\x25\x05\x67\x73\xb4\x5c\x2b\x52\x9a\xa9\xac\x20\xc4\xcd\x39\x1b\x19\x32\x87\x12\x6d\xd9\xb8\x57\xad\xc4\x60\x34\x83\x49\x6f\xfa\xf9\xbc\xf5\x54\xeb\xcc\xc9\x9e\xfd\x5d\x15\x5b\x70\x2b\xfe\x9a\x0d\x93\xe9\x75\xef\xeb\xcf\xc1\x68\xb6\x03\x6a\xed\x32\xde\x35\x19\xc8\xb2\x03\x4e\x71\x2f\xb9\x7b\xe2\x74\x57\x60\x6d\x52\xe6\x2c\x82\x94\x06\x3d\xa7\xe7\xa7\x20\xe7\xf0\xf0\xd0\x9e\x61\x88\xe4\xec\xe3\x23\xa4\xcc\x3e\x9e\x75\x3a\x4b\xe2\xb4\x98\xb7\xb5\xcb\x3b\x7c\x4f\x96\x7d\x70\xbf\x50\x73\xb3\x79\x1a\x88\x8a\xf2\xe1\x21\x68\xf3\xba\xf2\xc2\x4b\x29\x0b\xbf\x0c\xca\x60\x55\xde\x07\xd0\xcd\x2c\xed\xe9\xd8\x22\xf7\xeb\x7d\x8d\xb6\x10\xaf\x07\x57\x5c\xf6\x86\xa3\x69\x6f\x38\x1a\x24\xb0\x3a\x09\xab\x3b\x57\xc2\x7f\x9b\xc5\x07\x1f\x1c\x3b\x9b\x2b\xca\x2a\x4d\xff\xef\x3c\x6c\x28\x2c\x53\x8e\x50\x44\x0c\xb5\x4d\xf9\x9d\xa1\x00\xd2\xd7\x9a\x36\x77\x2b\x63\xaa\x3e\xc8\x0b\x68\x1c\x90\xe9\xb3\x64\x19\xa1\x7a\xc6\x4e\x4c\xe1\x56\x40\x05\x8f\xc5\x16\xa6\xe1\x68\xc7\x9d\x7f\x9f\x0f\x37\x8a\xd5\xd1\x7b\x81\xda\x72\x7b\xf1\xe6\xce\x66\xeb\xf7\x22\xb4\x89\xcb\x36\x02\xd5\x3e\xb2\xd1\xcd\xec\x5f\xdd\x13\xeb\x14\xd8\xd5\x9a\x41\x3b\xcb\xb8\x62\x71\x7d\x35\x48\x1a\xda\x7b\x83\xdf\x77\x7e\xfd\xe6\xa8\x2d\x82\xcb\xeb\xa6\x0f\x58\x92\x2b\x22\x50\xae\x96\x28\xfa\xe3\xc9\x0d\x48\x59\xb5\xcf\xeb\x9f\xca\x73\x1a\x76\x72\x21\x0e\xa0\xb7\x0d\xd1\x1b\x37\xfc\x49\x76\x06\xa3\x69\x72\x33\x19\x0f\x47\x53\xf8\xde\xaa\x05\xb6\x7e\xfc\x0e\x00\x00\xff\xff\x16\xea\xa4\x28\xf6\x04\x00\x00")

func DockerTemplatesDockerfile_alpineTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesDockerfile_alpineTmpl,
		".docker/templates/dockerfile_alpine.tmpl",
	)
}

func DockerTemplatesDockerfile_alpineTmpl() (*asset, error) {
	bytes, err := DockerTemplatesDockerfile_alpineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/dockerfile_alpine.tmpl", size: 1270, mode: os.FileMode(420), modTime: time.Unix(1576830973, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesDockerfile_debianSlimTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\x51\xeb\xda\x30\x14\xc5\xdf\xf3\x29\x2e\x0e\x7c\x98\xa4\x19\xf8\x30\x90\x39\x94\xad\x1b\x65\xb3\x95\xac\x2a\x32\xc7\x88\x6d\x96\x66\xb6\x49\x48\x6e\x9d\x22\x7e\xf7\x61\xf5\xcf\x5f\x7d\xca\x39\xe7\x47\xe0\x9c\xfb\x85\x67\x33\x70\x47\xac\xac\x19\x0d\xa3\xf7\x34\xd4\xba\xa1\x01\xbd\xc4\xa2\x22\x64\x36\x4d\xd2\x7c\x9a\xa4\x31\x87\xc3\x3b\x7f\xd8\xd9\x3d\x7c\xb8\x89\x89\xf3\x16\xad\x69\x84\xae\xa3\xc2\x36\x1f\x09\x99\xf2\xaf\x90\xaf\x92\x34\xff\xbd\x8c\xf9\x8f\x24\x4b\xc7\xa7\x53\xb4\x94\x3e\x68\x6b\xce\x67\x42\x3e\x65\xf3\x35\x94\xb6\xd8\x49\x4f\xa5\x41\x7f\x74\x56\x1b\x8c\x42\x05\xec\xc1\x12\xbe\x48\xa1\xa8\x1a\x5b\xc2\xe0\xf0\xcc\x3a\xb8\x21\xc2\x21\x55\x12\xa1\x75\xa5\x40\x09\xfd\xfe\x5d\xa6\x4d\x40\x51\xd7\x40\x8f\xb0\x21\x4a\xe3\xcb\x1f\xa7\xdd\xf0\x15\xd2\xd6\x29\x2f\x4a\x09\x54\x82\xd2\x38\xa8\x10\x5d\x18\x31\xa6\x34\x56\xed\xf6\x32\x89\xe1\x3f\x6d\xd0\x79\xfb\x57\x16\x78\x35\x91\xd2\x38\xb9\x5f\xf5\x46\x2a\x35\xee\xd0\x73\xb3\xa2\x96\xc2\x80\x68\xd1\x5e\x55\x57\xd1\x37\x40\xfd\x1f\x60\x7b\xe1\x59\xad\xb7\x4c\x38\x64\xb5\x0e\x18\xd8\x5b\x60\xd8\xb8\xcb\x73\x61\x9d\x24\x24\x4e\x73\xbe\x9e\x67\x49\x9a\xc3\xcf\xde\xe3\x21\x7a\xbf\xc8\x32\xfb\xbe\x98\xc5\x70\x6d\x46\x56\x19\xff\xf6\x39\xe1\xc0\x82\xdf\xdf\xa2\xff\x01\x00\x00\xff\xff\x15\xde\x8c\x62\xdc\x01\x00\x00")

func DockerTemplatesDockerfile_debianSlimTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesDockerfile_debianSlimTmpl,
		".docker/templates/dockerfile_debian-slim.tmpl",
	)
}

func DockerTemplatesDockerfile_debianSlimTmpl() (*asset, error) {
	bytes, err := DockerTemplatesDockerfile_debianSlimTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/dockerfile_debian-slim.tmpl", size: 476, mode: os.FileMode(420), modTime: time.Unix(1576745957, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesDockerfile_ubuntuTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xdf\x6e\xd3\x30\x18\xc5\xef\xfd\x14\x9f\x86\x34\x01\x53\x62\xa6\x5d\x31\x31\xb4\x0a\x02\x44\xd0\x64\xf2\xd2\x56\x15\x45\xc8\x49\x8c\x63\x48\x6c\xcb\xfe\x5c\x48\xab\x3e\x10\xcf\xc1\x8b\xa1\xa6\x9d\xfa\xe7\xca\xe7\x9c\x9f\x65\x9d\xe3\x0f\x2c\x1f\x43\x28\x83\xc6\x70\x7b\xfd\x3a\xbe\x7e\x45\xc8\x78\x94\x66\xc5\x28\xcd\x12\x06\x8f\xff\xfe\x96\xdc\xa3\x12\x1a\x3e\x99\xb0\x12\x08\xcf\x7b\x63\x56\xca\xc4\x95\xe9\x5e\xc0\x1b\x2f\xf6\xf8\xfe\x10\xbf\x25\x64\xc4\x3e\x42\x31\x4b\xb3\xe2\xfb\x34\x61\x8f\x69\x9e\xdd\xad\xd7\xf1\x54\x38\xaf\x8c\xde\x6c\x08\x79\x97\x3f\xcc\xa1\x36\xd5\x2f\xe1\x22\xa1\xd1\xf5\xd6\x28\x8d\xb1\x6f\x80\x9e\x58\xc2\x26\x19\x54\x4d\x67\x6a\xb8\xfa\x73\xce\x06\xb8\x20\xdc\x62\x24\x05\x42\xb0\x35\x47\x01\x97\x97\x47\x99\xd2\x1e\x79\xdb\x42\xd4\xc3\x82\x48\x85\xb0\x20\xb6\xc7\xc6\xe8\x9b\xc8\x2a\xfb\xf4\x82\x55\xf6\xe6\x70\x35\x0a\x56\x3a\x5e\x0b\x88\x04\x48\x85\x57\x0d\xa2\xf5\xb7\x94\x4a\x85\x4d\x28\xb7\x03\x29\xfe\x56\x1a\xad\x33\x3f\x45\x85\x3b\x13\x4b\x85\xf7\xc7\x1b\x9f\x09\x29\xef\x06\x74\xde\xb3\x6a\x05\xd7\xc0\x03\x9a\x9d\x1a\x0a\xbb\x0e\x22\xf7\x03\xe8\x92\x3b\xda\xaa\x92\x72\x8b\xb4\x55\x1e\x3d\x7d\x09\x14\x3b\xbb\x3d\xb6\x6c\x90\x84\x24\x59\xc1\xe6\x0f\x79\x9a\x15\xf0\xf5\xe2\xf4\x5b\x2e\xbe\x91\x69\xfe\x65\x32\x4e\x60\xd7\x8c\xcc\x72\xf6\xf9\x7d\xca\x80\x7a\xb7\xdc\x47\xff\x03\x00\x00\xff\xff\x4f\xe2\x8d\x91\xf4\x01\x00\x00")

func DockerTemplatesDockerfile_ubuntuTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesDockerfile_ubuntuTmpl,
		".docker/templates/dockerfile_ubuntu.tmpl",
	)
}

func DockerTemplatesDockerfile_ubuntuTmpl() (*asset, error) {
	bytes, err := DockerTemplatesDockerfile_ubuntuTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/dockerfile_ubuntu.tmpl", size: 500, mode: os.FileMode(420), modTime: time.Unix(1576745927, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesDockerignoreTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xce\xc9\xcc\xe5\x4a\xcc\x29\xc8\xcc\x4b\xe5\xf2\x4d\xcc\x4e\x4d\xcb\xcc\x49\xe5\x4a\xc9\x4f\xce\x4e\x2d\xd2\x4d\xce\xcf\x2d\xc8\x2f\x4e\xd5\xab\xcc\xcd\x41\x17\xd2\x02\x0b\xea\xa5\x67\x96\x80\x09\x7d\x08\xa9\x05\xa5\xb4\xb8\x00\x01\x00\x00\xff\xff\x78\xa9\x74\xfb\x57\x00\x00\x00")

func DockerTemplatesDockerignoreTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesDockerignoreTmpl,
		".docker/templates/dockerignore.tmpl",
	)
}

func DockerTemplatesDockerignoreTmpl() (*asset, error) {
	bytes, err := DockerTemplatesDockerignoreTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/dockerignore.tmpl", size: 87, mode: os.FileMode(420), modTime: time.Unix(1576745854, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesEnvTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\x41\x6b\x83\x30\x18\x86\xef\xfe\x8a\xd0\xee\xaa\xd3\x42\x27\x1d\xe4\xf0\x69\x83\x64\xb5\x2a\x26\xba\xde\xc4\x39\xc7\x02\xd3\x43\x23\x74\x10\xf2\xdf\x87\x2e\x5a\xe9\xed\x7b\x9f\x97\xf7\xe1\xdb\x22\x7e\x13\xfd\x60\xf1\x77\x9a\xf0\xaa\x0c\x59\x55\xe4\x31\x56\x0a\x39\x65\x23\x8b\x3c\x46\x5a\x9b\x8e\x43\x34\xf1\xa3\xb8\x22\xad\x95\x12\x5f\xc8\x09\x6a\xd9\x6a\x6d\x2b\xd5\xf6\x9f\x23\xfb\x27\xf7\x4d\x49\x72\x46\xd3\x64\xb5\x33\x45\x02\x67\xc2\x32\x08\xc9\x54\xa5\xb7\xbe\x5d\x95\x01\x30\x82\x87\xd5\x5b\xf4\x0c\x11\xc1\x4f\xea\x61\xab\x9f\x67\x34\x2e\xf4\xeb\x9c\x38\x44\xb3\x2a\xa2\x47\x7c\xf0\x7d\x93\x0a\x93\xac\x2d\x22\x3f\xb5\x1c\x44\x23\xdb\xfa\xda\x7c\x5b\x24\x06\xc6\x69\xb8\xfc\xeb\x3b\xae\xe3\x2d\x34\x4b\x73\x8e\x0f\x3b\xd7\x5d\xc8\x1b\x94\x50\xa5\x19\x67\x78\x63\x5f\x3a\xb9\xf7\x76\x1d\xb2\x2f\xdd\xef\x78\x6c\x46\xfd\x49\x7c\xd4\x7d\x6d\x9d\x68\x00\x09\x3c\x68\x0d\x9c\xac\xfb\x17\xd7\xfb\x0b\x00\x00\xff\xff\xd4\x69\x13\x70\x83\x01\x00\x00")

func DockerTemplatesEnvTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesEnvTmpl,
		".docker/templates/env.tmpl",
	)
}

func DockerTemplatesEnvTmpl() (*asset, error) {
	bytes, err := DockerTemplatesEnvTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/env.tmpl", size: 387, mode: os.FileMode(420), modTime: time.Unix(1576920292, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesMakefileTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x51\xdf\x0f\xd2\x30\x10\x7e\xee\xfd\x15\x97\x6c\x66\xf8\xb0\x15\x66\x82\xd0\x48\x82\x89\x44\x79\xe0\x47\x94\x68\x7c\xd1\xd4\xad\xb0\x65\xdd\xba\xac\x1d\xa8\xc0\xff\x6e\xda\x4d\xc7\x03\xbe\x5c\x7b\xdf\x7d\xfd\xee\xbb\xde\x7a\xf3\xf6\xfd\x0a\xd9\x02\x7f\x8e\x9b\xdf\x85\x3a\x53\x73\xc9\x2b\x13\xa6\x2a\x29\x44\x03\x9f\x57\x1f\x3f\xad\x77\x5b\xb6\x40\x7f\xa4\x33\x21\x25\x9e\x1a\x51\xe3\xe1\xcb\x7a\x7b\xf8\xde\x17\xf1\x9d\xe3\x1e\x73\x29\xf0\x86\xfc\x52\x60\x70\xad\x9b\xbc\x32\xe8\xc7\xf7\x00\x6f\x98\xb4\x06\xc3\x14\x83\x45\x80\xe1\x11\xe3\x97\x00\x9e\x87\x46\x68\x43\x08\x23\xf6\x8c\xc0\x46\x06\xc4\x34\xad\x70\xd5\xb3\x68\x74\xae\x2a\xc2\x48\x9a\xeb\x5a\xf2\x5f\x7f\x91\x08\xfa\x0b\x03\xb2\x14\x49\xa6\xd0\x1f\xf5\x3e\x3a\xdd\xbc\xe4\x27\x61\x85\x7f\xb4\xb9\x4c\xbb\x14\x79\x95\xa2\xe1\x27\x34\x99\x28\x23\x88\xf6\x1f\x76\xdb\xaf\xac\xab\x81\x8b\x56\xad\x1b\x19\xbb\x77\xa1\x41\xff\xea\xfe\xe6\xce\xfc\x6b\xdf\xe1\x8e\xd1\xc0\xb3\x7a\xcf\x28\xff\x30\xc9\xed\x54\xce\x54\xdd\xea\x2c\xec\x9c\x31\x62\x13\xec\x45\x1c\x36\x38\x1a\x78\x30\x5c\x1f\xbc\xb9\xa7\x4f\x9a\xfe\x8f\xf1\x60\x21\x13\xb2\xb6\xdf\xb2\x77\x9b\x49\x54\x59\xf2\x2a\xd5\x0e\x1e\xfa\xdb\x0c\x6c\x40\x86\x1b\x5e\x08\xbb\x53\x20\x4b\x2d\x52\x0c\x2b\x0c\x34\xfd\xe6\x79\x94\xd6\x01\xfa\x6f\x00\x3c\xcc\x8c\xa9\x35\xa3\x54\x1b\x9e\x14\xea\x2c\x9a\xa3\x54\x97\x28\x51\x25\xe5\x74\x1a\xbf\x7e\x35\x1b\xcf\xe9\x64\x16\x4f\x27\xe3\x39\xbc\xb0\x63\x30\xf8\x13\x00\x00\xff\xff\xa4\xdc\xfd\x9f\x6f\x02\x00\x00")

func DockerTemplatesMakefileTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesMakefileTmpl,
		".docker/templates/makefile.tmpl",
	)
}

func DockerTemplatesMakefileTmpl() (*asset, error) {
	bytes, err := DockerTemplatesMakefileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/makefile.tmpl", size: 623, mode: os.FileMode(420), modTime: time.Unix(1576745814, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesReadmeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\xdb\x6e\x1b\x37\x13\xbe\xdf\xa7\x98\xdf\x0a\x10\x19\x30\x65\xf9\x04\xc8\xc6\x9f\x02\xa9\xdd\xb4\x41\x03\x34\x8d\x55\xf4\xa2\x30\x90\xd9\xdd\x59\x2d\x2b\x2e\xb9\x25\xb9\x92\x55\xc3\x7d\x93\xa2\x37\x7d\xc0\x3e\x42\x31\xe4\x4a\x5a\x39\x71\x1b\xc4\x5b\x5f\x48\xe2\x70\x0e\x1f\xbf\x99\xe1\xc1\x03\xf0\x4b\xa9\xbd\xc8\x4d\x36\x27\x9b\x24\xff\xff\x9f\x10\x30\x35\x97\xe0\x3c\x5a\x0f\x42\x7c\x91\x0c\x06\x30\xc5\x54\x11\x98\x02\x2e\x8d\xf6\xa4\xbd\x4b\x8e\x46\xf0\xd3\x37\x66\x09\xde\x40\xe3\xe8\x66\x38\x28\xcd\x52\x78\x23\x1a\x47\xfb\x09\x00\x4f\xbf\x92\xd6\xf9\x76\xb6\xe0\xdf\xeb\x49\x88\xf3\x97\x96\xd0\x13\xa0\x06\x54\x12\xdd\xcd\x70\x90\x05\x89\x40\x2d\x82\x64\x3f\x44\xb9\x0a\xc8\xc4\xa5\xa9\x6a\x13\x7c\x45\xa8\x22\x8b\x82\x8e\xc3\xaf\x14\x3a\x2f\x33\x47\x68\xb3\x12\x50\xe7\xf0\xad\x4c\x51\xe3\xcd\x70\x40\xdd\x29\x81\x3a\x17\xf3\x30\xd5\xb1\x9e\x1a\x0b\xb5\x35\xb7\xab\x9b\xe1\xc0\x1b\x2b\xc2\xef\xae\xf7\x5b\xca\x1a\x4f\x30\x65\xbe\x20\x33\x55\x85\x3a\x67\xd7\x51\x2e\x22\x8f\xad\x7c\xc7\x0e\xab\x5a\x91\x63\xfa\xba\x56\x51\x2a\x4c\xd1\xb5\x09\x0b\x46\x1f\xd8\xc8\xf9\x3b\x72\xf0\xb2\xf1\xa5\xb1\x2c\xc4\xf8\x2b\x8a\xff\xfa\xe3\xcf\xdf\x43\x4a\xac\x4c\x1b\x2f\xf5\x8c\x39\xec\x0c\xa3\xd6\x75\x69\x96\xb0\x32\x8d\x05\xd7\xd4\xb5\xb1\xfe\x66\x38\x70\x9c\x2e\x96\x89\x56\xb6\xbf\xcd\x3c\xe9\x3c\xe4\x9d\x13\xbf\x4d\x31\x0f\x07\xb0\xc9\x69\x92\xbc\x32\x16\x8a\x76\x88\x33\x3a\xe0\x10\xa0\x89\x72\x36\x48\x1b\xa9\x72\x90\x15\xce\x08\x62\xbe\x46\x49\xf2\xfe\xfd\x7b\x57\x92\x52\x49\x94\x40\xdd\x28\x05\xb7\x63\xfb\xeb\xdc\x2c\x0e\x03\x7d\x17\x77\x77\x30\xba\x92\x16\xee\xef\x59\x3b\x49\x8c\x85\x0f\xec\xd6\x99\x8f\x41\xa2\xde\x80\xc1\x3d\x28\xa8\xad\x5d\x18\xc6\x42\x7f\xb1\xd7\x06\xb7\x8d\x06\xe1\x25\x08\x61\x2b\x10\x0b\x78\x36\xac\x97\xf9\xfe\x21\x93\x7e\x71\x68\x6a\x7f\x88\x75\x1d\x46\x8f\x22\xbc\xbb\x93\x05\x8c\xbe\x44\x47\xf7\xf7\xe2\xee\x8e\x74\xce\xb2\x28\x81\xfb\xfb\x3d\x8e\x0f\xbb\x7f\x91\xc4\xc7\xea\x34\x49\xae\x43\xd3\x31\xe1\x35\x90\xc2\x5d\xa5\x79\xab\xf4\x18\x1d\x4d\x0d\x22\x87\x9d\x4a\x5f\xdb\xac\x39\x82\x4d\x91\xef\xc6\xf2\xc6\xfe\x9b\x5f\x56\xd9\xb8\xf9\x68\x2f\x3c\xee\x21\x70\xbd\x80\x67\x6f\x7f\xbc\x6a\x59\x3c\x74\xb6\x25\x34\xa6\xa5\xeb\xfa\x83\x76\x49\x92\x97\x50\xd0\x12\x9c\xe4\x19\x58\xb7\x0e\x63\x2f\x49\xd5\xa1\xf0\x1a\x9d\x93\x75\x9e\x79\xf2\x25\x41\x8a\x4e\x66\xee\xe2\x29\x90\xda\x4f\xd1\x70\xb9\x5b\x8d\x15\x81\x80\xeb\xcc\x62\x4d\x80\x4a\x85\x30\xd3\x25\x91\x77\x50\x58\x53\x05\xad\xe7\x0e\xbc\xac\x48\x49\x4d\xa3\x7e\x22\x3a\xa8\xa5\x26\xac\x79\xe1\x3b\xe1\xbb\xa1\x19\xca\x83\xf0\xc0\xfb\x00\x4a\x2d\xf5\x6c\xeb\xe1\x09\x98\x76\x71\x5c\x1a\xa5\x28\xf3\x40\x0b\xb2\xab\x08\xe5\xa3\x01\x23\xbc\xa0\x65\x34\x3d\x77\x2d\xea\x9e\xb8\x11\x2b\x42\x0b\xc7\xe3\xa3\xd3\x0e\xa4\x96\x17\x5f\xa2\x87\x25\x59\x02\xcf\x02\xca\x21\xa5\xc2\x58\x0a\xea\x7d\xc5\x77\x52\x67\xc1\xe3\x99\x38\x3a\x16\xc7\xe3\x4f\x81\xf1\xd0\xa6\x27\x2c\x06\x0a\xa9\x68\xe4\x6f\xfd\xb6\x4c\x5a\x0c\xdc\x13\x0e\x17\xc4\xfd\xb2\x56\xea\x37\x6a\xe6\x16\x20\x44\xf8\x7c\x2c\x36\x3a\x40\x60\x8d\x60\xd0\x13\xff\x54\xa1\x54\x20\x44\x5d\x1a\x1d\xda\x83\x4f\xba\x2e\xf3\x95\x9c\x95\x1e\x4a\x06\x10\x75\x74\x53\xa5\x64\x1d\x18\x0b\xd1\x18\xf3\xdc\x92\x73\xf4\x94\x9a\x74\xb0\x77\x65\x34\xaa\x1c\xa6\xb6\xa9\xea\x3d\x10\x62\x41\x56\x16\x92\x72\x10\x70\x25\x5d\xad\x70\xb5\xc6\x95\xae\x60\x33\xc9\x4b\x69\x91\x4e\xdb\xf2\xc0\xd4\x34\x1e\xba\xee\x9e\x00\x6c\xf6\x62\xef\x74\x32\x9a\x4c\xc6\xe3\xd3\xc9\xc1\xf1\xe8\x64\x72\x76\x7e\x72\x7e\x70\x34\xaf\xf6\x3e\x21\x75\xa1\x75\x11\x2c\xe6\xb2\x09\xbb\xf1\xd1\xbc\x02\xb4\xa6\xd1\x39\x20\xd4\x0a\x33\x02\xa9\xe1\x2d\x5a\x19\xf3\x4c\xb7\x7c\x91\xe0\xcd\xa8\x82\xcf\x80\x6c\xfa\x2f\x10\x72\xa0\x4c\x86\xaa\x34\xce\x5f\x9c\x1f\x8f\xb9\x41\xbf\x6b\x7c\xdd\x6c\xfb\xd3\xec\x9e\xc6\xbd\xb6\xc5\xcf\xce\x68\x10\x22\x7e\xfd\x63\x63\x04\x95\x3e\x3b\x83\x2f\x2e\x29\xdf\x43\x7c\xdc\x70\xf3\x94\x11\x70\xbc\xed\xc2\x11\xae\xbf\x7f\x23\x3d\xc1\x5a\xb9\xaf\xe0\x85\x51\xca\x2c\xb9\xb4\xb7\x27\x16\x5f\x15\xbc\x27\xbb\x3e\xa7\x36\x3a\xfd\x06\xe5\xd3\x67\x13\x74\x59\x9a\x07\x81\xdb\xb0\xfd\x05\xc5\x85\xb1\xd2\x93\xeb\xec\xfc\x0f\xee\x06\x18\x03\x97\xe8\x60\xad\x9d\xff\x07\x6b\xe6\xb7\x95\x15\x05\xdf\xa8\xb7\x50\xc2\x30\x84\x97\xba\x30\xb6\x42\x2f\x8d\xe6\xe6\x25\x1b\x0a\x2e\x92\xd1\x13\x98\xda\x1a\xae\xe0\x35\x86\x1f\x1c\x67\xdd\x29\xb3\x3c\x80\xb4\xf1\x40\x45\x41\x99\x97\x0b\x82\x8a\x7c\x69\xc2\x2b\x61\x86\xbe\x24\xfb\x60\xc3\x69\x0b\xa4\x75\x07\xc3\xaf\x83\x92\x83\xdf\x4e\xb8\x81\xa3\xee\xc1\x67\x6c\x30\xaf\x75\xa6\x9a\x9c\xc9\x7a\x47\xb1\x2b\xf6\xfb\xca\x83\x6d\x1d\x6e\x96\xfd\x4b\x23\xb3\xf9\x87\x0b\xe5\xb2\xe0\xdd\x06\xce\x37\x2b\x81\x61\xd8\xfe\x65\x00\x47\x0e\xd6\xae\xf6\x3f\x4e\x47\x7f\x88\x5d\x53\x11\x1c\x8d\xcf\x4f\xc7\x27\x93\xf3\xb3\xc9\xc9\xf8\x6c\x02\x02\xde\x45\x39\x42\xfb\x7e\x08\xff\x05\x60\xce\x36\xf7\x4c\x57\x53\x16\x4f\xaf\x78\xed\x7b\x7d\x35\xda\x5e\xde\xc3\xd3\x35\x3e\x0d\xc3\xeb\x15\x66\xa4\xc9\xf2\xa3\x2c\x5d\x45\x1c\xf1\x99\x98\x21\x57\x1f\x9f\x24\x0c\x44\x79\x88\xd5\x98\x33\x47\xe1\x39\x90\xbc\x21\xcf\xab\xe6\x83\x53\x9b\x25\x5c\xec\xff\x1d\x00\x00\xff\xff\x62\x30\x13\x1d\xa9\x10\x00\x00")

func DockerTemplatesReadmeTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesReadmeTmpl,
		".docker/templates/readme.tmpl",
	)
}

func DockerTemplatesReadmeTmpl() (*asset, error) {
	bytes, err := DockerTemplatesReadmeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/readme.tmpl", size: 4265, mode: os.FileMode(420), modTime: time.Unix(1576923960, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesReadme_alpineTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\xdb\x6e\x1b\x37\x10\x7d\xdf\xaf\x98\x5a\x01\x22\x01\xa6\x2c\xdf\x00\xd9\x68\x0a\xa4\x76\xd3\x06\x0d\xd0\x34\x56\xd1\x87\xc2\x40\x46\xbb\xb3\x5a\x56\x5c\x72\x4b\x72\x25\xab\x86\xfb\x27\x45\x5f\xfa\x81\xfd\x84\x62\xc8\x95\xb4\x72\xec\x36\x88\x37\x7e\x90\xc4\xe1\x5c\x0e\xcf\xcc\xf0\xe2\x1e\xf8\xa5\xd4\x5e\x64\x26\x9d\x93\x4d\x92\x2f\xbf\x10\x02\x26\xe6\x02\x9c\x47\xeb\x41\x88\xaf\x92\x5e\x0f\x26\x38\x55\x04\x26\x87\x0b\xa3\x3d\x69\xef\x92\xc3\x21\xfc\xf2\x9d\x59\x82\x37\x50\x3b\xba\xee\xf7\x0a\xb3\x14\xde\x88\xda\xd1\x20\x01\xe0\xe9\x57\xd2\x3a\xdf\xcc\xe6\xfc\x7b\x3d\x09\x71\xfe\xc2\x12\x7a\x02\xd4\x80\x4a\xa2\xbb\xee\xf7\xd2\x20\x11\xa8\x45\x90\x0c\x42\x94\xcb\x80\x4c\x5c\x98\xb2\x32\xc1\x57\x84\x2a\xd2\x28\x68\x39\xfc\x46\xa1\xf3\x32\x75\x84\x36\x2d\x00\x75\x06\xdf\xcb\x29\x6a\xbc\xee\xf7\xa8\x3d\x25\x50\x67\x62\x1e\xa6\x5a\xd6\x13\x63\xa1\xb2\xe6\x66\x75\xdd\xef\x79\x63\x45\xf8\xdd\xf6\x7e\x43\x69\xed\x09\x26\xcc\x17\xa4\xa6\x2c\x51\x67\xec\x3a\xca\x45\xe4\xb1\x91\xef\xd8\x61\x59\x29\x72\x4c\x5f\xdb\x2a\x4a\x85\xc9\xdb\x36\x61\xc1\xe8\x03\x1b\x19\x7f\x47\x0e\x5e\xd6\xbe\x30\x96\x85\x18\x7f\x45\xf1\x3f\x7f\xfd\xfd\x67\x48\x89\x95\xd3\xda\x4b\x3d\xbb\xee\xf7\x58\x26\xd2\x96\x2c\xaa\x5e\x15\x66\x09\x2b\x53\x5b\x70\x75\x55\x19\xeb\xaf\xfb\x3d\xc7\x39\x63\x99\x68\x64\x83\x6d\xfa\x49\x67\x21\xf9\x9c\xfd\x6d\x9e\x79\xd8\x83\x4d\x62\x93\xe4\x95\xb1\x90\x37\x43\x9c\xd1\x3e\x87\x00\x4d\x94\xb1\xc1\xb4\x96\x2a\x03\x59\xe2\x8c\x20\x26\x6d\x98\x24\xef\xdf\xbf\x77\x05\x29\x95\x44\x09\x54\xb5\x52\x70\x33\xb2\xbf\xcf\xcd\xe2\x20\x70\x78\x7e\x7b\x0b\xc3\x4b\x69\xe1\xee\x8e\xb5\x93\xc4\x58\xf8\xc0\x6e\x9d\xfe\x18\x24\xea\xf5\x18\xdc\xbd\xaa\xda\xda\x85\x61\xac\xf6\x17\x7b\x4d\x70\x5b\x6b\x10\x5e\x82\x10\xb6\x04\xb1\x80\x67\xfd\x6a\x99\x0d\x0e\x98\xf9\xf3\x03\x53\xf9\x03\xac\xaa\x30\x7a\x14\xe1\xed\xad\xcc\x61\xf8\x35\x3a\xba\xbb\x13\xb7\xb7\xa4\x33\x96\x45\x09\xdc\xdd\xed\x71\x7c\xd8\xfd\x8b\x24\x3e\x56\xac\x49\x72\x15\x3a\x8f\x09\xaf\x80\x14\xee\x2a\xcd\x1b\xa5\xc7\xe8\xa8\x2b\x10\x19\xec\x94\xfb\xda\x66\xcd\x11\x6c\x2a\x7d\x37\x96\x37\xf6\xff\xfc\xb2\xca\xc6\xcd\x83\x0d\xf1\xb8\x87\xc0\xf5\x02\x9e\xbd\xfd\xf9\xb2\x61\x71\x97\xe1\x20\x6b\x7b\xff\xa0\x6d\x92\xe4\x25\xe4\xb4\x04\x27\x79\x06\xd6\x2d\xc4\xf0\x0b\x52\x55\xa8\xbd\x5a\x67\x64\x9d\x67\xaa\x7c\x41\x30\x45\x27\x53\x77\xfe\x44\x54\xcd\xa7\xa8\xb9\xe8\xad\xc6\x92\x40\xc0\x55\x6a\xb1\x22\x40\xa5\x42\xa4\xc9\x92\xc8\x3b\xc8\xad\x29\x83\xd6\x73\x07\x5e\x96\xa4\xa4\xa6\x61\x67\x41\x1d\x54\x52\x13\x56\xbc\xfc\x1d\x04\xed\xe8\x8c\xe6\x1e\x02\xe0\x0d\x01\xa5\x96\x7a\xb6\xf5\xf0\x34\x58\xbb\x50\x2e\x8c\x52\x94\x7a\xa0\x05\xd9\x55\x44\xf3\x60\xcc\x88\x30\x68\x19\x4d\xcf\x5d\x03\xbc\x3b\x86\xc4\x8a\xd0\xc2\xd1\xe8\xf0\xa4\x85\xaa\x61\xc7\x17\xe8\x61\x49\x96\xc0\xb3\x80\x32\x98\x52\x6e\x2c\x05\xf5\x0e\x21\x38\xa9\xd3\xe0\xf4\x54\x1c\x1e\x89\xa3\xd1\xc7\x20\xb9\x6f\xd3\x1d\x1c\x03\xb9\x54\x34\xf4\x37\x7e\x5b\x32\x0d\x0c\xee\x12\x87\x0b\xe2\x0e\x5a\x2b\x75\x1e\x38\x75\x0b\x10\x22\x7c\x3e\x16\x1e\x1d\x20\xb0\x46\x30\xe8\x2e\x11\x54\xa2\x54\x20\x44\x55\x18\x1d\x1a\x86\x0f\xc1\x76\x0a\x4a\x39\x2b\x3c\x14\x8c\x21\xea\xe8\xba\x9c\x92\x75\x60\x2c\x44\x63\xcc\x32\x4b\xce\xd1\x13\x4b\xd4\xc1\xde\xa5\xd1\xa8\x32\x98\xd8\xba\xac\xf6\x40\x88\x05\x59\x99\x4b\xca\x40\xc0\xa5\x74\x95\xc2\xd5\x1a\xda\x74\x05\x9b\x49\x5e\x4d\x03\x76\xd2\x94\x0a\x4e\x4d\xed\xa1\xed\xee\x69\xd8\x66\x2f\xf6\x4e\xc6\xc3\xf1\x78\x34\x3a\x19\xef\x1f\x0d\x8f\xc7\xa7\x67\xc7\x67\xfb\x87\xf3\x72\xef\x23\x72\x18\xfa\x19\xc1\x62\x26\xeb\xb0\x57\x1f\xce\x4b\x40\x6b\x6a\x9d\x01\x42\xa5\x30\x25\x90\x1a\xde\xa2\x95\x31\xe1\x74\xc3\x37\x0d\xde\xa4\x4a\xf8\x34\xd4\xe6\xb3\x14\x0b\x39\x50\x26\x45\x55\x18\xe7\xcf\xcf\x8e\x46\xdc\xb5\x3f\xd4\xbe\xaa\xb7\x4d\x6b\x76\x0f\xed\xae\x1b\xe5\x57\x67\x34\x08\x11\xbf\xfe\xb3\x55\x82\x4a\xc7\xbd\xc2\xd3\x53\xbe\xb4\xf8\xb8\x23\x67\x53\x06\xc1\x21\xb7\xcb\x47\xb8\xfa\xf1\x8d\xf4\x04\x6b\xe5\x0e\xe3\xe7\x46\x29\xb3\xe4\x62\xdf\x1e\x6c\x7c\xb5\xf0\x9e\xec\xfa\x38\xdb\xe8\x74\x1e\x97\x0f\xa9\x4d\xdc\x65\x61\xee\xc5\x6e\x22\x77\x1a\x17\x17\xc6\x4a\x4f\xae\x75\x3a\xdc\xbb\x4b\x60\x8c\x5d\xa0\x83\xb5\x76\xf6\x79\x56\xce\x2f\x33\x2b\x72\xbe\x8a\x6f\xd1\x84\x61\x40\x20\x75\x6e\x6c\x89\x5e\x1a\xcd\x4d\x4d\x36\xd4\x5f\xa4\xa4\x3b\x3c\x95\x35\x5c\xd3\x6b\x18\x3f\x39\xae\x00\xa7\xcc\x72\x1f\xa6\xb5\x07\xca\x73\x4a\xbd\x5c\x10\x94\xe4\x0b\x13\x5e\x18\x33\xf4\x05\xd9\x7b\x7b\x51\x53\x2c\x8d\x3b\xe8\x7f\x1b\x94\x1c\xfc\x71\xcc\x5d\x1d\x75\xf7\x3f\x6d\xef\x79\xad\x53\x55\x67\x4c\xd9\x3b\x8a\x7d\x32\xe8\x30\x21\xb6\xf1\xb9\x59\xfc\x6f\xb5\x4c\xe7\x1f\x2e\x97\x4b\x84\x37\x22\x38\xdb\xac\x07\xfa\xe1\x88\x90\x01\x1f\x39\x58\xbb\x1a\x3c\x4c\x4a\xa7\xa0\x5d\x5d\x12\x1c\x8e\xce\x4e\x46\xc7\xe3\xb3\xd3\xf1\xf1\xe8\x74\x0c\x02\xde\x45\x39\x42\xf3\x08\x09\xff\x4f\x60\xe6\x36\x77\x54\x57\x51\x1a\x0f\xb9\x78\x5f\x7c\x7d\x39\xdc\x5e\xff\xc3\x23\x38\xbe\x2f\xc3\x3b\x18\x66\xa4\xc9\xf2\xcb\x6e\xba\x8a\x38\xe2\x5b\x33\x45\xae\x44\x3e\x6d\x18\x88\xf2\x10\x2b\x33\x63\x9a\xc2\x83\x22\x79\x43\x9e\x17\xce\xe7\xab\x36\x4b\x38\x1f\xfc\x1b\x00\x00\xff\xff\x1c\xf6\x1c\xb1\xf3\x10\x00\x00")

func DockerTemplatesReadme_alpineTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesReadme_alpineTmpl,
		".docker/templates/readme_alpine.tmpl",
	)
}

func DockerTemplatesReadme_alpineTmpl() (*asset, error) {
	bytes, err := DockerTemplatesReadme_alpineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/readme_alpine.tmpl", size: 4339, mode: os.FileMode(420), modTime: time.Unix(1576923983, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DockerTemplatesTravisTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\xcd\x6e\xdb\x3a\x10\x85\xf7\x7c\x8a\x81\x60\x20\x8b\x0b\x8a\x37\x5d\x74\x21\x20\x8b\xa0\x31\x8a\xa0\x4d\x5c\xd8\x68\xba\x29\x60\xd0\xd2\x58\x62\x43\x91\xca\x0c\xa9\x18\x51\xfd\xee\x45\x24\x25\x56\xf3\xd7\xec\x04\x9d\xef\x1c\x72\x86\x33\x7a\x1b\x90\xd6\x9c\x93\x69\x42\x26\x00\x24\x14\x3e\xbf\x46\x02\x53\xeb\x12\x59\x08\x8e\x85\xcf\x80\xf0\x26\x1a\xc2\x42\x88\x0d\x6e\x3d\xe1\xda\x38\x0e\xda\xda\xc1\x72\xcf\x00\xd5\x20\xb7\xa0\x22\x93\xb2\x3e\xd7\x56\x6d\x8c\x53\x43\x98\x64\x6b\xea\xf7\x93\x92\xd1\xb1\xa7\xde\x90\x47\xb2\x20\xbf\x42\x15\x42\xc3\x99\x52\xa5\x09\x55\xdc\xa4\xb9\xaf\xa7\x8e\xbf\xbe\x09\x2d\x6a\x46\x56\x85\xbf\x75\xd6\xeb\x42\x1d\xa7\x1f\x3e\xa6\xc7\xaa\x30\x1c\xd6\xd6\xb8\xb8\x4b\x83\xa6\xb4\xbc\x03\x29\x7d\x0c\x4d\x0c\x30\xf1\x8f\x5a\x7f\x7a\xd0\x04\xbb\x76\xfb\x9a\x9c\x57\xb5\x2f\xe0\xbf\x1d\x1c\x92\x9f\x55\xfc\x0f\x66\x5a\x6b\xdf\x9c\xba\x7d\x85\x7c\xd2\xb0\x77\x38\xc6\xec\x97\x8c\x98\x57\x1e\x8e\xba\x04\x77\x0d\x92\xa9\xd1\x05\x6d\x93\x2c\x50\xc4\xfd\x11\xfc\x1e\x72\x03\x22\x28\x0c\xf9\x18\xa9\x0a\x8d\xb5\x77\xe9\x2f\xf6\x93\xc3\x19\xa9\x35\x39\x3e\x4c\x0d\x21\x07\x4d\xe1\x71\x4c\xa6\x83\x95\x17\x23\xb5\x35\x16\x59\x25\xb3\xab\xf9\x72\x75\xbe\xb8\x4c\x7a\xf5\xfc\xe2\xf4\xf3\xfc\x24\xd9\xfd\x4f\x77\xd7\xbe\x55\xe1\xd6\xb8\x90\xcd\xba\x91\x51\x3f\x95\x92\xfb\x44\x08\x74\x6d\xd6\x75\xa4\x5d\x89\x30\x6b\xb5\x85\xec\x04\xd2\x2b\x24\x36\xde\xf1\x7e\xdf\x47\x8d\x9e\x93\xae\xeb\x91\xf4\xcc\x10\xbc\x29\x29\x6d\x1b\xe3\xf0\x2d\xe2\xbe\x9f\x5d\x87\xae\xd8\xef\x85\xb0\xda\x95\x51\x97\x98\xc1\x46\x73\x25\xc4\xf3\xf5\x19\x5e\xac\x1d\xee\x35\x5d\xab\x4d\x34\xb6\x00\x29\xf9\x26\x6a\xae\x40\x06\x48\x66\x7d\xe9\x09\xa4\x87\xae\x4e\x43\x06\xc7\x03\xf5\x94\x01\x8a\x0e\x64\x30\x87\x18\x59\xbd\xb0\xc6\x93\x3f\xd6\x97\xc6\x81\x8c\x30\x3b\x5b\x7c\xfa\x32\x5f\xae\xbf\xaf\xe6\xcb\xcb\xd3\x8b\x39\xc8\xe6\xf1\xdf\xb7\xd3\xd5\xea\xc7\x62\x79\x36\x35\x36\x91\xab\xc3\x3d\xc4\xf8\xf2\x9c\x8d\xfa\x9f\x00\x00\x00\xff\xff\x8c\x58\x8b\xbc\x4a\x04\x00\x00")

func DockerTemplatesTravisTmplBytes() ([]byte, error) {
	return bindataRead(
		_DockerTemplatesTravisTmpl,
		".docker/templates/travis.tmpl",
	)
}

func DockerTemplatesTravisTmpl() (*asset, error) {
	bytes, err := DockerTemplatesTravisTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".docker/templates/travis.tmpl", size: 1098, mode: os.FileMode(420), modTime: time.Unix(1576921076, 0)}
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
	".docker/templates/docker-compose.tmpl":         DockerTemplatesDockerComposeTmpl,
	".docker/templates/docker-entrypoint.tmpl":      DockerTemplatesDockerEntrypointTmpl,
	".docker/templates/docker-sync.tmpl":            DockerTemplatesDockerSyncTmpl,
	".docker/templates/dockerfile_alpine.tmpl":      DockerTemplatesDockerfile_alpineTmpl,
	".docker/templates/dockerfile_debian-slim.tmpl": DockerTemplatesDockerfile_debianSlimTmpl,
	".docker/templates/dockerfile_ubuntu.tmpl":      DockerTemplatesDockerfile_ubuntuTmpl,
	".docker/templates/dockerignore.tmpl":           DockerTemplatesDockerignoreTmpl,
	".docker/templates/env.tmpl":                    DockerTemplatesEnvTmpl,
	".docker/templates/makefile.tmpl":               DockerTemplatesMakefileTmpl,
	".docker/templates/readme.tmpl":                 DockerTemplatesReadmeTmpl,
	".docker/templates/readme_alpine.tmpl":          DockerTemplatesReadme_alpineTmpl,
	".docker/templates/travis.tmpl":                 DockerTemplatesTravisTmpl,
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
	".docker": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"docker-compose.tmpl":         &bintree{DockerTemplatesDockerComposeTmpl, map[string]*bintree{}},
			"docker-entrypoint.tmpl":      &bintree{DockerTemplatesDockerEntrypointTmpl, map[string]*bintree{}},
			"docker-sync.tmpl":            &bintree{DockerTemplatesDockerSyncTmpl, map[string]*bintree{}},
			"dockerfile_alpine.tmpl":      &bintree{DockerTemplatesDockerfile_alpineTmpl, map[string]*bintree{}},
			"dockerfile_debian-slim.tmpl": &bintree{DockerTemplatesDockerfile_debianSlimTmpl, map[string]*bintree{}},
			"dockerfile_ubuntu.tmpl":      &bintree{DockerTemplatesDockerfile_ubuntuTmpl, map[string]*bintree{}},
			"dockerignore.tmpl":           &bintree{DockerTemplatesDockerignoreTmpl, map[string]*bintree{}},
			"env.tmpl":                    &bintree{DockerTemplatesEnvTmpl, map[string]*bintree{}},
			"makefile.tmpl":               &bintree{DockerTemplatesMakefileTmpl, map[string]*bintree{}},
			"readme.tmpl":                 &bintree{DockerTemplatesReadmeTmpl, map[string]*bintree{}},
			"readme_alpine.tmpl":          &bintree{DockerTemplatesReadme_alpineTmpl, map[string]*bintree{}},
			"travis.tmpl":                 &bintree{DockerTemplatesTravisTmpl, map[string]*bintree{}},
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
