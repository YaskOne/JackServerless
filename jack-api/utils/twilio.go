package utils

import (
	"github.com/sfreiberg/gotwilio"
)

var accountSid = "ACf8551f060f1478803011a13b0d973107"
var authToken = "aff47e64ab1d77833d39deba00b04391"
var from = "+33757902108"

func SendSMS(to string, message string) {
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	twilio.SendSMS(from, to, message, "", "")
}
