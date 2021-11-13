package slackapi

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// ResponseUsersInfo defines the JSON-encoded output for UsersInfo.
type ResponseUsersInfo struct {
	Response
	User User `json:"user"`
}

// ResponseUsersGetPresence defines the JSON-encoded output for UsersGetPresence.
type ResponseUsersGetPresence struct {
	Response
	UserPresence
}

// ResponseUsersList defines the JSON-encoded output for UsersList.
type ResponseUsersList struct {
	Response
	Members          []User           `json:"members"`
	CacheTS          int              `json:"cache_ts"`
	ResponseMetadata ResponseMetadata `json:"response_metadata"`
}

// ResponseUserPrefs defines the JSON-encoded output for UserPrefs.
type ResponseUserPrefs struct {
	Response
	Prefs UserPrefs `json:"prefs"`
}

// ResponseUserIdentity defines the JSON-encoded output for UserIdentity.
type ResponseUserIdentity struct {
	Response
	Profile UserProfile `json:"profile"`
}

// ResponseUserAvatar defines the JSON-encoded output for UserAvatar.
type ResponseUserAvatar struct {
	Response
	UploadID  string                   `json:"upload_id"`
	UploadURL string                   `json:"upload_url"`
	Profile   ResponseUserPhotoProfile `json:"profile"`
}

// ResponseUsersIdentity defines the JSON-encoded output for UsersIdentity.
type ResponseUsersIdentity struct {
	Response
	Team Team          `json:"team"`
	User UsersIdentity `json:"user"`
}

// ResponseUserPhoto defines the JSON-encoded output for UserPhoto.
type ResponseUserPhoto struct {
	Response
	Profile ResponseUserPhotoProfile `json:"profile"`
}

// ResponseUserPhotoProfile defines the JSON-encoded output for UserPhotoProfile.
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

// ResponseUserPhotoUpload defines the JSON-encoded output for UserPhotoUpload.
type ResponseUserPhotoUpload struct {
	Response
	ID  string `json:"id"`
	URL string `json:"url"`
}

// ResponseUsersCounts defines the JSON-encoded output for UsersCounts.
type ResponseUsersCounts struct {
	Response
	Channels        []CountChannel        `json:"channels"`
	Groups          []CountGroup          `json:"groups"`
	InstantMessages []CountInstantMessage `json:"ims"`
}

// UserPresence defines the expected data from the JSON-encoded API response.
type UserPresence struct {
	Presence        string `json:"presence"`
	ConnectionCount int    `json:"connection_count"`
	LastActivity    int    `json:"last_activity"`
	AutoAway        bool   `json:"auto_away"`
	ManualAway      bool   `json:"manual_away"`
	Online          bool   `json:"online"`
}

// User defines the expected data from the JSON-encoded API response.
type User struct {
	Color             string      `json:"color"`
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Presence          string      `json:"presence"`
	Profile           UserProfile `json:"profile"`
	RealName          string      `json:"real_name"`
	Status            string      `json:"status"`
	TeamID            string      `json:"team_id"`
	TwoFactorType     string      `json:"two_factor_type"`
	Tz                string      `json:"tz"`
	TzLabel           string      `json:"tz_label"`
	TzOffset          int         `json:"tz_offset"`
	Deleted           bool        `json:"deleted"`
	Has2fa            bool        `json:"has_2fa"`
	IsAdmin           bool        `json:"is_admin"`
	IsBot             bool        `json:"is_bot"`
	IsOwner           bool        `json:"is_owner"`
	IsPrimaryOwner    bool        `json:"is_primary_owner"`
	IsRestricted      bool        `json:"is_restricted"`
	IsUltraRestricted bool        `json:"is_ultra_restricted"`
}

// UsersIdentity defines the expected data from the JSON-encoded API response.
type UsersIdentity struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Image24  string `json:"image_24"`
	Image32  string `json:"image_32"`
	Image48  string `json:"image_48"`
	Image72  string `json:"image_72"`
	Image192 string `json:"image_192"`
}

// UserProfile defines the expected data from the JSON-encoded API response.
type UserProfile struct {
	APIAppID           string      `json:"api_app_id"`
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

// CountChannel defines the expected data from the JSON-encoded API response.
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

// CountGroup defines the expected data from the JSON-encoded API response.
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

// CountInstantMessage defines the expected data from the JSON-encoded API response.
type CountInstantMessage struct {
	DmCount   int    `json:"dm_count"`
	ID        string `json:"id"`
	IsOpen    bool   `json:"is_open"`
	IsStarred bool   `json:"is_starred"`
	Name      string `json:"name"`
	UserID    string `json:"user_id"`
}

// UserPrefs defines the expected data from the JSON-encoded API response.
type UserPrefs struct {
	AllNotificationsPrefs              interface{} `json:"all_notifications_prefs"`
	A11yFontSize                       string      `json:"a11y_font_size"`
	AtChannelSuppressedChannels        string      `json:"at_channel_suppressed_channels"`
	ChannelSort                        string      `json:"channel_sort"`
	ClientLogsPri                      string      `json:"client_logs_pri"`
	DndEndHour                         string      `json:"dnd_end_hour"`
	DndStartHour                       string      `json:"dnd_start_hour"`
	EmailAlerts                        string      `json:"email_alerts"`
	EmojiMode                          string      `json:"emoji_mode"`
	EmojiUse                           string      `json:"emoji_use"`
	FlannelServerPool                  string      `json:"flannel_server_pool"`
	FrecencyEntJumper                  string      `json:"frecency_ent_jumper"`
	FrecencyJumper                     string      `json:"frecency_jumper"`
	HighlightWords                     string      `json:"highlight_words"`
	LastSnippetType                    string      `json:"last_snippet_type"`
	LastTosAcknowledged                string      `json:"last_tos_acknowledged"`
	Locale                             string      `json:"locale"`
	LoudChannels                       string      `json:"loud_channels"`
	LoudChannelsSet                    string      `json:"loud_channels_set"`
	MacSsbBounce                       string      `json:"mac_ssb_bounce"`
	MessagesTheme                      string      `json:"messages_theme"`
	MutedChannels                      string      `json:"muted_channels"`
	NeverChannels                      string      `json:"never_channels"`
	NewMsgSnd                          string      `json:"new_msg_snd"`
	NewxpSeenLastMessage               string      `json:"newxp_seen_last_message"`
	PreferredSkinTone                  string      `json:"preferred_skin_tone"`
	PushAtChannelSuppressedChannels    string      `json:"push_at_channel_suppressed_channels"`
	PushLoudChannels                   string      `json:"push_loud_channels"`
	PushLoudChannelsSet                string      `json:"push_loud_channels_set"`
	PushMentionChannels                string      `json:"push_mention_channels"`
	PushSound                          string      `json:"push_sound"`
	SearchExcludeChannels              string      `json:"search_exclude_channels"`
	SearchSort                         string      `json:"search_sort"`
	SidebarBehavior                    string      `json:"sidebar_behavior"`
	SidebarTheme                       string      `json:"sidebar_theme"`
	SidebarThemeCustomValues           string      `json:"sidebar_theme_custom_values"`
	SsbSpaceWindow                     string      `json:"ssb_space_window"`
	TwoFactorBackupType                string      `json:"two_factor_backup_type"`
	TwoFactorType                      string      `json:"two_factor_type"`
	Tz                                 string      `json:"tz"`
	UserColors                         string      `json:"user_colors"`
	WinssbWindowFlashBehavior          string      `json:"winssb_window_flash_behavior"`
	DisplayRealNamesOverride           int         `json:"display_real_names_override"`
	EmailAlertsSleepUntil              int         `json:"email_alerts_sleep_until"`
	GrowthMsgLimitApproachingCtaCount  int         `json:"growth_msg_limit_approaching_cta_count"`
	GrowthMsgLimitApproachingCtaTs     int         `json:"growth_msg_limit_approaching_cta_ts"`
	GrowthMsgLimitLongReachedCtaCount  int         `json:"growth_msg_limit_long_reached_cta_count"`
	GrowthMsgLimitLongReachedCtaLastTs int         `json:"growth_msg_limit_long_reached_cta_last_ts"`
	GrowthMsgLimitReachedCtaCount      int         `json:"growth_msg_limit_reached_cta_count"`
	GrowthMsgLimitReachedCtaLastTs     int         `json:"growth_msg_limit_reached_cta_last_ts"`
	KKeyOmniboxAutoHideCount           int         `json:"k_key_omnibox_auto_hide_count"`
	LastSeenAtChannelWarning           int         `json:"last_seen_at_channel_warning"`
	OnboardingSlackbotConversationStep int         `json:"onboarding_slackbot_conversation_step"`
	PushIDleWait                       int         `json:"push_idle_wait"`
	WhatsNewRead                       int         `json:"whats_new_read"`
	A11yAnimations                     bool        `json:"a11y_animations"`
	AllChannelsLoud                    bool        `json:"all_channels_loud"`
	AllowCallsToSetCurrentStatus       bool        `json:"allow_calls_to_set_current_status"`
	AllUnreadsSortOrder                bool        `json:"all_unreads_sort_order"`
	ArrowHistory                       bool        `json:"arrow_history"`
	BoxEnabled                         bool        `json:"box_enabled"`
	ColorNamesInList                   bool        `json:"color_names_in_list"`
	ConfirmClearAllUnreads             bool        `json:"confirm_clear_all_unreads"`
	ConfirmShCallStart                 bool        `json:"confirm_sh_call_start"`
	ConfirmUserMarkedAway              bool        `json:"confirm_user_marked_away"`
	ConvertEmoticons                   bool        `json:"convert_emoticons"`
	DisplayDisplayNames                bool        `json:"display_display_names"`
	DndEnabled                         bool        `json:"dnd_enabled"`
	DropboxEnabled                     bool        `json:"dropbox_enabled"`
	EmailMisc                          bool        `json:"email_misc"`
	EmailWeekly                        bool        `json:"email_weekly"`
	EmojiAutocompleteBig               bool        `json:"emoji_autocomplete_big"`
	EnableReactEmojiPicker             bool        `json:"enable_react_emoji_picker"`
	EnableUnreadView                   bool        `json:"enable_unread_view"`
	EnhancedDebugging                  bool        `json:"enhanced_debugging"`
	EnterIsSpecialInTbt                bool        `json:"enter_is_special_in_tbt"`
	EnterpriseMigrationSeen            bool        `json:"enterprise_migration_seen"`
	ExpandInlineImages                 bool        `json:"expand_inline_imgs"`
	ExpandInternalInlineImages         bool        `json:"expand_internal_inline_imgs"`
	ExpandNonMediaAttachments          bool        `json:"expand_non_media_attachments"`
	ExpandSnippets                     bool        `json:"expand_snippets"`
	FKeySearch                         bool        `json:"f_key_search"`
	FullerTimestamps                   bool        `json:"fuller_timestamps"`
	FullTextExtracts                   bool        `json:"full_text_extracts"`
	GdriveAuthed                       bool        `json:"gdrive_authed"`
	GdriveEnabled                      bool        `json:"gdrive_enabled"`
	GraphicEmoticons                   bool        `json:"graphic_emoticons"`
	GrowlsEnabled                      bool        `json:"growls_enabled"`
	HasCreatedChannel                  bool        `json:"has_created_channel"`
	HasInvited                         bool        `json:"has_invited"`
	HasSearched                        bool        `json:"has_searched"`
	HasUploaded                        bool        `json:"has_uploaded"`
	HideHexSwatch                      bool        `json:"hide_hex_swatch"`
	HideUserGroupInfoPane              bool        `json:"hide_user_group_info_pane"`
	IntroToAppsMessageSeen             bool        `json:"intro_to_apps_message_seen"`
	Jumbomoji                          bool        `json:"jumbomoji"`
	KKeyOmnibox                        bool        `json:"k_key_omnibox"`
	LoadLato2                          bool        `json:"load_lato_2"`
	LsDisabled                         bool        `json:"ls_disabled"`
	MacSsbBullet                       bool        `json:"mac_ssb_bullet"`
	MarkMsgsReadImmediately            bool        `json:"mark_msgs_read_immediately"`
	MeasureCSSUsage                    bool        `json:"measure_css_usage"`
	MentionsExcludeAtChannels          bool        `json:"mentions_exclude_at_channels"`
	MentionsExcludeAtUserGroups        bool        `json:"mentions_exclude_at_user_groups"`
	MsgPreview                         bool        `json:"msg_preview"`
	MsgPreviewPersistent               bool        `json:"msg_preview_persistent"`
	MuteSounds                         bool        `json:"mute_sounds"`
	NoCreatedOverlays                  bool        `json:"no_created_overlays"`
	NoInvitesWidgetInSidebar           bool        `json:"no_invites_widget_in_sidebar"`
	NoJoinedOverlays                   bool        `json:"no_joined_overlays"`
	NoMacelectronBanner                bool        `json:"no_macelectron_banner"`
	NoMacssb1Banner                    bool        `json:"no_macssb1_banner"`
	NoMacssb2Banner                    bool        `json:"no_macssb2_banner"`
	NoOmniboxInChannels                bool        `json:"no_omnibox_in_channels"`
	NoTextInNotifications              bool        `json:"no_text_in_notifications"`
	NoWinssb1Banner                    bool        `json:"no_winssb1_banner"`
	ObeyInlineImageLimit               bool        `json:"obey_inline_img_limit"`
	OnboardingCancelled                bool        `json:"onboarding_cancelled"`
	OverloadedMessageEnabled           bool        `json:"overloaded_message_enabled"`
	PagekeysHandled                    bool        `json:"pagekeys_handled"`
	PostsFormattingGuide               bool        `json:"posts_formatting_guide"`
	PrevNextBtn                        bool        `json:"prev_next_btn"`
	PrivacyPolicySeen                  bool        `json:"privacy_policy_seen"`
	PromptedForEmailDisabling          bool        `json:"prompted_for_email_disabling"`
	PushDmAlert                        bool        `json:"push_dm_alert"`
	PushEverything                     bool        `json:"push_everything"`
	PushMentionAlert                   bool        `json:"push_mention_alert"`
	PushShowPreview                    bool        `json:"push_show_preview"`
	RequireAt                          bool        `json:"require_at"`
	SearchExcludeBots                  bool        `json:"search_exclude_bots"`
	SearchOnlyCurrentTeam              bool        `json:"search_only_current_team"`
	SearchOnlyMyChannels               bool        `json:"search_only_my_channels"`
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
	SnippetEditorWrapLongLines         bool        `json:"snippet_editor_wrap_long_lines"`
	SpacesNewXpBannerDismissed         bool        `json:"spaces_new_xp_banner_dismissed"`
	SsEmojis                           bool        `json:"ss_emojis"`
	StartScrollAtOldest                bool        `json:"start_scroll_at_oldest"`
	TabUIReturnSelects                 bool        `json:"tab_ui_return_selects"`
	ThreadsEverything                  bool        `json:"threads_everything"`
	Time24                             bool        `json:"time24"`
	TwoFactorAuthEnabled               bool        `json:"two_factor_auth_enabled"`
	WebappSpellcheck                   bool        `json:"webapp_spellcheck"`
	WelcomeMessageHidden               bool        `json:"welcome_message_hidden"`
	WinssbRunFromTray                  bool        `json:"winssb_run_from_tray"`
}

// UsersCounts count number of users in the team.
func (s *SlackAPI) UsersCounts() ResponseUsersCounts {
	var response ResponseUsersCounts
	s.getRequest(&response, "users.counts", nil)
	return response
}

// UsersDeletePhoto delete the user avatar.
func (s *SlackAPI) UsersDeletePhoto() Response {
	var response Response
	s.getRequest(&response, "users.deletePhoto", nil)
	return response
}

// UsersGetPresence gets user presence information.
func (s *SlackAPI) UsersGetPresence(query string) ResponseUsersGetPresence {
	var response ResponseUsersGetPresence
	s.getRequest(&response, "users.getPresence", struct {
		User string `json:"user"`
	}{query})
	return response
}

// UsersID gets user identifier from username.
func (s *SlackAPI) UsersID(query string, limit int) string {
	response := s.UsersList(limit, "")

	if response.Ok {
		/* allow user references: @username */
		if len(query) > 1 && query[0] == '@' {
			query = query[1:]
		}

		for _, user := range response.Members {
			if user.Name == query {
				return user.ID
			}
		}
	}

	return query
}

// UsersIdentity get a user's identity.
func (s *SlackAPI) UsersIdentity() ResponseUsersIdentity {
	var response ResponseUsersIdentity
	s.getRequest(&response, "users.identity", nil)
	return response
}

// UsersInfo gets information about a user.
func (s *SlackAPI) UsersInfo(query string) ResponseUsersInfo {
	var response ResponseUsersInfo
	s.getRequest(&response, "users.info", struct {
		User string `json:"user"`
	}{query})
	return response
}

// UsersList lists all users in a Slack team.
func (s *SlackAPI) UsersList(limit int, cursor string) ResponseUsersList {
	if s.teamUsers.Ok {
		return s.teamUsers
	}

	var response ResponseUsersList
	s.getRequest(&response, "users.list", struct {
		Limit  int    `json:"limit"`
		Cursor string `json:"cursor"`
	}{
		Limit:  limit,
		Cursor: cursor,
	})
	s.teamUsers = response

	return response
}

// UsersListWithPresence lists all users in a Slack team.
func (s *SlackAPI) UsersListWithPresence() ResponseUsersList {
	if s.teamUsers.Ok {
		return s.teamUsers
	}

	var response ResponseUsersList
	s.getRequest(&response, "users.list", struct {
		Presence bool `json:"presence"`
	}{true})
	s.teamUsers = response

	return response
}

// UsersLookupByEmail find a user with an email address.
func (s *SlackAPI) UsersLookupByEmail(email string) User {
	var response User
	s.getRequest(&response, "users.lookupByEmail", struct {
		Email string `json:"email"`
	}{email})
	return response
}

// UsersPrefsGet get user account preferences.
func (s *SlackAPI) UsersPrefsGet() ResponseUserPrefs {
	var response ResponseUserPrefs
	s.getRequest(&response, "users.prefs.get", nil)
	return response
}

// UsersPrefsSet set user account preferences.
func (s *SlackAPI) UsersPrefsSet(name string, value string) ResponseUserPrefs {
	var response ResponseUserPrefs
	s.postRequest(&response, "users.prefs.set", struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}{name, value})
	return response
}

// UsersPreparePhoto upload a picture to use as the avatar.
func (s *SlackAPI) UsersPreparePhoto(image string) ResponseUserPhotoUpload {
	var response ResponseUserPhotoUpload
	s.postRequest(&response, "users.preparePhoto", struct {
		Image string `json:"image"`
	}{"@" + image})
	return response
}

// UsersProfileGet retrieves a user's profile information.
func (s *SlackAPI) UsersProfileGet(query string) ResponseUserIdentity {
	var response ResponseUserIdentity
	s.getRequest(&response, "users.profile.get", struct {
		User          string `json:"user"`
		IncludeLabels bool   `json:"include_labels"`
	}{query, false})
	return response
}

// UsersProfileGetWithLabels retrieves a user's profile information.
func (s *SlackAPI) UsersProfileGetWithLabels(query string) ResponseUserIdentity {
	var response ResponseUserIdentity
	s.getRequest(&response, "users.profile.get", struct {
		User          string `json:"user"`
		IncludeLabels bool   `json:"include_labels"`
	}{query, true})
	return response
}

// UsersProfileSet set the profile information for a user.
func (s *SlackAPI) UsersProfileSet(name string, value string) ResponseUserIdentity {
	var response ResponseUserIdentity
	s.postRequest(&response, "users.profile.set", struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}{name, value})
	return response
}

// UsersProfileSetMultiple set the profile information for a user.
func (s *SlackAPI) UsersProfileSetMultiple(profile string) ResponseUserIdentity {
	var response ResponseUserIdentity
	s.postRequest(&response, "users.profile.set", struct {
		Profile string `json:"profile"`
	}{profile})
	return response
}

// UsersSetActive marks a user as active.
func (s *SlackAPI) UsersSetActive() Response {
	var response Response
	s.getRequest(&response, "users.setActive", nil)
	return response
}

// UsersAdminSetInactive deactivates an existing user account.
func (s *SlackAPI) UsersAdminSetInactive(user string) Response {
	in := url.Values{"user": {user}}
	var out Response
	if err := s.baseFormPOST("/api/users.admin.setInactive", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

// UsersAdminSetRegular activates an account as a regular user.
func (s *SlackAPI) UsersAdminSetRegular(user string) Response {
	in := url.Values{"user": {user}}
	var out Response
	if err := s.baseFormPOST("/api/users.admin.setRegular", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

// UsersSetAvatar upload a picture and set it as the avatar.
func (s *SlackAPI) UsersSetAvatar(image string) ResponseUserAvatar {
	var response ResponseUserAvatar
	upload := s.UsersPreparePhoto(image)
	result := s.UsersSetPhoto(upload.ID)
	response.UploadID = upload.ID
	response.UploadURL = upload.URL
	response.Profile = result.Profile
	return response
}

// UsersSetPhoto define which picture will be the avatar.
func (s *SlackAPI) UsersSetPhoto(imageid string) ResponseUserPhoto {
	var response ResponseUserPhoto
	s.postRequest(&response, "users.setPhoto", struct {
		CropX int    `json:"crop_x"`
		CropY int    `json:"crop_y"`
		CropW int    `json:"crop_w"`
		ID    string `json:"id"`
	}{0, 0, 1024, imageid})
	return response
}

// UsersSetPresence manually sets user presence.
func (s *SlackAPI) UsersSetPresence(value string) Response {
	var response Response
	s.postRequest(&response, "users.setPresence", struct {
		Presence string `json:"presence"`
	}{value})
	return response
}

// UsersSetStatus set the status message and emoji.
func (s *SlackAPI) UsersSetStatus(emoji string, text string) ResponseUserIdentity {
	profile, err := json.Marshal(map[string]string{
		"status_emoji": emoji,
		"status_text":  text,
	})

	if err != nil {
		return ResponseUserIdentity{}
	}

	return s.UsersProfileSetMultiple(string(profile))
}

type Invitation struct {
	Email string `json:"email"`
	Type  string `json:"type"`
	Mode  string `json:"mode"`
}

type InviteBulkInput struct {
	Invites                    []Invitation `json:"invites"`
	Source                     string       `json:"source"`
	Campaign                   string       `json:"campaign"`
	Mode                       string       `json:"mode"`
	Restricted                 bool         `json:"restricted"`
	UltraRestricted            bool         `json:"ultra_restricted"`
	EmailPasswordPolicyEnabled bool         `json:"email_password_policy_enabled"`
}

type UsersAdminInviteBulkResponse struct {
	Response
	Invites []InvitationResponse `json:"invites"`
}

type InvitationResponse struct {
	Response
	Email string `json:"email"`
}

// UsersAdminInviteBulk is https://cixtor.slack.com/admin/invites
func (s *SlackAPI) UsersAdminInviteBulk(input InviteBulkInput) UsersAdminInviteBulkResponse {
	var out UsersAdminInviteBulkResponse
	if err := s.baseJSONPOST("/api/users.admin.inviteBulk", input, &out); err != nil {
		return UsersAdminInviteBulkResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

type InvitesHistoryInput struct {
	// Type must be "pending" or "accepted".
	Type string
	// SortBy is any InvitationHistory property name, e.g. date_create
	SortBy string
	// SortDir must be "ASC" or "DESC"
	SortDir string
}

type Inviter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type InvitationHistory struct {
	ID         int      `json:"id"`
	Email      string   `json:"email"`
	DateCreate int      `json:"date_create"`
	DateResent int      `json:"date_resent"`
	Bouncing   bool     `json:"bouncing"`
	Channels   []string `json:"channels"`
	Inviter    Inviter  `json:"inviter"`
	InviteType string   `json:"invite_type"`
}

type InvitesHistoryResponse struct {
	Response
	Invites []InvitationHistory `json:"invites"`
}

// UsersAdminFetchInvitesHistory is https://cixtor.slack.com/admin/invites
func (s *SlackAPI) UsersAdminFetchInvitesHistory(input InvitesHistoryInput) InvitesHistoryResponse {
	in := url.Values{}
	if input.Type == "accepted" {
		in.Add("type", "accepted")
	} else {
		in.Add("type", "pending")
	}
	if input.SortBy == "" {
		in.Add("sort_by", "date_create")
	} else {
		in.Add("sort_by", input.SortBy)
	}
	if input.SortDir == "ASC" {
		in.Add("sort_dir", "ASC")
	} else {
		in.Add("sort_dir", "DESC")
	}
	var out InvitesHistoryResponse
	if err := s.baseFormPOST("/api/users.admin.fetchInvitesHistory", in, &out); err != nil {
		return InvitesHistoryResponse{Response: Response{Error: err.Error()}}
	}
	return out
}

// UsersAdminResendInvitation is https://cixtor.slack.com/admin/invites
func (s *SlackAPI) UsersAdminResendInvitation(inviteID int) Response {
	in := url.Values{"invite_id": {fmt.Sprintf("%d", inviteID)}}
	var out Response
	if err := s.baseFormPOST("/api/users.admin.resendInvitation", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

// UsersAdminRevokeInvitation is https://cixtor.slack.com/admin/invites
func (s *SlackAPI) UsersAdminRevokeInvitation(inviteID int) Response {
	in := url.Values{"invite_id": {fmt.Sprintf("%d", inviteID)}}
	var out Response
	if err := s.baseFormPOST("/api/users.admin.revokeInvitation", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

// AdminUsersSessionClearSettings is https://api.slack.com/methods/admin.users.session.clearSettings
func (s *SlackAPI) AdminUsersSessionClearSettings(users []string) Response {
	b, _ := json.Marshal(users)
	in := url.Values{"users": {string(b)}}
	var out Response
	if err := s.baseFormPOST("/api/admin.users.session.clearSettings", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}

// AdminUsersSessionGetSettings is https://api.slack.com/methods/admin.users.session.getSettings
func (s *SlackAPI) AdminUsersSessionGetSettings(users []string) Response {
	b, _ := json.Marshal(users)
	in := url.Values{"users": {string(b)}}
	var out Response
	if err := s.baseFormPOST("/api/admin.users.session.getSettings", in, &out); err != nil {
		return Response{Error: err.Error()}
	}
	return out
}
