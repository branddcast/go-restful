package article

type Article struct {
	Id      uint   `json:"Id,omitempty"`
	Title   string `json:"Title,omitempty"`
	Desc    string `json:"desc,omitempty"`
	Content string `json:"content,omitempty"`
}
