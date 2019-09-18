# Slack API [![GoReport](https://goreportcard.com/badge/github.com/cixtor/slackapi)](https://goreportcard.com/report/github.com/cixtor/slackapi) [![GoDoc](https://godoc.org/github.com/cixtor/slackapi?status.svg)](https://godoc.org/github.com/cixtor/slackapi)

Slack, the _"messaging app for teams"_ offers an API that has been used to build multiple projects around it, from bots to independent clients as well as integrations with other external services. This project aims to offer a low level experience for advanced users that want to either drop the web client or interact with the API for testing purpose.

## Installation CLI

```
go get -u github.com/cixtor/slackcli
```

## Installation Source

```
package main

import (
    "fmt"
    "github.com/cixtor/slackapi"
)

func main() {
    client := slackapi.New()
    fmt.Println(client.Version())
}
```

## Features

The client is built on top of the [Bot Users](https://api.slack.com/bot-users) documentation. Most if not all the methods available in the API are implemented and can be executed placing a colon character as the suffix of each method.

Note that the client runs with the same chat session of the user that is using the program, but technically speaking the interaction is similar to that of a bot. This offers some advantages, for example, like other APIs and integrations, bot users are free. Unlike regular users, the actions they can perform are somewhat limited. For teams on the Free Plan, each bot user counts as a separate integration.

## Implemented API Endpoints

- :large_blue_circle: means the method has been implemented
- :red_circle: means the method is not implemented yet
- :black_circle: means the method has been deprecated upstream

| :shipit: | Method | Description |
|----------|--------|-------------|
| :large_blue_circle: | `api.test` | Checks API calling code |
| :large_blue_circle: | `apps.list` | Lists associated applications |
| :red_circle: | `apps.permissions.info` | Returns list of permissions this app has on a team |
| :red_circle: | `apps.permissions.request` | Allows an app to request additional scopes |
| :large_blue_circle: | `auth.revoke` | Revokes a token |
| :large_blue_circle: | `auth.test` | Checks authentication and identity |
| :large_blue_circle: | `bots.info` | Gets information about a bot user |
| :large_blue_circle: | `channels.archive` | Archives a channel |
| :large_blue_circle: | `channels.create` | Creates a channel |
| :large_blue_circle: | `channels.history` | Fetches history of messages and events from a channel |
| :large_blue_circle: | `channels.id` | Gets channel identifier from readable name |
| :large_blue_circle: | `channels.info` | Gets information about a channel |
| :large_blue_circle: | `channels.invite` | Invites a user to a channel |
| :large_blue_circle: | `channels.join` | Joins a channel, creating it if needed |
| :large_blue_circle: | `channels.kick` | Removes a user from a channel |
| :large_blue_circle: | `channels.leave` | Leaves a channel |
| :large_blue_circle: | `channels.list` | Lists all channels in a Slack team |
| :large_blue_circle: | `channels.mark` | Sets the read cursor in a channel |
| :large_blue_circle: | `channels.myHistory` | Displays messages of the current user from a channel |
| :large_blue_circle: | `channels.purgeHistory` | Deletes history of messages and events from a channel |
| :large_blue_circle: | `channels.rename` | Renames a channel |
| :red_circle: | `channels.replies` | Retrieve a thread of messages posted to a channel |
| :large_blue_circle: | `channels.setPurpose` | Sets the purpose for a channel |
| :large_blue_circle: | `channels.setRetention` | Sets the retention time of the messages |
| :large_blue_circle: | `channels.setTopic` | Sets the topic for a channel |
| :large_blue_circle: | `channels.suggestions` | Prints a list of suggested channels to join |
| :large_blue_circle: | `channels.unarchive` | Unarchives a channel |
| :large_blue_circle: | `chat.delete` | Deletes a message |
| :red_circle: | `chat.getPermalink` | Retrieve a permalink URL for a specific extant message |
| :large_blue_circle: | `chat.meMessage` | Share a me message into a channel |
| :red_circle: | `chat.postEphemeral` | Sends an ephemeral message to a user in a channel |
| :large_blue_circle: | `chat.postMessage` | Sends a message to a channel |
| :large_blue_circle: | `chat.robotMessage` | Sends a message to a channel as a robot |
| :large_blue_circle: | `chat.session` | Starts a new chat session |
| :red_circle: | `chat.unfurl` | Provide custom unfurl behavior for user-posted URLs |
| :large_blue_circle: | `chat.update` | Updates a message |
| :red_circle: | `conversations.archive` | Archives a conversation |
| :red_circle: | `conversations.close` | Closes a direct message or multi-person direct message |
| :red_circle: | `conversations.create` | Initiates a public or private channel-based conversation |
| :red_circle: | `conversations.history` | Fetches a conversation's history of messages and events |
| :red_circle: | `conversations.info` | Retrieve information about a conversation |
| :red_circle: | `conversations.invite` | Invites users to a channel |
| :red_circle: | `conversations.join` | Joins an existing conversation |
| :red_circle: | `conversations.kick` | Removes a user from a conversation |
| :red_circle: | `conversations.leave` | Leaves a conversation |
| :red_circle: | `conversations.list` | Lists all channels in a Slack team |
| :red_circle: | `conversations.members` | Retrieve members of a conversation |
| :red_circle: | `conversations.open` | Opens or resumes a direct message or multi-person direct message |
| :red_circle: | `conversations.rename` | Renames a conversation |
| :red_circle: | `conversations.replies` | Retrieve a thread of messages posted to a conversation |
| :red_circle: | `conversations.setPurpose` | Sets the purpose for a conversation |
| :red_circle: | `conversations.setTopic` | Sets the topic for a conversation |
| :red_circle: | `conversations.unarchive` | Reverses conversation archival |
| :red_circle: | `dialog.open` | Open a dialog with a user |
| :large_blue_circle: | `dnd.endDnd` | Ends the current user's _"Do Not Disturb"_ session immediately |
| :large_blue_circle: | `dnd.endSnooze` | Ends the current user's snooze mode immediately |
| :large_blue_circle: | `dnd.info` | Retrieves a user's current _"Do Not Disturb"_ status |
| :large_blue_circle: | `dnd.setSnooze` | Turns on _"Do Not Disturb"_ mode for the current user, or changes its duration |
| :large_blue_circle: | `dnd.teamInfo` | Retrieves the _"Do Not Disturb"_ status for users on a team |
| :large_blue_circle: | `emoji.list` | Lists custom emoji for a team |
| :large_blue_circle: | `eventlog.history` | Lists all the events since the specified time |
| :large_blue_circle: | `files.comments.add` | Add a comment to an existing file |
| :large_blue_circle: | `files.comments.delete` | Deletes an existing comment on a file |
| :large_blue_circle: | `files.comments.edit` | Edit an existing file comment |
| :large_blue_circle: | `files.delete` | Deletes a file |
| :large_blue_circle: | `files.info` | Gets information about a team file |
| :large_blue_circle: | `files.list` | Lists and filters team files |
| :large_blue_circle: | `files.listAfterTime` | Lists and filters team files after this timestamp _(inclusive)_ |
| :large_blue_circle: | `files.listBeforeTime` | Lists and filters team files before this timestamp _(inclusive)_ |
| :large_blue_circle: | `files.listByChannel` | Lists and filters team files in a specific channel |
| :large_blue_circle: | `files.listByType` | Lists and filters team files by type: all, posts, snippets, images, gdocs, zips, pdfs |
| :large_blue_circle: | `files.listByUser` | Lists and filters team files created by a single user |
| :large_blue_circle: | `files.revokePublicURL` | Revokes public/external sharing access for a file |
| :large_blue_circle: | `files.sharedPublicURL` | Enables a file for public/external sharing |
| :large_blue_circle: | `files.upload` | Uploads or creates a file |
| :large_blue_circle: | `groups.archive` | Archives a private channel |
| :large_blue_circle: | `groups.close` | Closes a private channel |
| :large_blue_circle: | `groups.create` | Creates a private channel |
| :large_blue_circle: | `groups.createChild` | Clones and archives a private channel |
| :large_blue_circle: | `groups.history` | Fetches history of messages and events from a private channel |
| :large_blue_circle: | `groups.id` | Gets private channel identifier from readable name |
| :large_blue_circle: | `groups.info` | Gets information about a private channel |
| :large_blue_circle: | `groups.invite` | Invites a user to a private channel |
| :large_blue_circle: | `groups.kick` | Removes a user from a private channel |
| :large_blue_circle: | `groups.leave` | Leaves a private channel |
| :large_blue_circle: | `groups.list` | Lists private channels that the calling user has access to |
| :large_blue_circle: | `groups.mark` | Sets the read cursor in a private channel |
| :large_blue_circle: | `groups.myHistory` | Displays messages of the current user from a private channel |
| :large_blue_circle: | `groups.open` | Opens a private channel |
| :large_blue_circle: | `groups.purgeHistory` | Deletes history of messages and events from a private channel |
| :large_blue_circle: | `groups.rename` | Renames a private channel |
| :red_circle: | `groups.replies` | Retrieve a thread of messages posted to a private channel |
| :large_blue_circle: | `groups.setPurpose` | Sets the purpose for a private channel |
| :large_blue_circle: | `groups.setRetention` | Sets the retention time of the messages |
| :large_blue_circle: | `groups.setTopic` | Sets the topic for a private channel |
| :large_blue_circle: | `groups.unarchive` | Unarchives a private channel |
| :large_blue_circle: | `help` | Displays usage and program options |
| :large_blue_circle: | `help.issues.list` | List issues reported by the current user |
| :large_blue_circle: | `im.close` | Close a direct message channel |
| :large_blue_circle: | `im.history` | Fetches history of messages and events from direct message channel |
| :large_blue_circle: | `im.list` | Lists direct message channels for the calling user |
| :large_blue_circle: | `im.mark` | Sets the read cursor in a direct message channel |
| :large_blue_circle: | `im.myHistory` | Displays messages of the current user from direct message channel |
| :large_blue_circle: | `im.open` | Opens a direct message channel |
| :red_circle: | `im.replies` | Retrieve a thread of messages posted to a direct message conversation |
| :large_blue_circle: | `im.purgeHistory` | Deletes history of messages and events from direct message channel |
| :large_blue_circle: | `migration.exchange` | For Enterprise Grid workspaces, map local user IDs to global user IDs |
| :large_blue_circle: | `mpim.close` | Closes a multiparty direct message channel |
| :large_blue_circle: | `mpim.history` | Fetches history of messages and events from a multiparty direct message |
| :large_blue_circle: | `mpim.list` | Lists multiparty direct message channels for the calling user |
| :large_blue_circle: | `mpim.listSimple` | Lists ID and members in a multiparty direct message channels |
| :large_blue_circle: | `mpim.mark` | Sets the read cursor in a multiparty direct message channel |
| :large_blue_circle: | `mpim.myHistory` | Displays messages of the current user from multiparty direct message channel |
| :large_blue_circle: | `mpim.open` | This method opens a multiparty direct message |
| :large_blue_circle: | `mpim.purgeHistory` | Deletes history of messages and events from multiparty direct message channel |
| :red_circle: | `mpim.replies` | Retrieve a thread of messages posted to a direct message conversation from a multiparty direct message |
| :red_circle: | `oauth.access` | Exchanges a temporary OAuth code for an API token |
| :red_circle: | `oauth.token` | Exchanges a temporary OAuth verifier code for a workspace token |
| :large_blue_circle: | `pins.add` | Pins an item to a channel |
| :large_blue_circle: | `pins.list` | Lists items pinned to a channel |
| :large_blue_circle: | `pins.remove` | Un-pins an item from a channel |
| :large_blue_circle: | `reactions.add` | Adds a reaction to an item |
| :large_blue_circle: | `reactions.get` | Gets reactions for an item |
| :large_blue_circle: | `reactions.list` | Lists reactions made by a user |
| :large_blue_circle: | `reactions.remove` | Removes a reaction from an item |
| :red_circle: | `reminders.add` | Creates a reminder |
| :red_circle: | `reminders.complete` | Marks a reminder as complete |
| :red_circle: | `reminders.delete` | Deletes a reminder |
| :red_circle: | `reminders.info` | Gets information about a reminder |
| :red_circle: | `reminders.list` | Lists all reminders created by or for a given user |
| :red_circle: | `rtm.connect` | Starts a Real Time Messaging session |
| :large_blue_circle: | `rtm.start` | Starts a Real Time Messaging session |
| :large_blue_circle: | `rtm.events` | Prints the API events in real time |
| :large_blue_circle: | `search.all` | Searches for messages and files matching a query |
| :large_blue_circle: | `search.files` | Searches for files matching a query |
| :large_blue_circle: | `search.messages` | Searches for messages matching a query |
| :large_blue_circle: | `signup.checkEmail` | Checks if an email address is valid |
| :large_blue_circle: | `signup.confirmEmail` | Confirm an email address for signup |
| :large_blue_circle: | `stars.add` | Adds a star to an item |
| :large_blue_circle: | `stars.list` | Lists stars for a user |
| :large_blue_circle: | `stars.remove` | Removes a star from an item |
| :large_blue_circle: | `team.accessLogs` | Gets the access logs for the current team |
| :large_blue_circle: | `team.billableInfo` | Gets billable users information for the current team |
| :large_blue_circle: | `team.info` | Gets information about the current team |
| :red_circle: | `team.integrationLogs` | Gets the integration logs for the current team |
| :large_blue_circle: | `team.profile.get` | Retrieve a team's profile |
| :red_circle: | `usergroups.create` | Create a User Group |
| :red_circle: | `usergroups.disable` | Disable an existing User Group |
| :red_circle: | `usergroups.enable` | Enable a User Group |
| :red_circle: | `usergroups.list` | List all User Groups for a team |
| :red_circle: | `usergroups.update` | Update an existing User Group |
| :red_circle: | `usergroups.users.list` | List all users in a User Group |
| :red_circle: | `usergroups.users.update` | Update the list of users for a User Group |
| :red_circle: | `users.conversations` | List conversations the calling user may access |
| :large_blue_circle: | `users.counts` | Count number of users in the team |
| :large_blue_circle: | `users.deletePhoto` | Delete the user profile photo |
| :large_blue_circle: | `users.getPresence` | Gets user presence information |
| :large_blue_circle: | `users.id` | Gets user identifier from username |
| :large_blue_circle: | `users.identity` | Get a user's identity |
| :large_blue_circle: | `users.info` | Gets information about a user |
| :large_blue_circle: | `users.list` | Lists all users in a Slack team |
| :large_blue_circle: | `users.lookupByEmail` | Find a user with an email address |
| :large_blue_circle: | `users.prefs.get` | Get user account preferences |
| :large_blue_circle: | `users.prefs.set` | Set user account preferences |
| :large_blue_circle: | `users.preparePhoto` | Upload a picture to use as the avatar |
| :large_blue_circle: | `users.profile.get` | Retrieves a user's profile information |
| :large_blue_circle: | `users.profile.set` | Set the profile information for a user |
| :large_blue_circle: | `users.search` | Search users by name or email address |
| :black_circle: | `users.setActive` | Marked a user as active. **Deprecated and non-functional.** |
| :large_blue_circle: | `users.setAvatar` | Upload a picture and set it as the avatar |
| :large_blue_circle: | `users.setEmail` | Changes the email address without confirmation |
| :large_blue_circle: | `users.setPhoto` | Set the user profile photo |
| :large_blue_circle: | `users.setPresence` | Manually sets user presence |
| :large_blue_circle: | `users.setStatus` | Set the status message and emoji |
| :large_blue_circle: | `users.setUsername` | Changes the username without admin privileges |
| :large_blue_circle: | `version` | Displays the program version number |
