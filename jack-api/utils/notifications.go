package utils

import (
	"fmt"
	"log"
	"github.com/maddevsio/fcm"
)

var takewayKey = "AAAA0qrCnWQ:APA91bH0r1Fjcoa3qlEyDwyioj9GUVl3_Oex4gW6CxKchSprF4WjhhIDEzMzZMIWURHyMP2LilaM_ws2xZ_lDm8o2D2snKbuhW_FWdssdpdTNJD_XMM6KzTGD_mGoARtrREGQNvLU-iS"
var businessKey = "AAAAQnKVOOY:APA91bFne73gLfU6oF436Wg4zIu6i4-Jj9IFd0S9gKZKkqQIcQdbfu4eDN4JlTH606DcheH3J6YpHJHn7SLK98XeGsF2dXiWECxrytyBFpO-nAgAukK6k6LLrW5CDSgv_JJkmKqGC3J8"

var OrderUpdated = "OrderUpdated"
var NewOrder = "NewOrder"

func SendPushToClient(app string, token string, title string, message string, data map[string]string)  {

	if token == "" {
		return
	}

	c := fcm.NewFCM(getServerKey(app))

	response, err := c.Send(fcm.Message{
		Data:             data,
		RegistrationIDs:  []string{token},
		ContentAvailable: true,
		Priority:         fcm.PriorityHigh,
		Notification: fcm.Notification{
			Title: title,
			Body:  message,
			Sound: "default",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status Code   :", response.StatusCode)
	fmt.Println("Success       :", response.Success)
	fmt.Println("Fail          :", response.Fail)
	fmt.Println("Canonical_ids :", response.CanonicalIDs)
	fmt.Println("Topic MsgId   :", response.MsgID)
}

func getServerKey(app string) string {
	if app == "Takeway" {
		return takewayKey
	} else if app == "Business" {
		return businessKey
	}
	return ""
}