package utility

import (
	"fmt"
	"math/rand"
	"time"
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

func GenerateOTP() models.OneTimePassword {
	var oneTimePassword models.OneTimePassword
	rand.Seed(time.Now().UnixNano())
	otpStr := fmt.Sprintf("%06d", rand.Intn(1000000))
	expiry := time.Now().Add(5 * time.Minute)

	oneTimePassword.OTP = &otpStr
	oneTimePassword.OTPExpiresAt = &expiry

	return oneTimePassword
}
