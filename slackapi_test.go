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

func TestConversationsSetPurpose(t *testing.T) {
	s := New()
	x := s.ConversationsSetPurpose("channel", "purpose")
	y := `{"ok":false,"error":"not_authed","purpose":""}`
	CheckResponse(t, x, y)
}

func TestConversationsSetTopic(t *testing.T) {
	s := New()
	x := s.ConversationsSetTopic("channel", "topic")
	y := `{"ok":false,"error":"not_authed","topic":""}`
	CheckResponse(t, x, y)
}

func TestConversationsUnarchive(t *testing.T) {
	s := New()
	x := s.ConversationsUnarchive("channel")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestDNDEndDnd(t *testing.T) {
	s := New()
	x := s.DNDEndDnd()
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestDNDEndSnooze(t *testing.T) {
	s := New()
	x := s.DNDEndSnooze()
	y := `{"ok":false,"error":"not_authed","dnd_enabled":false,"next_dnd_start_ts":0,"next_dnd_end_ts":0,"snooze_debug":{}}`
	CheckResponse(t, x, y)
}

func TestDNDInfo(t *testing.T) {
	s := New()
	x := s.DNDInfo("admin")
	y := `{"ok":false,"error":"not_authed","dnd_enabled":false,"next_dnd_start_ts":0,"next_dnd_end_ts":0,"snooze_debug":{}}`
	CheckResponse(t, x, y)
}

func TestDNDSetSnooze(t *testing.T) {
	s := New()
	x := s.DNDSetSnooze(60)
	y := `{"ok":false,"error":"not_authed","snooze_debug":{}}`
	CheckResponse(t, x, y)
}

func TestDNDTeamInfo(t *testing.T) {
	s := New()
	x := s.DNDTeamInfo("admin")
	y := `{"ok":false,"error":"not_authed","cached":false,"users":null}`
	CheckResponse(t, x, y)
}

func TestEmojiList(t *testing.T) {
	s := New()
	x := s.EmojiList()
	y := `{"ok":false,"error":"not_authed","cache_ts":"","emoji":null}`
	CheckResponse(t, x, y)
}

func TestFilesCommentsAdd(t *testing.T) {
	s := New()
	x := s.FilesCommentsAdd("fileid", "comment")
	y := `{"ok":false,"error":"unknown_method","comment":{"comment":"","id":"","user":"","created":0,"timestamp":0,"is_intro":false}}`
	CheckResponse(t, x, y)
}

func TestFilesCommentsDelete(t *testing.T) {
	s := New()
	x := s.FilesCommentsDelete("fileid", "commentid")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestFilesCommentsEdit(t *testing.T) {
	s := New()
	x := s.FilesCommentsEdit("fileid", "commentid", "comment")
	y := `{"ok":false,"error":"unknown_method","comment":{"comment":"","id":"","user":"","created":0,"timestamp":0,"is_intro":false}}`
	CheckResponse(t, x, y)
}

func TestFilesDelete(t *testing.T) {
	s := New()
	x := s.FilesDelete("fileid")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestFilesInfo(t *testing.T) {
	s := New()
	x := s.FilesInfo("fileid", 1, 1)
	y := `{"ok":false,"error":"not_authed","file":{"channels":null,"comments_count":0,"created":0,"deanimate_gif":"","edit_link":"","external_type":"","filetype":"","groups":null,"id":"","image_exif_rotation":0,"initial_comment":{"comment":"","id":"","user":"","created":0,"timestamp":0,"is_intro":false},"ims":null,"lines":0,"lines_more":0,"mimetype":"","mode":"","name":"","num_stars":0,"original_h":0,"original_w":0,"permalink":"","permalink_public":"","pretty_type":"","preview":"","preview_highlight":"","reactions":null,"score":"","size":0,"thumb_160":"","thumb_360":"","thumb_360_gif":"","thumb_360_h":0,"thumb_360_w":0,"thumb_480":"","thumb_480_gif":"","thumb_480_h":0,"thumb_480_w":0,"thumb_64":"","thumb_80":"","timestamp":0,"title":"","url":"","url_download":"","url_private":"","url_private_download":"","user":"","username":"","display_as_bot":false,"editable":false,"is_external":false,"is_public":false,"is_starred":false,"public_url_shared":false,"top_file":false},"comments":null,"paging":{"count":0,"page":0,"pages":0,"total":0}}`
	CheckResponse(t, x, y)
}

func TestFilesList(t *testing.T) {
	s := New()
	x := s.FilesList(FileListArgs{})
	y := `{"ok":false,"error":"not_authed","files":null,"paging":{"count":0,"page":0,"pages":0,"total":0}}`
	CheckResponse(t, x, y)
}

func TestFilesRevokePublicURL(t *testing.T) {
	s := New()
	x := s.FilesRevokePublicURL("fileid")
	y := `{"error":"not_authed","ok":false}`
	CheckResponse(t, x, y)
}

func TestFilesSharedPublicURL(t *testing.T) {
	s := New()
	x := s.FilesSharedPublicURL("fileid")
	y := `{"ok":false,"error":"not_authed","file":{"channels":null,"comments_count":0,"created":0,"deanimate_gif":"","edit_link":"","external_type":"","filetype":"","groups":null,"id":"","image_exif_rotation":0,"initial_comment":{"comment":"","id":"","user":"","created":0,"timestamp":0,"is_intro":false},"ims":null,"lines":0,"lines_more":0,"mimetype":"","mode":"","name":"","num_stars":0,"original_h":0,"original_w":0,"permalink":"","permalink_public":"","pretty_type":"","preview":"","preview_highlight":"","reactions":null,"score":"","size":0,"thumb_160":"","thumb_360":"","thumb_360_gif":"","thumb_360_h":0,"thumb_360_w":0,"thumb_480":"","thumb_480_gif":"","thumb_480_h":0,"thumb_480_w":0,"thumb_64":"","thumb_80":"","timestamp":0,"title":"","url":"","url_download":"","url_private":"","url_private_download":"","user":"","username":"","display_as_bot":false,"editable":false,"is_external":false,"is_public":false,"is_starred":false,"public_url_shared":false,"top_file":false}}`
	CheckResponse(t, x, y)
}

func TestFilesUpload(t *testing.T) {
	s := New()
	x := s.FilesUpload(FileUploadArgs{})
	y := `{"ok":false,"error":"not_authed","file":{"channels":null,"comments_count":0,"created":0,"deanimate_gif":"","edit_link":"","external_type":"","filetype":"","groups":null,"id":"","image_exif_rotation":0,"initial_comment":{"comment":"","id":"","user":"","created":0,"timestamp":0,"is_intro":false},"ims":null,"lines":0,"lines_more":0,"mimetype":"","mode":"","name":"","num_stars":0,"original_h":0,"original_w":0,"permalink":"","permalink_public":"","pretty_type":"","preview":"","preview_highlight":"","reactions":null,"score":"","size":0,"thumb_160":"","thumb_360":"","thumb_360_gif":"","thumb_360_h":0,"thumb_360_w":0,"thumb_480":"","thumb_480_gif":"","thumb_480_h":0,"thumb_480_w":0,"thumb_64":"","thumb_80":"","timestamp":0,"title":"","url":"","url_download":"","url_private":"","url_private_download":"","user":"","username":"","display_as_bot":false,"editable":false,"is_external":false,"is_public":false,"is_starred":false,"public_url_shared":false,"top_file":false}}`
	CheckResponse(t, x, y)
}

func TestGroupsClose(t *testing.T) {
	s := New()
	x := s.GroupsClose("channel")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestGroupsID(t *testing.T) {
	s := New()
	x := s.GroupsID("channel")
	y := `"channel"`
	CheckResponse(t, x, y)
}

func TestGroupsMyHistory(t *testing.T) {
	s := New()
	x := s.GroupsMyHistory("channel", "1234567890")
	y := `{"Filtered":0,"Latest":"","Messages":null,"Oldest":"","Total":0,"Username":""}`
	CheckResponse(t, x, y)
}

func TestGroupsPurgeHistory(t *testing.T) {
	s := New()
	x := s.GroupsPurgeHistory("channel", "1234567890", true)
	y := `{"Deleted":0,"NotDeleted":0,"Messages":null}`
	CheckResponse(t, x, y)
}

func TestGroupsSetRetention(t *testing.T) {
	s := New()
	x := s.GroupsSetRetention("channel", 1)
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestInstantMessageClose(t *testing.T) {
	s := New()
	x := s.InstantMessageClose("channel")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestInstantMessageHistory(t *testing.T) {
	s := New()
	x := s.InstantMessageHistory(HistoryArgs{Channel: "channel", Latest: "1234567890"})
	y := `{"ok":false,"error":"not_authed","messages":null,"has_more":false,"pin_count":0,"unread_count_display":0,"response_metadata":{"next_cursor":""}}`
	CheckResponse(t, x, y)
}

func TestInstantMessageList(t *testing.T) {
	s := New()
	x := s.InstantMessageList()
	y := `{"ok":false,"error":"not_authed","ims":null}`
	CheckResponse(t, x, y)
}

func TestInstantMessageMark(t *testing.T) {
	s := New()
	x := s.InstantMessageMark("channel", "1234567890")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestInstantMessageMyHistory(t *testing.T) {
	s := New()
	x := s.InstantMessageMyHistory("channel", "1234567890")
	y := `{"Filtered":0,"Latest":"","Messages":null,"Oldest":"","Total":0,"Username":""}`
	CheckResponse(t, x, y)
}

func TestInstantMessageOpen(t *testing.T) {
	s := New()
	x := s.InstantMessageOpen("user")
	y := `{"ok":false,"error":"not_authed","already_open":false,"no_op":false,"channel":{"id":""}}`
	CheckResponse(t, x, y)
}

func TestInstantMessagePurgeHistory(t *testing.T) {
	s := New()
	x := s.InstantMessagePurgeHistory("channel", "1234567890", true)
	y := `{"Deleted":0,"NotDeleted":0,"Messages":null}`
	CheckResponse(t, x, y)
}

func TestEventlogHistory(t *testing.T) {
	s := New()
	x := s.EventlogHistory("1234567890")
	y := `{"ok":false,"error":"not_authed","events":null,"has_more":false,"total":0}`
	CheckResponse(t, x, y)
}

func TestHelpIssuesList(t *testing.T) {
	s := New()
	x := s.HelpIssuesList()
	y := `{"ok":false,"error":"not_authed","issues":null}`
	CheckResponse(t, x, y)
}

func TestMigrationExchange(t *testing.T) {
	s := New()
	x := s.MigrationExchange([]string{}, false)
	y := `{"ok":false,"error":"not_authed","team_id":"","enterprise_id":"","user_id_map":null,"invalid_user_ids":null}`
	CheckResponse(t, x, y)
}

func TestMultiPartyInstantMessageClose(t *testing.T) {
	s := New()
	x := s.MultiPartyInstantMessageClose("1234567890")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestMultiPartyInstantMessageHistory(t *testing.T) {
	s := New()
	x := s.MultiPartyInstantMessageHistory(HistoryArgs{Channel: "channel", Latest: "1234567890"})
	y := `{"ok":false,"error":"not_authed","messages":null,"has_more":false,"pin_count":0,"unread_count_display":0,"response_metadata":{"next_cursor":""}}`
	CheckResponse(t, x, y)
}
