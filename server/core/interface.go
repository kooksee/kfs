package core

import (
	"github.com/kooksee/kfs/types"
	"time"
)

type IApiCore interface {
	// file
	// 添加文件,把文件添加到文件系统中,同时pin到自己内容库
	FileAdd(fp string) error

	// 删除文件,删除该文件的的所有的备份以及meta以及历史记录
	FileRm(hash string) error

	// 列出所有的文件,只列出最新的文件名以及对应的hash
	// n=-1 列出所有的数据,n!=-1就列出n条数据
	FileList(n int64) []types.NameHash

	// 查看该文件的历史记录,n: 第几次的历史记录
	// 返回对应的内容
	// 当n=-1的时候是最新的内容
	FileHistory(hash string, n int64) []byte

	// 通过hash获得该文件,并放到自己的缓存中
	// 返回对应的内容
	FileGet(hash string) []byte

	// 把内容固化到自己的内容库中,同时把缓存的内容删除
	FilePin(hash string) error

	// 把hash所对应的内容分享出去
	FileShare(hash string) error

	// metadata
	// 查看内容的metadata
	MetadataList(hash string) types.Metadata

	// 参看该文件的metadata的历史记录
	MetadataHistory(hash string, n uint64) types.Metadata

	// 添加内容的metadata
	MetadataAdd(metadata types.Metadata) error

	// 删除内容的metadata,同时把该metadata对应的内容也删除
	MetadataRm(hash string) error

	// 更新metadata
	MetadataUpdate(metadata types.Metadata) error

	// peer
	// 列出所有的节点,或者列出n个节点
	//n=-1列出所有的节点
	PeerList(n uint64) []string

	// 根据节点的ID删除该节点
	PeerRm(nodeID string) error

	// 根据节点的地址添加该节点
	PeerAdd(nodeUrl string) error

	// object,内容对象操作
	// 添加内容
	ObjectAdd(hash string, data []byte) error

	// 列出对象
	ObjectList(n uint64) []types.NameHash

	// 删除对象
	ObjectRm(hash string) error

	// 获得该对象的历史记录
	// n=-1获得最新的
	ObjectHistory(hash string, n int64) []byte

	// 获得该对象的内容
	// 必须是有metadata的对象才行
	ObjectGet(hash string) []byte

	// 固化对象的内容,并向缓存系统广播自己缓存的内容
	ObjectPin(hash string) error

	// 把对象的内容分享出去,该内容必须是自己本地的数据而且网上不存在的
	// 分享出去的同时并把metadata信息存放到区块链上
	ObjectShare(hash string) error

	// 解锁本账户
	AccountUnLock(passwd string, duration time.Duration) error

	// 锁定本账户
	AccountLock() error
}

type ApiCore struct {
	IApiCore
}
