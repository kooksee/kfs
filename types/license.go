package types

type License struct {
	Type        string             `json:"type,omitempty"`
	Parameters  []LicenseParameter `json:"parameters,omitempty"`
	Description string             `json:"description,omitempty"`
}

type LicenseParameter struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Values      string `json:"values,omitempty"`
}
