# Notifications Micro Service

The project is a work in progress and really should not be used by anyone in it's current state.  The final product will be a microservice written in Go that pulls data from a queue, initially AWS SQS, and send out a notification, based on data in the message.

### Installation Locally
```sh
$ git clone git@github.com:jasonmichels/notifications.git
$ cd notifications
$ go get ./...
$ go run app.go
```

### Todo's
 - Make iOS notifications work
 - Added environment variables for config items
 - Make email send dynamic information
 - Add pulling down data from a queue
 - Make transporter resolver dynamic
 - Add docker file
 - Probably a few other things...

 ## Contributing

 This project is just getting started and a major work in progress.  Let me know if you would like to see anything

### License

The notification micro service is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT)