package models

type MedicalReport struct {
	ReportMetadata struct {
		ReportDate    string `json:"report_date"`
		ReportType    string `json:"report_type"`
		HospitalOrLab string `json:"hospital_or_lab_name"`
	} `json:"report_metadata"`
	Metrics             []Metric        `json:"metrics"`
	AbnormalFindings    []string        `json:"abnormal_findings"`
	SimpleExplanation   string          `json:"simple_explanation"`
	OverallRiskLevel    string          `json:"overall_risk_level"`
	Recommendations     Recommendations `json:"recommendations"`
	FollowUpSuggestions []string        `json:"follow_up_suggestions"`
}

type Metric struct {
	TestName       string `json:"test_name"`
	Value          string `json:"value"`
	Unit           string `json:"unit"`
	ReferenceRange string `json:"reference_range"`
	Status         string `json:"status"`
}

type Recommendations struct {
	Diet      []string `json:"diet"`
	Lifestyle []string `json:"lifestyle"`
}
