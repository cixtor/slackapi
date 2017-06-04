package main

import (
	"fmt"
	"github.com/cixtor/slackapi"
)

// MonitorRealTimeMessages prints all the supported websocket events.
func MonitorRealTimeMessages(client *slackapi.SlackAPI) {
	rtm, err := client.NewRTM()

	if err != nil {
		fmt.Println("RTM error;", err)
		return
	}

	go rtm.ManageEvents()

	for msg := range rtm.Events {
		switch event := msg.Data.(type) {
		case *slackapi.HelloEvent:
			fmt.Println("hello; connection established")

		case *slackapi.PresenceChangeEvent:
			fmt.Println("presence;", event.User, "=>", event.Presence)

		case *slackapi.MessageEvent:
			if event.Text == "stop" {
				rtm.Stop()
			} else {
				fmt.Printf(
					"message; %s@%s: %#v\n",
					event.User,
					event.Channel,
					event.Text)
			}

		case *slackapi.ErrorEvent:
			fmt.Println("error;", event.Text)

		case *slackapi.ReconnectURLEvent:
			fmt.Println("reconnect;", event.URL)

		default:
			fmt.Printf("%s; %#v\n", msg.Type, msg.Data)
		}
	}

	fmt.Println("stopped")
}
