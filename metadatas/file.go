package metadatas

type FileMeta struct {
	Metadata

	Name    string `json:"name,omitempty"`
	Size    int64  `json:"size,omitempty"`
	Mode    string `json:"node,omitempty"`
	ModTime int64  `json:"mod_time,omitempty"`
	IsDir   bool   `json:"is_dir,omitempty"`
}

type ImageMeta struct {
	FileMeta
}