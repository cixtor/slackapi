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
	x := s.AuthTest()
	y := `{"ok":false,"error":"not_authed","team":"","team_id":"","url":"","user":"","user_id":""}`
	CheckResponse(t, x, y)
}

func TestBotsInfo(t *testing.T) {
	s := New()
	x := s.BotsInfo("user")
	y := `{"ok":false,"error":"not_authed","bot":{"id":"","deleted":false,"name":"","icons":null}}`
	CheckResponse(t, x, y)
}

func TestChannelsArchive(t *testing.T) {
	s := New()
	x := s.ChannelsArchive("channel")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestChannelsCreate(t *testing.T) {
	s := New()
	x := s.ChannelsCreate("channel")
	y := `{"ok":false,"error":"not_authed","channel":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_general":false,"is_group":false,"is_member":false,"is_mpim":false,"is_open":false,"last_read":"","latest":{"text":"","ts":"","type":"","user":""},"members":null,"name":"","num_members":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unread_count":0,"unread_count_display":0}}`
	CheckResponse(t, x, y)
}

func TestChannelsHistory(t *testing.T) {
	s := New()
	x := s.ChannelsHistory("channel", "1234567890")
	y := `{"ok":false,"error":"not_authed","has_more":false,"messages":null,"unread_count_display":0}`
	CheckResponse(t, x, y)
}

func TestChannelsID(t *testing.T) {
	s := New()
	x := s.ChannelsID("channel")
	y := `"channel"`
	CheckResponse(t, x, y)
}

func TestChannelsInfo(t *testing.T) {
	s := New()
	x := s.ChannelsInfo("channel")
	y := `{"ok":false,"error":"not_authed","channel":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_general":false,"is_group":false,"is_member":false,"is_mpim":false,"is_open":false,"last_read":"","latest":{"text":"","ts":"","type":"","user":""},"members":null,"name":"","num_members":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unread_count":0,"unread_count_display":0}}`
	CheckResponse(t, x, y)
}

func TestChannelsInvite(t *testing.T) {
	s := New()
	x := s.ChannelsInvite("channel", "user")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestChannelsJoin(t *testing.T) {
	s := New()
	x := s.ChannelsJoin("channel")
	y := `{"ok":false,"error":"not_authed","already_in_channel":false,"channel":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_general":false,"is_group":false,"is_member":false,"is_mpim":false,"is_open":false,"last_read":"","latest":{"text":"","ts":"","type":"","user":""},"members":null,"name":"","num_members":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unread_count":0,"unread_count_display":0}}`
	CheckResponse(t, x, y)
}

func TestChannelsKick(t *testing.T) {
	s := New()
	x := s.ChannelsKick("channel", "user")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestChannelsLeave(t *testing.T) {
	s := New()
	x := s.ChannelsLeave("channel")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestChannelsList(t *testing.T) {
	s := New()
	x := s.ChannelsList()
	y := `{"ok":false,"error":"not_authed","channels":null}`
	CheckResponse(t, x, y)
}

func TestChannelsMark(t *testing.T) {
	s := New()
	x := s.ChannelsMark("channel", "1234567890")
	y := `{"ok":false,"error":"not_authed"}`
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

func TestChannelsRename(t *testing.T) {
	s := New()
	x := s.ChannelsRename("channel", "lennahc")
	y := `{"ok":false,"error":"not_authed","channel":{"id":"","is_channel":false,"is_group":false,"name":"","created":0}}`
	CheckResponse(t, x, y)
}

func TestChannelsSetPurpose(t *testing.T) {
	s := New()
	x := s.ChannelsSetPurpose("channel", "purpose")
	y := `{"ok":false,"error":"not_authed","purpose":""}`
	CheckResponse(t, x, y)
}

func TestChannelsSetRetention(t *testing.T) {
	s := New()
	x := s.ChannelsSetRetention("channel", 1)
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestChannelsSetTopic(t *testing.T) {
	s := New()
	x := s.ChannelsSetTopic("channel", "topic")
	y := `{"ok":false,"error":"not_authed","topic":""}`
	CheckResponse(t, x, y)
}

func TestChannelsSuggestions(t *testing.T) {
	s := New()
	x := s.ChannelsSuggestions()
	y := `{"ok":false,"error":"not_authed","status":{"ok":false},"suggestion_types_tried":null}`
	CheckResponse(t, x, y)
}

func TestChannelsUnarchive(t *testing.T) {
	s := New()
	x := s.ChannelsUnarchive("channel")
	y := `{"ok":false,"error":"not_authed"}`
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
	y := `{"ok":false,"error":"not_authed","channel":"","ts":"","message":{"display_as_bot":false,"pinned_to":null}}`
	CheckResponse(t, x, y)
}

func TestChatUpdate(t *testing.T) {
	s := New()
	x := s.ChatUpdate(MessageArgs{})
	y := `{"ok":false,"error":"not_authed","channel":"","ts":"","message":{"display_as_bot":false,"pinned_to":null}}`
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
	y := `{"ok":false,"error":"not_authed","comment":{"comment":"","created":0,"id":"","timestamp":0,"user":""}}`
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
	y := `{"ok":false,"error":"not_authed","comment":{"comment":"","created":0,"id":"","timestamp":0,"user":""}}`
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
	y := `{"ok":false,"error":"not_authed","file":{"channels":null,"comments_count":0,"created":0,"display_as_bot":false,"editable":false,"edit_link":"","external_type":"","filetype":"","groups":null,"id":"","image_exif_rotation":0,"ims":null,"initial_comment":{"comment":"","created":0,"id":"","timestamp":0,"user":""},"is_external":false,"is_public":false,"is_starred":false,"lines":0,"lines_more":0,"mimetype":"","mode":"","name":"","num_stars":0,"original_h":0,"original_w":0,"permalink":"","pretty_type":"","preview":"","preview_highlight":"","public_url_shared":false,"size":0,"thumb_160":"","thumb_360":"","thumb_360_gif":"","thumb_360_h":0,"thumb_360_w":0,"thumb_480":"","thumb_480_h":0,"thumb_480_w":0,"thumb_64":"","thumb_80":"","timestamp":0,"title":"","url":"","url_download":"","url_private":"","url_private_download":"","user":"","username":""},"comments":null,"paging":{"count":0,"page":0,"pages":0,"total":0}}`
	CheckResponse(t, x, y)
}

func TestFilesList(t *testing.T) {
	s := New()
	x := s.FilesList(FileListArgs{})
	y := `{"ok":false,"error":"not_authed","files":null,"paging":{"count":0,"page":0,"pages":0,"total":0}}`
	CheckResponse(t, x, y)
}

func TestFilesUpload(t *testing.T) {
	s := New()
	x := s.FilesUpload(FileUploadArgs{})
	y := `{"ok":false,"error":"not_authed","file":{"channels":null,"comments_count":0,"created":0,"display_as_bot":false,"editable":false,"edit_link":"","external_type":"","filetype":"","groups":null,"id":"","image_exif_rotation":0,"ims":null,"initial_comment":{"comment":"","created":0,"id":"","timestamp":0,"user":""},"is_external":false,"is_public":false,"is_starred":false,"lines":0,"lines_more":0,"mimetype":"","mode":"","name":"","num_stars":0,"original_h":0,"original_w":0,"permalink":"","pretty_type":"","preview":"","preview_highlight":"","public_url_shared":false,"size":0,"thumb_160":"","thumb_360":"","thumb_360_gif":"","thumb_360_h":0,"thumb_360_w":0,"thumb_480":"","thumb_480_h":0,"thumb_480_w":0,"thumb_64":"","thumb_80":"","timestamp":0,"title":"","url":"","url_download":"","url_private":"","url_private_download":"","user":"","username":""}}`
	CheckResponse(t, x, y)
}

func TestGroupsArchive(t *testing.T) {
	s := New()
	x := s.GroupsArchive("channel")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestGroupsClose(t *testing.T) {
	s := New()
	x := s.GroupsClose("channel")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestGroupsCreate(t *testing.T) {
	s := New()
	x := s.GroupsCreate("channel")
	y := `{"ok":false,"error":"not_authed","group":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_general":false,"is_group":false,"is_member":false,"is_mpim":false,"is_open":false,"last_read":"","latest":{"text":"","ts":"","type":"","user":""},"members":null,"name":"","num_members":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unread_count":0,"unread_count_display":0}}`
	CheckResponse(t, x, y)
}

func TestGroupsCreateChild(t *testing.T) {
	s := New()
	x := s.GroupsCreateChild("channel")
	y := `{"ok":false,"error":"not_authed","group":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_general":false,"is_group":false,"is_member":false,"is_mpim":false,"is_open":false,"last_read":"","latest":{"text":"","ts":"","type":"","user":""},"members":null,"name":"","num_members":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unread_count":0,"unread_count_display":0}}`
	CheckResponse(t, x, y)
}

func TestGroupsHistory(t *testing.T) {
	s := New()
	x := s.GroupsHistory("channel", "1234567890")
	y := `{"ok":false,"error":"not_authed","has_more":false,"messages":null,"unread_count_display":0}`
	CheckResponse(t, x, y)
}

func TestGroupsID(t *testing.T) {
	s := New()
	x := s.GroupsID("channel")
	y := `"channel"`
	CheckResponse(t, x, y)
}

func TestGroupsInfo(t *testing.T) {
	s := New()
	x := s.GroupsInfo("channel")
	y := `{"ok":false,"error":"not_authed","group":{"created":0,"creator":"","id":"","is_archived":false,"is_channel":false,"is_general":false,"is_group":false,"is_member":false,"is_mpim":false,"is_open":false,"last_read":"","latest":{"text":"","ts":"","type":"","user":""},"members":null,"name":"","num_members":0,"purpose":{"creator":"","last_set":0,"value":""},"topic":{"creator":"","last_set":0,"value":""},"unread_count":0,"unread_count_display":0}}`
	CheckResponse(t, x, y)
}

func TestGroupsInvite(t *testing.T) {
	s := New()
	x := s.GroupsInvite("channel", "user")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestGroupsKick(t *testing.T) {
	s := New()
	x := s.GroupsKick("channel", "user")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestGroupsLeave(t *testing.T) {
	s := New()
	x := s.GroupsLeave("channel")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestGroupsList(t *testing.T) {
	s := New()
	x := s.GroupsList()
	y := `{"ok":false,"error":"not_authed","groups":null}`
	CheckResponse(t, x, y)
}

func TestGroupsMark(t *testing.T) {
	s := New()
	x := s.GroupsMark("channel", "1234567890")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestGroupsMyHistory(t *testing.T) {
	s := New()
	x := s.GroupsMyHistory("channel", "1234567890")
	y := `{"Filtered":0,"Latest":"","Messages":null,"Oldest":"","Total":0,"Username":""}`
	CheckResponse(t, x, y)
}

func TestGroupsOpen(t *testing.T) {
	s := New()
	x := s.GroupsOpen("channel")
	y := `{"ok":false,"error":"not_authed","already_open":false,"no_op":false,"channel":{"id":""}}`
	CheckResponse(t, x, y)
}

func TestGroupsPurgeHistory(t *testing.T) {
	s := New()
	x := s.GroupsPurgeHistory("channel", "1234567890", true)
	y := `{"Deleted":0,"NotDeleted":0,"Messages":null}`
	CheckResponse(t, x, y)
}

func TestGroupsRename(t *testing.T) {
	s := New()
	x := s.GroupsRename("channel", "lennahc")
	y := `{"ok":false,"error":"not_authed","channel":{"id":"","is_channel":false,"is_group":false,"name":"","created":0}}`
	CheckResponse(t, x, y)
}

func TestGroupsSetPurpose(t *testing.T) {
	s := New()
	x := s.GroupsSetPurpose("channel", "purpose")
	y := `{"ok":false,"error":"not_authed","purpose":""}`
	CheckResponse(t, x, y)
}

func TestGroupsSetRetention(t *testing.T) {
	s := New()
	x := s.GroupsSetRetention("channel", 1)
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestGroupsSetTopic(t *testing.T) {
	s := New()
	x := s.GroupsSetTopic("channel", "topic")
	y := `{"ok":false,"error":"not_authed","topic":""}`
	CheckResponse(t, x, y)
}

func TestGroupsUnarchive(t *testing.T) {
	s := New()
	x := s.GroupsUnarchive("channel")
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
	x := s.InstantMessageHistory("channel", "1234567890")
	y := `{"ok":false,"error":"not_authed","has_more":false,"messages":null,"unread_count_display":0}`
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

func TestMultiPartyInstantMessageHistory(t *testing.T) {
	s := New()
	x := s.MultiPartyInstantMessageHistory("channel", "1234567890")
	y := `{"ok":false,"error":"not_authed","has_more":false,"messages":null,"unread_count_display":0}`
	CheckResponse(t, x, y)
}

func TestMultiPartyInstantMessageList(t *testing.T) {
	s := New()
	x := s.MultiPartyInstantMessageList()
	y := `{"ok":false,"error":"not_authed","groups":null}`
	CheckResponse(t, x, y)
}

func TestMultiPartyInstantMessageListSimple(t *testing.T) {
	s := New()
	x := s.MultiPartyInstantMessageListSimple()
	y := `{}`
	CheckResponse(t, x, y)
}

func TestMultiPartyInstantMessageMyHistory(t *testing.T) {
	s := New()
	x := s.MultiPartyInstantMessageMyHistory("channel", "1234567890")
	y := `{"Filtered":0,"Latest":"","Messages":null,"Oldest":"","Total":0,"Username":""}`
	CheckResponse(t, x, y)
}

func TestMultiPartyInstantMessagePurgeHistory(t *testing.T) {
	s := New()
	x := s.MultiPartyInstantMessagePurgeHistory("channel", "1234567890", true)
	y := `{"Deleted":0,"NotDeleted":0,"Messages":null}`
	CheckResponse(t, x, y)
}

func TestReactionsAdd(t *testing.T) {
	s := New()
	x := s.ReactionsAdd(ReactionArgs{})
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestReactionsGet(t *testing.T) {
	s := New()
	x := s.ReactionsGet(ReactionArgs{})
	y := `{"ok":false,"error":"not_authed","channel":"","file":{"channels":null,"comments_count":0,"created":0,"display_as_bot":false,"editable":false,"edit_link":"","external_type":"","filetype":"","groups":null,"id":"","image_exif_rotation":0,"ims":null,"initial_comment":{"comment":"","created":0,"id":"","timestamp":0,"user":""},"is_external":false,"is_public":false,"is_starred":false,"lines":0,"lines_more":0,"mimetype":"","mode":"","name":"","num_stars":0,"original_h":0,"original_w":0,"permalink":"","pretty_type":"","preview":"","preview_highlight":"","public_url_shared":false,"size":0,"thumb_160":"","thumb_360":"","thumb_360_gif":"","thumb_360_h":0,"thumb_360_w":0,"thumb_480":"","thumb_480_h":0,"thumb_480_w":0,"thumb_64":"","thumb_80":"","timestamp":0,"title":"","url":"","url_download":"","url_private":"","url_private_download":"","user":"","username":"","reactions":null},"file_comment":"","message":{"reactions":null,"text":"","ts":"","type":"","user":""},"type":"","ts":""}`
	CheckResponse(t, x, y)
}

func TestReactionsList(t *testing.T) {
	s := New()
	x := s.ReactionsList(ReactionListArgs{})
	y := `{"ok":false,"error":"not_authed","items":null,"paging":{"count":0,"page":0,"pages":0,"total":0}}`
	CheckResponse(t, x, y)
}

func TestReactionsRemove(t *testing.T) {
	s := New()
	x := s.ReactionsRemove(ReactionArgs{})
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestResourceArchive(t *testing.T) {
	s := New()
	x := s.ResourceArchive("action", "channel")
	y := `{"ok":false,"error":"unknown_method"}`
	CheckResponse(t, x, y)
}

func TestResourceHistory(t *testing.T) {
	s := New()
	x := s.ResourceHistory("action", HistoryArgs{})
	y := `{"ok":false,"error":"unknown_method","has_more":false,"messages":null,"unread_count_display":0}`
	CheckResponse(t, x, y)
}

func TestResourceInvite(t *testing.T) {
	s := New()
	x := s.ResourceInvite("action", "channel", "user")
	y := `{"ok":false,"error":"unknown_method"}`
	CheckResponse(t, x, y)
}

func TestResourceKick(t *testing.T) {
	s := New()
	x := s.ResourceKick("action", "channel", "user")
	y := `{"ok":false,"error":"unknown_method"}`
	CheckResponse(t, x, y)
}

func TestResourceLeave(t *testing.T) {
	s := New()
	x := s.ResourceLeave("action", "channel")
	y := `{"ok":false,"error":"unknown_method"}`
	CheckResponse(t, x, y)
}

func TestResourceMark(t *testing.T) {
	s := New()
	x := s.ResourceMark("action", "channel", "1234567890")
	y := `{"ok":false,"error":"unknown_method"}`
	CheckResponse(t, x, y)
}

func TestResourceMyHistory(t *testing.T) {
	s := New()
	x := s.ResourceMyHistory("action", "channel", "1234567890")
	y := `{"Filtered":0,"Latest":"","Messages":null,"Oldest":"","Total":0,"Username":""}`
	CheckResponse(t, x, y)
}

func TestResourcePurgeHistory(t *testing.T) {
	s := New()
	x := s.ResourcePurgeHistory("action", "channel", "1234567890", true)
	y := `{"Deleted":0,"NotDeleted":0,"Messages":null}`
	CheckResponse(t, x, y)
}

func TestResourceRename(t *testing.T) {
	s := New()
	x := s.ResourceRename("action", "channel", "lennahc")
	y := `{"ok":false,"error":"unknown_method","channel":{"id":"","is_channel":false,"is_group":false,"name":"","created":0}}`
	CheckResponse(t, x, y)
}

func TestResourceSetPurpose(t *testing.T) {
	s := New()
	x := s.ResourceSetPurpose("action", "channel", "purpose")
	y := `{"ok":false,"error":"unknown_method","purpose":""}`
	CheckResponse(t, x, y)
}

func TestResourceSetRetention(t *testing.T) {
	s := New()
	x := s.ResourceSetRetention("action", "channel", 1)
	y := `{"ok":false,"error":"unknown_method"}`
	CheckResponse(t, x, y)
}

func TestResourceSetTopic(t *testing.T) {
	s := New()
	x := s.ResourceSetTopic("action", "channel", "topic")
	y := `{"ok":false,"error":"unknown_method","topic":""}`
	CheckResponse(t, x, y)
}

func TestResourceUnarchive(t *testing.T) {
	s := New()
	x := s.ResourceUnarchive("action", "channel")
	y := `{"ok":false,"error":"unknown_method"}`
	CheckResponse(t, x, y)
}

func TestTeamAccessLogs(t *testing.T) {
	s := New()
	x := s.TeamAccessLogs(AccessLogArgs{})
	y := `{"ok":false,"error":"not_authed","logins":null,"needed":"","provided":""}`
	CheckResponse(t, x, y)
}

func TestTeamBillableInfo(t *testing.T) {
	s := New()
	x := s.TeamBillableInfo("user")
	y := `{"ok":false,"error":"not_authed","billable_info":null}`
	CheckResponse(t, x, y)
}

func TestTeamInfo(t *testing.T) {
	s := New()
	x := s.TeamInfo()
	y := `{"ok":false,"error":"not_authed","team":{"domain":"","email_domain":"","icon":{"image_102":"","image_132":"","image_34":"","image_44":"","image_68":"","image_88":"","image_original":""},"id":"","name":""}}`
	CheckResponse(t, x, y)
}

func TestTeamProfileGet(t *testing.T) {
	s := New()
	x := s.TeamProfileGet()
	y := `{"ok":false,"error":"not_authed","profile":{"fields":null}}`
	CheckResponse(t, x, y)
}

func TestUsersCounts(t *testing.T) {
	s := New()
	x := s.UsersCounts()
	y := `{"ok":false,"error":"not_authed","channels":null,"groups":null,"ims":null}`
	CheckResponse(t, x, y)
}

func TestUsersDeletePhoto(t *testing.T) {
	s := New()
	x := s.UsersDeletePhoto()
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestUsersGetPresence(t *testing.T) {
	s := New()
	x := s.UsersGetPresence("user")
	y := `{"ok":false,"error":"not_authed","auto_away":false,"connection_count":0,"last_activity":0,"manual_away":false,"online":false,"presence":""}`
	CheckResponse(t, x, y)
}

func TestUsersID(t *testing.T) {
	s := New()
	x := s.UsersID("user")
	y := `"user"`
	CheckResponse(t, x, y)
}

func TestUsersInfo(t *testing.T) {
	s := New()
	x := s.UsersInfo("user")
	y := `{"ok":false,"error":"not_authed","user":{"color":"","deleted":false,"has_2fa":false,"id":"","is_admin":false,"is_bot":false,"is_owner":false,"is_primary_owner":false,"is_restricted":false,"is_ultra_restricted":false,"name":"","presence":"","profile":{"api_app_id":"","bot_id":"","avatar_hash":"","email":"","fields":null,"first_name":"","image_1024":"","image_192":"","image_24":"","image_32":"","image_48":"","image_512":"","image_72":"","image_original":"","last_name":"","phone":"","real_name":"","real_name_normalized":"","status_text":"","status_emoji":"","skype":"","title":""},"real_name":"","status":"","team_id":"","two_factor_type":"","tz":"","tz_label":"","tz_offset":0}}`
	CheckResponse(t, x, y)
}

func TestUsersList(t *testing.T) {
	s := New()
	x := s.UsersList()
	y := `{"ok":false,"error":"not_authed","members":null}`
	CheckResponse(t, x, y)
}

func TestUsersListWithPresence(t *testing.T) {
	s := New()
	x := s.UsersListWithPresence()
	y := `{"ok":false,"error":"not_authed","members":null}`
	CheckResponse(t, x, y)
}

func TestUsersPrefsGet(t *testing.T) {
	s := New()
	x := s.UsersPrefsGet()
	y := `{"ok":false,"error":"not_authed","prefs":{"a11y_animations":false,"a11y_font_size":"","all_channels_loud":false,"all_notifications_prefs":null,"allow_calls_to_set_current_status":false,"all_unreads_sort_order":false,"arrow_history":false,"at_channel_suppressed_channels":"","box_enabled":false,"channel_sort":"","client_logs_pri":"","color_names_in_list":false,"confirm_clear_all_unreads":false,"confirm_sh_call_start":false,"confirm_user_marked_away":false,"convert_emoticons":false,"display_display_names":false,"display_real_names_override":0,"dnd_enabled":false,"dnd_end_hour":"","dnd_start_hour":"","dropbox_enabled":false,"email_alerts":"","email_alerts_sleep_until":0,"email_misc":false,"email_weekly":false,"emoji_autocomplete_big":false,"emoji_mode":"","emoji_use":"","enable_react_emoji_picker":false,"enable_unread_view":false,"enhanced_debugging":false,"enter_is_special_in_tbt":false,"enterprise_migration_seen":false,"expand_inline_imgs":false,"expand_internal_inline_imgs":false,"expand_non_media_attachments":false,"expand_snippets":false,"f_key_search":false,"flannel_server_pool":"","frecency_ent_jumper":"","frecency_jumper":"","fuller_timestamps":false,"full_text_extracts":false,"gdrive_authed":false,"gdrive_enabled":false,"graphic_emoticons":false,"growls_enabled":false,"growth_msg_limit_approaching_cta_count":0,"growth_msg_limit_approaching_cta_ts":0,"growth_msg_limit_long_reached_cta_count":0,"growth_msg_limit_long_reached_cta_last_ts":0,"growth_msg_limit_reached_cta_count":0,"growth_msg_limit_reached_cta_last_ts":0,"has_created_channel":false,"has_invited":false,"has_searched":false,"has_uploaded":false,"hide_hex_swatch":false,"hide_user_group_info_pane":false,"highlight_words":"","intro_to_apps_message_seen":false,"jumbomoji":false,"k_key_omnibox":false,"k_key_omnibox_auto_hide_count":0,"last_seen_at_channel_warning":0,"last_snippet_type":"","last_tos_acknowledged":"","load_lato_2":false,"locale":"","loud_channels":"","loud_channels_set":"","ls_disabled":false,"mac_ssb_bounce":"","mac_ssb_bullet":false,"mark_msgs_read_immediately":false,"measure_css_usage":false,"mentions_exclude_at_channels":false,"mentions_exclude_at_user_groups":false,"messages_theme":"","msg_preview":false,"msg_preview_persistent":false,"muted_channels":"","mute_sounds":false,"never_channels":"","new_msg_snd":"","newxp_seen_last_message":"","no_created_overlays":false,"no_invites_widget_in_sidebar":false,"no_joined_overlays":false,"no_macelectron_banner":false,"no_macssb1_banner":false,"no_macssb2_banner":false,"no_omnibox_in_channels":false,"no_text_in_notifications":false,"no_winssb1_banner":false,"obey_inline_img_limit":false,"onboarding_cancelled":false,"onboarding_slackbot_conversation_step":0,"overloaded_message_enabled":false,"pagekeys_handled":false,"posts_formatting_guide":false,"preferred_skin_tone":"","prev_next_btn":false,"privacy_policy_seen":false,"prompted_for_email_disabling":false,"push_at_channel_suppressed_channels":"","push_dm_alert":false,"push_everything":false,"push_idle_wait":0,"push_loud_channels":"","push_loud_channels_set":"","push_mention_alert":false,"push_mention_channels":"","push_show_preview":false,"push_sound":"","require_at":false,"search_exclude_bots":false,"search_exclude_channels":"","search_only_current_team":false,"search_only_my_channels":false,"search_sort":"","seen_calls_ss_main_coachmark":false,"seen_calls_ss_window_coachmark":false,"seen_calls_video_beta_coachmark":false,"seen_calls_video_ga_coachmark":false,"seen_custom_status_badge":false,"seen_custom_status_callout":false,"seen_domain_invite_reminder":false,"seen_gdrive_coachmark":false,"seen_guest_admin_slackbot_announcement":false,"seen_highlights_arrows_coachmark":false,"seen_highlights_coachmark":false,"seen_intl_channel_names_coachmark":false,"seen_member_invite_reminder":false,"seen_onboarding_channels":false,"seen_onboarding_direct_messages":false,"seen_onboarding_invites":false,"seen_onboarding_private_groups":false,"seen_onboarding_recent_mentions":false,"seen_onboarding_search":false,"seen_onboarding_slackbot_conversation":false,"seen_onboarding_starred_items":false,"seen_onboarding_start":false,"seen_replies_coachmark":false,"seen_single_emoji_msg":false,"seen_ssb_prompt":false,"seen_threads_notification_banner":false,"seen_unread_view_coachmark":false,"seen_welcome_2":false,"separate_private_channels":false,"separate_shared_channels":false,"show_all_skin_tones":false,"show_jumper_scores":false,"show_memory_instrument":false,"show_typing":false,"sidebar_behavior":"","sidebar_theme":"","sidebar_theme_custom_values":"","snippet_editor_wrap_long_lines":false,"spaces_new_xp_banner_dismissed":false,"ssb_space_window":"","ss_emojis":false,"start_scroll_at_oldest":false,"tab_ui_return_selects":false,"threads_everything":false,"time24":false,"two_factor_auth_enabled":false,"two_factor_backup_type":"","two_factor_type":"","tz":"","user_colors":"","webapp_spellcheck":false,"welcome_message_hidden":false,"whats_new_read":0,"winssb_run_from_tray":false,"winssb_window_flash_behavior":""}}`
	CheckResponse(t, x, y)
}

func TestUsersPrefsSet(t *testing.T) {
	s := New()
	x := s.UsersPrefsSet("name", "value")
	y := `{"ok":false,"error":"not_authed","prefs":{"a11y_animations":false,"a11y_font_size":"","all_channels_loud":false,"all_notifications_prefs":null,"allow_calls_to_set_current_status":false,"all_unreads_sort_order":false,"arrow_history":false,"at_channel_suppressed_channels":"","box_enabled":false,"channel_sort":"","client_logs_pri":"","color_names_in_list":false,"confirm_clear_all_unreads":false,"confirm_sh_call_start":false,"confirm_user_marked_away":false,"convert_emoticons":false,"display_display_names":false,"display_real_names_override":0,"dnd_enabled":false,"dnd_end_hour":"","dnd_start_hour":"","dropbox_enabled":false,"email_alerts":"","email_alerts_sleep_until":0,"email_misc":false,"email_weekly":false,"emoji_autocomplete_big":false,"emoji_mode":"","emoji_use":"","enable_react_emoji_picker":false,"enable_unread_view":false,"enhanced_debugging":false,"enter_is_special_in_tbt":false,"enterprise_migration_seen":false,"expand_inline_imgs":false,"expand_internal_inline_imgs":false,"expand_non_media_attachments":false,"expand_snippets":false,"f_key_search":false,"flannel_server_pool":"","frecency_ent_jumper":"","frecency_jumper":"","fuller_timestamps":false,"full_text_extracts":false,"gdrive_authed":false,"gdrive_enabled":false,"graphic_emoticons":false,"growls_enabled":false,"growth_msg_limit_approaching_cta_count":0,"growth_msg_limit_approaching_cta_ts":0,"growth_msg_limit_long_reached_cta_count":0,"growth_msg_limit_long_reached_cta_last_ts":0,"growth_msg_limit_reached_cta_count":0,"growth_msg_limit_reached_cta_last_ts":0,"has_created_channel":false,"has_invited":false,"has_searched":false,"has_uploaded":false,"hide_hex_swatch":false,"hide_user_group_info_pane":false,"highlight_words":"","intro_to_apps_message_seen":false,"jumbomoji":false,"k_key_omnibox":false,"k_key_omnibox_auto_hide_count":0,"last_seen_at_channel_warning":0,"last_snippet_type":"","last_tos_acknowledged":"","load_lato_2":false,"locale":"","loud_channels":"","loud_channels_set":"","ls_disabled":false,"mac_ssb_bounce":"","mac_ssb_bullet":false,"mark_msgs_read_immediately":false,"measure_css_usage":false,"mentions_exclude_at_channels":false,"mentions_exclude_at_user_groups":false,"messages_theme":"","msg_preview":false,"msg_preview_persistent":false,"muted_channels":"","mute_sounds":false,"never_channels":"","new_msg_snd":"","newxp_seen_last_message":"","no_created_overlays":false,"no_invites_widget_in_sidebar":false,"no_joined_overlays":false,"no_macelectron_banner":false,"no_macssb1_banner":false,"no_macssb2_banner":false,"no_omnibox_in_channels":false,"no_text_in_notifications":false,"no_winssb1_banner":false,"obey_inline_img_limit":false,"onboarding_cancelled":false,"onboarding_slackbot_conversation_step":0,"overloaded_message_enabled":false,"pagekeys_handled":false,"posts_formatting_guide":false,"preferred_skin_tone":"","prev_next_btn":false,"privacy_policy_seen":false,"prompted_for_email_disabling":false,"push_at_channel_suppressed_channels":"","push_dm_alert":false,"push_everything":false,"push_idle_wait":0,"push_loud_channels":"","push_loud_channels_set":"","push_mention_alert":false,"push_mention_channels":"","push_show_preview":false,"push_sound":"","require_at":false,"search_exclude_bots":false,"search_exclude_channels":"","search_only_current_team":false,"search_only_my_channels":false,"search_sort":"","seen_calls_ss_main_coachmark":false,"seen_calls_ss_window_coachmark":false,"seen_calls_video_beta_coachmark":false,"seen_calls_video_ga_coachmark":false,"seen_custom_status_badge":false,"seen_custom_status_callout":false,"seen_domain_invite_reminder":false,"seen_gdrive_coachmark":false,"seen_guest_admin_slackbot_announcement":false,"seen_highlights_arrows_coachmark":false,"seen_highlights_coachmark":false,"seen_intl_channel_names_coachmark":false,"seen_member_invite_reminder":false,"seen_onboarding_channels":false,"seen_onboarding_direct_messages":false,"seen_onboarding_invites":false,"seen_onboarding_private_groups":false,"seen_onboarding_recent_mentions":false,"seen_onboarding_search":false,"seen_onboarding_slackbot_conversation":false,"seen_onboarding_starred_items":false,"seen_onboarding_start":false,"seen_replies_coachmark":false,"seen_single_emoji_msg":false,"seen_ssb_prompt":false,"seen_threads_notification_banner":false,"seen_unread_view_coachmark":false,"seen_welcome_2":false,"separate_private_channels":false,"separate_shared_channels":false,"show_all_skin_tones":false,"show_jumper_scores":false,"show_memory_instrument":false,"show_typing":false,"sidebar_behavior":"","sidebar_theme":"","sidebar_theme_custom_values":"","snippet_editor_wrap_long_lines":false,"spaces_new_xp_banner_dismissed":false,"ssb_space_window":"","ss_emojis":false,"start_scroll_at_oldest":false,"tab_ui_return_selects":false,"threads_everything":false,"time24":false,"two_factor_auth_enabled":false,"two_factor_backup_type":"","two_factor_type":"","tz":"","user_colors":"","webapp_spellcheck":false,"welcome_message_hidden":false,"whats_new_read":0,"winssb_run_from_tray":false,"winssb_window_flash_behavior":""}}`
	CheckResponse(t, x, y)
}

func TestUsersPreparePhoto(t *testing.T) {
	s := New()
	x := s.UsersPreparePhoto("image.jpg")
	y := `{"ok":false,"error":"not_authed","id":"","url":""}`
	CheckResponse(t, x, y)
}

func TestUsersProfileGet(t *testing.T) {
	s := New()
	x := s.UsersProfileGet("user")
	y := `{"ok":false,"error":"not_authed","profile":{"api_app_id":"","bot_id":"","avatar_hash":"","email":"","fields":null,"first_name":"","image_1024":"","image_192":"","image_24":"","image_32":"","image_48":"","image_512":"","image_72":"","image_original":"","last_name":"","phone":"","real_name":"","real_name_normalized":"","status_text":"","status_emoji":"","skype":"","title":""}}`
	CheckResponse(t, x, y)
}

func TestUsersProfileGetWithLabels(t *testing.T) {
	s := New()
	x := s.UsersProfileGetWithLabels("user")
	y := `{"ok":false,"error":"not_authed","profile":{"api_app_id":"","bot_id":"","avatar_hash":"","email":"","fields":null,"first_name":"","image_1024":"","image_192":"","image_24":"","image_32":"","image_48":"","image_512":"","image_72":"","image_original":"","last_name":"","phone":"","real_name":"","real_name_normalized":"","status_text":"","status_emoji":"","skype":"","title":""}}`
	CheckResponse(t, x, y)
}

func TestUsersProfileSet(t *testing.T) {
	s := New()
	x := s.UsersProfileSet("name", "value")
	y := `{"ok":false,"error":"not_authed","profile":{"api_app_id":"","bot_id":"","avatar_hash":"","email":"","fields":null,"first_name":"","image_1024":"","image_192":"","image_24":"","image_32":"","image_48":"","image_512":"","image_72":"","image_original":"","last_name":"","phone":"","real_name":"","real_name_normalized":"","status_text":"","status_emoji":"","skype":"","title":""}}`
	CheckResponse(t, x, y)
}

func TestUsersProfileSetMultiple(t *testing.T) {
	s := New()
	x := s.UsersProfileSetMultiple("{}")
	y := `{"ok":false,"error":"not_authed","profile":{"api_app_id":"","bot_id":"","avatar_hash":"","email":"","fields":null,"first_name":"","image_1024":"","image_192":"","image_24":"","image_32":"","image_48":"","image_512":"","image_72":"","image_original":"","last_name":"","phone":"","real_name":"","real_name_normalized":"","status_text":"","status_emoji":"","skype":"","title":""}}`
	CheckResponse(t, x, y)
}

func TestUsersSearch(t *testing.T) {
	s := New()
	x := s.UsersSearch("user")
	y := `null`
	CheckResponse(t, x, y)
}

func TestUsersSetActive(t *testing.T) {
	s := New()
	x := s.UsersSetActive()
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestUsersSetAvatar(t *testing.T) {
	s := New()
	x := s.UsersSetAvatar("image.jpg")
	y := `{"ok":false,"upload_id":"","upload_url":"","profile":{"avatar_hash":"","image_1024":"","image_192":"","image_24":"","image_32":"","image_48":"","image_512":"","image_72":"","image_original":""}}`
	CheckResponse(t, x, y)
}

func TestUsersSetPhoto(t *testing.T) {
	s := New()
	x := s.UsersSetPhoto("imageid")
	y := `{"ok":false,"error":"not_authed","profile":{"avatar_hash":"","image_1024":"","image_192":"","image_24":"","image_32":"","image_48":"","image_512":"","image_72":"","image_original":""}}`
	CheckResponse(t, x, y)
}

func TestUsersSetPresence(t *testing.T) {
	s := New()
	x := s.UsersSetPresence("value")
	y := `{"ok":false,"error":"not_authed"}`
	CheckResponse(t, x, y)
}

func TestUsersSetStatus(t *testing.T) {
	s := New()
	x := s.UsersSetStatus(":slack:", "status")
	y := `{"ok":false,"error":"not_authed","profile":{"api_app_id":"","bot_id":"","avatar_hash":"","email":"","fields":null,"first_name":"","image_1024":"","image_192":"","image_24":"","image_32":"","image_48":"","image_512":"","image_72":"","image_original":"","last_name":"","phone":"","real_name":"","real_name_normalized":"","status_text":"","status_emoji":"","skype":"","title":""}}`
	CheckResponse(t, x, y)
}