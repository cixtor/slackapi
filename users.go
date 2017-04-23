package main

import (
	"strings"
)

type ResponseUsersInfo struct {
	Response
	User UserData `json:"user"`
}

type ResponseUsersGetPresence struct {
	Response
	UserPresence
}

type ResponseUsersList struct {
	Response
	Members []UserData `json:"members"`
}

type UserPresence struct {
	AutoAway        bool   `json:"auto_away"`
	ConnectionCount int    `json:"connection_count"`
	LastActivity    int    `json:"last_activity"`
	ManualAway      bool   `json:"manual_away"`
	Online          bool   `json:"online"`
	Presence        string `json:"presence"`
}

type UserData struct {
	Color             string      `json:"color"`
	Deleted           bool        `json:"deleted"`
	Has2fa            bool        `json:"has_2fa"`
	Id                string      `json:"id"`
	IsAdmin           bool        `json:"is_admin"`
	IsBot             bool        `json:"is_bot"`
	IsOwner           bool        `json:"is_owner"`
	IsPrimaryOwner    bool        `json:"is_primary_owner"`
	IsRestricted      bool        `json:"is_restricted"`
	IsUltraRestricted bool        `json:"is_ultra_restricted"`
	Name              string      `json:"name"`
	Presence          string      `json:"presence"`
	Profile           UserProfile `json:"profile"`
	RealName          string      `json:"real_name"`
	Status            string      `json:"status"`
	TeamId            string      `json:"team_id"`
	TwoFactorType     string      `json:"two_factor_type"`
	Tz                string      `json:"tz"`
	TzLabel           string      `json:"tz_label"`
	TzOffset          int         `json:"tz_offset"`
}

type UserProfile struct {
	ApiAppID           string      `json:"api_app_id"`
	BotID              string      `json:"bot_id"`
	AvatarHash         string      `json:"avatar_hash"`
	Email              string      `json:"email"`
	Fields             interface{} `json:"fields"`
	FirstName          string      `json:"first_name"`
	Image1024          string      `json:"image_1024"`
	Image192           string      `json:"image_192"`
	Image24            string      `json:"image_24"`
	Image32            string      `json:"image_32"`
	Image48            string      `json:"image_48"`
	Image512           string      `json:"image_512"`
	Image72            string      `json:"image_72"`
	ImageOriginal      string      `json:"image_original"`
	LastName           string      `json:"last_name"`
	Phone              string      `json:"phone"`
	RealName           string      `json:"real_name"`
	RealNameNormalized string      `json:"real_name_normalized"`
	StatusText         string      `json:"status_text"`
	StatusEmoji        string      `json:"status_emoji"`
	Skype              string      `json:"skype"`
	Title              string      `json:"title"`
}

type ResponseUserPrefs struct {
	Response
	Prefs UserPrefs `json:"prefs"`
}

type ResponseUserIdentity struct {
	Response
	Profile UserProfile `json:"profile"`
}

type ResponseUserAvatar struct {
	Response
	UploadID  string                   `json:"upload_id"`
	UploadURL string                   `json:"upload_url"`
	Profile   ResponseUserPhotoProfile `json:"profile"`
}

type ResponseUserPhoto struct {
	Response
	Profile ResponseUserPhotoProfile `json:"profile"`
}

type ResponseUserPhotoProfile struct {
	AvatarHash    string `json:"avatar_hash"`
	Image1024     string `json:"image_1024"`
	Image192      string `json:"image_192"`
	Image24       string `json:"image_24"`
	Image32       string `json:"image_32"`
	Image48       string `json:"image_48"`
	Image512      string `json:"image_512"`
	Image72       string `json:"image_72"`
	ImageOriginal string `json:"image_original"`
}

type ResponseUserPhotoUpload struct {
	Response
	ID  string `json:"id"`
	URL string `json:"url"`
}

type ResponseUsersCounts struct {
	Response
	Channels []CountChannel `json:"channels"`
	Groups   []CountGroup   `json:"groups"`
	Ims      []CountIm      `json:"ims"`
}

type CountChannel struct {
	ID                  string `json:"id"`
	IsArchived          bool   `json:"is_archived"`
	IsGeneral           bool   `json:"is_general"`
	IsMember            bool   `json:"is_member"`
	IsMuted             bool   `json:"is_muted"`
	IsStarred           bool   `json:"is_starred"`
	MentionCount        int    `json:"mention_count"`
	MentionCountDisplay int    `json:"mention_count_display"`
	Name                string `json:"name"`
	NameNormalized      string `json:"name_normalized"`
	UnreadCount         int    `json:"unread_count"`
	UnreadCountDisplay  int    `json:"unread_count_display"`
}

type CountGroup struct {
	ID                  string `json:"id"`
	IsArchived          bool   `json:"is_archived"`
	IsMpim              bool   `json:"is_mpim"`
	IsMuted             bool   `json:"is_muted"`
	IsOpen              bool   `json:"is_open"`
	IsStarred           bool   `json:"is_starred"`
	MentionCount        int    `json:"mention_count"`
	MentionCountDisplay int    `json:"mention_count_display"`
	Name                string `json:"name"`
	NameNormalized      string `json:"name_normalized"`
	UnreadCount         int    `json:"unread_count"`
	UnreadCountDisplay  int    `json:"unread_count_display"`
}

type CountIm struct {
	DmCount   int    `json:"dm_count"`
	ID        string `json:"id"`
	IsOpen    bool   `json:"is_open"`
	IsStarred bool   `json:"is_starred"`
	Name      string `json:"name"`
	UserID    string `json:"user_id"`
}

type UserPrefs struct {
	A11yAnimations                     bool        `json:"a11y_animations"`
	A11yFontSize                       string      `json:"a11y_font_size"`
	AllChannelsLoud                    bool        `json:"all_channels_loud"`
	AllNotificationsPrefs              interface{} `json:"all_notifications_prefs"`
	AllowCallsToSetCurrentStatus       bool        `json:"allow_calls_to_set_current_status"`
	AllUnreadsSortOrder                bool        `json:"all_unreads_sort_order"`
	ArrowHistory                       bool        `json:"arrow_history"`
	AtChannelSuppressedChannels        string      `json:"at_channel_suppressed_channels"`
	BoxEnabled                         bool        `json:"box_enabled"`
	ChannelSort                        string      `json:"channel_sort"`
	ClientLogsPri                      string      `json:"client_logs_pri"`
	ColorNamesInList                   bool        `json:"color_names_in_list"`
	ConfirmClearAllUnreads             bool        `json:"confirm_clear_all_unreads"`
	ConfirmShCallStart                 bool        `json:"confirm_sh_call_start"`
	ConfirmUserMarkedAway              bool        `json:"confirm_user_marked_away"`
	ConvertEmoticons                   bool        `json:"convert_emoticons"`
	DisplayDisplayNames                bool        `json:"display_display_names"`
	DisplayRealNamesOverride           int         `json:"display_real_names_override"`
	DndEnabled                         bool        `json:"dnd_enabled"`
	DndEndHour                         string      `json:"dnd_end_hour"`
	DndStartHour                       string      `json:"dnd_start_hour"`
	DropboxEnabled                     bool        `json:"dropbox_enabled"`
	EmailAlerts                        string      `json:"email_alerts"`
	EmailAlertsSleepUntil              int         `json:"email_alerts_sleep_until"`
	EmailMisc                          bool        `json:"email_misc"`
	EmailWeekly                        bool        `json:"email_weekly"`
	EmojiAutocompleteBig               bool        `json:"emoji_autocomplete_big"`
	EmojiMode                          string      `json:"emoji_mode"`
	EmojiUse                           string      `json:"emoji_use"`
	EnableReactEmojiPicker             bool        `json:"enable_react_emoji_picker"`
	EnableUnreadView                   bool        `json:"enable_unread_view"`
	EnhancedDebugging                  bool        `json:"enhanced_debugging"`
	EnterIsSpecialInTbt                bool        `json:"enter_is_special_in_tbt"`
	EnterpriseMigrationSeen            bool        `json:"enterprise_migration_seen"`
	ExpandInlineImgs                   bool        `json:"expand_inline_imgs"`
	ExpandInternalInlineImgs           bool        `json:"expand_internal_inline_imgs"`
	ExpandNonMediaAttachments          bool        `json:"expand_non_media_attachments"`
	ExpandSnippets                     bool        `json:"expand_snippets"`
	FKeySearch                         bool        `json:"f_key_search"`
	FlannelServerPool                  string      `json:"flannel_server_pool"`
	FrecencyEntJumper                  string      `json:"frecency_ent_jumper"`
	FrecencyJumper                     string      `json:"frecency_jumper"`
	FullerTimestamps                   bool        `json:"fuller_timestamps"`
	FullTextExtracts                   bool        `json:"full_text_extracts"`
	GdriveAuthed                       bool        `json:"gdrive_authed"`
	GdriveEnabled                      bool        `json:"gdrive_enabled"`
	GraphicEmoticons                   bool        `json:"graphic_emoticons"`
	GrowlsEnabled                      bool        `json:"growls_enabled"`
	GrowthMsgLimitApproachingCtaCount  int         `json:"growth_msg_limit_approaching_cta_count"`
	GrowthMsgLimitApproachingCtaTs     int         `json:"growth_msg_limit_approaching_cta_ts"`
	GrowthMsgLimitLongReachedCtaCount  int         `json:"growth_msg_limit_long_reached_cta_count"`
	GrowthMsgLimitLongReachedCtaLastTs int         `json:"growth_msg_limit_long_reached_cta_last_ts"`
	GrowthMsgLimitReachedCtaCount      int         `json:"growth_msg_limit_reached_cta_count"`
	GrowthMsgLimitReachedCtaLastTs     int         `json:"growth_msg_limit_reached_cta_last_ts"`
	HasCreatedChannel                  bool        `json:"has_created_channel"`
	HasInvited                         bool        `json:"has_invited"`
	HasSearched                        bool        `json:"has_searched"`
	HasUploaded                        bool        `json:"has_uploaded"`
	HideHexSwatch                      bool        `json:"hide_hex_swatch"`
	HideUserGroupInfoPane              bool        `json:"hide_user_group_info_pane"`
	HighlightWords                     string      `json:"highlight_words"`
	IntroToAppsMessageSeen             bool        `json:"intro_to_apps_message_seen"`
	Jumbomoji                          bool        `json:"jumbomoji"`
	KKeyOmnibox                        bool        `json:"k_key_omnibox"`
	KKeyOmniboxAutoHideCount           int         `json:"k_key_omnibox_auto_hide_count"`
	LastSeenAtChannelWarning           int         `json:"last_seen_at_channel_warning"`
	LastSnippetType                    string      `json:"last_snippet_type"`
	LastTosAcknowledged                string      `json:"last_tos_acknowledged"`
	LoadLato2                          bool        `json:"load_lato_2"`
	Locale                             string      `json:"locale"`
	LoudChannels                       string      `json:"loud_channels"`
	LoudChannelsSet                    string      `json:"loud_channels_set"`
	LsDisabled                         bool        `json:"ls_disabled"`
	MacSsbBounce                       string      `json:"mac_ssb_bounce"`
	MacSsbBullet                       bool        `json:"mac_ssb_bullet"`
	MarkMsgsReadImmediately            bool        `json:"mark_msgs_read_immediately"`
	MeasureCSSUsage                    bool        `json:"measure_css_usage"`
	MentionsExcludeAtChannels          bool        `json:"mentions_exclude_at_channels"`
	MentionsExcludeAtUserGroups        bool        `json:"mentions_exclude_at_user_groups"`
	MessagesTheme                      string      `json:"messages_theme"`
	MsgPreview                         bool        `json:"msg_preview"`
	MsgPreviewPersistent               bool        `json:"msg_preview_persistent"`
	MutedChannels                      string      `json:"muted_channels"`
	MuteSounds                         bool        `json:"mute_sounds"`
	NeverChannels                      string      `json:"never_channels"`
	NewMsgSnd                          string      `json:"new_msg_snd"`
	NewxpSeenLastMessage               string      `json:"newxp_seen_last_message"`
	NoCreatedOverlays                  bool        `json:"no_created_overlays"`
	NoInvitesWidgetInSidebar           bool        `json:"no_invites_widget_in_sidebar"`
	NoJoinedOverlays                   bool        `json:"no_joined_overlays"`
	NoMacelectronBanner                bool        `json:"no_macelectron_banner"`
	NoMacssb1Banner                    bool        `json:"no_macssb1_banner"`
	NoMacssb2Banner                    bool        `json:"no_macssb2_banner"`
	NoOmniboxInChannels                bool        `json:"no_omnibox_in_channels"`
	NoTextInNotifications              bool        `json:"no_text_in_notifications"`
	NoWinssb1Banner                    bool        `json:"no_winssb1_banner"`
	ObeyInlineImgLimit                 bool        `json:"obey_inline_img_limit"`
	OnboardingCancelled                bool        `json:"onboarding_cancelled"`
	OnboardingSlackbotConversationStep int         `json:"onboarding_slackbot_conversation_step"`
	OverloadedMessageEnabled           bool        `json:"overloaded_message_enabled"`
	PagekeysHandled                    bool        `json:"pagekeys_handled"`
	PostsFormattingGuide               bool        `json:"posts_formatting_guide"`
	PreferredSkinTone                  string      `json:"preferred_skin_tone"`
	PrevNextBtn                        bool        `json:"prev_next_btn"`
	PrivacyPolicySeen                  bool        `json:"privacy_policy_seen"`
	PromptedForEmailDisabling          bool        `json:"prompted_for_email_disabling"`
	PushAtChannelSuppressedChannels    string      `json:"push_at_channel_suppressed_channels"`
	PushDmAlert                        bool        `json:"push_dm_alert"`
	PushEverything                     bool        `json:"push_everything"`
	PushIdleWait                       int         `json:"push_idle_wait"`
	PushLoudChannels                   string      `json:"push_loud_channels"`
	PushLoudChannelsSet                string      `json:"push_loud_channels_set"`
	PushMentionAlert                   bool        `json:"push_mention_alert"`
	PushMentionChannels                string      `json:"push_mention_channels"`
	PushShowPreview                    bool        `json:"push_show_preview"`
	PushSound                          string      `json:"push_sound"`
	RequireAt                          bool        `json:"require_at"`
	SearchExcludeBots                  bool        `json:"search_exclude_bots"`
	SearchExcludeChannels              string      `json:"search_exclude_channels"`
	SearchOnlyCurrentTeam              bool        `json:"search_only_current_team"`
	SearchOnlyMyChannels               bool        `json:"search_only_my_channels"`
	SearchSort                         string      `json:"search_sort"`
	SeenCallsSsMainCoachmark           bool        `json:"seen_calls_ss_main_coachmark"`
	SeenCallsSsWindowCoachmark         bool        `json:"seen_calls_ss_window_coachmark"`
	SeenCallsVideoBetaCoachmark        bool        `json:"seen_calls_video_beta_coachmark"`
	SeenCallsVideoGaCoachmark          bool        `json:"seen_calls_video_ga_coachmark"`
	SeenCustomStatusBadge              bool        `json:"seen_custom_status_badge"`
	SeenCustomStatusCallout            bool        `json:"seen_custom_status_callout"`
	SeenDomainInviteReminder           bool        `json:"seen_domain_invite_reminder"`
	SeenGdriveCoachmark                bool        `json:"seen_gdrive_coachmark"`
	SeenGuestAdminSlackbotAnnouncement bool        `json:"seen_guest_admin_slackbot_announcement"`
	SeenHighlightsArrowsCoachmark      bool        `json:"seen_highlights_arrows_coachmark"`
	SeenHighlightsCoachmark            bool        `json:"seen_highlights_coachmark"`
	SeenIntlChannelNamesCoachmark      bool        `json:"seen_intl_channel_names_coachmark"`
	SeenMemberInviteReminder           bool        `json:"seen_member_invite_reminder"`
	SeenOnboardingChannels             bool        `json:"seen_onboarding_channels"`
	SeenOnboardingDirectMessages       bool        `json:"seen_onboarding_direct_messages"`
	SeenOnboardingInvites              bool        `json:"seen_onboarding_invites"`
	SeenOnboardingPrivateGroups        bool        `json:"seen_onboarding_private_groups"`
	SeenOnboardingRecentMentions       bool        `json:"seen_onboarding_recent_mentions"`
	SeenOnboardingSearch               bool        `json:"seen_onboarding_search"`
	SeenOnboardingSlackbotConversation bool        `json:"seen_onboarding_slackbot_conversation"`
	SeenOnboardingStarredItems         bool        `json:"seen_onboarding_starred_items"`
	SeenOnboardingStart                bool        `json:"seen_onboarding_start"`
	SeenRepliesCoachmark               bool        `json:"seen_replies_coachmark"`
	SeenSingleEmojiMsg                 bool        `json:"seen_single_emoji_msg"`
	SeenSsbPrompt                      bool        `json:"seen_ssb_prompt"`
	SeenThreadsNotificationBanner      bool        `json:"seen_threads_notification_banner"`
	SeenUnreadViewCoachmark            bool        `json:"seen_unread_view_coachmark"`
	SeenWelcome2                       bool        `json:"seen_welcome_2"`
	SeparatePrivateChannels            bool        `json:"separate_private_channels"`
	SeparateSharedChannels             bool        `json:"separate_shared_channels"`
	ShowAllSkinTones                   bool        `json:"show_all_skin_tones"`
	ShowJumperScores                   bool        `json:"show_jumper_scores"`
	ShowMemoryInstrument               bool        `json:"show_memory_instrument"`
	ShowTyping                         bool        `json:"show_typing"`
	SidebarBehavior                    string      `json:"sidebar_behavior"`
	SidebarTheme                       string      `json:"sidebar_theme"`
	SidebarThemeCustomValues           string      `json:"sidebar_theme_custom_values"`
	SnippetEditorWrapLongLines         bool        `json:"snippet_editor_wrap_long_lines"`
	SpacesNewXpBannerDismissed         bool        `json:"spaces_new_xp_banner_dismissed"`
	SsbSpaceWindow                     string      `json:"ssb_space_window"`
	SsEmojis                           bool        `json:"ss_emojis"`
	StartScrollAtOldest                bool        `json:"start_scroll_at_oldest"`
	TabUIReturnSelects                 bool        `json:"tab_ui_return_selects"`
	ThreadsEverything                  bool        `json:"threads_everything"`
	Time24                             bool        `json:"time24"`
	TwoFactorAuthEnabled               bool        `json:"two_factor_auth_enabled"`
	TwoFactorBackupType                string      `json:"two_factor_backup_type"`
	TwoFactorType                      string      `json:"two_factor_type"`
	Tz                                 string      `json:"tz"`
	UserColors                         string      `json:"user_colors"`
	WebappSpellcheck                   bool        `json:"webapp_spellcheck"`
	WelcomeMessageHidden               bool        `json:"welcome_message_hidden"`
	WhatsNewRead                       int         `json:"whats_new_read"`
	WinssbRunFromTray                  bool        `json:"winssb_run_from_tray"`
	WinssbWindowFlashBehavior          string      `json:"winssb_window_flash_behavior"`
}

func (s *SlackAPI) UsersCounts() ResponseUsersCounts {
	var response ResponseUsersCounts
	s.GetRequest(&response, "users.counts", "token")
	return response
}

func (s *SlackAPI) UsersDeletePhoto() Response {
	var response Response
	s.GetRequest(&response, "users.deletePhoto", "token")
	return response
}

func (s *SlackAPI) UsersGetPresence(query string) ResponseUsersGetPresence {
	var response ResponseUsersGetPresence
	s.GetRequest(&response, "users.getPresence", "token", "user="+query)
	return response
}

func (s *SlackAPI) UsersId(query string) string {
	response := s.UsersList()

	if response.Ok {
		for _, user := range response.Members {
			if user.Name == query {
				return user.Id
			}
		}
	}

	return query
}

func (s *SlackAPI) UsersInfo(query string) ResponseUsersInfo {
	query = s.UsersId(query)
	var response ResponseUsersInfo
	s.GetRequest(&response, "users.info", "token", "user="+query)
	return response
}

func (s *SlackAPI) UsersList() ResponseUsersList {
	if s.TeamUsers.Ok == true {
		return s.TeamUsers
	}

	var response ResponseUsersList
	s.GetRequest(&response, "users.list", "token", "presence=1")
	s.TeamUsers = response

	return response
}

func (s *SlackAPI) UsersPrefsSet(name string, value string) ResponseUserPrefs {
	var response ResponseUserPrefs
	s.GetRequest(&response,
		"users.prefs.set",
		"token",
		"name="+name,
		"value="+value)
	return response
}

func (s *SlackAPI) UsersPreparePhoto(image string) ResponseUserPhotoUpload {
	var response ResponseUserPhotoUpload
	s.PostRequest(&response,
		"users.preparePhoto",
		"token",
		"image=@"+image)
	return response
}

func (s *SlackAPI) UsersProfileGet(user string) ResponseUserIdentity {
	var response ResponseUserIdentity
	s.GetRequest(&response,
		"users.profile.get",
		"token",
		"user="+s.UsersId(user),
		"include_labels=1")
	return response
}

func (s *SlackAPI) UsersProfileSet(user string, name string, value string) ResponseUserIdentity {
	var response ResponseUserIdentity
	s.GetRequest(&response,
		"users.profile.set",
		"token",
		"user="+s.UsersId(user),
		"name="+name,
		"value="+value)
	return response
}

func (s *SlackAPI) UsersSearch(query string) []UserData {
	var matches []UserData
	response := s.UsersList()

	if response.Ok {
		for _, user := range response.Members {
			if strings.Contains(user.Name, query) ||
				strings.Contains(user.RealName, query) ||
				strings.Contains(user.Profile.Email, query) {
				matches = append(matches, user)
			}
		}
	}

	return matches
}

func (s *SlackAPI) UsersSetActive() Response {
	var response Response
	s.GetRequest(&response, "users.setActive", "token")
	return response
}

func (s *SlackAPI) UsersSetAvatar(image string) ResponseUserAvatar {
	var response ResponseUserAvatar
	upload := s.UsersPreparePhoto(image)
	result := s.UsersSetPhoto(upload.ID)
	response.UploadID = upload.ID
	response.UploadURL = upload.URL
	response.Profile = result.Profile
	return response
}

func (s *SlackAPI) UsersSetPhoto(imageid string) ResponseUserPhoto {
	var response ResponseUserPhoto
	s.GetRequest(&response,
		"users.setPhoto",
		"token",
		"crop_x=0",
		"crop_y=0",
		"crop_w=1000",
		"id="+imageid)
	return response
}

func (s *SlackAPI) UsersSetPresence(value string) Response {
	var response Response
	s.GetRequest(&response, "users.setPresence", "token", "presence="+value)
	return response
}
