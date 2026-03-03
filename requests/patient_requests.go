package requests

type CreatePatientRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	DOB       string `json:"dob"`
}
