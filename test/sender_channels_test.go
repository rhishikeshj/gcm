package gcm_test

import (
    "gcm"
    "testing"
)

const MessageCount = 10

func TestSenderChannels(t *testing.T) {
    // Create a Sender to send the message.
    sender, err := gcm.NewSender("AIzaSyCGPH_AXJbOXVMJmmDHg9cTIgsgPw75QPw", nil)
    if err != nil {
        panic(err)
    }

    inputChannel := sender.InputChannel
    responseChannel := sender.ResponseChannel
    var message_counter int
    go func() {
        for {
            <-responseChannel
            message_counter++
            if message_counter >= MessageCount {
                break
            }
        }
    }()

    go func() {
        for i := 0; i < MessageCount; i++ {
            data := map[string]interface{}{"content": "Notification : " + string(i),
                "title":            "Whoa ! Did you see this ?",
                "default-action":   "http://www.reddit.com",
                "primary_action":   "http://www.google.com",
                "secondary_action": "http://www.helpshift.com",
                "primary_button":   "google",
                "secondary_button": "helpshift"}
            regIDs := []string{"APA91bHVh_zsSeTIe7em27wWHrXjzRhpSCT3NALk8qgFJAxCmesSEgzbs2qvQYE_ZNTIGf5dnYzcWviIzw2K8MH-EqDxiM16plOM8x3FUN1Qe9gyjvhabEUyQOgB0C-nT1d2_2ou_Y-G4KkbOTigGuQQaHE1hx81bxOTcN67BMzA-RtOi1izCH4"}
            msg := gcm.NewMessage(data, regIDs...)
            msg.SetDryRun(true)
            inputChannel <- msg
        }
    }()
}
