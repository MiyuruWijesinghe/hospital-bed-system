package services

import "hospital/repositories"

type DashboardResponse struct {
	HospitalSummary repositories.HospitalSummary `json:"hospitalSummary"`
	WardSummary     []repositories.WardSummary   `json:"wardSummary"`
}

func GetDashboardData() (DashboardResponse, error) {

	hospitalSummary, err := repositories.GetHospitalSummary()
	if err != nil {
		return DashboardResponse{}, err
	}

	wardSummary, err := repositories.GetWardSummary()
	if err != nil {
		return DashboardResponse{}, err
	}

	return DashboardResponse{
		HospitalSummary: hospitalSummary,
		WardSummary:     wardSummary,
	}, nil
}
