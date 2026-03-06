package repositories

import "hospital/config"

type WardSummary struct {
	WardID    uint
	WardName  string
	TotalBeds int64
	Available int64
	Occupied  int64
	Cleaning  int64
}

type HospitalSummary struct {
	TotalWards       int64
	TotalRooms       int64
	TotalBeds        int64
	AvailableBeds    int64
	OccupiedBeds     int64
	CleaningBeds     int64
	ActiveAdmissions int64
}

func GetHospitalSummary() (HospitalSummary, error) {

	var summary HospitalSummary

	config.DB.Table("wards").Count(&summary.TotalWards)
	config.DB.Table("rooms").Count(&summary.TotalRooms)
	config.DB.Table("beds").Count(&summary.TotalBeds)
	config.DB.Table("beds").Where("status = ?", "AVAILABLE").Count(&summary.AvailableBeds)
	config.DB.Table("beds").Where("status = ?", "OCCUPIED").Count(&summary.OccupiedBeds)
	config.DB.Table("beds").Where("status = ?", "CLEANING").Count(&summary.CleaningBeds)
	config.DB.Table("admissions").Where("status = ?", "ACTIVE").Count(&summary.ActiveAdmissions)

	return summary, nil
}

func GetWardSummary() ([]WardSummary, error) {

	var wards []WardSummary

	err := config.DB.Raw(`
		SELECT 
			w.id AS ward_id,
			w.name AS ward_name,
			COUNT(b.id) AS total_beds,
			SUM(CASE WHEN b.status = 'AVAILABLE' THEN 1 ELSE 0 END) AS available,
			SUM(CASE WHEN b.status = 'OCCUPIED' THEN 1 ELSE 0 END) AS occupied,
			SUM(CASE WHEN b.status = 'CLEANING' THEN 1 ELSE 0 END) AS cleaning
		FROM wards w
		LEFT JOIN rooms r ON r.ward_id = w.id
		LEFT JOIN beds b ON b.room_id = r.id
		GROUP BY w.id, w.name
	`).Scan(&wards).Error

	return wards, err
}
