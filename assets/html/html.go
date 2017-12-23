// Code generated by go-bindata.
// sources:
// templates/html/spr.html
// DO NOT EDIT!

package html

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

var _templatesHtmlSprHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x4b\x73\xdb\x36\x10\x3e\x4b\xbf\x02\xe5\xa5\x97\x50\x98\x26\x39\xb4\x2e\xc9\x69\xc6\x89\x5b\x4f\x92\xc6\x13\x3b\x7d\x9c\x3a\x2b\x62\x4d\x22\x06\x01\x04\x00\x25\xb3\xa9\xff\x7b\x07\x7c\x09\xd4\xcb\x4a\xa7\xed\x25\x09\x16\xdf\x7e\xfb\x61\xb1\x0b\xae\x92\x7c\xf5\xf2\xdd\xf9\xcd\xef\x57\xaf\x48\xe9\x2a\x91\xcd\x13\xff\x17\x11\x20\x8b\x34\x42\x19\x65\x73\x42\x92\x12\x81\x65\xf3\x19\x21\x89\xe3\x4e\x60\xf6\xf9\x33\x59\x5c\x5f\xbd\x5f\xfc\x0c\x15\x92\x87\x07\xf2\x17\x19\x2c\x97\xac\x5b\xff\x5a\xaa\xaf\x2d\x79\x27\xc9\x05\x37\xd6\x25\xb4\xf3\x6b\x29\x2a\x74\x40\x4a\xe7\x74\x8c\x9f\x6a\xbe\x4a\xa3\x73\x25\x1d\x4a\x17\xdf\x34\x1a\x23\x92\x77\xab\x34\x72\x78\xef\xa8\x17\xf3\x3d\xc9\x4b\x30\x16\x5d\xfa\xe1\xe6\x22\xfe\x36\x0a\x68\x24\x54\x98\x46\x06\x6f\xd1\x18\x34\x81\xb3\x32\xbc\xe0\x32\x3a\x10\xf1\xb7\xf8\xc3\x8b\xf8\x5c\x55\x1a\x1c\x5f\x8a\x30\xe8\xe5\xab\xf4\xbb\x88\xd0\x9d\x10\xa0\xb5\xc0\xb8\x52\x4b\x2e\x30\x5e\xe3\x32\x06\xad\xe3\x1c\x34\x4c\xdd\x1b\xb4\x27\x7b\x5b\x07\xae\xb6\xf1\x12\x4c\x6c\x5d\x33\xa1\x59\x0a\xc8\xef\xf6\x11\xfd\x04\x92\x95\x28\xd8\x85\xe1\x28\x99\x68\xc2\x74\x99\x1a\xf7\xb9\xac\x38\xae\xb5\x32\x2e\x80\xae\x39\x73\x65\xca\x70\xc5\x73\x8c\xdb\xc5\x13\xc2\x25\x77\x1c\x44\x6c\x73\x10\x98\x7e\xf3\x84\x54\x70\xcf\xab\xba\x0a\x0c\x5c\x4e\x0d\xb5\x45\xd3\xae\x7c\x12\x52\xa9\xda\xe8\x9b\xf0\xda\x28\x8d\xc6\x35\x69\xe4\x2f\xf1\x4c\x80\x75\x95\x62\xfc\x96\x23\x0b\xb4\xf8\xc2\x79\x03\xd6\xbd\xed\xb7\xc8\xc3\xc3\x70\x8a\x7d\x54\xaa\x38\x73\xd3\x32\x01\xe3\x78\x2e\xb6\x8e\x3e\x71\xb0\xdc\xe1\x1f\x3e\x19\x81\xd7\xb4\x3e\x8f\x38\xb7\x85\xbb\xa5\x37\x2c\x7d\x6e\x09\x8c\xd5\x7f\x25\x20\x47\xaf\xaf\xdd\x91\xa3\xfd\x5c\xd5\xd2\x99\x66\x73\xb4\x3d\x81\x18\xda\xdc\x70\xed\xb8\x92\x61\x3d\x3a\xbb\xd5\x4b\xe4\xf2\xa5\x0f\x3a\x69\xb8\x23\xb4\xbc\x82\x22\xd4\x4f\x5b\x83\xa5\x6b\x75\x1b\xdb\x4f\xb1\xe6\xf2\x6e\xf1\x51\x17\x07\xef\xce\xad\xb9\x73\x68\xce\x72\x30\x2c\x22\x2b\x10\x35\xa6\x91\xad\xab\x0a\x4c\x73\x28\xec\xe0\xe3\x13\x1f\x84\xfe\x01\x84\x50\xb7\xae\x44\xed\xd3\x64\x1f\xf3\xce\x0d\x82\x53\xe6\x9f\x13\xd4\x46\x4c\x0e\xce\xe8\x49\x49\x1b\xdc\xff\x9f\x9b\x1f\xa2\xfd\x17\xd7\x3f\x70\x7f\x51\x0d\x10\x4f\x24\xb8\xbc\x23\x06\x45\x1a\xb5\x4f\x93\x2d\x11\x5d\x44\x4a\x83\xb7\x69\x44\x73\x6b\x69\x05\xfa\x4f\x94\x8b\x8f\x76\x91\x5b\xdb\x3f\xb3\xdd\x01\x88\x35\x79\x1a\xd1\x8f\xb0\x82\xce\x30\x60\x2b\xee\xf1\x51\x96\xd0\xce\x9e\xcd\x67\xb3\xf9\xec\xb0\x9b\x15\x5c\xeb\xa6\x02\xbd\xc8\x8d\xb2\xb6\x04\x6e\xec\x16\x41\xab\xf6\x71\xb1\xeb\x52\x59\x25\x6f\x7d\xfa\x16\xeb\xf5\x7a\x94\x3c\xff\x52\x67\xab\xcd\xe8\x7c\x50\xf7\xb6\xc3\xd6\x99\x89\xff\xa0\xd2\xee\x8b\x4a\x48\xb2\x54\xac\x21\x0c\x1c\xc4\x5d\x9a\x62\xd0\xfc\x0e\x9b\xae\xdc\xde\xb6\xa6\x17\x57\x97\xaf\xb1\xf1\xf7\xec\x3d\x08\x69\xff\x48\x18\x5f\x91\x5c\x80\xb5\x69\x14\x44\x8c\x08\x67\x13\x43\xdc\xf6\x4a\xd4\x85\x08\xed\x1e\xb7\x55\x45\x3b\x18\x0d\xae\xdc\xa0\xae\xc0\x95\x87\x70\xc6\x7f\xc0\x43\xca\xab\xd6\xd4\x11\x77\xb2\x7b\xcd\x1e\x54\x81\xf6\x39\x61\x7c\xe5\x0b\xae\xdb\x2b\x9f\xee\x8c\x15\x89\xad\x40\x88\xec\x68\xa3\x25\x56\x83\xcc\x76\xda\x2d\xa1\xad\x9d\x74\xca\x90\x91\x65\xb3\x05\x1d\x04\x8e\xd8\x84\x76\xe1\x12\x5a\x3e\x1d\x15\x97\xcf\xb2\xd3\x7a\x70\x10\x0b\x92\x11\xee\xac\x2f\x29\x70\x7c\x85\xe4\xc3\xfb\xcb\x10\xdc\x27\x31\x08\xf6\x6c\x93\x83\x21\x3f\x61\x6e\x0b\x54\x55\x9f\x72\xcf\xe8\x6a\x86\x9b\x2c\xbf\xe9\x2d\x9b\x6b\x11\x4a\x16\xdb\xa0\xc1\xb4\x41\x55\x5c\xee\x61\x7b\xcb\xe5\x40\x38\x45\xee\x72\x7a\xe8\x1e\x5a\xb8\xdf\x47\x0b\xf7\xbb\xb4\x1e\xb9\x87\x16\xee\x27\xb4\xfe\x9b\x94\x94\xcf\xb3\x1f\x51\x55\xe8\x4c\x93\xd0\xf2\x79\x36\x9f\xcd\x67\xfe\x52\x96\xaa\x96\x8c\xcb\x82\x2c\xd5\xbd\x4f\x71\x92\x2b\xb6\x99\x4d\x27\x87\x79\x42\x0e\x08\x0f\x36\x42\x95\x53\x73\x80\x4f\x68\x1b\x84\x0c\xd7\xac\x0d\x97\x39\xd7\x20\x48\x8e\xd2\x19\xc5\xd9\xae\x92\xe0\x92\x36\xbc\x7b\x48\x17\x43\x19\x74\x8d\xd1\x2d\x8e\x95\x46\x37\x41\x8e\x49\xba\x6e\x97\x63\x8a\x12\xe7\xe7\xb2\x7e\x6c\x37\x59\xe2\xca\xec\x42\x40\x91\x50\x57\xb6\x8b\x01\xde\x2f\x7f\xf1\x9f\xf6\x71\x75\xee\x63\x30\x94\x79\x6f\xa2\xce\x04\x4c\x2c\x3b\xaf\x8d\x6f\x9f\x84\x3a\xd6\xae\xc7\x5e\xb0\xfd\xce\xe2\xda\x19\x2e\x0b\x1f\xb1\x3d\xe1\x41\xdc\xa5\xbd\x31\x35\x3e\x86\x79\x2d\xd5\x5a\x76\xa0\x2d\x21\x08\x16\xd9\x3e\xdf\x76\xe3\x71\x19\x1d\xec\xa8\x8a\x01\x72\x50\xc4\x4b\xd4\x06\x73\x70\x7b\x85\x6c\x36\x1f\x15\x13\x40\x8f\x09\x9a\xc0\x0e\x8a\xba\xb4\xd7\xb5\x46\x63\x91\xed\x95\xb5\xd9\x7c\x54\x56\x00\x3d\x26\x6b\x02\x3b\x41\x16\x97\xc5\x11\x16\x2e\x8b\x93\x85\x79\xec\x29\xca\x3a\xdc\x44\xda\x8c\xb4\xbf\x34\x12\xda\x77\xcb\xa4\x05\x83\xe6\xcb\xba\x1e\xbb\xea\x86\x2b\x8e\x7d\x9f\xcd\x67\x7b\x3b\x53\x8f\xb0\x7e\x36\x3a\x0e\x8a\x0d\xac\xc7\x0f\xe2\x09\x70\x6d\xd0\xb9\x66\xf0\xe8\x0f\x10\x3c\x1a\x87\x5e\x90\xf1\xa5\x50\xb5\xc9\xb1\x7d\x89\x37\xcf\x45\x2d\xba\xd8\x82\x67\x09\xf4\x13\x90\xff\xd1\x6c\xcf\x28\xf5\xc8\x45\x38\xd8\x28\x53\xd0\xed\xb9\x20\xfb\x12\x74\x42\x21\x4b\xa8\xe0\x87\x63\x16\xdc\x95\xf5\x72\x91\xab\x2a\x1c\xa9\x62\x4f\x3e\x92\xbd\x47\xad\xc8\xc3\x03\x5d\x0a\xb5\xa4\x15\x58\x87\x86\x4e\x00\x3b\xda\xfe\x4d\xd6\xe0\x0c\x09\x6d\xd3\xb7\x9b\xff\xe0\x9f\xed\xe0\xe7\xe7\xbd\x6c\x9e\xd0\xee\x3f\x5b\xfe\x0e\x00\x00\xff\xff\xa6\x57\xf2\xb2\x7d\x11\x00\x00")

func templatesHtmlSprHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHtmlSprHtml,
		"templates/html/spr.html",
	)
}

func templatesHtmlSprHtml() (*asset, error) {
	bytes, err := templatesHtmlSprHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/html/spr.html", size: 4477, mode: os.FileMode(420), modTime: time.Unix(1514045244, 0)}
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
	"templates/html/spr.html": templatesHtmlSprHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"html": &bintree{nil, map[string]*bintree{
			"spr.html": &bintree{templatesHtmlSprHtml, map[string]*bintree{}},
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

