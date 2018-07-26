package core

import "github.com/kooksee/kfs/types"

// 把文件添加到kfs中
func FileAdd(f string) error {

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

// 查看文件的metadata
func FileList(hash string) error {

	return nil
}
