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
4. `$ SLACK_TOKEN=xoxs-token slackapi api.test # test token integrity`

### Progress

- [x] `api.test` - Checks API calling code.
- [x] `auth.test` - Checks authentication &amp; identity.
- [x] `channels.history` - Fetches history of messages and events from a channel.
- [x] `channels.info` - Gets information about a channel.
- [x] `channels.list` - Lists all channels in a Slack team.
- [x] `channels.mark` - Sets the read cursor in a channel.
- [x] `channels.setPurpose` - Sets the purpose for a channel.
- [x] `channels.setTopic` - Sets the topic for a channel.
- [x] `chat.delete` - Deletes a message.
- [x] `chat.postMessage` - Sends a message to a channel.
- [x] `chat.session` - Starts a new chat session.
- [ ] `chat.update` - Updates a message.
- [x] `emoji.list` - Lists custom emoji for a team.
- [ ] `files.delete` - Deletes a file.
- [ ] `files.upload` - Uploads or creates a file.
- [x] `groups.close` - Closes a private channel.
- [x] `groups.history` - Fetches history of messages and events from a private channel.
- [x] `groups.info` - Gets information about a private channel.
- [x] `groups.list` - Lists private channels that the calling user has access to.
- [ ] `groups.mark` - Sets the read cursor in a private channel.
- [ ] `groups.open` - Opens a private channel.
- [ ] `groups.setPurpose` - Sets the purpose for a private channel.
- [ ] `groups.setTopic` - Sets the topic for a private channel.
- [x] `im.close` - Close a direct message channel.
- [ ] `im.history` - Fetches history of messages and events from direct message channel.
- [x] `im.list` - Lists direct message channels for the calling user.
- [ ] `im.mark` - Sets the read cursor in a direct message channel.
- [x] `im.open` - Opens a direct message channel.
- [ ] `mpim.close` - Closes a multiparty direct message channel.
- [ ] `mpim.history` - Fetches history of messages and events from a multiparty direct message.
- [x] `mpim.list` - Lists multiparty direct message channels for the calling user.
- [ ] `mpim.mark` - Sets the read cursor in a multiparty direct message channel.
- [ ] `mpim.open` - This method opens a multiparty direct message.
- [ ] `pins.add` - Pins an item to a channel.
- [ ] `pins.remove` - Un-pins an item from a channel.
- [x] `reactions.add` - Adds a reaction to an item.
- [x] `reactions.get` - Gets reactions for an item.
- [x] `reactions.list` - Lists reactions made by a user.
- [x] `reactions.remove` - Removes a reaction from an item.
- [ ] `rtm.start` - Starts a Real Time Messaging session.
- [ ] `stars.add` - Adds a star to an item.
- [ ] `stars.remove` - Removes a star from an item.
- [x] `team.info` - Gets information about the current team.
- [x] `users.getPresence` - Gets user presence information.
- [x] `users.info` - Gets information about a user.
- [x] `users.list` - Lists all users in a Slack team.
- [x] `users.search` - Search users by name or email address.
- [x] `users.setActive` - Marks a user as active.
- [x] `users.setPresence` - Manually sets user presence.

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
