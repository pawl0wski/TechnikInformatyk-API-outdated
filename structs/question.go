package structs

type Question struct {
	Uuid          string   `json:"uuid"`
	Content       string   `json:"content"`
	HaveImage     bool     `json:"haveImage"`
	AnswerA       string   `json:"answerA"`
	AnswerB       string   `json:"answerB"`
	AnswerC       string   `json:"answerC"`
	AnswerD       string   `json:"answerD"`
	CorrectAnswer int8     `json:"correctAnswer"`
	ExamUuids     []string `json:"examUuids"`
}
