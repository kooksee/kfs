package types

type Metadata struct {
	IMetadata

	ID        string            `json:"id,omitempty"`
	Abstract  string            `json:"abstract,omitempty"`
	Category  []string          `json:"category,omitempty"`
	Tag       []string          `json:"tag,omitempty"`
	DNA       string            `json:"dna,omitempty"`
	ChunkNum  uint64            `json:"chunk_num,omitempty"`
	ParentDna string            `json:"parent_dna,omitempty"`
	Extra     map[string]string `json:"extra,omitempty"`
	Source    string            `json:"source,omitempty"`
	Title     string            `json:"title,omitempty"`
	Include   []string          `json:"include,omitempty"`

	ContentHash string `json:"content_hash,omitempty"`
	CreateTime  int64  `json:"create_time,omitempty"`
	UpdateTime  int64  `json:"update_time,omitempty"`
	Status      Status `json:"status,omitempty"`
	Language    string `json:"language,omitempty"`
	Signature   string `json:"signature,omitempty"`
	Type        string `json:"type,omitempty"`
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
