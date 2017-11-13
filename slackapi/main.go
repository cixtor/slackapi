package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cixtor/slackapi"
)

func main() {
	var command string

	client := slackapi.New()

	flag.Usage = func() {
		fmt.Println("Slack API Client")
		fmt.Println("  http://cixtor.com/")
		fmt.Println("  https://api.slack.com/")
		fmt.Println("  https://github.com/cixtor/slackapi")
		fmt.Println("  version", client.Version())
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
		fmt.Println("  slackapi auth.revoke                                     Revokes a token")
		fmt.Println("  slackapi auth.test                                       Checks authentication and identity")
		fmt.Println("  slackapi bots.info [bot]                                 Gets information about a bot user")
		fmt.Println("  slackapi channels.archive [channel]                      Archives a channel")
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
		fmt.Println("  slackapi channels.suggestions                            Prints a list of suggested channels to join")
		fmt.Println("  slackapi channels.unarchive [channel]                    Unarchives a channel")
		fmt.Println("  slackapi chat.delete [channel] [time]                    Deletes a message")
		fmt.Println("  slackapi chat.meMessage [channel] [text]                 Share a me message into a channel")
		fmt.Println("  slackapi chat.postAttachment [channel] [json]            Sends an attachment to a channel")
		fmt.Println("  slackapi chat.postMessage [channel] [text]               Sends a message to a channel")
		fmt.Println("  slackapi chat.robotMessage [channel] [text]              Sends a message to a channel as a robot")
		fmt.Println("  slackapi chat.session                                    Starts a new chat session")
		fmt.Println("  slackapi chat.update [channel] [time] [text]             Updates a message")
		fmt.Println("  slackapi dnd.endDnd                                      Ends the current user's \"Do Not Disturb\" session immediately")
		fmt.Println("  slackapi dnd.endSnooze                                   Ends the current user's snooze mode immediately")
		fmt.Println("  slackapi dnd.info [user]                                 Retrieves a user's current \"Do Not Disturb\" status")
		fmt.Println("  slackapi dnd.setSnooze                                   Ends the current user's snooze mode immediately")
		fmt.Println("  slackapi dnd.teamInfo [users]                            Retrieves the \"Do Not Disturb\" status for users on a team")
		fmt.Println("  slackapi emoji.list                                      Lists custom emoji for a team")
		fmt.Println("  slackapi eventlog.history [time]                         Lists all the events since the specified time")
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
		fmt.Println("  slackapi files.upload [channel] [filename]               Uploads or creates a file from local data")
		fmt.Println("  slackapi groups.archive [channel]                        Archives a private channel")
		fmt.Println("  slackapi groups.close [channel]                          Closes a private channel")
		fmt.Println("  slackapi groups.create [channel]                         Creates a private channel")
		fmt.Println("  slackapi groups.createChild [channel]                    Clones and archives a private channel")
		fmt.Println("  slackapi groups.history [channel] [time]                 Fetches history of messages and events from a private channel")
		fmt.Println("  slackapi groups.id [channel]                             Gets private channel identifier from readable name")
		fmt.Println("  slackapi groups.info [channel]                           Gets information about a private channel")
		fmt.Println("  slackapi groups.invite [channel] [user]                  Invites a user to a private channel")
		fmt.Println("  slackapi groups.kick [channel] [user]                    Removes a user from a private channel")
		fmt.Println("  slackapi groups.leave [channel]                          Leaves a private channel")
		fmt.Println("  slackapi groups.list                                     Lists private channels that the calling user has access to")
		fmt.Println("  slackapi groups.mark [channel] [time]                    Sets the read cursor in a private channel")
		fmt.Println("  slackapi groups.myHistory [channel] [time]               Displays messages of the current user from a private channel")
		fmt.Println("  slackapi groups.open [group]                             Opens a private channel")
		fmt.Println("  slackapi groups.rename [channel] [name]                  Renames a private channel")
		fmt.Println("  slackapi groups.purgeHistory [channel] [time]            Deletes history of messages and events from a private channel")
		fmt.Println("  slackapi groups.setPurpose [channel] [purpose]           Sets the purpose for a private channel")
		fmt.Println("  slackapi groups.setRetention [channel] [duration]        Sets the retention time of the messages")
		fmt.Println("  slackapi groups.setTopic [channel] [topic]               Sets the topic for a private channel")
		fmt.Println("  slackapi groups.unarchive [channel]                      Unarchives a private channel")
		fmt.Println("  slackapi help.issues.list                                List issues reported by the current user")
		fmt.Println("  slackapi im.close [channel]                              Close a direct message channel")
		fmt.Println("  slackapi im.history [channel] [time]                     Fetches history of messages and events from direct message channel")
		fmt.Println("  slackapi im.list                                         Lists direct message channels for the calling user")
		fmt.Println("  slackapi im.mark [channel] [time]                        Sets the read cursor in a direct message channel")
		fmt.Println("  slackapi im.myHistory [channel] [time]                   Displays messages of the current user from direct message channel")
		fmt.Println("  slackapi im.open [user]                                  Opens a direct message channel")
		fmt.Println("  slackapi im.purgeHistory [channel] [time]                Deletes history of messages and events from direct message channel")
		fmt.Println("  slackapi mpim.close [channel]                            Closes a multiparty direct message channel")
		fmt.Println("  slackapi mpim.history [channel] [time]                   Fetches history of messages and events from a multiparty direct message")
		fmt.Println("  slackapi mpim.list                                       Lists multiparty direct message channels for the calling user")
		fmt.Println("  slackapi mpim.listSimple                                 Lists ID and members in a multiparty direct message channels")
		fmt.Println("  slackapi mpim.mark [channel] [time]                      Sets the read cursor in a multiparty direct message channel")
		fmt.Println("  slackapi mpim.myHistory [channel] [time]                 Displays messages of the current user from multiparty direct message channel")
		fmt.Println("  slackapi mpim.open [user1,user2,etc]                     This method opens a multiparty direct message")
		fmt.Println("  slackapi mpim.purgeHistory [channel] [time]              Deletes history of messages and events from multiparty direct message channel")
		fmt.Println("  slackapi reactions.add [channel] [time] [name]           Adds a reaction to an item")
		fmt.Println("  slackapi reactions.get [channel] [time]                  Gets reactions for an item")
		fmt.Println("  slackapi reactions.list [user]                           Lists reactions made by a user")
		fmt.Println("  slackapi reactions.remove [channel] [time] [name]        Removes a reaction from an item")
		fmt.Println("  slackapi rtm.start                                       Starts a Real Time Messaging session")
		fmt.Println("  slackapi rtm.events                                      Prints the API events in real time")
		fmt.Println("  slackapi team.accessLogs [count] [page]                  Gets the access logs for the current team")
		fmt.Println("  slackapi team.billableInfo [user]                        Gets billable users information for the current team")
		fmt.Println("  slackapi team.info                                       Gets information about the current team")
		fmt.Println("  slackapi team.profile.get                                Retrieve a team's profile")
		fmt.Println("  slackapi users.counts                                    Count number of users in the team")
		fmt.Println("  slackapi users.deletePhoto                               Delete the user avatar")
		fmt.Println("  slackapi users.getPresence [user]                        Gets user presence information")
		fmt.Println("  slackapi users.id [user]                                 Gets user identifier from username")
		fmt.Println("  slackapi users.identity                                  Get a user's identity")
		fmt.Println("  slackapi users.info [user]                               Gets information about a user")
		fmt.Println("  slackapi users.list                                      Lists all users in a Slack team")
		fmt.Println("  slackapi users.prefs.get                                 Get user account preferences")
		fmt.Println("  slackapi users.prefs.set [name] [value]                  Set user account preferences")
		fmt.Println("  slackapi users.preparePhoto [image]                      Upload a picture to use as the avatar")
		fmt.Println("  slackapi users.profile.get [user]                        Retrieves a user's profile information")
		fmt.Println("  slackapi users.profile.set [name] [value]                Set the profile information for a user")
		fmt.Println("  slackapi users.search [user]                             Search users by name or email address")
		fmt.Println("  slackapi users.setActive                                 Marks a user as active")
		fmt.Println("  slackapi users.setAvatar [image]                         Upload a picture and set it as the avatar")
		fmt.Println("  slackapi users.setEmail [email]                          Changes the email address without confirmation")
		fmt.Println("  slackapi users.setPhoto [image_id]                       Define which picture will be the avatar")
		fmt.Println("  slackapi users.setPresence [presence]                    Manually sets user presence")
		fmt.Println("  slackapi users.setStatus [emoji] [text]                  Set the status message and emoji")
		fmt.Println("  slackapi users.setUsername [username]                    Changes the username without admin privileges")
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
		fmt.Println("  :token       Sets the token for the chat session")
		fmt.Println("  :update      Updates the latest chat session message")
		fmt.Println("  :userid      Displays the unique identifier of an user")
		fmt.Println("  :userlist    Displays the information of all the users")
		fmt.Println("  :usersearch  Searches the information of a specific user")
		fmt.Println()
		fmt.Println("Usage (environment):")
		fmt.Println("  SLACK_TOKEN        Set global token for all actions")
		fmt.Println("  SLACK_VERBOSE      Print debug messages and HTTP requests")
		fmt.Println("  SLACK_ROBOT_NAME   Name to identify user with chat.robotMessage")
		fmt.Println("  SLACK_ROBOT_IMAGE  Avatar to identify user with chat.robotMessage")

		flag.PrintDefaults()
	}

	flag.Parse()
	client.AutoConfigure()
	command = flag.Arg(0)

	if command == "" {
		command = "help"
	}

	if command == "chat.session" {
		session := NewSession()
		session.AutoConfigure()
		session.StartSession()
		os.Exit(0)
	}

	switch command {

	case "api.test":
		PrintAndExit(client.APITest())

	case "apps.list":
		PrintAndExit(client.AppsList())

	case "auth.revoke":
		PrintAndExit(client.AuthRevoke())

	case "auth.test":
		PrintAndExit(client.AuthTest())

	case "bots.info":
		PrintAndExit(client.BotsInfo(flag.Arg(1)))

	case "channels.archive":
		PrintAndExit(client.ChannelsArchive(flag.Arg(1)))

	case "channels.create":
		PrintAndExit(client.ChannelsCreate(flag.Arg(1)))

	case "channels.history":
		PrintAndExit(client.ChannelsHistory(slackapi.HistoryArgs{
			Channel: flag.Arg(1),
			Latest:  flag.Arg(2),
		}))

	case "channels.id":
		PrintAndExit(client.ChannelsID(flag.Arg(1)))

	case "channels.info":
		PrintAndExit(client.ChannelsInfo(flag.Arg(1)))

	case "channels.invite":
		PrintAndExit(client.ChannelsInvite(flag.Arg(1), flag.Arg(2)))

	case "channels.join":
		PrintAndExit(client.ChannelsJoin(flag.Arg(1)))

	case "channels.kick":
		PrintAndExit(client.ChannelsKick(flag.Arg(1), flag.Arg(2)))

	case "channels.leave":
		PrintAndExit(client.ChannelsLeave(flag.Arg(1)))

	case "channels.list":
		PrintAndExit(client.ChannelsList())

	case "channels.mark":
		PrintAndExit(client.ChannelsMark(flag.Arg(1), flag.Arg(2)))

	case "channels.myHistory":
		PrintAndExit(client.ChannelsMyHistory(flag.Arg(1), flag.Arg(2)))

	case "channels.purgeHistory":
		client.ChannelsPurgeHistory(flag.Arg(1), flag.Arg(2), true)

	case "channels.rename":
		PrintAndExit(client.ChannelsRename(flag.Arg(1), flag.Arg(2)))

	case "channels.setPurpose":
		PrintAndExit(client.ChannelsSetPurpose(flag.Arg(1), flag.Arg(2)))

	case "channels.setRetention":
		num, err := strconv.Atoi(flag.Arg(2))
		if err != nil {
			fmt.Println("duration format;", err)
			os.Exit(1)
		}
		PrintAndExit(client.ChannelsSetRetention(flag.Arg(1), num))

	case "channels.setTopic":
		PrintAndExit(client.ChannelsSetTopic(flag.Arg(1), flag.Arg(2)))

	case "channels.suggestions":
		PrintAndExit(client.ChannelsSuggestions())

	case "channels.unarchive":
		PrintAndExit(client.ChannelsUnarchive(flag.Arg(1)))

	case "chat.delete":
		PrintAndExit(client.ChatDelete(slackapi.MessageArgs{
			Channel: flag.Arg(1),
			Ts:      flag.Arg(2),
		}))

	case "chat.meMessage":
		PrintAndExit(client.ChatMeMessage(slackapi.MessageArgs{
			Channel: flag.Arg(1),
			Text:    flag.Arg(2),
		}))

	case "chat.postAttachment":
		var data slackapi.Attachment
		if err := json.Unmarshal([]byte(flag.Arg(2)), &data); err != nil {
			fmt.Printf("json unmarshal; %s\n", err.Error())
			os.Exit(1)
		}
		PrintAndExit(client.ChatPostMessage(slackapi.MessageArgs{
			Channel:     flag.Arg(1),
			Attachments: []slackapi.Attachment{data},
		}))

	case "chat.postMessage":
		PrintAndExit(client.ChatPostMessage(slackapi.MessageArgs{
			Channel: flag.Arg(1),
			Text:    flag.Arg(2),
		}))

	case "chat.robotMessage":
		robotName := os.Getenv("SLACK_ROBOT_NAME")
		robotImage := os.Getenv("SLACK_ROBOT_IMAGE")
		data := slackapi.MessageArgs{
			Channel: flag.Arg(1),
			Text:    flag.Arg(2),
			AsUser:  false,
		}
		if robotName == "" {
			robotName = "foobar"
		}
		if robotImage == "" {
			robotImage = ":slack:"
		}
		data.Username = robotName
		if robotImage[0] == ':' {
			data.IconEmoji = robotImage
		} else {
			data.IconURL = robotImage
		}
		PrintAndExit(client.ChatPostMessage(data))

	case "chat.update":
		PrintAndExit(client.ChatUpdate(slackapi.MessageArgs{
			Channel: flag.Arg(1),
			Ts:      flag.Arg(2),
			Text:    flag.Arg(3),
		}))

	case "dnd.endDnd":
		PrintAndExit(client.DNDEndDnd())

	case "dnd.endSnooze":
		PrintAndExit(client.DNDEndSnooze())

	case "dnd.info":
		PrintAndExit(client.DNDInfo(flag.Arg(1)))

	case "dnd.setSnooze":
		num, err := strconv.Atoi(flag.Arg(1))
		if err != nil {
			fmt.Println("minutes format;", err)
			os.Exit(1)
		}
		PrintAndExit(client.DNDSetSnooze(num))

	case "dnd.teamInfo":
		PrintAndExit(client.DNDTeamInfo(flag.Arg(1)))

	case "emoji.list":
		PrintAndExit(client.EmojiList())

	case "eventlog.history":
		PrintAndExit(client.EventlogHistory(flag.Arg(1)))

	case "files.comments.add":
		PrintAndExit(client.FilesCommentsAdd(flag.Arg(1), flag.Arg(2)))

	case "files.comments.delete":
		PrintAndExit(client.FilesCommentsDelete(flag.Arg(1), flag.Arg(2)))

	case "files.comments.edit":
		PrintAndExit(client.FilesCommentsEdit(flag.Arg(1), flag.Arg(2), flag.Arg(3)))

	case "files.delete":
		PrintAndExit(client.FilesDelete(flag.Arg(1)))

	case "files.info":
		numc, err := strconv.Atoi(flag.Arg(2))
		if err != nil {
			fmt.Println("count format;", err)
			os.Exit(1)
		}
		nump, err := strconv.Atoi(flag.Arg(3))
		if err != nil {
			fmt.Println("page format;", err)
			os.Exit(1)
		}
		PrintAndExit(client.FilesInfo(flag.Arg(1), numc, nump))

	case "files.list":
		var data slackapi.FileListArgs
		if count := flag.Arg(1); count != "" {
			numc, err := strconv.Atoi(count)
			if err != nil {
				fmt.Println("count format;", err)
				os.Exit(1)
			}
			data.Count = numc
		}
		if page := flag.Arg(2); page != "" {
			nump, err := strconv.Atoi(page)
			if err != nil {
				fmt.Println("page format;", err)
				os.Exit(1)
			}
			data.Page = nump
		}
		PrintAndExit(client.FilesList(data))

	case "files.listAfterTime":
		numc, err := strconv.Atoi(flag.Arg(2))
		if err != nil {
			fmt.Println("count format;", err)
			os.Exit(1)
		}
		nump, err := strconv.Atoi(flag.Arg(3))
		if err != nil {
			fmt.Println("page format;", err)
			os.Exit(1)
		}
		PrintAndExit(client.FilesList(slackapi.FileListArgs{
			TsFrom: flag.Arg(1),
			Count:  numc,
			Page:   nump,
		}))

	case "files.listBeforeTime":
		numc, err := strconv.Atoi(flag.Arg(2))
		if err != nil {
			fmt.Println("count format;", err)
			os.Exit(1)
		}
		nump, err := strconv.Atoi(flag.Arg(3))
		if err != nil {
			fmt.Println("page format;", err)
			os.Exit(1)
		}
		PrintAndExit(client.FilesList(slackapi.FileListArgs{
			TsTo:  flag.Arg(1),
			Count: numc,
			Page:  nump,
		}))

	case "files.listByChannel":
		numc, err := strconv.Atoi(flag.Arg(2))
		if err != nil {
			fmt.Println("count format;", err)
			os.Exit(1)
		}
		nump, err := strconv.Atoi(flag.Arg(3))
		if err != nil {
			fmt.Println("page format;", err)
			os.Exit(1)
		}
		PrintAndExit(client.FilesList(slackapi.FileListArgs{
			Channel: flag.Arg(1),
			Count:   numc,
			Page:    nump,
		}))

	case "files.listByType":
		numc, err := strconv.Atoi(flag.Arg(2))
		if err != nil {
			fmt.Println("count format;", err)
			os.Exit(1)
		}
		nump, err := strconv.Atoi(flag.Arg(3))
		if err != nil {
			fmt.Println("page format;", err)
			os.Exit(1)
		}
		PrintAndExit(client.FilesList(slackapi.FileListArgs{
			Types: flag.Arg(1),
			Count: numc,
			Page:  nump,
		}))

	case "files.listByUser":
		numc, err := strconv.Atoi(flag.Arg(2))
		if err != nil {
			fmt.Println("count format;", err)
			os.Exit(1)
		}
		nump, err := strconv.Atoi(flag.Arg(3))
		if err != nil {
			fmt.Println("page format;", err)
			os.Exit(1)
		}
		PrintAndExit(client.FilesList(slackapi.FileListArgs{
			User:  flag.Arg(1),
			Count: numc,
			Page:  nump,
		}))

	case "files.upload":
		var data slackapi.FileUploadArgs
		data.Channels = flag.Arg(1)
		data.File = "@" + flag.Arg(2)
		// grab last part of the file path.
		if strings.Contains(data.File, "/") {
			index := strings.LastIndex(data.File, "/")
			data.Filename = data.File[index+1 : len(data.File)]
		} else {
			data.Filename = data.File
		}
		// convert file name into a human title.
		index := strings.Index(data.Filename, ".")
		data.Title = data.Filename[0:index]
		data.Title = strings.Replace(data.Title, "-", "\x20", -1)
		PrintAndExit(client.FilesUpload(data))

	case "groups.archive":
		PrintAndExit(client.GroupsArchive(flag.Arg(1)))

	case "groups.close":
		PrintAndExit(client.GroupsClose(flag.Arg(1)))

	case "groups.create":
		PrintAndExit(client.GroupsCreate(flag.Arg(1)))

	case "groups.createChild":
		PrintAndExit(client.GroupsCreateChild(flag.Arg(1)))

	case "groups.history":
		PrintAndExit(client.GroupsHistory(slackapi.HistoryArgs{
			Channel: flag.Arg(1),
			Latest:  flag.Arg(2),
		}))

	case "groups.id":
		PrintAndExit(client.GroupsID(flag.Arg(1)))

	case "groups.info":
		PrintAndExit(client.GroupsInfo(flag.Arg(1)))

	case "groups.invite":
		PrintAndExit(client.GroupsInvite(flag.Arg(1), flag.Arg(2)))

	case "groups.kick":
		PrintAndExit(client.GroupsKick(flag.Arg(1), flag.Arg(2)))

	case "groups.leave":
		PrintAndExit(client.GroupsLeave(flag.Arg(1)))

	case "groups.list":
		PrintAndExit(client.GroupsList())

	case "groups.mark":
		PrintAndExit(client.GroupsMark(flag.Arg(1), flag.Arg(2)))

	case "groups.myHistory":
		PrintAndExit(client.GroupsMyHistory(flag.Arg(1), flag.Arg(2)))

	case "groups.open":
		PrintAndExit(client.GroupsOpen(flag.Arg(1)))

	case "groups.rename":
		PrintAndExit(client.GroupsRename(flag.Arg(1), flag.Arg(2)))

	case "groups.purgeHistory":
		client.GroupsPurgeHistory(flag.Arg(1), flag.Arg(2), true)

	case "groups.setPurpose":
		PrintAndExit(client.GroupsSetPurpose(flag.Arg(1), flag.Arg(2)))

	case "groups.setRetention":
		num, err := strconv.Atoi(flag.Arg(2))
		if err != nil {
			fmt.Println("duration format;", err)
			os.Exit(1)
		}
		PrintAndExit(client.GroupsSetRetention(flag.Arg(1), num))

	case "groups.setTopic":
		PrintAndExit(client.GroupsSetTopic(flag.Arg(1), flag.Arg(2)))

	case "groups.unarchive":
		PrintAndExit(client.GroupsUnarchive(flag.Arg(1)))

	case "help.issues.list":
		PrintAndExit(client.HelpIssuesList())

	case "im.close":
		PrintAndExit(client.InstantMessageClose(flag.Arg(1)))

	case "im.history":
		PrintAndExit(client.InstantMessageHistory(slackapi.HistoryArgs{
			Channel: flag.Arg(1),
			Latest:  flag.Arg(2),
		}))

	case "im.list":
		PrintAndExit(client.InstantMessageList())

	case "im.mark":
		PrintAndExit(client.InstantMessageMark(flag.Arg(1), flag.Arg(2)))

	case "im.myHistory":
		PrintAndExit(client.InstantMessageMyHistory(flag.Arg(1), flag.Arg(2)))

	case "im.open":
		PrintAndExit(client.InstantMessageOpen(flag.Arg(1)))

	case "im.purgeHistory":
		client.InstantMessagePurgeHistory(flag.Arg(1), flag.Arg(2), true)

	case "mpim.close":
		PrintAndExit(client.MultiPartyInstantMessageClose(flag.Arg(1)))

	case "mpim.history":
		PrintAndExit(client.MultiPartyInstantMessageHistory(slackapi.HistoryArgs{
			Channel: flag.Arg(1),
			Latest:  flag.Arg(2),
		}))

	case "mpim.list":
		PrintAndExit(client.MultiPartyInstantMessageList())

	case "mpim.listSimple":
		PrintAndExit(client.MultiPartyInstantMessageListSimple())

	case "mpim.mark":
		PrintAndExit(client.MultiPartyInstantMessageMark(flag.Arg(1), flag.Arg(2)))

	case "mpim.myHistory":
		PrintAndExit(client.MultiPartyInstantMessageMyHistory(flag.Arg(1), flag.Arg(2)))

	case "mpim.open":
		PrintAndExit(client.MultiPartyInstantMessageOpen(strings.Split(flag.Arg(1), ",")))

	case "mpim.purgeHistory":
		client.MultiPartyInstantMessagePurgeHistory(flag.Arg(1), flag.Arg(2), true)

	case "reactions.add":
		PrintAndExit(client.ReactionsAdd(slackapi.ReactionArgs{
			Channel:   flag.Arg(1),
			Timestamp: flag.Arg(2),
			Name:      flag.Arg(3),
		}))

	case "reactions.get":
		PrintAndExit(client.ReactionsGet(slackapi.ReactionArgs{
			Channel:   flag.Arg(1),
			Timestamp: flag.Arg(2),
		}))

	case "reactions.list":
		PrintAndExit(client.ReactionsList(slackapi.ReactionListArgs{
			User: flag.Arg(1),
		}))

	case "reactions.remove":
		PrintAndExit(client.ReactionsRemove(slackapi.ReactionArgs{
			Channel:   flag.Arg(1),
			Timestamp: flag.Arg(2),
			Name:      flag.Arg(3),
		}))

	case "rtm.events":
		MonitorRealTimeMessages(client)

	case "team.accessLogs":
		numc, err := strconv.Atoi(flag.Arg(2))
		if err != nil {
			fmt.Println("count format;", err)
			os.Exit(1)
		}
		nump, err := strconv.Atoi(flag.Arg(3))
		if err != nil {
			fmt.Println("page format;", err)
			os.Exit(1)
		}
		PrintAndExit(client.TeamAccessLogs(slackapi.AccessLogArgs{
			Count: numc,
			Page:  nump,
		}))

	case "team.billableInfo":
		PrintAndExit(client.TeamBillableInfo(flag.Arg(1)))

	case "team.info":
		PrintAndExit(client.TeamInfo())

	case "team.profile.get":
		PrintAndExit(client.TeamProfileGet())

	case "users.counts":
		PrintAndExit(client.UsersCounts())

	case "users.deletePhoto":
		PrintAndExit(client.UsersDeletePhoto())

	case "users.getPresence":
		PrintAndExit(client.UsersGetPresence(flag.Arg(1)))

	case "users.id":
		PrintAndExit(client.UsersID(flag.Arg(1)))

	case "users.identity":
		PrintAndExit(client.UsersIdentity())

	case "users.info":
		PrintAndExit(client.UsersInfo(flag.Arg(1)))

	case "users.list":
		PrintAndExit(client.UsersList())

	case "users.prefs.get":
		PrintAndExit(client.UsersPrefsGet())

	case "users.prefs.set":
		PrintAndExit(client.UsersPrefsSet(flag.Arg(1), flag.Arg(2)))

	case "users.preparePhoto":
		PrintAndExit(client.UsersPreparePhoto(flag.Arg(1)))

	case "users.profile.get":
		PrintAndExit(client.UsersProfileGet(flag.Arg(1)))

	case "users.profile.set":
		PrintAndExit(client.UsersProfileSet(flag.Arg(1), flag.Arg(2)))

	case "users.search":
		PrintAndExit(client.UsersSearch(flag.Arg(1)))

	case "users.setActive":
		PrintAndExit(client.UsersSetActive())

	case "users.setAvatar":
		PrintAndExit(client.UsersSetAvatar(flag.Arg(1)))

	case "users.setEmail":
		PrintAndExit(client.UsersProfileSet("email", flag.Arg(1)))

	case "users.setPhoto":
		PrintAndExit(client.UsersSetPhoto(flag.Arg(1)))

	case "users.setPresence":
		PrintAndExit(client.UsersSetPresence(flag.Arg(1)))

	case "users.setStatus":
		PrintAndExit(client.UsersSetStatus(flag.Arg(1), flag.Arg(2)))

	case "users.setUsername":
		PrintAndExit(client.UsersProfileSet("username", flag.Arg(1)))

	case "version":
		fmt.Println(client.Version())

	case "help":
		flag.Usage()
	}

	os.Exit(0)
}
