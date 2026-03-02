package utility

import (
	"encoding/json"
	"vita-track-ai/models"
)

// Helper function to safely get strings from claims
func GetClaim(key string, claims map[string]interface{}) string {

	if val, ok := claims[key]; ok && val != nil {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func ParseResponse(respBody []byte) (*models.MedicalReport, error) {
	// Step 1: Unmarshal outer JSON
	var outer struct {
		JSON string `json:"json"`
	}
	if err := json.Unmarshal(respBody, &outer); err != nil {
		return nil, err
	}

	// Step 2: Unmarshal inner JSON string
	var report models.MedicalReport
	if err := json.Unmarshal([]byte(outer.JSON), &report); err != nil {
		return nil, err
	}

	return &report, nil
}
