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
    fmt.Println("Sending email to jmichels@fisdap.net")

    username := os.Getenv("SENDGRID_USERNAME")
    password := os.Getenv("SENDGRID_PASSWORD")

    // @TODO Will need to get the username and password from the environment
    sg := sendgrid.NewSendGridClient(username, password)
    
    message := sendgrid.NewMail()
    message.AddTo("jmichels@fisdap.net")
    message.AddToName("Jason Michels")
    message.SetSubject("SendGrid Testing")
    message.SetText("WIN")
    message.SetFrom("jmichels@fisdap.net")

    if r := sg.Send(message); r == nil {
        fmt.Println("Email sent!")
    } else {
        fmt.Println(r)
    }
}