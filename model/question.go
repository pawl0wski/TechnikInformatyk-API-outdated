package model

type Question struct {
	Uuid          string   `json:"uuid"`
	Content       string   `json:"content"`
	HaveImage     bool     `json:"haveImage"`
	AnswerA       string   `json:"answerA"`
	AnswerB       string   `json:"answerB"`
	AnswerC       string   `json:"answerC"`
	AnswerD       string   `json:"answerD"`
	CorrectAnswer uint8    `json:"correctAnswer"`
	ExamUuids     []string `json:"examUuids"`
}

func (q Question) IsEmpty() bool {
	return q.Uuid == "" && q.Content == "" && q.AnswerA == "" && q.AnswerB == "" && q.AnswerC == "" && q.AnswerD == ""
}
