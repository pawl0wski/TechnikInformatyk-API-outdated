package structs

type Question struct {
	Uuid         string   `json:"uuid"`
	Content      string   `json:"content"`
	HaveImage    bool     `json:"haveImage"`
	AnswerA      string   `json:"answerA"`
	AnswerB      string   `json:"answerB"`
	AnswerC      string   `json:"answerC"`
	AnswerD      string   `json:"answerD"`
	CorretAnswer int8     `json:"corretAnswer"`
	ExamUuids    []string `json:"examUuids"`
}
