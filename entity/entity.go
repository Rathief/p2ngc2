package entity

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type Hero struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
	ImgURL   string `json:"imgurl"`
}
