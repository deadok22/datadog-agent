// Code generated by go-bindata. DO NOT EDIT.
// source: pkg/ebpf/bytecode/build/oom-kill-kern-user.h
// +build ebpf_bindata

package bindata

import (
	"os"
	"time"
)

var _bindataOomkillkernuserH = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x51\x5d\x6b\x2a\x31\x14\x7c\xcf\xaf\x18\xd8\x97\x7b\x7d\xb8\x8b\xde\x8b\x70\xb1\x14\x4a\x11\x5a\xfc\x2a\x6a\x9f\x44\x42\x4c\x4e\xd6\xe0\x26\x59\x92\x2c\xad\x94\xfe\xf7\xe2\xfa\x51\xac\xfa\x98\x9c\x99\x39\x73\x66\x32\xa3\x9d\x22\x8d\xc9\x64\xc4\x07\xcf\xc3\x21\x1f\xf4\xa7\x63\xfe\x3a\xeb\x4f\xf9\x13\xcb\x14\x69\xe3\xe8\xfa\x90\x65\xc6\xc9\xb2\x56\x84\xbb\xd2\xb8\xfa\x3d\x4f\xdb\x8a\xe2\x9f\xf5\xfd\x6e\xb2\x17\x9d\x3f\xcc\x06\xfc\x71\x32\x1a\xf1\x61\x7f\x7c\x52\x3b\xfb\x45\xbb\xcb\x32\x72\xca\x68\xc6\x62\x0a\xb5\x4c\xf0\xde\xf2\x98\x44\x8a\xf8\x60\x80\x5c\x8b\x00\x59\x04\x5f\x57\xdc\x09\x4b\x8b\x76\xe7\xff\xb2\xc7\x80\x3c\xc7\x8b\x51\xf0\x1a\x29\x98\xa2\xa0\x60\x5c\x81\x2a\x78\x49\x31\x32\x80\xf3\xfa\x6f\x07\x95\x51\xe7\xd8\x8d\x29\x4b\x52\x17\xb8\xf4\x0d\x1c\x0b\x4b\x37\x55\x1b\x37\x5a\x7a\x6b\x17\x67\x67\x2c\x7f\x90\x2f\xd6\x34\xc4\x74\x9b\x38\xf7\x49\x94\x70\xb5\x5d\x51\xd8\x09\x54\xa2\xa0\x83\xbd\xee\xbf\xfd\xeb\x88\x0c\x42\x6e\x22\x8c\x46\x5a\x37\xd5\x34\xcb\xf0\x26\xe2\xd1\x32\x29\xac\xb6\x10\x87\xd4\x4e\x37\x5a\xb2\xb2\xe0\xde\xdb\x1e\xfb\xec\xb1\x43\xea\xc8\x5b\xd8\xf7\xa2\x7e\x5d\xab\xf9\x37\x5a\x39\xfb\x0a\x00\x00\xff\xff\x50\xd5\xb0\x84\x25\x02\x00\x00")

func bindataOomkillkernuserHBytes() ([]byte, error) {
	return bindataRead(
		_bindataOomkillkernuserH,
		"oom-kill-kern-user.h",
	)
}

func bindataOomkillkernuserH() (*asset, error) {
	bytes, err := bindataOomkillkernuserHBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "oom-kill-kern-user.h",
		size:        549,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}
