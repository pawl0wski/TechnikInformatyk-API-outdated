package structs

type Player struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
	Lvl  int64  `json:"lvl"`
}
