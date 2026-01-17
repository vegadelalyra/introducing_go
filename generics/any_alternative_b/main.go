package main

import "fmt"

const (
	TypeEmail = "EMAIL"
	TypeSMS   = "SMS"
	TypePush  = "PUSH"
)

type Notification struct {
	Kind         string
	EmailAddress string
	PhoneNumber  string
	Message      string
	IsSilent     bool
}

func SendNotification(n Notification) {
	switch n.Kind {
	case TypeEmail:
		fmt.Printf("Sending Email to %s: %s\n", n.EmailAddress, n.Message)
	case TypeSMS:
		fmt.Printf("Sending SMS to %s: %s\n", n.EmailAddress, n.Message)
	case TypePush:
		status := "loud"
		if n.IsSilent {
			status = "silent"
		}
		fmt.Printf("Sending %s Push %s\n", status, n.Message)
	}
}

func main() {
    // This struct lives on the STACK. 
    // There is no pointer-chasing to find the data. 
    notif := Notification{
        Kind: TypeEmail,
        EmailAddress: "user@example.com",
        Message: "Hello Go!",
    }

    SendNotification(notif)
}
