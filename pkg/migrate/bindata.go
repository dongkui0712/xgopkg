// Code generated by go-bindata. DO NOT EDIT.
// sources:
// README.MD
// README_zh-CN.MD
// migrations/201803041120_add_package.sql
package migrate

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

var _readmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x70\x49\x2c\x49\x4c\x4a\x2c\x4e\x55\x48\xc9\x4f\x2e\xcd\x4d\xcd\x2b\x01\x04\x00\x00\xff\xff\x57\x6d\xe5\x94\x13\x00\x00\x00")

func readmeMdBytes() ([]byte, error) {
	return bindataRead(
		_readmeMd,
		"README.MD",
	)
}

func readmeMd() (*asset, error) {
	bytes, err := readmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README.MD", size: 19, mode: os.FileMode(420), modTime: time.Unix(1520139157, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _readme_zhCnMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x78\x36\x75\xc3\xb3\xde\x75\x4f\x77\x4d\x7e\x36\xad\xfd\xd9\xc2\xc5\x5c\x5c\xca\x48\x62\x2f\xf6\x37\x3e\x5f\xbe\x1b\x10\x00\x00\xff\xff\xc8\xfa\xce\xd5\x25\x00\x00\x00")

func readme_zhCnMdBytes() ([]byte, error) {
	return bindataRead(
		_readme_zhCnMd,
		"README_zh-CN.MD",
	)
}

func readme_zhCnMd() (*asset, error) {
	bytes, err := readme_zhCnMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README_zh-CN.MD", size: 37, mode: os.FileMode(420), modTime: time.Unix(1520139141, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations201803041120_add_packageSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcc\x31\x8b\xc2\x30\x18\x87\xf1\xbd\x9f\xe2\x3f\x26\x70\x43\x39\x38\xb8\xa5\x43\xda\xbe\x6a\x30\x4d\x25\xa6\x43\xa7\x26\xd6\x1a\x4b\x31\x16\xa9\xe0\xc7\x97\x22\xb8\x3f\xbf\xa7\x30\x24\x2c\xc1\x8a\x5c\x11\xdc\x2b\xcc\xdd\xec\xfb\xc9\x87\xc1\x81\x25\x80\x1b\xcf\x0e\xa7\x31\x8c\x71\x61\xbf\x29\x87\xae\x2d\x74\xa3\x14\x44\x63\xeb\x4e\xea\xc2\x50\x45\xda\xfe\xac\xe9\x3c\x85\x2e\xfa\xdb\xe0\xd0\x5f\xfd\x83\xfd\xa5\x1c\x25\x6d\x44\xa3\x3e\x64\x6d\x0e\x46\x56\xc2\xb4\xd8\x53\x0b\xb6\xbe\x79\xc2\x41\x7a\x2b\x35\x65\x32\xc6\x7b\x99\x7f\x49\xb1\x13\xe6\x48\x36\x7b\x2e\x97\xff\x77\x00\x00\x00\xff\xff\xc9\x8c\x4f\xb2\xa5\x00\x00\x00")

func migrations201803041120_add_packageSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations201803041120_add_packageSql,
		"migrations/201803041120_add_package.sql",
	)
}

func migrations201803041120_add_packageSql() (*asset, error) {
	bytes, err := migrations201803041120_add_packageSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/201803041120_add_package.sql", size: 165, mode: os.FileMode(420), modTime: time.Unix(1520135929, 0)}
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
	"README.MD":                               readmeMd,
	"README_zh-CN.MD":                         readme_zhCnMd,
	"migrations/201803041120_add_package.sql": migrations201803041120_add_packageSql,
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
	"README.MD":       {readmeMd, map[string]*bintree{}},
	"README_zh-CN.MD": {readme_zhCnMd, map[string]*bintree{}},
	"migrations": {nil, map[string]*bintree{
		"201803041120_add_package.sql": {migrations201803041120_add_packageSql, map[string]*bintree{}},
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
