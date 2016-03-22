package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.3.22"

func main() {
	var client SlackAPI
	var command string

	flag.Usage = func() {
		fmt.Println("Slack API Client")
		fmt.Println("  http://cixtor.com/")
		fmt.Println("  https://api.slack.com/")
		fmt.Println("  https://github.com/cixtor/slackapi")
		fmt.Println("  version", version)
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
		fmt.Println("  slackapi channels.myHistory [channel] [time]       Displays messages of the current user from a channel")
		fmt.Println("  slackapi channels.purgeHistory [channel] [time]    Deletes history of messages and events from a channel")
		fmt.Println("  slackapi channels.setPurpose [channel] [purpose]   Sets the purpose for a channel")
		fmt.Println("  slackapi channels.setTopic [channel] [topic]       Sets the topic for a channel")
		fmt.Println("  slackapi chat.delete [channel] [time]              Deletes a message")
		fmt.Println("  slackapi chat.postMessage [channel] [text]         Sends a message to a channel")
		fmt.Println("  slackapi chat.session                              Starts a new chat session")
		fmt.Println("  slackapi chat.update [channel] [time] [text]       Updates a message")
		fmt.Println("  slackapi emoji.list                                Lists custom emoji for a team")
		fmt.Println("  slackapi files.comments.add [file] [text]          Add a comment to an existing file")
		fmt.Println("  slackapi files.comments.delete [file] [fcid]       Deletes an existing comment on a file")
		fmt.Println("  slackapi files.comments.edit [file] [fcid] [text]  Edit an existing file comment")
		fmt.Println("  slackapi files.delete [file]                       Deletes a file and associated comments")
		fmt.Println("  slackapi files.upload [channel] [fpath]            Uploads or creates a file from local data")
		fmt.Println("  slackapi groups.close [channel]                    Closes a private channel")
		fmt.Println("  slackapi groups.history [channel] [time]           Fetches history of messages and events from a private channel")
		fmt.Println("  slackapi groups.id [channel]                       Gets private channel identifier from readable name")
		fmt.Println("  slackapi groups.info [channel]                     Gets information about a private channel")
		fmt.Println("  slackapi groups.list                               Lists private channels that the calling user has access to")
		fmt.Println("  slackapi groups.mark [channel] [time]              Sets the read cursor in a private channel")
		fmt.Println("  slackapi groups.myHistory [channel] [time]         Displays messages of the current user from a private channel")
		fmt.Println("  slackapi groups.open [group]                       Opens a private channel")
		fmt.Println("  slackapi groups.purgeHistory [channel] [time]      Deletes history of messages and events from a private channel")
		fmt.Println("  slackapi groups.setPurpose [channel] [purpose]     Sets the purpose for a private channel")
		fmt.Println("  slackapi groups.setTopic [channel] [topic]         Sets the topic for a private channel")
		fmt.Println("  slackapi im.close [channel]                        Close a direct message channel")
		fmt.Println("  slackapi im.history [channel] [time]               Fetches history of messages and events from direct message channel")
		fmt.Println("  slackapi im.list                                   Lists direct message channels for the calling user")
		fmt.Println("  slackapi im.mark [channel] [time]                  Sets the read cursor in a direct message channel")
		fmt.Println("  slackapi im.myHistory [channel] [time]             Displays messages of the current user from direct message channel")
		fmt.Println("  slackapi im.open [user]                            Opens a direct message channel")
		fmt.Println("  slackapi im.purgeHistory [channel] [time]          Deletes history of messages and events from direct message channel")
		fmt.Println("  slackapi mpim.list                                 Lists multiparty direct message channels for the calling user")
		fmt.Println("  slackapi reactions.add [name] [channel] [time]     Adds a reaction to an item")
		fmt.Println("  slackapi reactions.get [channel] [time]            Gets reactions for an item")
		fmt.Println("  slackapi reactions.list [user]                     Lists reactions made by a user")
		fmt.Println("  slackapi reactions.remove [name] [channel] [time]  Removes a reaction from an item")
		fmt.Println("  slackapi team.info                                 Gets information about the current team")
		fmt.Println("  slackapi users.getPresence [user]                  Gets user presence information")
		fmt.Println("  slackapi users.id [user]                           Gets user identifier from username")
		fmt.Println("  slackapi users.info [user]                         Gets information about a user")
		fmt.Println("  slackapi users.list                                Lists all users in a Slack team")
		fmt.Println("  slackapi users.search [user]                       Search users by name or email address")
		fmt.Println("  slackapi users.setActive                           Marks a user as active")
		fmt.Println("  slackapi users.setPresence [presence]              Manually sets user presence")
		fmt.Println("  slackapi version                                   Displays the program version number")
		fmt.Println("  slackapi help                                      Displays usage and program options")
		fmt.Println()
		fmt.Println("Usage (chat.session):")
		fmt.Println("  :close       Close current chat session")
		fmt.Println("  :delete      Deletes the latest message in the session history")
		fmt.Println("  :exec        Executes and sends the output of a local command")
		fmt.Println("  :execv       Same as :exec but includes the executed command")
		fmt.Println("  :exit        Exits the program without closing chat sessions")
		fmt.Println("  :flush       Deletes all the messages in the session history")
		fmt.Println("  :history     Fetches messages and events in current session")
		fmt.Println("  :messages    Displays the messages in the current session")
		fmt.Println("  :myhistory   Fetches messages and events from current user")
		fmt.Println("  :open        Opens a new session with a user, channel, or group")
		fmt.Println("  :owner       Displays account information of the user in session")
		fmt.Println("  :purge       Deletes the messages in the current session")
		fmt.Println("  :robotimage  Sets the avatar for the robot")
		fmt.Println("  :robotinfo   Displays the configuration of the robot")
		fmt.Println("  :robotname   Sets the user name of the robot")
		fmt.Println("  :robotoff    Deactivates the robot to send normal messages")
		fmt.Println("  :roboton     Activates the robot to send 3rd-party messages")
		fmt.Println("  :token       Sets the token for the chat session")
		fmt.Println("  :userid      Displays the unique identifier of an user")
		fmt.Println("  :userlist    Displays the information of all the users")
		fmt.Println("  :usersearch  Searches the information of a specific user")
		flag.PrintDefaults()
		os.Exit(2)
	}

	flag.Parse()
	client.AutoConfigure()
	command = flag.Arg(0)

	if command == "" {
		command = "help"
	}

	switch command {
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
	case "channels.myHistory":
		client.ChannelsMyHistory(flag.Arg(1), flag.Arg(2))
	case "channels.purgeHistory":
		client.ChannelsPurgeHistory(flag.Arg(1), flag.Arg(2))
	case "channels.setPurpose":
		client.ChannelsSetPurpose(flag.Arg(1), flag.Arg(2))
	case "channels.setTopic":
		client.ChannelsSetTopic(flag.Arg(1), flag.Arg(2))
	case "chat.delete":
		client.PrintChatDelete(flag.Arg(1), flag.Arg(2))
	case "chat.postMessage":
		client.PrintChatPostMessage(flag.Arg(1), flag.Arg(2))
	case "chat.session":
		client.ChatSession()
	case "chat.update":
		client.PrintChatUpdate(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "emoji.list":
		client.PrintEmojiList()
	case "files.comments.add":
		client.PrintFilesCommentsAdd(flag.Arg(1), flag.Arg(2))
	case "files.comments.delete":
		client.PrintFilesCommentsDelete(flag.Arg(1), flag.Arg(2))
	case "files.comments.edit":
		client.PrintFilesCommentsEdit(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "files.delete":
		client.PrintFilesDelete(flag.Arg(1))
	case "files.upload":
		client.PrintFilesUpload(flag.Arg(1), flag.Arg(2))
	case "groups.close":
		client.PrintGroupsClose(flag.Arg(1))
	case "groups.history":
		client.PrintGroupsHistory(flag.Arg(1), flag.Arg(2))
	case "groups.id":
		client.PrintGroupsId(flag.Arg(1))
	case "groups.info":
		client.PrintGroupsInfo(flag.Arg(1))
	case "groups.list":
		client.PrintGroupsList()
	case "groups.mark":
		client.PrintGroupsMark(flag.Arg(1), flag.Arg(2))
	case "groups.myHistory":
		client.PrintGroupsMyHistory(flag.Arg(1), flag.Arg(2))
	case "groups.open":
		client.PrintGroupsOpen(flag.Arg(1))
	case "groups.purgeHistory":
		client.PrintGroupsPurgeHistory(flag.Arg(1), flag.Arg(2))
	case "groups.setPurpose":
		client.PrintGroupsSetPurpose(flag.Arg(1), flag.Arg(2))
	case "groups.setTopic":
		client.PrintGroupsSetTopic(flag.Arg(1), flag.Arg(2))
	case "im.close":
		client.PrintInstantMessagingClose(flag.Arg(1))
	case "im.history":
		client.PrintInstantMessagingHistory(flag.Arg(1), flag.Arg(2))
	case "im.list":
		client.PrintInstantMessagingList()
	case "im.mark":
		client.PrintInstantMessagingMark(flag.Arg(1), flag.Arg(2))
	case "im.myHistory":
		client.PrintInstantMessagingMyHistory(flag.Arg(1), flag.Arg(2))
	case "im.open":
		client.PrintInstantMessagingOpen(flag.Arg(1))
	case "im.purgeHistory":
		client.PrintInstantMessagingPurgeHistory(flag.Arg(1), flag.Arg(2))
	case "mpim.list":
		client.PrintMultiPartyInstantMessagingList()
	case "reactions.add":
		client.PrintReactionsAdd(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "reactions.get":
		client.PrintReactionsGet(flag.Arg(1), flag.Arg(2))
	case "reactions.list":
		client.PrintReactionsList(flag.Arg(1))
	case "reactions.remove":
		client.PrintReactionsRemove(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "team.info":
		client.PrintTeamInfo()
	case "users.getPresence":
		client.PrintUsersGetPresence(flag.Arg(1))
	case "users.id":
		client.PrintUsersId(flag.Arg(1))
	case "users.info":
		client.PrintUsersInfo(flag.Arg(1))
	case "users.list":
		client.PrintUsersList()
	case "users.search":
		client.PrintUsersSearch(flag.Arg(1))
	case "users.setActive":
		client.PrintUsersSetActive()
	case "users.setPresence":
		client.PrintUsersSetPresence(flag.Arg(1))
	case "version":
		fmt.Println(version)
	case "help":
		flag.Usage()
	}

	os.Exit(0)
}
