package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Message struct {
	Text  string
	Title string
}

var msg = Message{}

// messageQueue is the queue that holds new messages until they are processed
var messageQueue chan Message

func init() { // need the init function to initialize the channel, and the listeners
	// initialize the queue, choosing the buffer size as 8 (number of messages the channel can hold at once)
	messageQueue = make(chan Message, 8)

	// start a goroutine that listens on the queue/channel for new messages
	go listenForMessages()
}

func listenForMessages() {
	// whenever we detect a message in the queue, append it to Messages
	for m := range messageQueue {
		msg = m
	}
}

func messagePost(c http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/clear" && req.Method == "POST" {
		msg = Message{}
		c.WriteHeader(http.StatusAccepted)
		return
	}
	if req.URL.Path != "/" {
		http.NotFound(c, req)
		return
	}
	switch req.Method {
	case "POST":
		decoder := json.NewDecoder(req.Body)
		var m Message
		err := decoder.Decode(&m)
		if err != nil {
			c.WriteHeader(http.StatusBadRequest)
		} else {
			// add the message to the channel, it'll only wait if the channel is full
			messageQueue <- m

			c.WriteHeader(http.StatusCreated)
		}
	case "GET":
		data, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		c.Header().Set("Content-Type", "application/json")
		c.WriteHeader(http.StatusOK)
		c.Write(data)
	}
}

func main() {
	http.HandleFunc("/", messagePost)
	if err := http.ListenAndServe(":"+os.Args[1], nil); err != nil {
		panic(err)
	}
}
