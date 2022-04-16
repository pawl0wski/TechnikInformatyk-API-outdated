package structs

type Exam struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Icon        string     `json:"icon"`
	Questions   []Question `json:"questions"`
}
