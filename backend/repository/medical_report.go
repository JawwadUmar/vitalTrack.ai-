package repository

import (
	"vita-track-ai/database"
	"vita-track-ai/models"
)

func CreateMedicalReport(reportDb *models.MedicalReportDB) error {
	return database.DB.Create(reportDb).Error
}

func GetMedicalReportByID(id string) (*models.MedicalReportDB, error) {
	var medicalReport models.MedicalReportDB
	err := database.DB.First(&medicalReport, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &medicalReport, nil
}
