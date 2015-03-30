package transports

import ( 
	"fmt"
	apns "github.com/anachronistic/apns"
)

type IosTransport struct {

}

// The iOS transporter is example code and needs to actually be implemented
func (transport *IosTransport) Testing() {
	payload := apns.NewPayload()
	payload.Alert = "Hello, world!"
	payload.Badge = 42
	payload.Sound = "bingbong.aiff"

	pn := apns.NewPushNotification()
	pn.DeviceToken = "YOUR_DEVICE_TOKEN_HERE"
	pn.AddPayload(payload)

	client := apns.NewClient("gateway.sandbox.push.apple.com:2195", "YOUR_CERT_PEM", "YOUR_KEY_NOENC_PEM")
	resp := client.Send(pn)

	alert, _ := pn.PayloadString()
	fmt.Println("  Alert:", alert)
	fmt.Println("Success:", resp.Success)
	fmt.Println("  Error:", resp.Error)
}