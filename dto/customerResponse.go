package dto

// CustomerResponse Exported
type CustomerResponse struct {
	ID          int    `json:"customer_id"`
	Name        string `json:"fullName"`
	City        string `json:"city"`
	Zipcode     int    `json:"zipCode"`
	DateOfBirth string `json:"dateOfBirth"`
	Status      string `json:"status"`
}
