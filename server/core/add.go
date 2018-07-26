package core

import "github.com/kooksee/kfs/types"

type ApiCore struct {
	IApiCore
}

// 把文件添加到kfs中
func (a *ApiCore) FileAdd(f string) error {

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
