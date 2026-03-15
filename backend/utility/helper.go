package utility

import (
	"fmt"
	"math/rand"
	"time"
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

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
