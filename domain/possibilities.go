package domain

// possibilities struct
type Possibilities struct {
	PID        string `json:"id"` //uuid string
	PDay       string `json:"PDay"`
	Start_Time int
	End_Time   int
}
