package types

import (
	"os"
	"path"
)

func NewFileMeta(fp string) (*FileMeta, error) {
	f1, err := os.Open(fp)
	if err != nil {
		return nil, err
	}

	stat, err := f1.Stat()
	if err != nil {
		return nil, err
	}

	f := &FileMeta{}
	f.Name = stat.Name()
	f.Size = stat.Size()
	f.ModTime = stat.ModTime().Unix()
	f.IsDir = stat.IsDir()
	f.Ext = path.Ext(fp)
	return f, nil
}

type FileMeta struct {
	Name string `json:"name,omitempty"`
	Size int64  `json:"size,omitempty"`

	// modification time
	ModTime int64  `json:"mod_time,omitempty"`
	IsDir   bool   `json:"is_dir,omitempty"`
	Ext     string `json:"ext,omitempty"`
}

type ImageMeta struct {
	FileMeta

	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`

	// 缩略图地址
	Thumb  string `json:"thumb,omitempty"`
}

type NameHash struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}
