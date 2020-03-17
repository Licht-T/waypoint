// Package google Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// data/terraform-cloud-run-0/main.tf
// data/terraform-cloud-run-0/outputs.tf
// data/terraform-cloud-run-0/variables.tf
package google

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _terraformCloudRun0MainTf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\x51\x6e\xac\x30\x0c\xfc\xcf\x29\x46\xf9\x7e\xe2\xa9\x07\xe0\x18\xfd\xaa\x2a\xe4\x0d\x5e\x9a\x36\xd8\x28\x09\x54\x55\xb5\x77\xaf\x92\x0d\xab\x55\xdb\x55\x17\x09\x8c\x8d\x3d\x9e\xcc\x10\x39\xe9\x1a\x1d\xc3\x4e\xaa\x53\xe0\xc1\x05\x5d\xc7\x21\xae\x32\x24\x8e\x9b\x77\x6c\x61\x47\x3e\xd2\x1a\xb2\xc5\xa7\x01\x84\x66\x46\xb9\x7a\x6c\x14\xbb\x92\x1a\x20\xa8\xa3\xec\x55\x5a\x75\x4f\x0d\xb0\x44\x7d\x65\x97\xf7\xfe\x96\x1a\x03\x64\x9e\x97\x40\x99\x2b\x2c\x30\x73\xa6\x91\x32\xb5\x14\x20\x11\xcd\x15\x26\xa1\xbf\x54\x81\x91\xb7\x63\xd0\x77\xf4\x78\x68\xb5\x93\x39\x3f\x6b\x48\x0b\xbb\x4b\xb7\x53\xc9\xe4\x85\x63\xba\x02\xf0\x33\x4d\xdc\xf8\xd4\xf7\x6f\x30\xe5\x3e\x19\x53\xc9\xec\xba\x78\x9a\x87\x45\x83\x77\x1f\x16\x56\x94\xd6\xfc\x72\xd6\xe3\xe0\x65\xf4\x32\x35\xf8\xa8\xa1\x20\xdb\x12\xd3\xff\xb8\x4a\xe7\x65\xd3\x37\x8e\xb6\x9d\x71\x3e\x14\x2e\x3d\x9e\xda\x4e\x4b\x21\x3c\x26\x8e\xc9\xfe\xab\x95\xe7\x7d\xfb\xdf\xce\xdc\xe6\x74\x71\xa3\xda\x74\x6b\xbe\x6b\xbe\xfe\xea\xd6\x7d\x93\xbb\x9b\x40\xfb\x72\xef\x60\xfd\x6d\xca\xbe\x4a\x7f\xa8\x52\xf7\x28\xa1\xfb\x21\x78\x77\x3e\x5a\x77\xd5\x6b\x4e\xe6\x2b\x00\x00\xff\xff\xc5\xb4\x80\x10\xba\x02\x00\x00"

func terraformCloudRun0MainTfBytes() ([]byte, error) {
	return bindataRead(
		_terraformCloudRun0MainTf,
		"terraform-cloud-run-0/main.tf",
	)
}

func terraformCloudRun0MainTf() (*asset, error) {
	bytes, err := terraformCloudRun0MainTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "terraform-cloud-run-0/main.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _terraformCloudRun0OutputsTf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xce\x41\x0a\x02\x31\x0c\x85\xe1\x7d\x4f\x11\xe6\x00\x83\x17\xf0\x24\x22\x25\x4e\x9f\xa5\x10\x1b\x49\x93\xd9\x88\x77\x97\x2e\xdc\xcd\x4a\xf7\x8f\xff\x7b\x1a\xfe\x0c\xa7\x45\x6f\x03\xb6\xa3\xe4\x8a\x0e\x63\x6f\xda\x17\x7a\x25\xa2\x9d\x25\x40\x67\xaa\xaa\x55\x90\x37\xd1\x28\xd9\xa2\xe7\xb9\x6f\x1b\xd6\x82\x3b\x87\xf8\x3a\x9c\x3d\xc6\xe5\x74\x5d\x0f\x5a\xe9\x9d\xd2\x97\x2a\x18\xcd\xfe\x91\x1e\x70\x2e\xec\x3c\xad\x63\x22\x4c\x7e\x7e\x1f\x26\x33\xf5\x09\x00\x00\xff\xff\xe7\x2c\x88\x26\x19\x01\x00\x00"

func terraformCloudRun0OutputsTfBytes() ([]byte, error) {
	return bindataRead(
		_terraformCloudRun0OutputsTf,
		"terraform-cloud-run-0/outputs.tf",
	)
}

func terraformCloudRun0OutputsTf() (*asset, error) {
	bytes, err := terraformCloudRun0OutputsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "terraform-cloud-run-0/outputs.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _terraformCloudRun0VariablesTf = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xcc\x31\x0a\x02\x41\x0c\x05\xd0\x3e\xa7\xf8\xa4\xb7\xf0\x00\x7b\x98\x38\xc6\x25\x92\xcd\x2c\xd9\xac\x20\x32\x77\xb7\xb4\x11\xa6\x7f\xbc\x97\xa4\xc9\xcd\x15\x1c\xb2\x29\xe3\x43\x40\xbd\x77\xc5\x82\xa3\xd2\x62\xa5\x41\xf4\x43\x7b\xf6\xa7\xb6\x9a\x3a\xdb\x64\x9d\x6f\xde\x9b\x94\xf5\xf8\x07\x81\xbb\x3e\xe4\xf4\xc2\x02\x3e\x8f\x4b\xd3\xa8\x14\xbf\x32\x0d\xfa\x06\x00\x00\xff\xff\x53\xca\x0c\x11\xb5\x00\x00\x00"

func terraformCloudRun0VariablesTfBytes() ([]byte, error) {
	return bindataRead(
		_terraformCloudRun0VariablesTf,
		"terraform-cloud-run-0/variables.tf",
	)
}

func terraformCloudRun0VariablesTf() (*asset, error) {
	bytes, err := terraformCloudRun0VariablesTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "terraform-cloud-run-0/variables.tf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"terraform-cloud-run-0/main.tf":      terraformCloudRun0MainTf,
	"terraform-cloud-run-0/outputs.tf":   terraformCloudRun0OutputsTf,
	"terraform-cloud-run-0/variables.tf": terraformCloudRun0VariablesTf,
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
	"terraform-cloud-run-0": &bintree{nil, map[string]*bintree{
		"main.tf":      &bintree{terraformCloudRun0MainTf, map[string]*bintree{}},
		"outputs.tf":   &bintree{terraformCloudRun0OutputsTf, map[string]*bintree{}},
		"variables.tf": &bintree{terraformCloudRun0VariablesTf, map[string]*bintree{}},
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