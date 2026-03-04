package responses

type PatientResponse struct {
	ID        uint   `json:"ID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	FullName  string `json:"FullName"`
	Gender    string `json:"Gender"`
	DOB       string `json:"DOB"`
	Age       int    `json:"Age"`
	AgeGroup  string `json:"AgeGroup"`
}
