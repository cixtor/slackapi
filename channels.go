package slackapi

// ResponseChannelsInfo defines the JSON-encoded output for ChannelsInfo.
//
// Example:
//
//   {
//       "ok": true,
//       "channel": {
//           "id": "C012AB3CD",
//           "name": "general",
//           "is_channel": true,
//           "is_group": false,
//           "is_im": false,
//           "created": 1449252889,
//           "creator": "W012A3BCD",
//           "is_archived": false,
//           "is_general": true,
//           "unlinked": 0,
//           "name_normalized": "general",
//           "is_read_only": false,
//           "is_shared": false,
//           "is_ext_shared": false,
//           "is_org_shared": false,
//           "pending_shared": [],
//           "is_pending_ext_shared": false,
//           "is_member": true,
//           "is_private": false,
//           "is_mpim": false,
//           "last_read": "1502126650.228446",
//           "topic": {
//               "value": "For public discussion of generalities",
//               "creator": "W012A3BCD",
//               "last_set": 1449709364
//           },
//           "purpose": {
//               "value": "This part of the workspace is for fun. Make fun here.",
//               "creator": "W012A3BCD",
//               "last_set": 1449709364
//           },
//           "previous_names": [
//               "specifics",
//               "abstractions",
//               "etc"
//           ],
//           "num_members": 23,
//           "locale": "en-US"
//       }
//   }
type ResponseChannelsInfo struct {
	Response
	Channel Channel `json:"channel"`
}

type ResponseChannelsGenericInfo struct {
	Response
	Channels []Channel `json:"channels"`
}

// ResponseChannelsJoin defines the JSON-encoded output for ChannelsJoin.
type ResponseChannelsJoin struct {
	Response
	AlreadyInChannel bool    `json:"already_in_channel"`
	Channel          Channel `json:"channel"`
}

// ResponseChannelsList defines the JSON-encoded output for ChannelsList.
type ResponseChannelsList struct {
	Response
	Channels []Channel `json:"channels"`
}

// ChannelSuggestions defines the expected data from the JSON-encoded API response.
type ChannelSuggestions struct {
	Response
	Status               Response `json:"status"`
	SuggestionTypesTried []string `json:"suggestion_types_tried"`
}

// ChannelsMyHistory displays messages of the current user from a channel.
func (s *SlackAPI) ChannelsMyHistory(channel string, latest string) MyHistory {
	return s.ResourceMyHistory("channels.history", channel, latest)
}

// ChannelsPurgeHistory deletes history of messages and events from a channel.
func (s *SlackAPI) ChannelsPurgeHistory(channel string, latest string, verbose bool) DeletedHistory {
	return s.ResourcePurgeHistory("channels.history", channel, latest, verbose)
}

// ChannelsSetRetention sets the retention time of the messages.
func (s *SlackAPI) ChannelsSetRetention(channel string, duration int) Response {
	return s.ResourceSetRetention("channels.setRetention", channel, duration)
}

// ChannelsSuggestions prints a list of suggested channels to join.
func (s *SlackAPI) ChannelsSuggestions() ChannelSuggestions {
	var response ChannelSuggestions
	s.getRequest(&response, "channels.suggestions", nil)
	return response
}
