package core

import (
	"github.com/kooksee/kfs/types"
	"github.com/kooksee/kfs/cmn"
)

// 把文件添加到kfs中
func (a *ApiCore) FileAdd(f string) error {
	// 获得文件地址
	// 获得文件的metadata
	// 获得文件分片
	// 存储文件的分片和metadata
	// 签名，活的DNA

	if !cmn.FileExists(f) {
		return cmn.Err("文件%s不存在", f)
	}

	meta, err := types.CreateFileMeta(kvDb, f)
	if err != nil {
		return err
	}

	d, err := meta.Encode()
	if err != nil {
		return err
	}

	return metaDb.Set([]byte(meta.(*types.Metadata).ID), d)
}

// 删除文件,删除历史以及metadata
func (a *ApiCore) FileRm(hash string) error {
	return nil
}
