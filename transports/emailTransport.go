package transports

import ( 
    "os"
    "fmt"
    "github.com/sendgrid/sendgrid-go"
)

type EmailTransport struct {
}

// Send email using sendgrid, but this could be swapped out for any email implementation
func (transport *EmailTransport) Send() {
    fmt.Println("Send email from email transport")

    username := os.Getenv("SENDGRID_USERNAME")
    password := os.Getenv("SENDGRID_PASSWORD")

    // @TODO Will need to get the username and password from the environment
    sg := sendgrid.NewSendGridClient(username, password)
    
    message := sendgrid.NewMail()
    message.AddTo("test@jasonmichels.com")
    message.AddToName("Jason")
    message.SetSubject("SendGrid Testing")
    message.SetText("WIN")
    message.SetFrom("test@jasonmichels.com")

    if r := sg.Send(message); r == nil {
        fmt.Println("Email sent!")
    } else {
        fmt.Println(r)
    }
}