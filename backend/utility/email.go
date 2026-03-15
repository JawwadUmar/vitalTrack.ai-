package utility

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to string, otp string) error {

	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Subject: Verify Your Email\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body := fmt.Sprintf(`
	<html>
	<body style="font-family: Arial, sans-serif; background-color:#f4f4f4; padding:20px;">
	
	<div style="max-width:600px; margin:auto; background:white; padding:30px; border-radius:10px; text-align:center;">
		
		<h2 style="color:#333;">Email Verification</h2>
		
		<p style="color:#555; font-size:16px;">
			Thank you for signing up at VitaTrack.AI. Please use the OTP below to verify your email address.
		</p>

		<div style="font-size:30px; letter-spacing:6px; font-weight:bold; 
		            background:#f1f1f1; padding:15px; border-radius:8px; 
		            margin:20px 0; display:inline-block;">
			%s
		</div>

		<p style="color:#888; font-size:14px;">
			This OTP will expire in 5 minutes.
		</p>

		<hr style="margin:30px 0;">

		<p style="color:#aaa; font-size:12px;">
			If you did not request this email, please ignore it.
		</p>

	</div>

	</body>
	</html>
	`, otp)

	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		[]string{to},
		message,
	)

	return err
}
