package core

type ApiCore interface {
	// file
	// 添加文件,把文件添加到文件系统中,同时pin到自己内容库
	FileAdd(fp string)

	// 删除文件,删除该文件的的所有的备份以及meta以及历史记录
	FileRm(hash string)

	// 列出所有的文件,只列出最新的
	FileList()

	// 查看该文件的历史记录,n: 第几次的历史记录
	FileHistory(hash string, n uint64)

	// 通过hash获得该文件,并放到自己的缓存中
	FileGet(hash string)

	// 把内容固化到自己的内容库中,同时把缓存的内容删除
	FilePin(hash string)

	// metadata
	// 查看内容的metadata
	MetadataList(hash string)

	// 添加内容的metadata
	MetadataAdd()

	// 删除内容的metadata,同时把该metadata对应的内容也删除
	MetadataRm()
	MetadataUpdate()

	PeerList()
	PeerRm()
	PeerAdd()
}
