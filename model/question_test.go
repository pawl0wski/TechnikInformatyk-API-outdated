package model

import "testing"

func TestQuestionIsEmptyIfIs(t *testing.T) {
	newQuestion := Question{}
	if newQuestion.IsEmpty() == false {
		t.Error("IsEmpty() returns false if Question is empty")
	}
}

func TestQuestionIsEmptyIfIsNot(t *testing.T) {
	newQuestion := Question{}
	newQuestion.Content = "abc"
	if newQuestion.IsEmpty() == true {
		t.Error("IsEmpty() returns true if Question is not empty")
	}
}
