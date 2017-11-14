### SlackAPI (Go Package / CLI Client)

Slack, the _"messaging app for teams"_ offers an API that has been used to build multiple projects around it, from bots to independent clients as well as integrations with other external services. This project aims to offer a low level experience for advanced users that want to either drop the web client or interact with the API for testing purpose.

### Installation

Install the CLI with this command:

```
go get -u github.com/cixtor/slackapi/slackapi
```

Import the package into your project with this:

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

### Usage

Use a [session token](https://api.slack.com/web#authentication) to authenticate the HTTP requests against the API service. Slack automatically generates a token for your when you open a new session [here](https://slack.com/messages/); you can see this token in the JavaScript console of your web browser if you type `boot_data.api_token` but be aware that it will expire once you close the session, consider to use a [legacy token](https://api.slack.com/custom-integrations/legacy-tokens) instead.

```
$ export SLACK_TOKEN=xoxs-token
$ slackapi help
$ slackapi auth.test
$ slackapi chat.session
username:channel> :token xoxs-token
username:channel> :owner
username:channel> :exit
```

You can also export an environment variable `SLACK_VERBOSE=true` to print additional information during the execution of certain operations to troubleshoot issues with either the communication with th API or the program in itself.

### Features

The client is built on top of the [Bot Users](https://api.slack.com/bot-users) documentation. Most if not all the methods available in the API are implemented and can be executed placing a colon character as the suffix of each method.

Note that the client runs with the same chat session of the user that is using the program, but technically speaking the interaction is similar to that of a bot. This offers some advantages, for example, like other APIs and integrations, bot users are free. Unlike regular users, the actions they can perform are somewhat limited. For teams on the Free Plan, each bot user counts as a separate integration.

### Chat Session Commands

- [x] `:close` - Close current chat session.
- [x] `:delete` - Deletes the latest message in the session history.
- [x] `:exec` - Executes and sends the output of a local command.
- [x] `:execv` - Same as :exec but includes the executed command.
- [x] `:exit` - Exits the program without closing chat sessions.
- [x] `:flush` - Deletes all the messages in the session history.
- [x] `:history` - Fetches messages and events in current session.
- [x] `:messages` - Displays the messages in the current session.
- [x] `:myhistory` - Fetches messages and events from current user.
- [x] `:open` - Opens a new session with a user, channel, or group.
- [x] `:owner` - Displays account information of the user in session.
- [x] `:purge` - Deletes the messages in the current session.
- [x] `:token` - Sets the token for the chat session.
- [x] `:update` - Updates the latest chat session message.
- [x] `:userid` - Displays the unique identifier of an user.
- [x] `:userlist` - Displays the information of all the users.
- [x] `:usersearch` - Searches the information of a specific user.

### Official Client Methods

- [x] `api.test` - Checks API calling code.
- [x] `apps.list` - Lists associated applications.
- [ ] `apps.permissions.info` - Returns list of permissions this app has on a team.
- [ ] `apps.permissions.request` - Allows an app to request additional scopes
- [x] `auth.revoke` - Revokes a token.
- [x] `auth.test` - Checks authentication and identity.
- [x] `bots.info` - Gets information about a bot user.
- [x] `channels.archive` - Archives a channel.
- [x] `channels.create` - Creates a channel.
- [x] `channels.history` - Fetches history of messages and events from a channel.
- [x] `channels.id` - Gets channel identifier from readable name.
- [x] `channels.info` - Gets information about a channel.
- [x] `channels.invite` - Invites a user to a channel.
- [x] `channels.join` - Joins a channel, creating it if needed.
- [x] `channels.kick` - Removes a user from a channel.
- [x] `channels.leave` - Leaves a channel.
- [x] `channels.list` - Lists all channels in a Slack team.
- [x] `channels.mark` - Sets the read cursor in a channel.
- [x] `channels.myHistory` - Displays messages of the current user from a channel.
- [x] `channels.purgeHistory` - Deletes history of messages and events from a channel.
- [x] `channels.rename` - Renames a channel.
- [ ] `channels.replies` - Retrieve a thread of messages posted to a channel
- [x] `channels.setPurpose` - Sets the purpose for a channel.
- [x] `channels.setRetention` - Sets the retention time of the messages.
- [x] `channels.setTopic` - Sets the topic for a channel.
- [x] `channels.suggestions` - Prints a list of suggested channels to join.
- [x] `channels.unarchive` - Unarchives a channel.
- [x] `chat.delete` - Deletes a message.
- [ ] `chat.getPermalink` - Retrieve a permalink URL for a specific extant message
- [x] `chat.meMessage` - Share a me message into a channel.
- [ ] `chat.postEphemeral` - Sends an ephemeral message to a user in a channel.
- [x] `chat.postMessage` - Sends a message to a channel.
- [x] `chat.robotMessage` - Sends a message to a channel as a robot.
- [x] `chat.session` - Starts a new chat session.
- [ ] `chat.unfurl` - Provide custom unfurl behavior for user-posted URLs
- [x] `chat.update` - Updates a message.
- [ ] `conversations.archive` - Archives a conversation.
- [ ] `conversations.close` - Closes a direct message or multi-person direct message.
- [ ] `conversations.create` - Initiates a public or private channel-based conversation
- [ ] `conversations.history` - Fetches a conversation's history of messages and events.
- [ ] `conversations.info` - Retrieve information about a conversation.
- [ ] `conversations.invite` - Invites users to a channel.
- [ ] `conversations.join` - Joins an existing conversation.
- [ ] `conversations.kick` - Removes a user from a conversation.
- [ ] `conversations.leave` - Leaves a conversation.
- [ ] `conversations.list` - Lists all channels in a Slack team.
- [ ] `conversations.members` - Retrieve members of a conversation.
- [ ] `conversations.open` - Opens or resumes a direct message or multi-person direct message.
- [ ] `conversations.rename` - Renames a conversation.
- [ ] `conversations.replies` - Retrieve a thread of messages posted to a conversation
- [ ] `conversations.setPurpose` - Sets the purpose for a conversation.
- [ ] `conversations.setTopic` - Sets the topic for a conversation.
- [ ] `conversations.unarchive` - Reverses conversation archival.
- [ ] `dialog.open` - Open a dialog with a user
- [x] `dnd.endDnd` - Ends the current user's _"Do Not Disturb"_ session immediately.
- [x] `dnd.endSnooze` - Ends the current user's snooze mode immediately.
- [x] `dnd.info` - Retrieves a user's current _"Do Not Disturb"_ status.
- [x] `dnd.setSnooze` - Turns on _"Do Not Disturb"_ mode for the current user, or changes its duration.
- [x] `dnd.teamInfo` - Retrieves the _"Do Not Disturb"_ status for users on a team.
- [x] `emoji.list` - Lists custom emoji for a team.
- [x] `eventlog.history` - Lists all the events since the specified time.
- [x] `files.comments.add` - Add a comment to an existing file.
- [x] `files.comments.delete` - Deletes an existing comment on a file.
- [x] `files.comments.edit` - Edit an existing file comment.
- [x] `files.delete` - Deletes a file.
- [x] `files.info` - Gets information about a team file.
- [x] `files.list` - Lists and filters team files.
- [x] `files.listAfterTime` - Lists and filters team files after this timestamp _(inclusive)_.
- [x] `files.listBeforeTime` - Lists and filters team files before this timestamp _(inclusive)_.
- [x] `files.listByChannel` - Lists and filters team files in a specific channel.
- [x] `files.listByType` - Lists and filters team files by type: all, posts, snippets, images, gdocs, zips, pdfs.
- [x] `files.listByUser` - Lists and filters team files created by a single user.
- [ ] `files.revokePublicURL` - Revokes public/external sharing access for a file.
- [ ] `files.sharedPublicURL` - Enables a file for public/external sharing.
- [x] `files.upload` - Uploads or creates a file.
- [x] `groups.archive` - Archives a private channel.
- [x] `groups.close` - Closes a private channel.
- [x] `groups.create` - Creates a private channel.
- [x] `groups.createChild` - Clones and archives a private channel.
- [x] `groups.history` - Fetches history of messages and events from a private channel.
- [x] `groups.id` - Gets private channel identifier from readable name.
- [x] `groups.info` - Gets information about a private channel.
- [x] `groups.invite` - Invites a user to a private channel.
- [x] `groups.kick` - Removes a user from a private channel.
- [x] `groups.leave` - Leaves a private channel.
- [x] `groups.list` - Lists private channels that the calling user has access to.
- [x] `groups.mark` - Sets the read cursor in a private channel.
- [x] `groups.myHistory` - Displays messages of the current user from a private channel.
- [x] `groups.open` - Opens a private channel.
- [x] `groups.purgeHistory` - Deletes history of messages and events from a private channel.
- [x] `groups.rename` - Renames a private channel.
- [ ] `groups.replies` - Retrieve a thread of messages posted to a private channel
- [x] `groups.setPurpose` - Sets the purpose for a private channel.
- [x] `groups.setRetention` - Sets the retention time of the messages.
- [x] `groups.setTopic` - Sets the topic for a private channel.
- [x] `groups.unarchive` - Unarchives a private channel.
- [x] `help` - Displays usage and program options.
- [x] `help.issues.list` - List issues reported by the current user.
- [x] `im.close` - Close a direct message channel.
- [x] `im.history` - Fetches history of messages and events from direct message channel.
- [x] `im.list` - Lists direct message channels for the calling user.
- [x] `im.mark` - Sets the read cursor in a direct message channel.
- [x] `im.myHistory` - Displays messages of the current user from direct message channel.
- [x] `im.open` - Opens a direct message channel.
- [ ] `im.replies` - Retrieve a thread of messages posted to a direct message conversation
- [x] `im.purgeHistory` - Deletes history of messages and events from direct message channel.
- [ ] `migration.exchange` - For Enterprise Grid workspaces, map local user IDs to global user IDs
- [x] `mpim.close` - Closes a multiparty direct message channel.
- [x] `mpim.history` - Fetches history of messages and events from a multiparty direct message.
- [x] `mpim.list` - Lists multiparty direct message channels for the calling user.
- [x] `mpim.listSimple` - Lists ID and members in a multiparty direct message channels.
- [x] `mpim.mark` - Sets the read cursor in a multiparty direct message channel.
- [x] `mpim.myHistory` - Displays messages of the current user from multiparty direct message channel.
- [x] `mpim.open` - This method opens a multiparty direct message.
- [x] `mpim.purgeHistory` - Deletes history of messages and events from multiparty direct message channel.
- [ ] `mpim.replies` - Retrieve a thread of messages posted to a direct message conversation from a multiparty direct message.
- [ ] `oauth.access` - Exchanges a temporary OAuth code for an API token.
- [ ] `oauth.token` - Exchanges a temporary OAuth verifier code for a workspace token.
- [ ] `pins.add` - Pins an item to a channel.
- [x] `pins.list` - Lists items pinned to a channel.
- [ ] `pins.remove` - Un-pins an item from a channel.
- [x] `reactions.add` - Adds a reaction to an item.
- [x] `reactions.get` - Gets reactions for an item.
- [x] `reactions.list` - Lists reactions made by a user.
- [x] `reactions.remove` - Removes a reaction from an item.
- [ ] `reminders.add` - Creates a reminder.
- [ ] `reminders.complete` - Marks a reminder as complete.
- [ ] `reminders.delete` - Deletes a reminder.
- [ ] `reminders.info` - Gets information about a reminder.
- [ ] `reminders.list` - Lists all reminders created by or for a given user.
- [ ] `rtm.connect` - Starts a Real Time Messaging session.
- [x] `rtm.start` - Starts a Real Time Messaging session.
- [x] `rtm.events` - Prints the API events in real time
- [ ] `search.all` - Searches for messages and files matching a query.
- [ ] `search.files` - Searches for files matching a query.
- [ ] `search.messages` - Searches for messages matching a query.
- [ ] `stars.add` - Adds a star to an item.
- [ ] `stars.list` - Lists stars for a user.
- [ ] `stars.remove` - Removes a star from an item.
- [x] `team.accessLogs` - Gets the access logs for the current team.
- [x] `team.billableInfo` - Gets billable users information for the current team.
- [x] `team.info` - Gets information about the current team.
- [ ] `team.integrationLogs` - Gets the integration logs for the current team.
- [x] `team.profile.get` - Retrieve a team's profile.
- [ ] `usergroups.create` - Create a User Group.
- [ ] `usergroups.disable` - Disable an existing User Group.
- [ ] `usergroups.enable` - Enable a User Group.
- [ ] `usergroups.list` - List all User Groups for a team.
- [ ] `usergroups.update` - Update an existing User Group.
- [ ] `usergroups.users.list` - List all users in a User Group.
- [ ] `usergroups.users.update` - Update the list of users for a User Group.
- [ ] `users.conversations` - List conversations the calling user may access.
- [x] `users.counts` - Count number of users in the team.
- [x] `users.deletePhoto` - Delete the user profile photo
- [x] `users.getPresence` - Gets user presence information.
- [x] `users.id` - Gets user identifier from username.
- [x] `users.identity` - Get a user's identity.
- [x] `users.info` - Gets information about a user.
- [x] `users.list` - Lists all users in a Slack team.
- [ ] `users.lookupByEmail` - Find a user with an email address.
- [x] `users.prefs.get` - Get user account preferences.
- [x] `users.prefs.set` - Set user account preferences.
- [x] `users.preparePhoto` - Upload a picture to use as the avatar.
- [x] `users.profile.get` - Retrieves a user's profile information.
- [x] `users.profile.set` - Set the profile information for a user.
- [x] `users.search` - Search users by name or email address.
- [x] `users.setActive` - Marked a user as active. **Deprecated and non-functional.**
- [x] `users.setAvatar` - Upload a picture and set it as the avatar.
- [x] `users.setEmail` - Changes the email address without confirmation.
- [x] `users.setPhoto` - Set the user profile photo.
- [x] `users.setPresence` - Manually sets user presence.
- [x] `users.setStatus` - Set the status message and emoji.
- [x] `users.setUsername` - Changes the username without admin privileges.
- [x] `version` - Displays the program version number.
