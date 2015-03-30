package transports

import ( 
    "fmt"
    "github.com/sendgrid/sendgrid-go"
)

type EmailTransport struct {
}

// Send email using sendgrid, but this could be swapped out for any email implementation
func (transport *EmailTransport) Send() {
    fmt.Println("Send email from email transport")

    // @TODO Will need to get the username and password from the environment
    sg := sendgrid.NewSendGridClient("username", "password")
    
    message := sendgrid.NewMail()
    message.AddTo("receiver@email.com")
    message.AddToName("Receipient Name")
    message.SetSubject("SendGrid Testing")
    message.SetText("WIN")
    message.SetFrom("sender@email.com")

    if r := sg.Send(message); r == nil {
        fmt.Println("Email sent!")
    } else {
        fmt.Println(r)
    }
}