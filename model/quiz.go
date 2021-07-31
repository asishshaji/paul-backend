package model

type Quiz struct {
	Type     string   `json:"type"`
	Options  []string `json:"options"`
	Question string   `json:"question"`
	ImageUrl string   `json:"imageUrl"`
	Answer   string   `json:"answer"`
}
