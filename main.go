package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.2.23"

func main() {
	var client SlackAPI

	flag.Usage = func() {
		fmt.Println("Slack API Client")
		fmt.Println("  http://cixtor.com/")
		fmt.Println("  https://api.slack.com/")
		fmt.Println("  https://github.com/cixtor/slackapi")
		fmt.Println()
		fmt.Println("Description:")
		fmt.Println("  Low level Slack API client with custom commands. Slack, the 'messaging app for")
		fmt.Println("  teams' offers an API that has been used to build multiple projects around it,")
		fmt.Println("  from bots to independent clients as well as integrations with other external")
		fmt.Println("  services. This project aims to offer a low level experience for advanced users")
		fmt.Println("  that want to either drop the web client or interact with the API for testing")
		fmt.Println("  purpose.")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  slackapi api.test                                  Checks API calling code")
		fmt.Println("  slackapi auth.test                                 Checks authentication and identity")
		fmt.Println("  slackapi channels.history [channel] [time]         Fetches history of messages and events from a channel")
		fmt.Println("  slackapi channels.info [channel]                   Gets information about a channel")
		fmt.Println("  slackapi channels.list                             Lists all channels in a Slack team")
		fmt.Println("  slackapi channels.mark [channel] [time]            Sets the read cursor in a channel")
		fmt.Println("  slackapi channels.setPurpose [channel] [purpose]   Sets the purpose for a channel")
		fmt.Println("  slackapi channels.setTopic [channel] [topic]       Sets the topic for a channel")
		fmt.Println("  slackapi chat.delete [channel] [time]              Deletes a message")
		fmt.Println("  slackapi chat.postMessage [channel] [text]         Sends a message to a channel")
		fmt.Println("  slackapi chat.session                              Starts a new chat session")
		fmt.Println("  slackapi chat.update [channel] [time] [text]       Updates a message")
		fmt.Println("  slackapi emoji.list                                Lists custom emoji for a team")
		fmt.Println("  slackapi groups.close [channel]                    Closes a private channel")
		fmt.Println("  slackapi groups.history [channel] [time]           Fetches history of messages and events from a private channel")
		fmt.Println("  slackapi groups.info [channel]                     Gets information about a private channel")
		fmt.Println("  slackapi groups.list                               Lists private channels that the calling user has access to")
		fmt.Println("  slackapi groups.mark [channel] [time]              Sets the read cursor in a private channel")
		fmt.Println("  slackapi groups.open [group]                       Opens a private channel")
		fmt.Println("  slackapi groups.setPurpose [channel] [purpose]     Sets the purpose for a private channel")
		fmt.Println("  slackapi groups.setTopic [channel] [topic]         Sets the topic for a private channel")
		fmt.Println("  slackapi im.close [channel]                        Close a direct message channel")
		fmt.Println("  slackapi im.history [channel] [time]               Fetches history of messages and events from direct message channel")
		fmt.Println("  slackapi im.list                                   Lists direct message channels for the calling user")
		fmt.Println("  slackapi im.mark [channel] [time]                  Sets the read cursor in a direct message channel")
		fmt.Println("  slackapi im.open [user]                            Opens a direct message channel")
		fmt.Println("  slackapi mpim.list                                 Lists multiparty direct message channels for the calling user")
		fmt.Println("  slackapi reactions.add [name] [channel] [time]     Adds a reaction to an item")
		fmt.Println("  slackapi reactions.get [channel] [time]            Gets reactions for an item")
		fmt.Println("  slackapi reactions.list [user]                     Lists reactions made by a user")
		fmt.Println("  slackapi reactions.remove [name] [channel] [time]  Removes a reaction from an item")
		fmt.Println("  slackapi team.info                                 Gets information about the current team")
		fmt.Println("  slackapi users.getPresence [user]                  Gets user presence information")
		fmt.Println("  slackapi users.info [user]                         Gets information about a user")
		fmt.Println("  slackapi users.list                                Lists all users in a Slack team")
		fmt.Println("  slackapi users.search [user]                       Search users by name or email address")
		fmt.Println("  slackapi users.setActive                           Marks a user as active")
		fmt.Println("  slackapi users.setPresence [presence]              Manually sets user presence")
		fmt.Println("  slackapi version                                   Displays the program version number")
		fmt.Println("  slackapi help                                      Displays usage and program options")
		fmt.Println()
		fmt.Println("Usage (chat.session):")
		fmt.Println("  :token       Sets the token for the chat session")
		fmt.Println("  :history     Displays the messages in the current session")
		fmt.Println("  :open        Opens a new session with a user, channel, or group")
		fmt.Println("  :delete      Deletes the latest message in the session history")
		fmt.Println("  :flush       Deletes all the messages in the session history")
		fmt.Println("  :exec        Executes and sends the output of a local command")
		fmt.Println("  :execv       Same as :exec but includes the executed command")
		fmt.Println("  :boton       Activates the robot to send 3rd-party messages")
		fmt.Println("  :botoff      Deactivates the robot to send normal messages")
		fmt.Println("  :botinfo     Displays the configuration of the robot")
		fmt.Println("  :botname     Sets the user name of the robot")
		fmt.Println("  :botimage    Sets the avatar for the robot")
		fmt.Println("  :userid      Displays the unique identifier of an user")
		fmt.Println("  :userlist    Displays the information of all the users")
		fmt.Println("  :usersearch  Searches the information of a specific user")
		flag.PrintDefaults()
	}

	flag.Parse()
	client.AutoConfigure()

	switch flag.Arg(0) {
	case "api.test":
		client.ApiTest()
	case "auth.test":
		client.AuthTestVerbose()
	case "channels.history":
		client.ChannelsHistory(flag.Arg(1), flag.Arg(2))
	case "channels.info":
		client.ChannelsInfo(flag.Arg(1))
	case "channels.list":
		client.ChannelsListVerbose()
	case "channels.mark":
		client.ChannelsMark(flag.Arg(1), flag.Arg(2))
	case "channels.setPurpose":
		client.ChannelsSetPurpose(flag.Arg(1), flag.Arg(2))
	case "channels.setTopic":
		client.ChannelsSetTopic(flag.Arg(1), flag.Arg(2))
	case "chat.delete":
		client.ChatDeleteVerbose(flag.Arg(1), flag.Arg(2))
	case "chat.postMessage":
		client.ChatPostMessageVerbose(flag.Arg(1), flag.Arg(2))
	case "chat.session":
		client.ChatSession()
	case "chat.update":
		client.ChatUpdateVerbose(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "emoji.list":
		client.EmojiList()
	case "groups.close":
		client.GroupsClose(flag.Arg(1))
	case "groups.history":
		client.GroupsHistory(flag.Arg(1), flag.Arg(2))
	case "groups.info":
		client.GroupsInfo(flag.Arg(1))
	case "groups.list":
		client.GroupsListVerbose()
	case "groups.mark":
		client.GroupsMark(flag.Arg(1), flag.Arg(2))
	case "groups.open":
		client.GroupsOpenVerbose(flag.Arg(1))
	case "groups.setPurpose":
		client.GroupsSetPurpose(flag.Arg(1), flag.Arg(2))
	case "groups.setTopic":
		client.GroupsSetTopic(flag.Arg(1), flag.Arg(2))
	case "im.close":
		client.InstantMessagingCloseVerbose(flag.Arg(1))
	case "im.history":
		client.InstantMessagingHistory(flag.Arg(1), flag.Arg(2))
	case "im.list":
		client.InstantMessagingList()
	case "im.mark":
		client.InstantMessagingMark(flag.Arg(1), flag.Arg(2))
	case "im.open":
		client.InstantMessagingOpenVerbose(flag.Arg(1))
	case "mpim.list":
		client.MultiPartyInstantMessagingList()
	case "reactions.add":
		client.ReactionsAdd(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "reactions.get":
		client.ReactionsGet(flag.Arg(1), flag.Arg(2))
	case "reactions.list":
		client.ReactionsList(flag.Arg(1))
	case "reactions.remove":
		client.ReactionsRemove(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "team.info":
		client.TeamInfo()
	case "users.getPresence":
		client.UsersGetPresence(flag.Arg(1))
	case "users.info":
		client.UsersInfo(flag.Arg(1))
	case "users.list":
		client.UsersListVerbose()
	case "users.search":
		client.UsersSearchVerbose(flag.Arg(1))
	case "users.setActive":
		client.UsersSetActive()
	case "users.setPresence":
		client.UsersSetPresence(flag.Arg(1))
	case "version":
		fmt.Println(version)
	case "help":
		flag.Usage()
	}

	os.Exit(0)
}
