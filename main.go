package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var client SlackAPI

	flag.Usage = func() {
		fmt.Println("Slack API Client")
		fmt.Println("  http://cixtor.com/")
		fmt.Println("  https://api.slack.com/")
		fmt.Println("  https://github.com/cixtor/slackapi")
		fmt.Println()
		fmt.Println("Low level Slack API client with custom commands. Slack, the 'messaging app for")
		fmt.Println("teams' offers an API that has been used to build multiple projects around it,")
		fmt.Println("from bots to independent clients as well as integrations with other external")
		fmt.Println("services. This project aims to offer a low level experience for advanced users")
		fmt.Println("that want to either drop the web client or interact with the API for testing")
		fmt.Println("purpose.")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  slackapi api.test")
		fmt.Println("  slackapi auth.test")
		fmt.Println("  slackapi channels.history [channel] [time]")
		fmt.Println("  slackapi channels.info [channel]")
		fmt.Println("  slackapi channels.list")
		fmt.Println("  slackapi channels.mark [channel] [time]")
		fmt.Println("  slackapi channels.setPurpose [channel] [purpose]")
		fmt.Println("  slackapi channels.setTopic [channel] [topic]")
		fmt.Println("  slackapi chat.delete [channel] [time]")
		fmt.Println("  slackapi chat.postMessage [channel] [text]")
		fmt.Println("  slackapi chat.session")
		fmt.Println("  slackapi chat.update [channel] [time] [text]")
		fmt.Println("  slackapi emoji.list")
		fmt.Println("  slackapi groups.close [channel]")
		fmt.Println("  slackapi groups.history [channel] [time]")
		fmt.Println("  slackapi groups.info [channel]")
		fmt.Println("  slackapi groups.list")
		fmt.Println("  slackapi groups.mark [channel] [time]")
		fmt.Println("  slackapi groups.open [group]")
		fmt.Println("  slackapi groups.setPurpose [channel] [purpose]")
		fmt.Println("  slackapi groups.setTopic [channel] [topic]")
		fmt.Println("  slackapi im.close [channel]")
		fmt.Println("  slackapi im.history [channel] [time]")
		fmt.Println("  slackapi im.list")
		fmt.Println("  slackapi im.mark [channel] [time]")
		fmt.Println("  slackapi im.open [user]")
		fmt.Println("  slackapi mpim.list")
		fmt.Println("  slackapi reactions.add [name] [channel] [time]")
		fmt.Println("  slackapi reactions.get [channel] [time]")
		fmt.Println("  slackapi reactions.list [user]")
		fmt.Println("  slackapi reactions.remove [name] [channel] [time]")
		fmt.Println("  slackapi team.info")
		fmt.Println("  slackapi users.getPresence [user]")
		fmt.Println("  slackapi users.info [user]")
		fmt.Println("  slackapi users.list")
		fmt.Println("  slackapi users.search [user]")
		fmt.Println("  slackapi users.setActive")
		fmt.Println("  slackapi users.setPresence [presence]")
		fmt.Println("  slackapi -help")
		fmt.Println()
		fmt.Println("Usage (chat.session):")
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
		client.AuthTest()
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
	case "-help":
		flag.Usage()
	}

	os.Exit(0)
}
