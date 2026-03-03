package responses

type PatientResponse struct {
	ID        uint   `json:"ID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Gender    string `json:"Gender"`
	DOB       string `json:"DOB"`
	Age       int    `json:"Age"`
}
