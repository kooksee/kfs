package sp2p

type IHandler func(*SP2p, *KMsg)

type IMessage interface {
	// 获取类型
	T() byte
	// 描述信息
	String() string
	// 业务处理
	OnHandle(*SP2p, *KMsg)
}

type ITable interface {
	// 路由表大小
	Size() int
	// 获得节点列表,把节点列表转换为[sp2p://<hex node id>@10.3.58.6:30303?discport=30301]的方式
	GetRawNodes() []string
	// 添加节点
	AddNode(*Node)
	// 更新节点
	UpdateNode(*Node)
	// 删除节点
	DeleteNode(Hash)
	// 随机得到路由表中的n个节点
	FindRandomNodes(int) []*Node
	// 查找距离最近的n个节点
	FindMinDisNodes(Hash, int) []*Node
	// 查找目标相比本节点更近的节点
	FindNodeWithTargetBySelf(Hash) []*Node
	// 查找目标相比另一个节点的更近的节点
	FindNodeWithTarget(Hash, Hash) []*Node
}
