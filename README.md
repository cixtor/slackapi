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
- :skull: means the method has been deprecated without a replacement

| :shipit: | Method | Description |
|----------|--------|-------------|
| :red_circle: | `admin.apps.approve` | Approve an app for installation on a workspace |
| :red_circle: | `admin.apps.restrict` | Restrict an app for installation on a workspace |
| :red_circle: | `admin.apps.approved.list` | List approved apps for an org or workspace |
| :red_circle: | `admin.apps.requests.list` | List app requests for a team/workspace |
| :red_circle: | `admin.apps.restricted.list` | List restricted apps for an org or workspace |
| :red_circle: | `admin.conversations.restrictAccess.addGroup` | Add an allowlist of IDP groups for accessing a channel |
| :red_circle: | `admin.conversations.restrictAccess.listGroups` | List all IDP Groups linked to a channel |
| :red_circle: | `admin.conversations.restrictAccess.removeGroup` | Remove a linked IDP group linked from a private channel |
| :red_circle: | `admin.conversations.setTeams` | Set the workspaces in an Enterprise grid org that connect to a channel |
| :red_circle: | `admin.emoji.add` | Add an emoji |
| :red_circle: | `admin.emoji.addAlias` | Add an emoji alias |
| :red_circle: | `admin.emoji.list` | List emoji for an Enterprise Grid organization |
| :red_circle: | `admin.emoji.remove` | Remove an emoji across an Enterprise Grid organization |
| :red_circle: | `admin.emoji.rename` | Rename an emoji |
| :red_circle: | `admin.inviteRequests.approve` | Approve a workspace invite request |
| :red_circle: | `admin.inviteRequests.deny` | Deny a workspace invite request |
| :red_circle: | `admin.inviteRequests.list` | List all pending workspace invite requests |
| :red_circle: | `admin.inviteRequests.approved.list` | List all approved workspace invite requests |
| :red_circle: | `admin.inviteRequests.denied.list` | List all denied workspace invite requests |
| :red_circle: | `admin.teams.admins.list` | List all of the admins on a given workspace |
| :red_circle: | `admin.teams.create` | Create an Enterprise team |
| :red_circle: | `admin.teams.list` | List all teams on an Enterprise organization |
| :red_circle: | `admin.teams.owners.list` | List all of the owners on a given workspace |
| :red_circle: | `admin.teams.settings.info` | Fetch information about settings in a workspace |
| :red_circle: | `admin.teams.settings.setDefaultChannels` | Set the default channels of a workspace |
| :red_circle: | `admin.teams.settings.setDescription` | Set the description of a given workspace |
| :red_circle: | `admin.teams.settings.setDiscoverability` | An API method that allows admins to set the discoverability of a given workspace |
| :red_circle: | `admin.teams.settings.setIcon` | Sets the icon of a workspace |
| :red_circle: | `admin.teams.settings.setName` | Set the name of a given workspace |
| :red_circle: | `admin.usergroups.addChannels` | Add up to one hundred default channels to an IDP group |
| :red_circle: | `admin.usergroups.addTeams` | Associate one or more default workspaces with an organization-wide IDP group |
| :red_circle: | `admin.usergroups.listChannels` | List the channels linked to an org-level IDP group (user group) |
| :red_circle: | `admin.usergroups.removeChannels` | Remove one or more default channels from an org-level IDP group (user group) |
| :red_circle: | `admin.users.assign` | Add an Enterprise user to a workspace |
| :red_circle: | `admin.users.invite` | Invite a user to a workspace |
| :red_circle: | `admin.users.list` | List users on a workspace |
| :red_circle: | `admin.users.remove` | Remove a user from a workspace |
| :red_circle: | `admin.users.session.clearSettings` | Clear user-specific session settings—the session duration and what happens when the client closes—for a list of users |
| :red_circle: | `admin.users.session.getSettings` | Get user-specific session settings—the session duration and what happens when the client closes—given a list of users |
| :red_circle: | `admin.users.session.invalidate` | Revoke a single session for a user. The user will be forced to login to Slack |
| :red_circle: | `admin.users.session.list` | List active user sessions for an organization |
| :red_circle: | `admin.users.setAdmin` | Set an existing guest, regular user, or owner to be an admin user |
| :red_circle: | `admin.users.setExpiration` | Set an expiration for a guest user |
| :red_circle: | `admin.users.setOwner` | Set an existing guest, regular user, or admin user to be a workspace owner |
| :red_circle: | `admin.users.setRegular` | Set an existing guest user, admin user, or owner to be a regular user |
| :red_circle: | `admin.users.session.reset` | Wipes all valid sessions on all devices for a given user |
| :large_blue_circle: | `api.test` | Checks API calling code |
| :large_blue_circle: | `apps.list` | Lists associated applications |
| :large_blue_circle: | `apps.connections.open` | Generate a temporary Socket Mode WebSocket URL that your app can connect to in order to receive events and interactive payloads over |
| :large_blue_circle: | `apps.event.authorizations.list` | Get a list of authorizations for the given event context. Each authorization represents an app installation that the event is visible to |
| :large_blue_circle: | `apps.manifest.create` | Create an app from an app manifest |
| :large_blue_circle: | `apps.manifest.delete` | Permanently deletes an app created through app manifests |
| :large_blue_circle: | `apps.manifest.export` | Export an app manifest from an existing app |
| :large_blue_circle: | `apps.manifest.update` | Update an app from an app manifest |
| :large_blue_circle: | `apps.manifest.validate` | Validate an app manifest |
| :large_blue_circle: | `apps.uninstall` | Uninstalls your app from a workspace |
| :large_blue_circle: | `auth.revoke` | Revokes a token |
| :large_blue_circle: | `auth.teams.list` | List the workspaces a token can access |
| :large_blue_circle: | `auth.test` | Checks authentication and identity |
| :large_blue_circle: | `bots.info` | Gets information about a bot user |
| :red_circle: | `calls.add` | Registers a new Call |
| :red_circle: | `calls.end` | Ends a Call |
| :red_circle: | `calls.info` | Returns information about a Call |
| :red_circle: | `calls.update` | Updates information about a Call |
| :red_circle: | `calls.participants.add` | Registers new participants added to a Call |
| :red_circle: | `calls.participants.remove` | Registers participants removed from a Call |
| :large_blue_circle: | `chat.delete` | Deletes a message |
| :large_blue_circle: | `chat.deleteAttachment` | Deletes a message attachment |
| :red_circle: | `chat.deleteScheduledMessage` | Deletes a pending scheduled message from the queue |
| :red_circle: | `chat.getPermalink` | Retrieve a permalink URL for a specific extant message |
| :large_blue_circle: | `chat.meMessage` | Share a me message into a channel |
| :red_circle: | `chat.postEphemeral` | Sends an ephemeral message to a user in a channel |
| :large_blue_circle: | `chat.postMessage` | Sends a message to a channel |
| :red_circle: | `chat.scheduleMessage` | Schedules a message to be sent to a channel |
| :red_circle: | `chat.scheduledMessages.list` | Returns a list of scheduled messages |
| :red_circle: | `chat.unfurl` | Provide custom unfurl behavior for user-posted URLs |
| :large_blue_circle: | `chat.update` | Updates a message |
| :large_blue_circle: | `client.counts` | List mentions in different conversations |
| :large_blue_circle: | `client.shouldReload` | Determine if the Slack client must reload or not |
| :large_blue_circle: | `conversations.acceptSharedInvite` | Accepts an invitation to a Slack Connect channel |
| :large_blue_circle: | `conversations.approveSharedInvite` | Approves an invitation to a Slack Connect channel |
| :large_blue_circle: | `conversations.archive` | Archives a conversation |
| :large_blue_circle: | `conversations.close` | Closes a direct message or multi-person direct message |
| :large_blue_circle: | `conversations.create` | Initiates a public or private channel-based conversation |
| :large_blue_circle: | `conversations.declineSharedInvite` | Declines a Slack Connect channel invite |
| :large_blue_circle: | `conversations.delete` | Delete a public or private channel |
| :large_blue_circle: | `conversations.genericInfo` | Retrieve information about various channels |
| :large_blue_circle: | `conversations.history` | Fetches a conversation's history of messages and events |
| :large_blue_circle: | `conversations.info` | Retrieve information about a conversation |
| :large_blue_circle: | `conversations.invite` | Invites users to a channel |
| :large_blue_circle: | `conversations.inviteShared` | Sends an invitation to a Slack Connect channel |
| :large_blue_circle: | `conversations.join` | Joins an existing conversation |
| :large_blue_circle: | `conversations.kick` | Removes a user from a conversation |
| :large_blue_circle: | `conversations.leave` | Leaves a conversation |
| :large_blue_circle: | `conversations.list` | Lists all channels in a Slack team |
| :large_blue_circle: | `conversations.listConnectInvites` | Lists shared channel invites that have been generated or received but have not been approved by all parties |
| :large_blue_circle: | `conversations.mark` | Sets the read cursor in a channel |
| :large_blue_circle: | `conversations.members` | Retrieve members of a conversation |
| :large_blue_circle: | `conversations.open` | Opens or resumes a direct message or multi-person direct message |
| :large_blue_circle: | `conversations.rename` | Renames a conversation |
| :large_blue_circle: | `conversations.replies` | Retrieve a thread of messages posted to a conversation |
| :large_blue_circle: | `conversations.setPurpose` | Sets the purpose for a conversation |
| :large_blue_circle: | `conversations.setTopic` | Sets the topic for a conversation |
| :large_blue_circle: | `conversations.suggestions` | List Slack suggestions to join conversations |
| :large_blue_circle: | `conversations.unarchive` | Reverses conversation archival |
| :red_circle: | `dialog.open` | Open a dialog with a user |
| :large_blue_circle: | `dnd.endDnd` | Ends the current user's _"Do Not Disturb"_ session immediately |
| :large_blue_circle: | `dnd.endSnooze` | Ends the current user's snooze mode immediately |
| :large_blue_circle: | `dnd.info` | Retrieves a user's current _"Do Not Disturb"_ status |
| :large_blue_circle: | `dnd.setSnooze` | Turns on _"Do Not Disturb"_ mode for the current user, or changes its duration |
| :large_blue_circle: | `dnd.teamInfo` | Retrieves the _"Do Not Disturb"_ status for up to 50 users on a team |
| :large_blue_circle: | `emoji.add` | Uploads and registers a new custom emoji |
| :large_blue_circle: | `emoji.addAlias` | Creates an alias for an existing emoji |
| :large_blue_circle: | `emoji.getInfo` | Retrieves information about a custom emoji |
| :large_blue_circle: | `emoji.list` | Lists custom emoji for a team |
| :large_blue_circle: | `emoji.remove` | Remove an emoji from a team |
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
| :red_circle: | `files.remote.add` | Adds a file from a remote service |
| :red_circle: | `files.remote.info` | Retrieve information about a remote file added to Slack |
| :red_circle: | `files.remote.list` | Retrieve information about a remote file added to Slack |
| :red_circle: | `files.remote.remove` | Remove a remote file |
| :red_circle: | `files.remote.share` | Share a remote file into a channel |
| :red_circle: | `files.remote.update` | Updates an existing remote file |
| :large_blue_circle: | `help.issues.list` | List issues reported by the current user |
| :large_blue_circle: | `migration.exchange` | For Enterprise Grid workspaces, map local user IDs to global user IDs |
| :red_circle: | `oauth.access` | Exchanges a temporary OAuth code for an API token |
| :red_circle: | `oauth.token` | Exchanges a temporary OAuth verifier code for a workspace token |
| :red_circle: | `oauth.v2.access` | Exchanges a temporary OAuth verifier code for an access token |
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
| :large_blue_circle: | `rtm.start` | Deprecated: Starts a Real Time Messaging session. Use rtm.connect instead |
| :large_blue_circle: | `rtm.events` | Prints the API events in real time |
| :large_blue_circle: | `search.all` | Searches for messages and files matching a query |
| :large_blue_circle: | `search.files` | Searches for files matching a query |
| :large_blue_circle: | `search.messages` | Searches for messages matching a query |
| :large_blue_circle: | `search.modules` | Searches for modules matching a query |
| :large_blue_circle: | `search.users` | Search users by name or email address |
| :large_blue_circle: | `signup.checkEmail` | Checks if an email address is valid |
| :large_blue_circle: | `signup.confirmEmail` | Confirm an email address for signup |
| :large_blue_circle: | `slackbot.responses.add` | Adds a new automatic Slackbot response |
| :large_blue_circle: | `slackbot.responses.edit` | Edits an existing automatic Slackbot response |
| :large_blue_circle: | `slackbot.responses.delete` | Deletes an existing automatic Slackbot response |
| :large_blue_circle: | `stars.add` | Adds a star to an item |
| :large_blue_circle: | `stars.list` | Lists stars for a user |
| :large_blue_circle: | `stars.remove` | Removes a star from an item |
| :large_blue_circle: | `team.accessLogs` | Gets the access logs for the current team |
| :large_blue_circle: | `team.billableInfo` | Gets billable users information for the current team |
| :large_blue_circle: | `team.billing.info` | Reads a workspace's billing plan information |
| :large_blue_circle: | `team.channels.info` | Retrieve a list of channels in a specific team |
| :large_blue_circle: | `team.channels.membership` | Retrieve membership information about a team |
| :large_blue_circle: | `team.info` | Gets information about the current team |
| :large_blue_circle: | `team.integrationLogs` | Gets the integration logs for the current team |
| :large_blue_circle: | `team.listExternal` | List external teams and their corresponding information |
| :large_blue_circle: | `team.preferences.list` | Retrieve a list of a workspace's team preferences |
| :large_blue_circle: | `team.profile.get` | Retrieve a team's profile |
| :red_circle: | `usergroups.create` | Create a User Group |
| :red_circle: | `usergroups.disable` | Disable an existing User Group |
| :red_circle: | `usergroups.enable` | Enable a User Group |
| :red_circle: | `usergroups.list` | List all User Groups for a team |
| :red_circle: | `usergroups.update` | Update an existing User Group |
| :red_circle: | `usergroups.users.list` | List all users in a User Group |
| :red_circle: | `usergroups.users.update` | Update the list of users for a User Group |
| :large_blue_circle: | `users.admin.fetchInvitesHistory` | Gets invitations to join your workspace |
| :large_blue_circle: | `users.admin.inviteBulk` | Invite others to join your workspace |
| :large_blue_circle: | `users.admin.resendInvitation` | Resend an invitation to join your workspace |
| :large_blue_circle: | `users.admin.revokeInvitation` | Revoke an invitation to join your workspace |
| :large_blue_circle: | `users.admin.setInactive` | Deactivates an existing user account |
| :large_blue_circle: | `users.admin.setRegular` | Activates an account as a regular user |
| :red_circle: | `users.conversations` | List conversations the calling user may access |
| :large_blue_circle: | `users.deletePhoto` | Delete the user profile photo |
| :large_blue_circle: | `users.getPresence` | Gets user presence information |
| :large_blue_circle: | `users.identity` | Get a user's identity |
| :large_blue_circle: | `users.info` | Gets information about a user |
| :large_blue_circle: | `users.list` | Lists all users in a Slack team |
| :large_blue_circle: | `users.lookupByEmail` | Find a user with an email address |
| :large_blue_circle: | `users.prefs.get` | Get user account preferences |
| :large_blue_circle: | `users.prefs.set` | Set user account preferences |
| :large_blue_circle: | `users.preparePhoto` | Upload a picture to use as the avatar |
| :large_blue_circle: | `users.profile.get` | Retrieves a user's profile information |
| :large_blue_circle: | `users.profile.set` | Set the profile information for a user |
| :large_blue_circle: | `users.setAvatar` | Upload a picture and set it as the avatar |
| :large_blue_circle: | `users.setEmail` | Changes the email address without confirmation |
| :large_blue_circle: | `users.setPhoto` | Set the user profile photo |
| :large_blue_circle: | `users.setPresence` | Manually sets user presence |
| :large_blue_circle: | `users.setStatus` | Set the status message and emoji |
| :large_blue_circle: | `users.setUsername` | Changes the username without admin privileges |
| :large_blue_circle: | `version` | Displays the program version number |
| :red_circle: | `views.open` | Open a view for a user |
| :red_circle: | `views.publish` | Publish a static view for a User |
| :red_circle: | `views.push` | Push a view onto the stack of a root view |
| :red_circle: | `views.update` | Update an existing view |
| :large_blue_circle: | `workflows.stepCompleted` | Indicate that an app's step in a workflow completed execution |
| :large_blue_circle: | `workflows.stepFailed` | Indicate that an app's step in a workflow failed to execute |
| :large_blue_circle: | `workflows.updateStep` | Update the configuration for a workflow step |
