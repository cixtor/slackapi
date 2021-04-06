package slackapi

import (
	"encoding/json"
	"testing"
)

func CheckResponse(t *testing.T, x interface{}, y string) {
	out, err := json.Marshal(x)
	if err != nil {
		t.Fatal("json fromat;", err)
	}
	if string(out) != y {
		t.Fatalf("invalid json response;\n- %s\n+ %s\n", y, out)
	}
}

func TestAPITest(t *testing.T) {
	s := New()
	x := s.APITest()
	y := `{"ok":true}`
	CheckResponse(t, x, y)
}

func TestAppsList(t *testing.T) {
	s := New()
	x := s.AppsList()
	y := `{"ok":false,"error":"not_authed","apps":null,"cache_ts":""}`
	CheckResponse(t, x, y)
}

func TestAuthRevoke(t *testing.T) {
	s := New()
	x := s.AuthRevoke()
	y := `{"ok":false,"error":"not_authed","revoked":false}`
	CheckResponse(t, x, y)
}

func TestAuthTest(t *testing.T) {
	s := New()
	x, err := s.AuthTest()
	if err != nil {
		t.Fatal(err)
	}
	y := `{"ok":false,"error":"not_authed","team":"","team_id":"","url":"","user":"","user_id":""}`
	CheckResponse(t, x, y)
}

func TestBotsInfo(t *testing.T) {
	s := New()
	x := s.BotsInfo("user")
	y := `{"ok":false,"error":"not_authed","bot":{"id":"","deleted":false,"name":"","icons":null}}`
	CheckResponse(t, x, y)
}

func TestChannelsID(t *testing.T) {
	s := New()
	x := s.ChannelsID("channel")
	y := `"channel"`
	CheckResponse(t, x, y)
}

func TestChannelsMyHistory(t *testing.T) {
	s := New()
	x := s.ChannelsMyHistory("channel", "1234567890")
	y := `{"Filtered":0,"Latest":"","Messages":null,"Oldest":"","Total":0,"Username":""}`
	CheckResponse(t, x, y)
}

func TestChannelsPurgeHistory(t *testing.T) {
	s := New()
	x := s.ChannelsPurgeHistory("channel", "1234567890", true)
	y := `{"Deleted":0,"NotDeleted":0,"Messages":null}`
	CheckResponse(t, x, y)
}

func TestChannelsSetRetention(t *testing.T) {
	s := New()
	x := s.ChannelsSetRetention("channel", 1)
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestChannelsSuggestions(t *testing.T) {
	s := New()
	x := s.ChannelsSuggestions()
	y := `{"ok":false,"error":"not_authed","status":{"ok":false},"suggestion_types_tried":null}`
	CheckResponse(t, x, y)
}

func TestChatDelete(t *testing.T) {
	s := New()
	x := s.ChatDelete(MessageArgs{})
	y := `{"ok":false,"error":"not_authed","channel":"","text":"","ts":""}`
	CheckResponse(t, x, y)
}

func TestChatMeMessage(t *testing.T) {
	s := New()
	x := s.ChatMeMessage(MessageArgs{})
	y := `{"ok":false,"error":"not_authed","channel":"","text":"","ts":""}`
	CheckResponse(t, x, y)
}

func TestChatPostMessage(t *testing.T) {
	s := New()
	x := s.ChatPostMessage(MessageArgs{})
	y := `{"ok":false,"error":"not_authed","channel":"","ts":"","message":{"display_as_bot":false}}`
	CheckResponse(t, x, y)
}

func TestChatUpdate(t *testing.T) {
	s := New()
	x := s.ChatUpdate(MessageArgs{})
	y := `{"ok":false,"error":"not_authed","channel":"","ts":"","message":{"display_as_bot":false}}`
	CheckResponse(t, x, y)
}

func TestConversationsArchive(t *testing.T) {
	s := New()
	x := s.ConversationsArchive("channel")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestConversationsCreate(t *testing.T) {
	s := New()
	x := s.ConversationsCreate("channel")
	y := `{"ok":false,"error":"not_authed","channel":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_general":false,"is_group":false,"is_member":false,"is_mpim":false,"is_open":false,"last_read":"","latest":{"text":"","ts":"","type":"","user":""},"members":null,"name":"","name_normalized":"","num_members":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unread_count":0,"unread_count_display":0}}`
	CheckResponse(t, x, y)
}

func TestConversationsHistory(t *testing.T) {
	s := New()
	x := s.ConversationsHistory(ConversationsHistoryInput{Channel: "channel", Latest: "1234567890"})
	y := `{"ok":false,"error":"not_authed","messages":null,"has_more":false,"pin_count":0,"unread_count_display":0,"response_metadata":{"next_cursor":""}}`
	CheckResponse(t, x, y)
}

func TestConversationsInfo(t *testing.T) {
	s := New()
	x := s.ConversationsInfo("channel")
	y := `{"ok":false,"error":"not_authed","channel":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_general":false,"is_group":false,"is_member":false,"is_mpim":false,"is_open":false,"last_read":"","latest":{"text":"","ts":"","type":"","user":""},"members":null,"name":"","name_normalized":"","num_members":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unread_count":0,"unread_count_display":0}}`
	CheckResponse(t, x, y)
}

func TestConversationsJoin(t *testing.T) {
	s := New()
	x := s.ConversationsJoin("channel")
	y := `{"ok":false,"error":"not_authed","channel":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_general":false,"is_group":false,"is_member":false,"is_mpim":false,"is_open":false,"last_read":"","latest":{"text":"","ts":"","type":"","user":""},"members":null,"name":"","name_normalized":"","num_members":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unread_count":0,"unread_count_display":0}}`
	CheckResponse(t, x, y)
}

func TestConversationsInvite(t *testing.T) {
	s := New()
	x := s.ConversationsInvite("channel", "user1", "user2", "user3")
	y := `{"ok":false,"error":"not_authed","channel":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_general":false,"is_group":false,"is_member":false,"is_mpim":false,"is_open":false,"last_read":"","latest":{"text":"","ts":"","type":"","user":""},"members":null,"name":"","name_normalized":"","num_members":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unread_count":0,"unread_count_display":0}}`
	CheckResponse(t, x, y)
}

func TestConversationsKick(t *testing.T) {
	s := New()
	x := s.ConversationsKick("channel", "user")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestConversationsLeave(t *testing.T) {
	s := New()
	x := s.ConversationsLeave("channel")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestConversationsList(t *testing.T) {
	s := New()
	x := s.ConversationsList(ConversationsListInput{})
	y := `{"ok":false,"error":"not_authed","channels":null}`
	CheckResponse(t, x, y)
}

func TestConversationsRename(t *testing.T) {
	s := New()
	x := s.ConversationsRename("channel", "lennahc")
	y := `{"ok":false,"error":"not_authed","channel":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_general":false,"is_group":false,"is_member":false,"is_mpim":false,"is_open":false,"last_read":"","latest":{"text":"","ts":"","type":"","user":""},"members":null,"name":"","name_normalized":"","num_members":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unread_count":0,"unread_count_display":0}}`
	CheckResponse(t, x, y)
}

func TestConversationsReplies(t *testing.T) {
	s := New()
	x := s.ConversationsReplies(ConversationsRepliesInput{Channel: "general", Timestamp: "1234567890.123456"})
	y := `{"ok":false,"error":"not_authed","messages":null,"has_more":false,"pin_count":0,"unread_count_display":0,"response_metadata":{"next_cursor":""}}`
	CheckResponse(t, x, y)
}
