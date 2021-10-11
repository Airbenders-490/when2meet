package domain

// availibilities struct
type Availibilities struct {
	AID        string `json:"AID"` //uuid string
	ADay       string `json:"ADay"`
	Start_Time int
	End_Time   int
}
