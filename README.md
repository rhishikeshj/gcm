gcm
===

The Android SDK provides a nice convenience library ([com.google.android.gcm.server](http://developer.android.com/reference/com/google/android/gcm/server/package-summary.html)) that greatly simplifies the interaction between Java-based application servers and Google's GCM servers. However, Google has not provided much support for application servers implemented in languages other than Java, specifically those written in the Go programming language. The `gcm` package helps to fill in this gap, providing a simple interface for sending GCM messages and automatically retrying requests in case of service unavailability.

Documentation: http://godoc.org/github.com/rhishikeshj/gcm

Getting Started
---------------

To install gcm, use `go get`:

```bash
go get github.com/rhishikeshj/gcm
```

Import gcm with the following:

```go
import "github.com/rhishikeshj/gcm"
```

Sample Usage
------------

Here is a quick sample illustrating how to send a message to the GCM server:

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/rhishikeshj/gcm"
)

func main() {
	// Create the message to be sent.
	data := map[string]interface{}{"score": "5x1", "time": "15:10"}
	regIDs := []string{"4", "8", "15", "16", "23", "42"}
	msg := gcm.NewMessage(data, regIDs...)
    msg.SetDryRun(true)

    // Create a Sender to send the message.
    sender, err := gcm.NewSender("<api_key>", nil)
    if err != nil {
        panic(err)
    }

    // Obtain the input and response channels from the Sender object
    inputChannel := sender.InputChannel
    responseChannel := sender.ResponseChannel
    var message_counter int
    go func() {
        for {
            nextResponse <-responseChannel
        }
    }()

    go func() {
        for i := 0; i < MessageCount; i++ {
            inputChannel <- msg
        }
    }()

	/* ... */
}
```

Note for Google AppEngine users
-------------------------------

If your application server runs on Google AppEngine, you must import the `appengine/urlfetch` package and create the `Sender` as follows:

```go
package sample

import (
	"appengine"
	"appengine/urlfetch"

	"github.com/rhishikeshj/gcm"
)

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	sender := &gcm.Sender{ApiKey: "sample_api_key", Http: client}

	/* ... */
}
```
