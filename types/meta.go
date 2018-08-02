package types

type Metadata struct {
	IMetadata

	// 用于metadata的全局id,metadata修改时ID不变
	ID string `json:"id,omitempty"`

	// 对该内容的描述
	Abstract string `json:"abstract,omitempty"`

	// 分组,只能归档到一个组
	Category []string `json:"category,omitempty"`

	// 标签,可以有多个标签
	Tag []string `json:"tag,omitempty"`

	// 根据签名生成的唯一的内容识别码
	DNA string `json:"dna,omitempty"`

	// 内容分片的数量
	ChunkNum uint64 `json:"chunk_num,omitempty"`

	// 修改前的dna
	ParentDna string `json:"parent_dna,omitempty"`

	// 扩展的额外的数据
	Extra map[string]string `json:"extra,omitempty"`

	// 内容原地址
	Source string `json:"source,omitempty"`

	// 内容的标题说明
	Title string `json:"title,omitempty"`

	// 组件扩展附加,比如内容的点赞评论和转发以及投票等组件
	Extend []string `json:"extend,omitempty"`

	// 其他类型的数据
	Data ISubMetadata `json:"data,omitempty"`

	// 内容hash,keecak256 hash
	ContentHash string `json:"content_hash,omitempty"`

	// 内容的创建时间
	CreateTime int64 `json:"create_time,omitempty"`

	// 内容的修改时间
	UpdateTime int64 `json:"update_time,omitempty"`

	// 创建和修改状态
	Status string `json:"status,omitempty"`

	// 内容的语言
	Language string `json:"language,omitempty"`

	// 内容的签名
	Signature string `json:"signature,omitempty"`

	// 内容的许可证
	License struct {
		Type   string            `json:"type,omitempty"`
		Params map[string]string `json:"parameters,omitempty"`
	} `json:"license,omitempty"`
}

func (m *Metadata) Decode(data []byte) error {
	return json.Unmarshal(data, m)
}

func (m *Metadata) Encode() ([]byte, error) {
	return json.Marshal(m)
}
