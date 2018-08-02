package types

type User struct {
	Email       []string `json:"email,omitempty"`
	Phone       []string `json:"phone,omitempty"`
	Name        string   `json:"name,omitempty"`
	NickName    string   `json:"nick_name,omitempty"`
	IDCard      string   `json:"id_card,omitempty"`
	CreditScore int64    `json:"credit_score,omitempty"`
	Game        []string `json:"game,omitempty"`
	Social      []string `json:"social,omitempty"`
	Job         []string `json:"job,omitempty"`
	Office      []string `json:"office,omitempty"`
}
