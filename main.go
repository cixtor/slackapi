package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var client SlackAPI

	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("  slackapi api.test")
		fmt.Println("  slackapi auth.test")
		fmt.Println("  slackapi users.list")
		fmt.Println("  slackapi users.search [query]")
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
		client.ChannelsList()
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
		client.GroupsList()
	case "groups.mark":
		client.GroupsMark(flag.Arg(1), flag.Arg(2))
	case "groups.open":
		client.GroupsOpen(flag.Arg(1))
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
