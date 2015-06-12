// Code generated by go-bindata.
// sources:
// NormativeTypes\nodes\tosca.nodes.BlockStorage
// NormativeTypes\nodes\tosca.nodes.Compute
// NormativeTypes\nodes\tosca.nodes.Container.Application
// NormativeTypes\nodes\tosca.nodes.Container.Runtime
// NormativeTypes\nodes\tosca.nodes.DBMS
// NormativeTypes\nodes\tosca.nodes.Database
// NormativeTypes\nodes\tosca.nodes.LoadBalancer
// NormativeTypes\nodes\tosca.nodes.ObjectStorage
// NormativeTypes\nodes\tosca.nodes.Root
// NormativeTypes\nodes\tosca.nodes.SoftwareComponent
// NormativeTypes\nodes\tosca.nodes.WebApplication
// NormativeTypes\nodes\tosca.nodes.WebServer
// DO NOT EDIT!

package toscalib

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

var _nodesToscaNodesBlockstorage = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\xd0\xb1\x4e\xc4\x30\x0c\x06\xe0\xfd\x9e\xc2\x2f\x70\x91\x58\xb3\x71\x3b\x0b\x3c\x40\x64\x12\x5f\xcf\x22\x8d\x83\xed\x56\x82\xa7\xa7\xa5\x05\x0a\x1b\x5b\x14\x7f\xff\x1f\xc5\x2e\x96\x31\x34\x29\x64\xe1\x52\x25\xbf\x3c\xb9\x28\x0e\x14\x4f\x00\x85\x94\x67\x2a\xe9\xaa\x32\x46\x38\xca\x47\x11\x5f\x40\x57\xe9\xa4\xce\x64\x2b\x07\x30\x7e\xa7\xed\x04\xe0\x6f\x9d\x22\x2c\x91\x8a\x7a\x9e\x1a\x7b\x58\xa7\xfb\x30\x4b\x33\x57\xe4\xe6\xf6\xe5\x01\xce\x30\x28\xa1\x93\x26\xd1\x44\xaf\x13\xd6\x08\x77\xf0\x70\xf9\x04\xb3\xd4\x69\xa4\xc4\xe5\x4f\xbf\x2b\xb7\x61\xbf\xd2\x25\xc4\x4a\x25\xc2\x15\xab\x6d\x6f\x59\xc3\x6e\x37\xf1\xff\x26\x33\x76\x7c\xe6\xca\x3f\x9f\x43\x77\xcc\xb7\x91\x9a\xff\x2e\xda\xf6\x72\xf4\xe1\xfe\x9b\x9e\x3e\x02\x00\x00\xff\xff\xf7\xa4\x4f\xb1\x60\x01\x00\x00")

func nodesToscaNodesBlockstorageBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesBlockstorage,
		"nodes/tosca.nodes.BlockStorage",
	)
}

func nodesToscaNodesBlockstorage() (*asset, error) {
	bytes, err := nodesToscaNodesBlockstorageBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.BlockStorage", size: 352, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _nodesToscaNodesCompute = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x92\xc1\x8e\x9b\x30\x10\x86\xef\x79\x0a\x3f\x40\x8b\x7a\xce\x6d\xb3\xbb\x87\x5e\xb2\x55\x69\x4e\x55\x84\x06\x7b\x12\xac\x18\x8f\x3b\x1e\x12\xf1\xf6\x35\x60\x28\xa9\x2a\x71\xe9\x09\x61\x7d\xff\x37\x9e\x19\x0b\x45\x0d\x85\x27\x83\xb1\x78\xa5\x36\x74\x82\xfb\x9d\x52\x06\xd9\xde\xd1\x54\x17\xa6\x76\xaf\xd6\xd0\x77\x22\x49\x00\x88\xb0\xad\x13\x1d\x07\x5c\xa9\x90\x78\x10\xac\xc0\x18\xc6\x98\x0f\x95\x92\x3e\xe0\x5e\xc5\xc4\xfa\xeb\xc4\x75\xb5\xb3\x7a\x13\xf3\x28\x0f\xe2\xdb\x5f\x40\x0b\x21\xff\xa3\x17\xee\xab\xa8\x1b\x6c\x61\x66\x66\x6a\xba\xad\x01\x81\xe1\x3f\x16\xd9\x55\x1c\xa7\xef\x57\x7f\xa1\xe9\x2a\xc4\xf2\x5f\x0b\x7c\x4b\xc2\x6c\x67\xfc\xd5\x59\xc6\x36\x69\x72\x89\xcf\xca\x91\x06\x57\x45\x21\x86\x2b\xfe\x71\x6a\x08\x50\x5b\x67\xa5\x9f\xcd\xcb\x89\x4d\xf2\x17\x11\xd0\xcd\x20\x5a\x12\xc3\x22\x9e\x97\x72\x48\xee\x5b\x39\x99\x17\x8c\xd1\x81\x58\xf2\xb1\xb1\x61\xc6\xd7\x67\xb3\x1b\xe3\x0f\x5a\x42\xa4\x75\xc7\x8c\x5e\xa7\xc5\xaa\x9f\x5f\x3e\xa9\xd3\xf1\xf0\x71\x3a\xbe\xbd\xbf\x9d\xd5\x6e\x75\x59\x3b\x2f\xbe\xa1\x28\xcf\x43\xfc\x47\x13\xaf\xe4\x05\xac\x47\xce\xe0\x1d\x9c\x35\x55\xa4\x8e\x35\x56\xe3\x10\x53\xb1\x75\x43\x25\x5d\xe4\x01\x8c\xc3\x93\x24\x9f\x9a\x3f\xef\xa6\xad\x98\x40\xd6\x6f\x17\x7c\xcf\x60\xf1\x62\x5a\xeb\x47\x9a\xe2\x66\xea\x23\x20\xa7\xf1\xf8\x6b\xd9\x47\xc1\x76\xc4\x13\xe4\xa0\x76\xb8\x19\x2e\x33\x38\x72\xb5\xf5\x26\x79\x36\x43\xf3\xd3\x39\x24\x7e\x0c\xff\x0e\x00\x00\xff\xff\x77\xc5\xfd\xae\x90\x03\x00\x00")

func nodesToscaNodesComputeBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesCompute,
		"nodes/tosca.nodes.Compute",
	)
}

func nodesToscaNodesCompute() (*asset, error) {
	bytes, err := nodesToscaNodesComputeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.Compute", size: 912, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _nodesToscaNodesContainerApplication = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x8e\x41\x0a\x02\x31\x0c\x45\xf7\x73\x8a\x5c\xc0\x1e\x60\x76\xe2\xc6\x9d\xe0\x05\xa4\xb6\x91\x09\x74\x92\x9a\x44\xc1\xdb\x3b\x55\x66\xa8\x98\x55\xf8\xfc\xf7\xff\x77\xb1\x14\x03\x4b\x46\x0b\x07\x61\x8f\xc4\xa8\x61\x5f\x6b\xa1\x14\x9d\x84\xc7\x01\x20\xa3\xd2\x13\xf3\xe5\xa6\x32\x8f\xd0\x23\x67\x11\x5f\x0c\x8a\xf7\x07\x29\xce\xc8\x6e\x0d\x00\xd8\xc1\x24\xe6\xdf\xbf\x5d\x8a\x35\x5e\xa9\x90\xbf\xd6\x80\x4d\xa1\xbe\x7a\x03\x5a\xfe\x6f\xd7\xbf\x47\xb1\x7c\x36\xda\x44\x75\xf5\xf6\x9a\x85\xe3\x32\x02\xf3\x89\x87\x77\x00\x00\x00\xff\xff\x32\xf2\xb2\xce\xe8\x00\x00\x00")

func nodesToscaNodesContainerApplicationBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesContainerApplication,
		"nodes/tosca.nodes.Container.Application",
	)
}

func nodesToscaNodesContainerApplication() (*asset, error) {
	bytes, err := nodesToscaNodesContainerApplicationBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.Container.Application", size: 232, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _nodesToscaNodesContainerRuntime = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xcd\xcd\x0d\xc2\x30\x0c\x86\xe1\x7b\xa7\xf0\x04\x19\xa0\xd7\x6e\x40\x07\x40\x6e\xf3\x55\x58\x4a\xec\x28\x31\x20\xb6\xa7\xfc\x89\xdc\x7a\xb3\xe4\xe7\xb5\xdd\xda\xca\x41\x2d\xa2\x85\xc9\xd4\x59\x14\x35\x9c\xae\xea\x92\x31\x0e\x44\x11\x55\x6e\x88\xe7\xad\x5a\x1e\xa9\xe7\xb3\x6d\x7e\xe7\x8a\xc9\x72\x31\x85\xfa\xae\x57\x2e\xbc\x48\x12\x17\xb4\x57\x4d\x74\xb1\xe6\x9f\x89\xc8\x1f\x05\xbf\x1b\xbd\xfc\x7f\x7e\xc3\x7d\x9d\x78\x49\x38\xcc\xe6\x2f\x1c\x9e\x01\x00\x00\xff\xff\xf3\x08\xa0\x97\xc6\x00\x00\x00")

func nodesToscaNodesContainerRuntimeBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesContainerRuntime,
		"nodes/tosca.nodes.Container.Runtime",
	)
}

func nodesToscaNodesContainerRuntime() (*asset, error) {
	bytes, err := nodesToscaNodesContainerRuntimeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.Container.Runtime", size: 198, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _nodesToscaNodesDbms = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x91\x31\x6e\xc3\x30\x0c\x45\xf7\x9c\x82\x27\xf0\x01\x3c\x36\x59\x3b\x65\x2c\x0a\x83\xb1\xe8\x84\x80\x42\xaa\x24\x63\xa3\xb7\xaf\x2c\xa7\x45\xd2\x4e\xdd\x0c\x9a\x7c\xef\xe3\x2b\xd4\x47\xec\x44\x13\x79\x77\x78\x79\x3d\xf6\x3b\x80\x44\xc6\x33\xa5\x61\x32\xbd\xf6\xf0\xb8\x71\xd4\x29\x16\x34\xda\xeb\xb5\xa8\x90\x44\xdd\x2e\xa6\x85\x2c\x98\x7c\xbd\x05\x30\xd5\x18\x0a\xba\x2f\x6a\x69\x1b\x01\xc4\x67\xa1\x1e\x3c\x8c\xe5\x7c\x1f\x19\x7d\xdc\xd8\x28\xf5\x30\x61\x76\xba\x4f\xab\x65\x34\x2e\xc1\x2a\x55\x7d\x21\xd0\xf6\x8d\xb9\x71\xe1\x9b\x0b\x93\x5a\xfb\xbd\x66\x06\x27\x9b\x79\xdc\x10\x45\x2d\x9e\xad\x2c\x41\x67\xb2\xff\x69\x57\xcc\x1f\x01\x2c\x9c\x33\x64\xf6\x20\xa9\xbd\xb4\x10\x09\x03\x01\x25\x35\x30\x79\x78\x25\x8e\x58\xf0\xc4\x99\x5b\x29\x95\xde\x24\x17\xf5\x5f\xc1\xb6\x66\x1f\x97\xbb\xbd\x4a\x20\xcb\x4f\xda\x19\x33\xa7\xc1\xf5\x66\x23\x0d\xeb\x55\xe5\xbd\x3d\x3d\xc9\xa1\xfa\x4f\xe8\x04\xef\xbb\xaf\x00\x00\x00\xff\xff\xf8\x4d\x4a\x93\xcc\x01\x00\x00")

func nodesToscaNodesDbmsBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesDbms,
		"nodes/tosca.nodes.DBMS",
	)
}

func nodesToscaNodesDbms() (*asset, error) {
	bytes, err := nodesToscaNodesDbmsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.DBMS", size: 460, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _nodesToscaNodesDatabase = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x52\x4b\x6e\xe3\x30\x0c\xdd\xe7\x14\xbc\xc0\xf8\x00\x5e\x66\x32\xc0\x6c\x06\x03\xb4\x07\x08\x18\x8b\x4e\x08\xc8\xa2\x4a\xd1\x09\x72\xfb\xd2\x8a\x6d\x38\x40\xb3\x68\x77\x02\xf5\x3e\xfc\x3c\x93\xd2\x61\x93\x24\x50\x69\x0e\x68\x78\xc2\x42\xed\x0e\x20\x90\xf2\x95\xc2\xb1\x57\x19\x5a\xd8\xa2\xde\x44\xcc\x01\x59\x25\x93\x1a\x53\x99\xe0\x00\x09\x07\x7a\xbc\x00\xec\x9e\xa9\x85\x62\xca\xe9\x3c\x97\x9c\xd9\x29\x67\x63\x49\x2e\x77\x21\x88\x72\xe6\x0e\x63\xe5\x81\xf4\xb5\x16\xe6\x06\x2a\x27\x8b\xda\xb3\x20\x27\xa3\x33\xe9\x2b\xc5\x89\x50\x1f\x63\xf2\xee\xe3\xdd\xcd\x57\x45\x28\xa4\x57\xee\x08\x6e\x1c\x23\x44\x2e\x46\xc9\xa7\x82\x5e\xb4\x62\xaa\xe6\xe8\xa0\xef\x8c\x20\xf5\xed\x33\x4c\x44\xc0\xae\x93\x31\xd9\x63\xa0\x49\xf7\xb0\x07\x0c\x03\x27\x37\x53\x9c\x90\xb3\x8e\xd2\xc7\xc8\x4a\xa1\x85\x1e\xe3\x32\x2c\x96\x72\x13\x0d\x3f\xb2\x5f\xc8\xd5\x75\xfa\x71\xe7\x6d\x4b\x2f\x7d\xe7\xca\x40\xc9\xe6\x2b\xfe\x82\x8b\x94\x75\xed\x00\x1d\x66\x3c\x71\x64\xbb\x2f\x21\x58\x2b\x7e\xf9\xe6\xb7\x24\x43\x4e\xeb\x4d\x3c\x06\x9e\x91\xe7\xbc\x1c\xf6\xff\xde\xd7\x6f\xa5\x58\x57\x51\x2e\x9c\x17\xd8\xb6\x56\x9a\xbf\xee\x4f\xe1\xff\xb4\xac\xad\xd5\xa3\xa5\xe5\x9c\x47\x4a\x21\x8b\xe7\xe1\x79\x5f\x5f\x74\xf8\x67\x06\xae\xe1\xde\x7d\x06\x00\x00\xff\xff\x07\x78\x7f\xa8\xf2\x02\x00\x00")

func nodesToscaNodesDatabaseBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesDatabase,
		"nodes/tosca.nodes.Database",
	)
}

func nodesToscaNodesDatabase() (*asset, error) {
	bytes, err := nodesToscaNodesDatabaseBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.Database", size: 754, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _nodesToscaNodesLoadbalancer = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x92\x41\x4e\xeb\x40\x0c\x86\xf7\x3d\x85\xa5\xb7\x79\x4f\x7a\x44\xac\xb3\x2c\x2d\x12\x12\x2a\x15\x6a\x57\x08\xa1\xe9\x8c\xdb\x5a\x4c\xc6\x83\xc7\x01\xba\xe3\x1a\x5c\x8f\x93\x30\x49\x86\x12\x58\x20\x76\x91\xfd\xdb\xff\xff\x39\xa3\x9c\xac\xa9\x02\x3b\x4c\xd5\x25\x1b\x37\x35\xde\x04\x8b\x52\x4f\x00\x1c\x0a\x3d\xa2\xbb\xdb\x0a\x37\x35\x8c\x95\xd7\xcc\x9a\x05\x51\x38\xa2\x28\x61\xea\xe4\x00\x7f\x60\x35\x9d\xf5\x5f\xc6\xef\x58\x48\xf7\xcd\xd0\x00\xd0\x43\xc4\x1a\x92\x0a\x85\x5d\x29\x09\x3e\xb4\x24\xe8\x6a\xd8\x1a\x9f\xb0\x54\x93\x1a\x6d\x53\x0d\xf8\x9c\x57\x53\x83\x41\x8d\xcf\x2d\x6b\xa2\xd9\x90\xa7\x4f\x33\xeb\x29\x37\xbf\xee\x1f\x32\x8e\xb5\xd5\x3c\xb8\xc8\x14\xb4\x5a\xb6\x1b\x4f\xb6\xc8\xd9\xda\x56\x04\x33\x69\xb6\xba\x39\xfd\x0f\xeb\xc5\xf4\x6a\xbd\x98\xcd\x67\xb7\x50\x24\x99\xd3\x0a\x45\x25\x0e\x79\xf1\x1e\xe1\xdc\xb3\xd1\x1c\x1f\xfe\x5e\x2c\xff\x15\xfb\xb7\x97\xd7\x04\x1c\xfa\x7e\xec\x0d\x20\xa0\x3e\xb1\xdc\xe7\xc4\x01\x2c\x87\x80\x56\x73\xae\xc9\x91\xb7\x43\x2a\x08\x27\x60\x62\xcc\x33\xa6\x37\x29\xbe\x23\xd6\xc3\x8f\x44\x47\xbd\xa0\xef\x37\xa4\x3d\xc5\x8f\x89\x71\xad\xfb\x5f\xad\x62\x5a\xf1\x71\xe4\x17\xfc\xdf\x2e\x70\x36\xa0\x50\x07\xcb\x19\x19\x81\x05\x1a\x16\x84\x7c\x16\x07\x9b\xe1\xdd\xb8\x31\x51\x9a\xbc\x07\x00\x00\xff\xff\xd9\xff\xac\xaf\x5e\x02\x00\x00")

func nodesToscaNodesLoadbalancerBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesLoadbalancer,
		"nodes/tosca.nodes.LoadBalancer",
	)
}

func nodesToscaNodesLoadbalancer() (*asset, error) {
	bytes, err := nodesToscaNodesLoadbalancerBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.LoadBalancer", size: 606, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _nodesToscaNodesObjectstorage = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x90\x41\x4e\xc3\x30\x10\x45\xf7\x39\xc5\x5c\x20\x16\x6b\x2f\x91\x10\x4b\x24\x38\x80\x35\xb1\x87\x68\x50\xe2\x31\xe3\x09\x02\x4e\xdf\xa4\x4e\xab\xa8\x5d\x77\x67\xf9\xbf\xe7\x2f\x7f\x93\x1a\xd1\x65\x49\x54\xdd\xdb\xf0\x45\xd1\x3e\x4c\x14\x47\xf2\x5d\x07\x90\x48\xf9\x87\x52\xf8\x54\x99\x3d\x1c\xd9\x77\x11\xdb\x88\xa2\x52\x48\x8d\xa9\x9e\x05\x80\x8c\x33\xed\x47\x00\xfb\x2b\xe4\xa1\x9a\x72\x1e\xdb\x5d\xe5\xff\xbb\x38\xe2\x84\xda\x2f\x99\xcd\x6d\xf1\x25\x8d\x92\x57\x13\x39\x5b\xbd\x1a\x00\x3d\x8c\x4a\x68\xa4\x41\x34\xd0\xf7\x82\x93\x87\x27\x78\x7d\x6e\xc4\x8c\xbf\x0f\x6d\x88\x58\x70\xe0\x89\x0f\x1f\xae\x6d\xaf\x40\x39\x15\x59\x9f\xba\xe9\x6e\xa3\x1d\x3d\xf7\xb2\x93\xdd\x29\x00\x00\xff\xff\xf2\x45\xe0\x5d\x7d\x01\x00\x00")

func nodesToscaNodesObjectstorageBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesObjectstorage,
		"nodes/tosca.nodes.ObjectStorage",
	)
}

func nodesToscaNodesObjectstorage() (*asset, error) {
	bytes, err := nodesToscaNodesObjectstorageBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.ObjectStorage", size: 381, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _nodesToscaNodesRoot = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\x41\x4e\xf3\x40\x0c\x85\xf7\x3d\x85\x0f\xf0\xff\x23\xd6\xdd\x01\x65\xdb\x4a\xb4\x5d\x21\x84\xdc\x89\x43\x2d\x4d\x66\x82\xc7\x41\xca\xed\x71\x92\x4e\x13\x01\x65\x56\x71\xfc\xf9\xd9\xef\x69\xca\x1e\x5d\x4c\x15\x65\xf7\x9c\x92\xae\x57\x00\xf6\xed\x85\x5b\xe5\x14\xd7\x70\x38\x13\x1c\x76\xfb\xc7\x7b\xd8\x1a\x04\x87\xbe\x25\xc0\x10\x20\xe9\x99\xe4\xd2\x39\x61\xa6\xb9\x9d\x4d\x40\xf8\x93\xa0\x96\xd4\x98\x1c\xaa\x0a\x9f\x3a\xa5\x3c\x88\x03\x8c\x2b\xdf\xb8\x9a\x2a\xab\x6d\x66\x0d\xd9\xa0\xf8\xbe\x00\x22\x36\x74\x13\xc9\x8a\x7a\xab\xeb\xb1\xc5\x13\x07\x56\x2e\x1b\x6b\x42\xed\xe4\x1b\x3f\x39\x5f\xc2\x6e\xb0\x60\x8c\xd0\x47\xc7\x42\x0d\x45\xbd\x08\xfc\x37\x4b\x2d\xc5\x8a\xa2\xef\x8b\xca\x62\x51\xff\x97\xda\xf4\x86\x84\x0b\x35\xa7\x7d\x6d\x5f\x9f\x50\xc0\x21\xf8\x7c\xe6\xb6\xf0\xcb\x7f\xd9\x6d\xc6\x4b\xf2\x2e\xfe\x1c\x4e\xde\x77\x22\x76\xa4\x19\x87\x17\xb8\xfb\x07\xc7\xed\xc3\xee\xb8\xdd\x3c\x6d\xe0\xd5\x70\x8e\x4a\x52\xa3\x2f\xb9\xec\x15\x63\x85\x52\xfd\x16\xcc\xcc\x8e\xf7\xba\xc0\x35\xf9\xde\x07\x72\x65\x6a\xf5\x15\x00\x00\xff\xff\x96\x43\xb6\xdc\x3c\x02\x00\x00")

func nodesToscaNodesRootBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesRoot,
		"nodes/tosca.nodes.Root",
	)
}

func nodesToscaNodesRoot() (*asset, error) {
	bytes, err := nodesToscaNodesRootBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.Root", size: 572, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _nodesToscaNodesSoftwarecomponent = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x90\x31\x6e\xc3\x30\x0c\x45\xf7\x9c\x82\x40\xe7\xf8\x00\x5e\xb3\x74\x2b\xd0\x1e\xc0\x60\x25\x1a\x21\x60\x89\x2a\xc9\xa4\xc8\xed\x2b\x25\xb6\x60\x0f\x1d\xf5\xf0\xf8\x45\x7e\x17\x0b\x38\x64\x89\x64\xc3\x97\xcc\xfe\x8b\x4a\x17\x49\x45\x32\x65\x1f\x4f\x00\x91\x94\xef\x14\xa7\x59\x25\x8d\xb0\xd7\x3f\x45\xbc\x0a\x45\xa5\x90\x3a\x93\x35\x1d\xe0\x0d\xa2\x24\xe4\x7c\xb6\x42\x81\x67\x0e\x60\x6b\x2e\x84\x2d\x18\xee\xa4\xc6\x92\x9f\x03\x9d\x4e\x2b\x7d\xe5\x00\xf8\xa3\xd0\x78\x50\x01\x94\x7e\x6e\xac\x14\x47\x98\x71\x31\x7a\x52\x8c\x89\xf3\x14\x2a\xad\x21\x8c\xcb\x71\xfe\xb5\x72\x44\xc7\xf6\xb6\xe1\xd2\xbd\x7f\x23\x57\x92\xaa\xb6\x1e\x75\x86\xab\x98\x6f\xc1\x75\x67\x2c\xf8\xcd\x0b\xfb\x63\xfb\xa0\x13\x6e\x7f\x48\xf6\x5a\x01\x69\x1f\x68\x95\x1d\xeb\x6b\x2d\xdf\x9c\xba\xa1\xb4\xa0\xd7\x43\xed\xca\x65\x33\xf7\xcc\x86\xf7\xba\x02\xc5\x8f\x7c\xfa\x0b\x00\x00\xff\xff\x80\x99\x40\x57\xb5\x01\x00\x00")

func nodesToscaNodesSoftwarecomponentBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesSoftwarecomponent,
		"nodes/tosca.nodes.SoftwareComponent",
	)
}

func nodesToscaNodesSoftwarecomponent() (*asset, error) {
	bytes, err := nodesToscaNodesSoftwarecomponentBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.SoftwareComponent", size: 437, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _nodesToscaNodesWebapplication = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\x41\x4e\x43\x31\x0c\x44\xf7\x3d\x85\x2f\x40\x0e\xf0\x77\x08\x21\xb1\x43\x82\x05\xcb\xaf\x34\x19\xa8\xa5\x5f\xdb\x38\xa6\xa2\xb7\x27\xa5\xe5\x13\x50\xb7\xe3\x37\x63\x7b\x42\x5b\xc9\x49\xb4\xa2\xa5\x17\x6c\x6f\xcd\x16\x2e\x39\x58\x65\xda\x10\x55\x38\x1f\x50\xe7\x57\xd7\xfd\x44\x23\xfb\xa4\x1a\x1d\x30\x57\x83\x07\xa3\x9d\x70\xa2\xa2\x12\xf8\x8c\xd9\xfb\xf8\xac\x10\xc5\xd1\x30\x51\x0b\x67\x79\xeb\x52\xc9\x96\xb7\xbc\xf0\xaf\x29\x9b\xcd\x90\x6a\xca\xf2\xcf\x74\xde\x38\x3a\xd2\xfd\x05\xec\x9c\xe3\xfd\x83\x1d\x7b\x48\x5c\x92\x6e\x68\xa7\x6d\xcd\x18\x76\x1d\xaf\x66\xdd\xf5\x6b\x33\x0b\x7c\x35\x9c\x9e\xfb\xfb\x68\x2f\xe5\x19\x7e\x18\x18\xc7\xf2\x5d\x50\xdb\xb1\xfd\xb0\xa3\xd6\xd2\x43\x3f\x02\xf5\x51\x36\x5f\x01\x00\x00\xff\xff\x8c\x92\x03\x78\x5e\x01\x00\x00")

func nodesToscaNodesWebapplicationBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesWebapplication,
		"nodes/tosca.nodes.WebApplication",
	)
}

func nodesToscaNodesWebapplication() (*asset, error) {
	bytes, err := nodesToscaNodesWebapplicationBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.WebApplication", size: 350, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _nodesToscaNodesWebserver = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\xd0\xb1\x6a\xc4\x30\x0c\x80\xe1\xfd\x9e\x42\xd0\xb5\x78\xea\x94\xed\x38\xba\x17\x6e\xe8\x50\x8a\x51\x6c\x85\x0a\x1c\xc9\xc8\x6a\x4a\xde\xbe\x49\x9a\x40\x28\x1d\xba\xda\xdf\x2f\x1b\xb9\xb6\x84\x41\x34\x53\x0b\xaf\xd4\xdf\xc9\x26\xb2\xee\x02\x90\xc9\x78\xa2\x1c\x07\xd3\xb1\x83\x33\xbb\xeb\xe0\x5f\x68\x74\xd3\xb1\xaa\x90\xf8\xa2\x13\x56\xec\xb9\xb0\x33\xb5\xb5\x06\x78\x80\x97\x65\x00\x3a\x3d\x42\xc1\x99\x0c\x9e\x80\x24\x57\x65\xf1\xb6\x81\x8c\x8e\xf1\x38\x3a\x5e\x38\xcf\x09\xcf\xfb\xe5\xc6\x31\x8f\x2c\xff\xf3\xe1\xba\xda\xad\xfa\xd0\xe6\x3f\xff\x01\xf0\xb9\xd2\x9f\xdd\x4d\xc5\x91\x85\x6c\x87\x13\x16\xce\xb1\xe9\xa7\x25\x8a\x6b\xd5\x3a\x78\x83\x5f\x9b\xba\xd6\x5a\x38\xa1\xb3\x0a\xbc\x5f\xbe\x03\x00\x00\xff\xff\x5a\x8e\x4d\x0f\x47\x01\x00\x00")

func nodesToscaNodesWebserverBytes() ([]byte, error) {
	return bindataRead(
		_nodesToscaNodesWebserver,
		"nodes/tosca.nodes.WebServer",
	)
}

func nodesToscaNodesWebserver() (*asset, error) {
	bytes, err := nodesToscaNodesWebserverBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "nodes/tosca.nodes.WebServer", size: 327, mode: os.FileMode(438), modTime: time.Unix(1433921141, 0)}
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
	"nodes/tosca.nodes.BlockStorage": nodesToscaNodesBlockstorage,
	"nodes/tosca.nodes.Compute": nodesToscaNodesCompute,
	"nodes/tosca.nodes.Container.Application": nodesToscaNodesContainerApplication,
	"nodes/tosca.nodes.Container.Runtime": nodesToscaNodesContainerRuntime,
	"nodes/tosca.nodes.DBMS": nodesToscaNodesDbms,
	"nodes/tosca.nodes.Database": nodesToscaNodesDatabase,
	"nodes/tosca.nodes.LoadBalancer": nodesToscaNodesLoadbalancer,
	"nodes/tosca.nodes.ObjectStorage": nodesToscaNodesObjectstorage,
	"nodes/tosca.nodes.Root": nodesToscaNodesRoot,
	"nodes/tosca.nodes.SoftwareComponent": nodesToscaNodesSoftwarecomponent,
	"nodes/tosca.nodes.WebApplication": nodesToscaNodesWebapplication,
	"nodes/tosca.nodes.WebServer": nodesToscaNodesWebserver,
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
	"nodes": &bintree{nil, map[string]*bintree{
		"tosca.nodes.BlockStorage": &bintree{nodesToscaNodesBlockstorage, map[string]*bintree{
		}},
		"tosca.nodes.Compute": &bintree{nodesToscaNodesCompute, map[string]*bintree{
		}},
		"tosca.nodes.Container.Application": &bintree{nodesToscaNodesContainerApplication, map[string]*bintree{
		}},
		"tosca.nodes.Container.Runtime": &bintree{nodesToscaNodesContainerRuntime, map[string]*bintree{
		}},
		"tosca.nodes.DBMS": &bintree{nodesToscaNodesDbms, map[string]*bintree{
		}},
		"tosca.nodes.Database": &bintree{nodesToscaNodesDatabase, map[string]*bintree{
		}},
		"tosca.nodes.LoadBalancer": &bintree{nodesToscaNodesLoadbalancer, map[string]*bintree{
		}},
		"tosca.nodes.ObjectStorage": &bintree{nodesToscaNodesObjectstorage, map[string]*bintree{
		}},
		"tosca.nodes.Root": &bintree{nodesToscaNodesRoot, map[string]*bintree{
		}},
		"tosca.nodes.SoftwareComponent": &bintree{nodesToscaNodesSoftwarecomponent, map[string]*bintree{
		}},
		"tosca.nodes.WebApplication": &bintree{nodesToscaNodesWebapplication, map[string]*bintree{
		}},
		"tosca.nodes.WebServer": &bintree{nodesToscaNodesWebserver, map[string]*bintree{
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

