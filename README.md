### SlackAPI Client

Slack, the _"messaging app for teams"_ offers an API that has been used to build multiple projects around it, from bots to independent clients as well as integrations with other external services. This project aims to offer a low level experience for advanced users that want to either drop the web client or interact with the API for testing purpose.

### Features

The client is built on top of the [Bot Users](https://api.slack.com/bot-users) documentation. Most if not all the methods available in the API are implemented and can be executed placing a colon character as the suffix of each method.

Note that the client runs with the same chat session of the user that is using the program, but technically speaking the interaction is similar to that of a bot. This offers some advantages, for example, like other APIs and integrations, bot users are free. Unlike regular users, the actions they can perform are somewhat limited. For teams on the Free Plan, each bot user counts as a separate integration.

### Usage

Since this is a client you will need to give access to the perform HTTP requests against the API service, for that you will need to specify your [chat session token](https://api.slack.com/web#authentication). Alternatively, instead of generate a new token for the client you can use the same token issue for your user account when you log into the web interface, this key is not intended to be used as an external token because it expires once you logout, but if you keep the session alive you can work with it.

1. Open the messages board of your team [here](https://slack.com/messages/)
2. Press `Ctrl + Shift + J` and enter the code `boot_data.api_token`
3. Copy the token and use the terminal to interact with the service

```
$ SLACK_TOKEN=xoxs-token slackapi auth.test
$ slackapi chat.session
username:channel> :token xoxs-token
username:channel> :owner
username:channel> :exit
```

Also, you can pass an environment variable `VERBOSE=true` to print additional information during the execution of certain operations to troubleshoot issues with either the communication with the Slack API service or the program in itself.

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
- [x] `:robotimage` - Sets the avatar for the robot.
- [x] `:robotinfo` - Displays the configuration of the robot.
- [x] `:robotname` - Sets the user name of the robot.
- [x] `:robotoff` - Deactivates the robot to send normal messages.
- [x] `:roboton` - Activates the robot to send 3rd-party messages.
- [x] `:token` - Sets the token for the chat session.
- [x] `:update` - Updates the latest chat session message.
- [x] `:userid` - Displays the unique identifier of an user.
- [x] `:userlist` - Displays the information of all the users.
- [x] `:usersearch` - Searches the information of a specific user.

### Non-Official Client Methods

- [x] `channels.id` - Gets channel identifier from readable name.
- [x] `channels.myHistory` - Displays messages of the current user from a channel.
- [x] `channels.purgeHistory` - Deletes history of messages and events from a channel.
- [x] `chat.session` - Starts a new chat session.
- [x] `files.listAfterTime` - Lists and filters team files after this timestamp _(inclusive)_.
- [x] `files.listBeforeTime` - Lists and filters team files before this timestamp _(inclusive)_.
- [x] `files.listByChannel` - Lists and filters team files in a specific channel.
- [x] `files.listByType` - Lists and filters team files by type: all, posts, snippets, images, gdocs, zips, pdfs.
- [x] `files.listByUser` - Lists and filters team files created by a single user.
- [x] `groups.id` - Gets private channel identifier from readable name.
- [x] `groups.myHistory` - Displays messages of the current user from a private channel.
- [x] `groups.purgeHistory` - Deletes history of messages and events from a private channel.
- [x] `im.myHistory` - Displays messages of the current user from direct message channel.
- [x] `im.purgeHistory` - Deletes history of messages and events from direct message channel.
- [x] `users.id` - Gets user identifier from username.
- [x] `users.search` - Search users by name or email address.
- [x] `version` - Displays the program version number.
- [x] `help` - Displays usage and program options.

### Official Client Methods

- [x] `api.test` - Checks API calling code.
- [x] `apps.list` - Lists associated applications.
- [ ] `auth.revoke` - Revokes a token.
- [x] `auth.test` - Checks authentication and identity.
- [ ] `bots.info` - Gets information about a bot user.
- [x] `channels.archive` - Archives a channel.
- [x] `channels.create` - Creates a channel.
- [x] `channels.history` - Fetches history of messages and events from a channel.
- [x] `channels.info` - Gets information about a channel.
- [x] `channels.invite` - Invites a user to a channel.
- [x] `channels.join` - Joins a channel, creating it if needed.
- [x] `channels.kick` - Removes a user from a channel.
- [x] `channels.leave` - Leaves a channel.
- [x] `channels.list` - Lists all channels in a Slack team.
- [x] `channels.mark` - Sets the read cursor in a channel.
- [x] `channels.rename` - Renames a channel.
- [x] `channels.setPurpose` - Sets the purpose for a channel.
- [x] `channels.setRetention` - Sets the retention time of the messages.
- [x] `channels.setTopic` - Sets the topic for a channel.
- [x] `channels.unarchive` - Unarchives a channel.
- [x] `chat.delete` - Deletes a message.
- [ ] `chat.meMessage` - Share a me message into a channel.
- [x] `chat.postMessage` - Sends a message to a channel.
- [x] `chat.update` - Updates a message.
- [ ] `dnd.endDnd` - Ends the current user's _"Do Not Disturb"_ session immediately.
- [ ] `dnd.endSnooze` - Ends the current user's snooze mode immediately.
- [ ] `dnd.info` - Retrieves a user's current _"Do Not Disturb"_ status.
- [ ] `dnd.setSnooze` - Turns on _"Do Not Disturb"_ mode for the current user, or changes its duration.
- [ ] `dnd.teamInfo` - Retrieves the _"Do Not Disturb"_ status for users on a team.
- [x] `emoji.list` - Lists custom emoji for a team.
- [x] `files.comments.add` - Add a comment to an existing file.
- [x] `files.comments.delete` - Deletes an existing comment on a file.
- [x] `files.comments.edit` - Edit an existing file comment.
- [x] `files.delete` - Deletes a file.
- [x] `files.info` - Gets information about a team file.
- [x] `files.list` - Lists and filters team files.
- [ ] `files.revokePublicURL` - Revokes public/external sharing access for a file
- [ ] `files.sharedPublicURL` - Enables a file for public/external sharing.
- [x] `files.upload` - Uploads or creates a file.
- [x] `groups.archive` - Archives a private channel.
- [x] `groups.close` - Closes a private channel.
- [ ] `groups.create` - Creates a private channel.
- [ ] `groups.createChild` - Clones and archives a private channel.
- [x] `groups.history` - Fetches history of messages and events from a private channel.
- [x] `groups.info` - Gets information about a private channel.
- [ ] `groups.invite` - Invites a user to a private channel.
- [x] `groups.kick` - Removes a user from a private channel.
- [x] `groups.leave` - Leaves a private channel.
- [x] `groups.list` - Lists private channels that the calling user has access to.
- [x] `groups.mark` - Sets the read cursor in a private channel.
- [x] `groups.open` - Opens a private channel.
- [x] `groups.rename` - Renames a private channel.
- [x] `groups.setPurpose` - Sets the purpose for a private channel.
- [x] `groups.setRetention` - Sets the retention time of the messages.
- [x] `groups.setTopic` - Sets the topic for a private channel.
- [x] `groups.unarchive` - Unarchives a private channel.
- [x] `im.close` - Close a direct message channel.
- [x] `im.history` - Fetches history of messages and events from direct message channel.
- [x] `im.list` - Lists direct message channels for the calling user.
- [x] `im.mark` - Sets the read cursor in a direct message channel.
- [x] `im.open` - Opens a direct message channel.
- [ ] `mpim.close` - Closes a multiparty direct message channel.
- [ ] `mpim.history` - Fetches history of messages and events from a multiparty direct message.
- [x] `mpim.list` - Lists multiparty direct message channels for the calling user.
- [ ] `mpim.mark` - Sets the read cursor in a multiparty direct message channel.
- [ ] `mpim.open` - This method opens a multiparty direct message.
- [ ] `oauth.access` - Exchanges a temporary OAuth code for an API token.
- [ ] `pins.add` - Pins an item to a channel.
- [ ] `pins.list` - Lists items pinned to a channel.
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
- [ ] `rtm.start` - Starts a Real Time Messaging session.
- [ ] `search.all` - Searches for messages and files matching a query.
- [ ] `search.files` - Searches for files matching a query.
- [ ] `search.messages` - Searches for messages matching a query.
- [ ] `stars.add` - Adds a star to an item.
- [ ] `stars.list` - Lists stars for a user.
- [ ] `stars.remove` - Removes a star from an item.
- [x] `team.accessLogs` - Gets the access logs for the current team.
- [ ] `team.billableInfo` - Gets billable users information for the current team.
- [x] `team.info` - Gets information about the current team.
- [ ] `team.integrationLogs` - Gets the integration logs for the current team.
- [ ] `team.profile.get` - Retrieve a team's profile.
- [ ] `usergroups.create` - Create a User Group
- [ ] `usergroups.disable` - Disable an existing User Group
- [ ] `usergroups.enable` - Enable a User Group
- [ ] `usergroups.list` - List all User Groups for a team
- [ ] `usergroups.update` - Update an existing User Group
- [ ] `usergroups.users.list` - List all users in a User Group
- [ ] `usergroups.users.update` - Update the list of users for a User Group
- [x] `users.getPresence` - Gets user presence information.
- [ ] `users.identity` - Get a user's identity.
- [x] `users.info` - Gets information about a user.
- [x] `users.list` - Lists all users in a Slack team.
- [x] `users.setActive` - Marks a user as active.
- [x] `users.setPresence` - Manually sets user presence.
- [ ] `users.profile.get` - Retrieves a user's profile information.
- [ ] `users.profile.set` - Set the profile information for a user.

### License

```
The MIT License (MIT)

Copyright (c) 2015 CIXTOR

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
