package model

type Player struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Lvl  uint64 `json:"lvl"`
}
