// Code generated by go-bindata.
// sources:
// assets/index.html
// DO NOT EDIT!

package dashboard

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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x59\x5b\x8f\xdb\xba\xf1\x7f\xcf\xa7\xe0\x5f\xff\x02\xf1\x02\x2b\xc9\x97\xbd\xba\xb6\x8b\x9c\x24\x2d\x16\x68\xce\x2e\x9a\xb4\xe8\x79\xa4\xa4\xd1\x8a\x1b\x4a\xd4\x21\x29\x7b\x8d\x85\xbf\x7b\x87\x94\x6c\xcb\xb6\x6c\xcb\x39\x0e\xd0\x87\x26\xc0\x5a\x22\xe7\xc6\x99\x1f\x87\xc3\xd1\xe8\xff\x3e\x3d\x7e\xfc\xf6\xdb\xd3\x67\x92\xe8\x94\x4f\xde\x8d\xcc\x0f\xe1\x34\x7b\x1e\x3b\x90\x39\x93\x77\x84\x8c\x12\xa0\x91\x79\xc0\xc7\x14\x34\x25\x61\x42\xa5\x02\x3d\x76\x0a\x1d\xbb\x77\x4e\x7d\x2a\xd1\x3a\x77\xe1\xf7\x82\x4d\xc7\xce\xbf\xdd\x7f\x7e\x70\x3f\x8a\x34\xa7\x9a\x05\x1c\x1c\x12\x8a\x4c\x43\x86\x7c\x0f\x9f\xc7\x10\x3d\xc3\x06\x67\x46\x53\x18\x3b\x53\x06\xb3\x5c\x48\x5d\x23\x9e\xb1\x48\x27\xe3\x08\xa6\x2c\x04\xd7\xbe\x5c\x12\x96\x31\xcd\x28\x77\x55\x48\x39\x8c\x7b\x28\xa8\x94\xa4\x99\xe6\x30\xf9\xac\x13\x90\xe4\xc3\xd3\x83\x22\x9f\xa8\x4a\x02\x41\x65\x34\xf2\xcb\xb9\x8a\x90\xb3\xec\x3b\x91\xc0\xc7\x8e\xd2\x73\x0e\x2a\x01\x40\x9d\x89\x84\x78\xec\x98\x35\xa8\xa1\xef\x87\x51\xf6\xa2\xbc\x90\x8b\x22\x8a\x39\x95\xe0\x85\x22\xf5\xe9\x0b\x7d\xf5\x39\x0b\x94\x1f\x08\xa1\xd5\x8c\xea\x30\xf1\x07\xde\xc0\xbb\xf1\x15\xa7\x1a\xca\x61\x2d\x69\xee\xa5\x2c\xf3\x42\xa5\x96\xab\x3c\x83\xca\x18\x7d\xe2\xd2\x19\x28\x91\x82\x7f\xe5\x5d\x7b\x5d\x1f\x15\x6c\x0c\xd7\xb4\x96\x6a\x55\x28\x59\xae\x89\x92\x61\x6b\x35\x12\x68\xa8\xfd\xae\xd7\xbb\xf2\x6e\xcb\x17\xef\x05\x05\x8e\xfc\x52\xd6\xe4\x8c\x82\xdd\x48\xa4\x67\x12\x1e\xd0\x00\xb8\x1b\x0a\x09\xfe\xb5\x77\xe7\xf5\x07\x7e\x20\xc5\x4c\x81\xb4\x2e\x39\x8f\x8e\x97\xdf\x0b\x90\x73\xbf\xef\xf5\xbc\x5e\xf5\x72\x46\xe9\x29\xc6\x2f\x33\xbe\x36\x0a\x7a\x5e\x7f\x39\x70\x3e\x0d\x7a\xc6\xb4\x06\xe9\xae\x40\x5a\x41\xf7\x45\x6d\xe1\x76\x5b\xdd\xc8\x5f\x26\x81\x51\x20\xa2\x39\xb1\x10\x1e\x3b\x39\x8d\x22\x96\x3d\x0f\xc9\xcd\x55\xfe\x4a\xba\xe6\xff\x12\xee\x11\x9b\x12\x16\x8d\x9d\x6a\x1f\x1b\x71\x38\xb4\x89\xca\x92\x94\x10\x63\x20\x19\x93\xb8\xc8\x42\xcd\x44\xd6\x29\x24\xbf\x24\xaa\x08\x43\x50\xea\x82\xbc\x55\x54\x84\xfc\xc9\x33\x94\x9d\x37\x9c\x1f\x12\x4b\x14\x51\x4d\xbf\xcd\x73\x18\x92\xf7\x2f\x4a\x64\xef\x2f\x49\x48\xc3\x04\x5f\x63\xca\x15\xac\x84\x0c\x97\x0f\x97\x2b\x59\x84\x80\x94\x42\x0e\xd7\x5a\x5f\x13\x89\x0c\x9a\xea\x42\x5d\x9a\xc9\xba\x66\xf3\x0f\x57\xa2\x04\x07\xcf\xf2\x55\x36\xae\xa9\x3d\x2d\xbe\x6a\x89\xce\xe8\x5c\x5c\xfc\xb9\xc6\xb7\xf0\x02\x96\x45\x1d\x9d\x30\x75\xb1\x1a\x5e\xac\x48\x16\xa5\x3f\x56\xae\xde\x08\xad\xc6\x95\x8d\x1d\x0d\xaf\xba\x44\xb7\xb3\x74\xd8\x94\x4a\x62\xb3\x1c\x14\xe9\x13\x80\x54\xe8\xbd\x7f\xd8\x8d\x1a\xe2\xb6\xd2\xf0\x91\x53\xa5\x3a\x6b\xf3\x31\xc9\x48\xcc\x38\x96\xb4\xb6\xe2\xcd\x15\x5a\xdf\x1a\x3b\xbd\x5c\x8a\x5c\x79\x76\x85\x2b\x5a\xe3\x69\xa4\x27\x96\x00\xb3\xff\x57\x5c\x3a\x74\xde\xcc\xf0\xd0\x86\xc1\xe3\x90\x3d\xeb\x04\x97\xb6\xb1\xe6\x9a\x33\x16\x6b\xef\x3f\x83\x7e\x28\x73\xb8\x95\xb3\xd7\x26\x09\xba\x90\x19\xa9\xd4\x74\x17\x8d\xc2\x10\xe3\xb9\xc8\x10\x64\x9f\x58\xf4\x45\x14\x99\xde\x2b\xce\x1a\x5f\x77\x46\x67\x23\x56\xca\x58\x85\xfb\x63\x4a\x79\x67\x87\xf4\x92\xd4\x7c\x53\x4d\x34\x2f\x4e\x42\x16\x81\x3c\xb6\xa6\xce\x48\xe5\x34\x23\xa1\x09\xd5\xaf\xf6\xf0\xcb\xe8\x34\xa0\xd2\x35\xf1\xc6\xed\xc2\xea\x53\x31\x45\x44\xbb\x52\xd9\x6d\xc9\x26\xe4\xad\x0c\x83\xf1\x9d\x67\x7c\xb3\x20\xb9\xb1\x11\x71\x84\x32\x27\x75\xb3\xde\xad\x11\xd7\x00\x9e\x5f\xb8\x08\xbf\xb7\x44\x4f\x49\xfb\x83\xf0\x09\x0c\x73\x03\x7e\xec\xf8\x90\xd8\x9f\x9f\x03\x9d\x4a\xc3\x5b\x56\xa4\x81\x89\x4a\x17\x03\xc9\x52\x40\xdf\xa5\xb9\x81\xd4\x39\x31\x55\xba\xa8\x1d\xa8\x4a\xda\xb3\xa3\x6a\x23\x65\x59\x88\x4d\x36\x86\xaa\xc1\x93\x70\x67\x10\x16\x50\x05\x4b\xf0\xe5\xa6\xfe\xc3\x45\x75\x6a\x28\xb4\x6e\xf6\x4a\x27\x5f\x92\xde\xcd\xc5\xa2\x0c\xea\x12\x93\x7f\xd8\x88\xd0\x48\x73\xc5\xd2\x86\xea\x6c\x2c\x32\xf6\xba\x6b\xc6\x2a\xc0\x17\x5e\x2c\x45\xfa\xab\x98\x75\x2e\x16\x4d\x86\xec\x8e\x1d\xdf\x3b\x4f\x52\xbc\xce\x8f\x6e\x99\x1f\x08\x96\x39\x32\x6b\xcb\x0e\x05\x77\xf9\xb3\x7b\xe5\xec\xf8\x6e\x93\xf0\xad\x86\x20\x91\x61\x99\x09\xe4\x2f\x04\x8f\xe6\x0c\x38\xb1\x7f\xdd\xea\xf0\x73\xc8\x70\x73\x3c\x82\x98\x16\x5c\x3b\x8b\x6d\x0d\xbb\xc6\x94\x0c\xa6\x08\xc0\x23\x6e\xc7\x22\xcb\x91\x0c\x76\x19\x6c\xd1\xed\x4c\xea\x26\x9a\x42\x7f\x31\x24\xcd\x56\x3f\xda\x27\x6b\xe8\x63\x1c\xdb\x67\x8c\x5b\x32\x68\x30\xb0\x2c\x27\xda\xd9\x6d\x6a\x96\x46\xa3\xcb\x58\x1a\xe8\x28\xf2\x2c\x88\x49\x8a\x9e\xe7\xb5\x53\xd6\x30\xb8\x33\x74\xd1\x06\x4c\x0c\x8e\x67\xe0\x83\x70\x62\x31\xa9\x67\xde\x84\x45\xb0\x5d\xbc\x54\x88\xcb\x0a\xce\xeb\xc5\x49\x1b\x4c\x1e\x81\x9f\x83\x55\x76\x83\x6f\xf7\xa0\xb9\xd7\xdf\x07\x9e\xc9\x13\x9d\x9b\x2d\x4d\xf2\xd2\x25\x27\x84\xbd\x79\xb0\x9d\x99\x25\x02\xec\xed\xf3\xcd\xf9\x1b\x08\xcc\x1f\x54\x0b\x49\x22\x48\x85\xb3\x20\x25\x34\x71\x8f\xc9\x02\x16\xfe\x31\x7e\xac\x00\x01\x9f\x76\xb8\x6d\x35\xba\xcb\x7e\x16\x04\xfd\xcb\x6c\xe1\xff\xe1\x67\x8d\x9f\xa9\x71\x48\x33\x7a\xf6\xe4\x87\xad\x74\xd8\xc4\x77\x72\x66\xb1\x2c\x78\x6c\x72\xa8\x33\x95\x03\xf6\xaf\x8b\xb7\x2d\x96\x43\x54\xbd\x25\x62\x8a\xb8\xd9\x2b\xca\x08\x5b\xf7\x61\xf6\x51\xc8\x43\xd3\x56\xc4\xe4\x2b\x56\x22\x2c\x84\x91\x8f\xcf\xc7\x89\x8b\xc0\x5c\x45\xf0\x48\x6f\x47\xff\x38\x83\xa8\x1d\xe5\xdf\x59\xca\x74\x3b\xd2\x6f\x42\x53\x13\x22\xd6\x56\x34\x55\x1a\xd1\x88\xb7\x73\x75\x54\x03\xce\x1f\xf0\x99\xe1\x3e\xe8\xf3\x91\x36\xf1\xff\x83\x31\x89\x26\x5b\x69\x07\xd5\x1e\x8c\x73\xc9\xd4\x7d\x1d\xdc\x77\x07\xf4\xfa\x1a\xba\x77\xdd\x7e\xdc\xed\xf5\xba\xb7\xb7\x41\x30\x18\x0c\xee\xfa\x77\xfd\xfe\xfd\xe0\x36\xba\xbf\x82\x38\xb8\x8d\xdb\xc9\x1b\xf4\xae\xc8\x5f\x59\x96\xc1\xbc\x1d\xfd\x75\xb7\x7b\x12\x7d\xcf\xeb\x95\xd7\x8e\x96\xe2\x49\xca\xb2\x42\xe3\xd1\x48\x9f\x8f\x7a\xe4\x70\x18\x7f\x6e\x14\xee\xe3\xf0\x26\x86\x38\xba\x8d\x07\x83\x90\xf6\xef\x01\x7a\xb7\x71\x3f\x88\x21\xbc\xbf\xba\xba\xb9\xbf\x86\xeb\xb8\x7f\x1f\xd2\x96\x51\x38\xd1\x4b\xbd\xee\x69\xd4\x27\x09\xc7\xb2\x1b\x66\x78\x7f\x09\x45\x16\x9d\x21\x0c\x38\x7b\x68\xb7\xe0\xb4\x49\x84\x7b\x92\x6f\x73\xb5\xb7\xbf\x0c\xfc\xe1\x82\xed\x68\xf9\xff\x85\xca\xef\xf0\x5f\x7d\xe0\x9e\x7a\xbe\x9d\x5a\xeb\x1f\xaa\xf4\x3f\x3c\x3d\x54\x1e\x6a\x3a\x81\x5b\x15\x6a\x07\x0f\xd6\xdf\x80\x26\x97\x64\x2e\x0a\x32\x63\x2a\xc1\x5a\x9d\x0c\x9f\x7e\x4a\x98\x57\x5f\x0e\x8e\x44\xfa\xd4\x76\x04\x43\x60\xbc\xae\xfa\x93\xa9\x75\xd5\x90\xd8\xf2\xb2\xf1\xda\xcf\x05\x8d\x1e\x2a\x9e\x03\x0d\x88\x75\x47\xe5\x80\x82\xe6\xc6\x82\xd1\xf0\xa5\x22\x3b\x49\x85\x91\xb9\xd6\x50\xd6\xb8\x67\xec\x5d\x34\x41\x25\xa3\xd3\xdd\xa6\x01\xa9\x7a\x07\x31\x7b\x85\xc8\xd5\x22\x6f\x55\x4e\x66\x9a\x62\x6d\x2e\x9b\xcb\xc9\x2d\xea\x4a\x81\xd9\x12\x7b\x38\x90\x87\x36\x70\x04\x92\x66\xd1\xf2\x2b\xcf\xff\x3b\x78\x21\xf8\xc8\x59\xf8\xbd\xba\xb3\xaf\x62\xbb\xd8\xf3\xcd\x8a\x9e\x90\xf2\x6a\x3d\xf7\x52\xb9\xd3\x60\x0f\x56\xd1\x9c\xe6\x0a\x6b\xd2\xea\x61\xdf\x62\x0a\xbe\xc5\xbd\x74\x32\xfe\xec\xaf\x77\x39\x9b\xa0\x1b\x0e\xac\xb6\xc4\xd9\xa2\xb1\xc1\xa3\x12\x91\xe7\x98\x71\xdc\x80\x2a\x24\xaa\x1a\x3d\xf5\x84\x42\x71\x08\x55\xec\x39\x21\x0a\xde\xae\x70\xaf\x16\x22\xd9\x73\xb2\xaf\xd6\x47\xa6\xcd\xf6\x7a\x21\xf9\xd8\xf1\x69\xce\x7c\xa8\xc6\x7d\xdb\x70\x75\x96\xcd\xd1\xf1\x5b\xaf\xdb\xed\x36\xdd\x10\xb7\xe4\x55\x1d\xd7\x06\x81\x06\x5d\x2d\xe5\x9d\xe9\x38\x44\x4f\x1c\x4b\xc6\x87\xf6\xc9\x68\xd9\xbc\x30\xe7\xd7\xb8\xde\x92\xb6\x09\xa2\xf1\xba\x5c\xde\x56\x4f\x60\xa8\x8e\xdb\x1d\x8e\x32\xed\xfc\xe0\x9d\xba\x39\xfb\x97\x4f\x36\xd9\x7f\x7a\xfc\xe2\x95\x49\xab\x33\x5a\x9f\x04\xfe\xe4\x92\x44\x22\x2c\x6c\xe7\x11\x13\xff\x67\x0e\xe6\xf1\x97\xf9\x43\xd4\x79\x5f\x7d\xe8\x7a\xbf\x6c\x5b\x6f\x7e\x3e\x2b\x2b\x1f\x3c\x13\xed\x57\xf6\xff\x04\x00\x00\xff\xff\xc2\xd4\xbf\xaa\x76\x1f\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 8054, mode: os.FileMode(436), modTime: time.Unix(1455135660, 0)}
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
	"index.html": indexHtml,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
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

