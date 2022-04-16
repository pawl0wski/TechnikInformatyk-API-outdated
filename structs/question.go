package structs

type Question struct {
	Uuid         string `json:"uuid"`
	Content      string `json:"content"`
	ImageId      int64  `json:"imageId"`
	AnswerA      string `json:"answerA"`
	AnswerB      string `json:"answerB"`
	AnswerC      string `json:"answerC"`
	AnswerD      string `json:"answerD"`
	CorretAnswer int8   `json:"corretAnswer"`
}
