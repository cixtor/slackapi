package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "1.5.18"

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
		fmt.Println("  slackapi api.test                                        Checks API calling code")
		fmt.Println("  slackapi apps.list                                       Lists associated applications")
		fmt.Println("  slackapi auth.test                                       Checks authentication and identity")
		fmt.Println("  slackapi channels.create [channel]                       Creates a channel if authorized")
		fmt.Println("  slackapi channels.history [channel] [time]               Fetches history of messages and events from a channel")
		fmt.Println("  slackapi channels.id [channel]                           Gets channel identifier from readable name")
		fmt.Println("  slackapi channels.info [channel]                         Gets information about a channel")
		fmt.Println("  slackapi channels.invite [channel] [user]                Invites a user to a channel")
		fmt.Println("  slackapi channels.join [channel]                         Joins a channel, creates it if it does not exist")
		fmt.Println("  slackapi channels.kick [channel] [user]                  Removes a user from a channel")
		fmt.Println("  slackapi channels.leave [channel]                        Leaves a channel")
		fmt.Println("  slackapi channels.list                                   Lists all channels in a Slack team")
		fmt.Println("  slackapi channels.mark [channel] [time]                  Sets the read cursor in a channel")
		fmt.Println("  slackapi channels.myHistory [channel] [time]             Displays messages of the current user from a channel")
		fmt.Println("  slackapi channels.purgeHistory [channel] [time]          Deletes history of messages and events from a channel")
		fmt.Println("  slackapi channels.rename [channel] [name]                Renames a channel")
		fmt.Println("  slackapi channels.setPurpose [channel] [purpose]         Sets the purpose for a channel")
		fmt.Println("  slackapi channels.setRetention [channel] [duration]      Sets the retention time of the messages")
		fmt.Println("  slackapi channels.setTopic [channel] [topic]             Sets the topic for a channel")
		fmt.Println("  slackapi chat.delete [channel] [time]                    Deletes a message")
		fmt.Println("  slackapi chat.postMessage [channel] [text]               Sends a message to a channel")
		fmt.Println("  slackapi chat.session                                    Starts a new chat session")
		fmt.Println("  slackapi chat.update [channel] [time] [text]             Updates a message")
		fmt.Println("  slackapi emoji.list                                      Lists custom emoji for a team")
		fmt.Println("  slackapi files.comments.add [file] [text]                Add a comment to an existing file")
		fmt.Println("  slackapi files.comments.delete [file] [fcid]             Deletes an existing comment on a file")
		fmt.Println("  slackapi files.comments.edit [file] [fcid] [text]        Edit an existing file comment")
		fmt.Println("  slackapi files.delete [file]                             Deletes a file and associated comments")
		fmt.Println("  slackapi files.info [file] [count] [page]                Gets information about a team file")
		fmt.Println("  slackapi files.list [count] [page]                       Lists and filters team files")
		fmt.Println("  slackapi files.listAfterTime [time] [count] [page]       Lists and filters team files after this timestamp (inclusive)")
		fmt.Println("  slackapi files.listBeforeTime [time] [count] [page]      Lists and filters team files before this timestamp (inclusive)")
		fmt.Println("  slackapi files.listByChannel [channel] [count] [page]    Lists and filters team files in a specific channel")
		fmt.Println("  slackapi files.listByType [type] [count] [page]          Lists and filters team files by type: all, posts, snippets, images, gdocs, zips, pdfs")
		fmt.Println("  slackapi files.listByUser [user] [count] [page]          Lists and filters team files created by a single user")
		fmt.Println("  slackapi files.upload [channel] [fpath]                  Uploads or creates a file from local data")
		fmt.Println("  slackapi groups.close [channel]                          Closes a private channel")
		fmt.Println("  slackapi groups.history [channel] [time]                 Fetches history of messages and events from a private channel")
		fmt.Println("  slackapi groups.id [channel]                             Gets private channel identifier from readable name")
		fmt.Println("  slackapi groups.info [channel]                           Gets information about a private channel")
		fmt.Println("  slackapi groups.list                                     Lists private channels that the calling user has access to")
		fmt.Println("  slackapi groups.mark [channel] [time]                    Sets the read cursor in a private channel")
		fmt.Println("  slackapi groups.myHistory [channel] [time]               Displays messages of the current user from a private channel")
		fmt.Println("  slackapi groups.open [group]                             Opens a private channel")
		fmt.Println("  slackapi groups.purgeHistory [channel] [time]            Deletes history of messages and events from a private channel")
		fmt.Println("  slackapi groups.setPurpose [channel] [purpose]           Sets the purpose for a private channel")
		fmt.Println("  slackapi groups.setRetention [channel] [duration]        Sets the retention time of the messages")
		fmt.Println("  slackapi groups.setTopic [channel] [topic]               Sets the topic for a private channel")
		fmt.Println("  slackapi im.close [channel]                              Close a direct message channel")
		fmt.Println("  slackapi im.history [channel] [time]                     Fetches history of messages and events from direct message channel")
		fmt.Println("  slackapi im.list                                         Lists direct message channels for the calling user")
		fmt.Println("  slackapi im.mark [channel] [time]                        Sets the read cursor in a direct message channel")
		fmt.Println("  slackapi im.myHistory [channel] [time]                   Displays messages of the current user from direct message channel")
		fmt.Println("  slackapi im.open [user]                                  Opens a direct message channel")
		fmt.Println("  slackapi im.purgeHistory [channel] [time]                Deletes history of messages and events from direct message channel")
		fmt.Println("  slackapi mpim.list                                       Lists multiparty direct message channels for the calling user")
		fmt.Println("  slackapi reactions.add [name] [channel] [time]           Adds a reaction to an item")
		fmt.Println("  slackapi reactions.get [channel] [time]                  Gets reactions for an item")
		fmt.Println("  slackapi reactions.list [user]                           Lists reactions made by a user")
		fmt.Println("  slackapi reactions.remove [name] [channel] [time]        Removes a reaction from an item")
		fmt.Println("  slackapi team.accessLogs [count] [page]                  Gets the access logs for the current team")
		fmt.Println("  slackapi team.info                                       Gets information about the current team")
		fmt.Println("  slackapi users.getPresence [user]                        Gets user presence information")
		fmt.Println("  slackapi users.id [user]                                 Gets user identifier from username")
		fmt.Println("  slackapi users.info [user]                               Gets information about a user")
		fmt.Println("  slackapi users.list                                      Lists all users in a Slack team")
		fmt.Println("  slackapi users.search [user]                             Search users by name or email address")
		fmt.Println("  slackapi users.setActive                                 Marks a user as active")
		fmt.Println("  slackapi users.setPresence [presence]                    Manually sets user presence")
		fmt.Println("  slackapi version                                         Displays the program version number")
		fmt.Println("  slackapi help                                            Displays usage and program options")
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
		fmt.Println("  :update      Updates the latest chat session message")
		fmt.Println("  :userid      Displays the unique identifier of an user")
		fmt.Println("  :userlist    Displays the information of all the users")
		fmt.Println("  :usersearch  Searches the information of a specific user")
		fmt.Println()
		fmt.Println("Usage (debug):")
		fmt.Println("  VERBOSE=true Environment variable to print debug messages")
		flag.PrintDefaults()
		os.Exit(2)
	}

	flag.Parse()
	client.AutoConfigure()
	command = flag.Arg(0)

	if command == "" {
		command = "help"
	}

	if command == "chat.session" {
		var session ChatSession
		session.AutoConfigure()
		session.StartChatSession()
		os.Exit(0)
	}

	switch command {
	case "api.test":
		client.PrintApiTest()
	case "apps.list":
		client.PrintAppsList()
	case "auth.test":
		client.PrintAuthTest()
	case "channels.create":
		client.PrintChannelsCreate(flag.Arg(1))
	case "channels.history":
		client.PrintChannelsHistory(flag.Arg(1), flag.Arg(2))
	case "channels.id":
		client.PrintChannelsId(flag.Arg(1))
	case "channels.info":
		client.PrintChannelsInfo(flag.Arg(1))
	case "channels.invite":
		client.PrintChannelsInvite(flag.Arg(1), flag.Arg(2))
	case "channels.join":
		client.PrintChannelsJoin(flag.Arg(1))
	case "channels.kick":
		client.PrintChannelsKick(flag.Arg(1), flag.Arg(2))
	case "channels.leave":
		client.PrintChannelsLeave(flag.Arg(1))
	case "channels.list":
		client.PrintChannelsList()
	case "channels.mark":
		client.PrintChannelsMark(flag.Arg(1), flag.Arg(2))
	case "channels.myHistory":
		client.PrintChannelsMyHistory(flag.Arg(1), flag.Arg(2))
	case "channels.purgeHistory":
		client.PrintChannelsPurgeHistory(flag.Arg(1), flag.Arg(2))
	case "channels.rename":
		client.PrintChannelsRename(flag.Arg(1), flag.Arg(2))
	case "channels.setPurpose":
		client.PrintChannelsSetPurpose(flag.Arg(1), flag.Arg(2))
	case "channels.setRetention":
		client.PrintChannelsSetRetention(flag.Arg(1), flag.Arg(2))
	case "channels.setTopic":
		client.PrintChannelsSetTopic(flag.Arg(1), flag.Arg(2))
	case "chat.delete":
		client.PrintChatDelete(flag.Arg(1), flag.Arg(2))
	case "chat.postMessage":
		client.PrintChatPostMessage(flag.Arg(1), flag.Arg(2))
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
	case "files.info":
		client.PrintFilesInfo(flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "files.list":
		client.PrintFilesList("none", "", flag.Arg(1), flag.Arg(2))
	case "files.listAfterTime":
		client.PrintFilesList("ts_from", flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "files.listBeforeTime":
		client.PrintFilesList("ts_to", flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "files.listByChannel":
		client.PrintFilesList("channel", flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "files.listByType":
		client.PrintFilesList("types", flag.Arg(1), flag.Arg(2), flag.Arg(3))
	case "files.listByUser":
		client.PrintFilesList("user", flag.Arg(1), flag.Arg(2), flag.Arg(3))
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
	case "groups.setRetention":
		client.PrintGroupsSetRetention(flag.Arg(1), flag.Arg(2))
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
	case "team.accessLogs":
		client.PrintTeamAccessLogs(flag.Arg(1), flag.Arg(2))
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
