// Auto-generated by avdl-compiler v1.3.24 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/chat_ui.avdl

package chat1

import (
	"errors"
	gregor1 "github.com/keybase/client/go/protocol/gregor1"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type UIPagination struct {
	Next     string `codec:"next" json:"next"`
	Previous string `codec:"previous" json:"previous"`
	Num      int    `codec:"num" json:"num"`
	Last     bool   `codec:"last" json:"last"`
}

func (o UIPagination) DeepCopy() UIPagination {
	return UIPagination{
		Next:     o.Next,
		Previous: o.Previous,
		Num:      o.Num,
		Last:     o.Last,
	}
}

type UnverifiedInboxUIItemMetadata struct {
	ChannelName       string   `codec:"channelName" json:"channelName"`
	Headline          string   `codec:"headline" json:"headline"`
	Snippet           string   `codec:"snippet" json:"snippet"`
	SnippetDecoration string   `codec:"snippetDecoration" json:"snippetDecoration"`
	WriterNames       []string `codec:"writerNames" json:"writerNames"`
	ResetParticipants []string `codec:"resetParticipants" json:"resetParticipants"`
}

func (o UnverifiedInboxUIItemMetadata) DeepCopy() UnverifiedInboxUIItemMetadata {
	return UnverifiedInboxUIItemMetadata{
		ChannelName:       o.ChannelName,
		Headline:          o.Headline,
		Snippet:           o.Snippet,
		SnippetDecoration: o.SnippetDecoration,
		WriterNames: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.WriterNames),
		ResetParticipants: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.ResetParticipants),
	}
}

type UnverifiedInboxUIItem struct {
	ConvID        string                         `codec:"convID" json:"convID"`
	TopicType     TopicType                      `codec:"topicType" json:"topicType"`
	Name          string                         `codec:"name" json:"name"`
	Visibility    keybase1.TLFVisibility         `codec:"visibility" json:"visibility"`
	Status        ConversationStatus             `codec:"status" json:"status"`
	MembersType   ConversationMembersType        `codec:"membersType" json:"membersType"`
	MemberStatus  ConversationMemberStatus       `codec:"memberStatus" json:"memberStatus"`
	TeamType      TeamType                       `codec:"teamType" json:"teamType"`
	Notifications *ConversationNotificationInfo  `codec:"notifications,omitempty" json:"notifications,omitempty"`
	Time          gregor1.Time                   `codec:"time" json:"time"`
	Version       ConversationVers               `codec:"version" json:"version"`
	MaxMsgID      MessageID                      `codec:"maxMsgID" json:"maxMsgID"`
	ReadMsgID     MessageID                      `codec:"readMsgID" json:"readMsgID"`
	LocalMetadata *UnverifiedInboxUIItemMetadata `codec:"localMetadata,omitempty" json:"localMetadata,omitempty"`
	FinalizeInfo  *ConversationFinalizeInfo      `codec:"finalizeInfo,omitempty" json:"finalizeInfo,omitempty"`
	Supersedes    []ConversationMetadata         `codec:"supersedes" json:"supersedes"`
	SupersededBy  []ConversationMetadata         `codec:"supersededBy" json:"supersededBy"`
}

func (o UnverifiedInboxUIItem) DeepCopy() UnverifiedInboxUIItem {
	return UnverifiedInboxUIItem{
		ConvID:       o.ConvID,
		TopicType:    o.TopicType.DeepCopy(),
		Name:         o.Name,
		Visibility:   o.Visibility.DeepCopy(),
		Status:       o.Status.DeepCopy(),
		MembersType:  o.MembersType.DeepCopy(),
		MemberStatus: o.MemberStatus.DeepCopy(),
		TeamType:     o.TeamType.DeepCopy(),
		Notifications: (func(x *ConversationNotificationInfo) *ConversationNotificationInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Notifications),
		Time:      o.Time.DeepCopy(),
		Version:   o.Version.DeepCopy(),
		MaxMsgID:  o.MaxMsgID.DeepCopy(),
		ReadMsgID: o.ReadMsgID.DeepCopy(),
		LocalMetadata: (func(x *UnverifiedInboxUIItemMetadata) *UnverifiedInboxUIItemMetadata {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.LocalMetadata),
		FinalizeInfo: (func(x *ConversationFinalizeInfo) *ConversationFinalizeInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.FinalizeInfo),
		Supersedes: (func(x []ConversationMetadata) []ConversationMetadata {
			if x == nil {
				return nil
			}
			ret := make([]ConversationMetadata, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Supersedes),
		SupersededBy: (func(x []ConversationMetadata) []ConversationMetadata {
			if x == nil {
				return nil
			}
			ret := make([]ConversationMetadata, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.SupersededBy),
	}
}

type UnverifiedInboxUIItems struct {
	Items      []UnverifiedInboxUIItem `codec:"items" json:"items"`
	Pagination *UIPagination           `codec:"pagination,omitempty" json:"pagination,omitempty"`
	Offline    bool                    `codec:"offline" json:"offline"`
}

func (o UnverifiedInboxUIItems) DeepCopy() UnverifiedInboxUIItems {
	return UnverifiedInboxUIItems{
		Items: (func(x []UnverifiedInboxUIItem) []UnverifiedInboxUIItem {
			if x == nil {
				return nil
			}
			ret := make([]UnverifiedInboxUIItem, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Items),
		Pagination: (func(x *UIPagination) *UIPagination {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Pagination),
		Offline: o.Offline,
	}
}

type InboxUIItem struct {
	ConvID            string                        `codec:"convID" json:"convID"`
	TopicType         TopicType                     `codec:"topicType" json:"topicType"`
	IsEmpty           bool                          `codec:"isEmpty" json:"isEmpty"`
	Name              string                        `codec:"name" json:"name"`
	Snippet           string                        `codec:"snippet" json:"snippet"`
	SnippetDecoration string                        `codec:"snippetDecoration" json:"snippetDecoration"`
	Channel           string                        `codec:"channel" json:"channel"`
	Headline          string                        `codec:"headline" json:"headline"`
	Visibility        keybase1.TLFVisibility        `codec:"visibility" json:"visibility"`
	Participants      []string                      `codec:"participants" json:"participants"`
	FullNames         map[string]string             `codec:"fullNames" json:"fullNames"`
	ResetParticipants []string                      `codec:"resetParticipants" json:"resetParticipants"`
	Status            ConversationStatus            `codec:"status" json:"status"`
	MembersType       ConversationMembersType       `codec:"membersType" json:"membersType"`
	MemberStatus      ConversationMemberStatus      `codec:"memberStatus" json:"memberStatus"`
	TeamType          TeamType                      `codec:"teamType" json:"teamType"`
	Time              gregor1.Time                  `codec:"time" json:"time"`
	Notifications     *ConversationNotificationInfo `codec:"notifications,omitempty" json:"notifications,omitempty"`
	CreatorInfo       *ConversationCreatorInfoLocal `codec:"creatorInfo,omitempty" json:"creatorInfo,omitempty"`
	Version           ConversationVers              `codec:"version" json:"version"`
	MaxMsgID          MessageID                     `codec:"maxMsgID" json:"maxMsgID"`
	ReadMsgID         MessageID                     `codec:"readMsgID" json:"readMsgID"`
	ConvRetention     *RetentionPolicy              `codec:"convRetention,omitempty" json:"convRetention,omitempty"`
	TeamRetention     *RetentionPolicy              `codec:"teamRetention,omitempty" json:"teamRetention,omitempty"`
	ConvSettings      *ConversationSettingsLocal    `codec:"convSettings,omitempty" json:"convSettings,omitempty"`
	FinalizeInfo      *ConversationFinalizeInfo     `codec:"finalizeInfo,omitempty" json:"finalizeInfo,omitempty"`
	Supersedes        []ConversationMetadata        `codec:"supersedes" json:"supersedes"`
	SupersededBy      []ConversationMetadata        `codec:"supersededBy" json:"supersededBy"`
}

func (o InboxUIItem) DeepCopy() InboxUIItem {
	return InboxUIItem{
		ConvID:            o.ConvID,
		TopicType:         o.TopicType.DeepCopy(),
		IsEmpty:           o.IsEmpty,
		Name:              o.Name,
		Snippet:           o.Snippet,
		SnippetDecoration: o.SnippetDecoration,
		Channel:           o.Channel,
		Headline:          o.Headline,
		Visibility:        o.Visibility.DeepCopy(),
		Participants: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Participants),
		FullNames: (func(x map[string]string) map[string]string {
			if x == nil {
				return nil
			}
			ret := make(map[string]string, len(x))
			for k, v := range x {
				kCopy := k
				vCopy := v
				ret[kCopy] = vCopy
			}
			return ret
		})(o.FullNames),
		ResetParticipants: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.ResetParticipants),
		Status:       o.Status.DeepCopy(),
		MembersType:  o.MembersType.DeepCopy(),
		MemberStatus: o.MemberStatus.DeepCopy(),
		TeamType:     o.TeamType.DeepCopy(),
		Time:         o.Time.DeepCopy(),
		Notifications: (func(x *ConversationNotificationInfo) *ConversationNotificationInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Notifications),
		CreatorInfo: (func(x *ConversationCreatorInfoLocal) *ConversationCreatorInfoLocal {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.CreatorInfo),
		Version:   o.Version.DeepCopy(),
		MaxMsgID:  o.MaxMsgID.DeepCopy(),
		ReadMsgID: o.ReadMsgID.DeepCopy(),
		ConvRetention: (func(x *RetentionPolicy) *RetentionPolicy {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ConvRetention),
		TeamRetention: (func(x *RetentionPolicy) *RetentionPolicy {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.TeamRetention),
		ConvSettings: (func(x *ConversationSettingsLocal) *ConversationSettingsLocal {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ConvSettings),
		FinalizeInfo: (func(x *ConversationFinalizeInfo) *ConversationFinalizeInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.FinalizeInfo),
		Supersedes: (func(x []ConversationMetadata) []ConversationMetadata {
			if x == nil {
				return nil
			}
			ret := make([]ConversationMetadata, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Supersedes),
		SupersededBy: (func(x []ConversationMetadata) []ConversationMetadata {
			if x == nil {
				return nil
			}
			ret := make([]ConversationMetadata, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.SupersededBy),
	}
}

type InboxUIItemError struct {
	Typ               ConversationErrorType   `codec:"typ" json:"typ"`
	Message           string                  `codec:"message" json:"message"`
	UnverifiedTLFName string                  `codec:"unverifiedTLFName" json:"unverifiedTLFName"`
	RekeyInfo         *ConversationErrorRekey `codec:"rekeyInfo,omitempty" json:"rekeyInfo,omitempty"`
	RemoteConv        UnverifiedInboxUIItem   `codec:"remoteConv" json:"remoteConv"`
}

func (o InboxUIItemError) DeepCopy() InboxUIItemError {
	return InboxUIItemError{
		Typ:               o.Typ.DeepCopy(),
		Message:           o.Message,
		UnverifiedTLFName: o.UnverifiedTLFName,
		RekeyInfo: (func(x *ConversationErrorRekey) *ConversationErrorRekey {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.RekeyInfo),
		RemoteConv: o.RemoteConv.DeepCopy(),
	}
}

type InboxUIItems struct {
	Items      []InboxUIItem `codec:"items" json:"items"`
	Pagination *UIPagination `codec:"pagination,omitempty" json:"pagination,omitempty"`
	Offline    bool          `codec:"offline" json:"offline"`
}

func (o InboxUIItems) DeepCopy() InboxUIItems {
	return InboxUIItems{
		Items: (func(x []InboxUIItem) []InboxUIItem {
			if x == nil {
				return nil
			}
			ret := make([]InboxUIItem, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Items),
		Pagination: (func(x *UIPagination) *UIPagination {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Pagination),
		Offline: o.Offline,
	}
}

type UIChannelNameMention struct {
	Name   string `codec:"name" json:"name"`
	ConvID string `codec:"convID" json:"convID"`
}

func (o UIChannelNameMention) DeepCopy() UIChannelNameMention {
	return UIChannelNameMention{
		Name:   o.Name,
		ConvID: o.ConvID,
	}
}

type UIAssetUrlInfo struct {
	PreviewUrl          string  `codec:"previewUrl" json:"previewUrl"`
	FullUrl             string  `codec:"fullUrl" json:"fullUrl"`
	FullUrlCached       bool    `codec:"fullUrlCached" json:"fullUrlCached"`
	MimeType            string  `codec:"mimeType" json:"mimeType"`
	VideoDuration       *string `codec:"videoDuration,omitempty" json:"videoDuration,omitempty"`
	InlineVideoPlayable bool    `codec:"inlineVideoPlayable" json:"inlineVideoPlayable"`
}

func (o UIAssetUrlInfo) DeepCopy() UIAssetUrlInfo {
	return UIAssetUrlInfo{
		PreviewUrl:    o.PreviewUrl,
		FullUrl:       o.FullUrl,
		FullUrlCached: o.FullUrlCached,
		MimeType:      o.MimeType,
		VideoDuration: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.VideoDuration),
		InlineVideoPlayable: o.InlineVideoPlayable,
	}
}

type UIMessageValid struct {
	MessageID             MessageID              `codec:"messageID" json:"messageID"`
	Ctime                 gregor1.Time           `codec:"ctime" json:"ctime"`
	OutboxID              *string                `codec:"outboxID,omitempty" json:"outboxID,omitempty"`
	MessageBody           MessageBody            `codec:"messageBody" json:"messageBody"`
	SenderUsername        string                 `codec:"senderUsername" json:"senderUsername"`
	SenderDeviceName      string                 `codec:"senderDeviceName" json:"senderDeviceName"`
	SenderDeviceType      string                 `codec:"senderDeviceType" json:"senderDeviceType"`
	Superseded            bool                   `codec:"superseded" json:"superseded"`
	AssetUrlInfo          *UIAssetUrlInfo        `codec:"assetUrlInfo,omitempty" json:"assetUrlInfo,omitempty"`
	SenderDeviceRevokedAt *gregor1.Time          `codec:"senderDeviceRevokedAt,omitempty" json:"senderDeviceRevokedAt,omitempty"`
	AtMentions            []string               `codec:"atMentions" json:"atMentions"`
	ChannelMention        ChannelMention         `codec:"channelMention" json:"channelMention"`
	ChannelNameMentions   []UIChannelNameMention `codec:"channelNameMentions" json:"channelNameMentions"`
	IsEphemeral           bool                   `codec:"isEphemeral" json:"isEphemeral"`
	IsEphemeralExpired    bool                   `codec:"isEphemeralExpired" json:"isEphemeralExpired"`
	ExplodedBy            *string                `codec:"explodedBy,omitempty" json:"explodedBy,omitempty"`
	Etime                 gregor1.Time           `codec:"etime" json:"etime"`
	Reactions             ReactionMap            `codec:"reactions" json:"reactions"`
	HasPairwiseMacs       bool                   `codec:"hasPairwiseMacs" json:"hasPairwiseMacs"`
}

func (o UIMessageValid) DeepCopy() UIMessageValid {
	return UIMessageValid{
		MessageID: o.MessageID.DeepCopy(),
		Ctime:     o.Ctime.DeepCopy(),
		OutboxID: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.OutboxID),
		MessageBody:      o.MessageBody.DeepCopy(),
		SenderUsername:   o.SenderUsername,
		SenderDeviceName: o.SenderDeviceName,
		SenderDeviceType: o.SenderDeviceType,
		Superseded:       o.Superseded,
		AssetUrlInfo: (func(x *UIAssetUrlInfo) *UIAssetUrlInfo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.AssetUrlInfo),
		SenderDeviceRevokedAt: (func(x *gregor1.Time) *gregor1.Time {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.SenderDeviceRevokedAt),
		AtMentions: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.AtMentions),
		ChannelMention: o.ChannelMention.DeepCopy(),
		ChannelNameMentions: (func(x []UIChannelNameMention) []UIChannelNameMention {
			if x == nil {
				return nil
			}
			ret := make([]UIChannelNameMention, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.ChannelNameMentions),
		IsEphemeral:        o.IsEphemeral,
		IsEphemeralExpired: o.IsEphemeralExpired,
		ExplodedBy: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.ExplodedBy),
		Etime:           o.Etime.DeepCopy(),
		Reactions:       o.Reactions.DeepCopy(),
		HasPairwiseMacs: o.HasPairwiseMacs,
	}
}

type UIMessageOutbox struct {
	State       OutboxState     `codec:"state" json:"state"`
	OutboxID    string          `codec:"outboxID" json:"outboxID"`
	MessageType MessageType     `codec:"messageType" json:"messageType"`
	Body        string          `codec:"body" json:"body"`
	Ctime       gregor1.Time    `codec:"ctime" json:"ctime"`
	Ordinal     float64         `codec:"ordinal" json:"ordinal"`
	Filename    string          `codec:"filename" json:"filename"`
	Title       string          `codec:"title" json:"title"`
	Preview     *MakePreviewRes `codec:"preview,omitempty" json:"preview,omitempty"`
}

func (o UIMessageOutbox) DeepCopy() UIMessageOutbox {
	return UIMessageOutbox{
		State:       o.State.DeepCopy(),
		OutboxID:    o.OutboxID,
		MessageType: o.MessageType.DeepCopy(),
		Body:        o.Body,
		Ctime:       o.Ctime.DeepCopy(),
		Ordinal:     o.Ordinal,
		Filename:    o.Filename,
		Title:       o.Title,
		Preview: (func(x *MakePreviewRes) *MakePreviewRes {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Preview),
	}
}

type MessageUnboxedState int

const (
	MessageUnboxedState_VALID       MessageUnboxedState = 1
	MessageUnboxedState_ERROR       MessageUnboxedState = 2
	MessageUnboxedState_OUTBOX      MessageUnboxedState = 3
	MessageUnboxedState_PLACEHOLDER MessageUnboxedState = 4
)

func (o MessageUnboxedState) DeepCopy() MessageUnboxedState { return o }

var MessageUnboxedStateMap = map[string]MessageUnboxedState{
	"VALID":       1,
	"ERROR":       2,
	"OUTBOX":      3,
	"PLACEHOLDER": 4,
}

var MessageUnboxedStateRevMap = map[MessageUnboxedState]string{
	1: "VALID",
	2: "ERROR",
	3: "OUTBOX",
	4: "PLACEHOLDER",
}

func (e MessageUnboxedState) String() string {
	if v, ok := MessageUnboxedStateRevMap[e]; ok {
		return v
	}
	return ""
}

type UIMessage struct {
	State__       MessageUnboxedState        `codec:"state" json:"state"`
	Valid__       *UIMessageValid            `codec:"valid,omitempty" json:"valid,omitempty"`
	Error__       *MessageUnboxedError       `codec:"error,omitempty" json:"error,omitempty"`
	Outbox__      *UIMessageOutbox           `codec:"outbox,omitempty" json:"outbox,omitempty"`
	Placeholder__ *MessageUnboxedPlaceholder `codec:"placeholder,omitempty" json:"placeholder,omitempty"`
}

func (o *UIMessage) State() (ret MessageUnboxedState, err error) {
	switch o.State__ {
	case MessageUnboxedState_VALID:
		if o.Valid__ == nil {
			err = errors.New("unexpected nil value for Valid__")
			return ret, err
		}
	case MessageUnboxedState_ERROR:
		if o.Error__ == nil {
			err = errors.New("unexpected nil value for Error__")
			return ret, err
		}
	case MessageUnboxedState_OUTBOX:
		if o.Outbox__ == nil {
			err = errors.New("unexpected nil value for Outbox__")
			return ret, err
		}
	case MessageUnboxedState_PLACEHOLDER:
		if o.Placeholder__ == nil {
			err = errors.New("unexpected nil value for Placeholder__")
			return ret, err
		}
	}
	return o.State__, nil
}

func (o UIMessage) Valid() (res UIMessageValid) {
	if o.State__ != MessageUnboxedState_VALID {
		panic("wrong case accessed")
	}
	if o.Valid__ == nil {
		return
	}
	return *o.Valid__
}

func (o UIMessage) Error() (res MessageUnboxedError) {
	if o.State__ != MessageUnboxedState_ERROR {
		panic("wrong case accessed")
	}
	if o.Error__ == nil {
		return
	}
	return *o.Error__
}

func (o UIMessage) Outbox() (res UIMessageOutbox) {
	if o.State__ != MessageUnboxedState_OUTBOX {
		panic("wrong case accessed")
	}
	if o.Outbox__ == nil {
		return
	}
	return *o.Outbox__
}

func (o UIMessage) Placeholder() (res MessageUnboxedPlaceholder) {
	if o.State__ != MessageUnboxedState_PLACEHOLDER {
		panic("wrong case accessed")
	}
	if o.Placeholder__ == nil {
		return
	}
	return *o.Placeholder__
}

func NewUIMessageWithValid(v UIMessageValid) UIMessage {
	return UIMessage{
		State__: MessageUnboxedState_VALID,
		Valid__: &v,
	}
}

func NewUIMessageWithError(v MessageUnboxedError) UIMessage {
	return UIMessage{
		State__: MessageUnboxedState_ERROR,
		Error__: &v,
	}
}

func NewUIMessageWithOutbox(v UIMessageOutbox) UIMessage {
	return UIMessage{
		State__:  MessageUnboxedState_OUTBOX,
		Outbox__: &v,
	}
}

func NewUIMessageWithPlaceholder(v MessageUnboxedPlaceholder) UIMessage {
	return UIMessage{
		State__:       MessageUnboxedState_PLACEHOLDER,
		Placeholder__: &v,
	}
}

func (o UIMessage) DeepCopy() UIMessage {
	return UIMessage{
		State__: o.State__.DeepCopy(),
		Valid__: (func(x *UIMessageValid) *UIMessageValid {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Valid__),
		Error__: (func(x *MessageUnboxedError) *MessageUnboxedError {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Error__),
		Outbox__: (func(x *UIMessageOutbox) *UIMessageOutbox {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Outbox__),
		Placeholder__: (func(x *MessageUnboxedPlaceholder) *MessageUnboxedPlaceholder {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Placeholder__),
	}
}

type UIMessages struct {
	Messages   []UIMessage   `codec:"messages" json:"messages"`
	Pagination *UIPagination `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

func (o UIMessages) DeepCopy() UIMessages {
	return UIMessages{
		Messages: (func(x []UIMessage) []UIMessage {
			if x == nil {
				return nil
			}
			ret := make([]UIMessage, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Messages),
		Pagination: (func(x *UIPagination) *UIPagination {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Pagination),
	}
}

type ChatSearchHit struct {
	BeforeMessages []UIMessage `codec:"beforeMessages" json:"beforeMessages"`
	HitMessage     UIMessage   `codec:"hitMessage" json:"hitMessage"`
	AfterMessages  []UIMessage `codec:"afterMessages" json:"afterMessages"`
	Matches        []string    `codec:"matches" json:"matches"`
}

func (o ChatSearchHit) DeepCopy() ChatSearchHit {
	return ChatSearchHit{
		BeforeMessages: (func(x []UIMessage) []UIMessage {
			if x == nil {
				return nil
			}
			ret := make([]UIMessage, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.BeforeMessages),
		HitMessage: o.HitMessage.DeepCopy(),
		AfterMessages: (func(x []UIMessage) []UIMessage {
			if x == nil {
				return nil
			}
			ret := make([]UIMessage, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.AfterMessages),
		Matches: (func(x []string) []string {
			if x == nil {
				return nil
			}
			ret := make([]string, len(x))
			for i, v := range x {
				vCopy := v
				ret[i] = vCopy
			}
			return ret
		})(o.Matches),
	}
}

type ChatAttachmentDownloadStartArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatAttachmentDownloadProgressArg struct {
	SessionID     int   `codec:"sessionID" json:"sessionID"`
	BytesComplete int64 `codec:"bytesComplete" json:"bytesComplete"`
	BytesTotal    int64 `codec:"bytesTotal" json:"bytesTotal"`
}

type ChatAttachmentDownloadDoneArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type ChatInboxUnverifiedArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Inbox     string `codec:"inbox" json:"inbox"`
}

type ChatInboxConversationArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Conv      string `codec:"conv" json:"conv"`
}

type ChatInboxFailedArg struct {
	SessionID int              `codec:"sessionID" json:"sessionID"`
	ConvID    ConversationID   `codec:"convID" json:"convID"`
	Error     InboxUIItemError `codec:"error" json:"error"`
}

type ChatThreadCachedArg struct {
	SessionID int     `codec:"sessionID" json:"sessionID"`
	Thread    *string `codec:"thread,omitempty" json:"thread,omitempty"`
}

type ChatThreadFullArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Thread    string `codec:"thread" json:"thread"`
}

type ChatSearchHitArg struct {
	SessionID int           `codec:"sessionID" json:"sessionID"`
	SearchHit ChatSearchHit `codec:"searchHit" json:"searchHit"`
}

type ChatSearchDoneArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
	NumHits   int `codec:"numHits" json:"numHits"`
}

type ChatConfirmChannelDeleteArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Channel   string `codec:"channel" json:"channel"`
}

type ChatUiInterface interface {
	ChatAttachmentDownloadStart(context.Context, int) error
	ChatAttachmentDownloadProgress(context.Context, ChatAttachmentDownloadProgressArg) error
	ChatAttachmentDownloadDone(context.Context, int) error
	ChatInboxUnverified(context.Context, ChatInboxUnverifiedArg) error
	ChatInboxConversation(context.Context, ChatInboxConversationArg) error
	ChatInboxFailed(context.Context, ChatInboxFailedArg) error
	ChatThreadCached(context.Context, ChatThreadCachedArg) error
	ChatThreadFull(context.Context, ChatThreadFullArg) error
	ChatSearchHit(context.Context, ChatSearchHitArg) error
	ChatSearchDone(context.Context, ChatSearchDoneArg) error
	ChatConfirmChannelDelete(context.Context, ChatConfirmChannelDeleteArg) (bool, error)
}

func ChatUiProtocol(i ChatUiInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "chat.1.chatUi",
		Methods: map[string]rpc.ServeHandlerDescription{
			"chatAttachmentDownloadStart": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentDownloadStartArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentDownloadStartArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentDownloadStartArg)(nil), args)
						return
					}
					err = i.ChatAttachmentDownloadStart(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatAttachmentDownloadProgress": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentDownloadProgressArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentDownloadProgressArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentDownloadProgressArg)(nil), args)
						return
					}
					err = i.ChatAttachmentDownloadProgress(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"chatAttachmentDownloadDone": {
				MakeArg: func() interface{} {
					ret := make([]ChatAttachmentDownloadDoneArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatAttachmentDownloadDoneArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatAttachmentDownloadDoneArg)(nil), args)
						return
					}
					err = i.ChatAttachmentDownloadDone(ctx, (*typedArgs)[0].SessionID)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatInboxUnverified": {
				MakeArg: func() interface{} {
					ret := make([]ChatInboxUnverifiedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatInboxUnverifiedArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatInboxUnverifiedArg)(nil), args)
						return
					}
					err = i.ChatInboxUnverified(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatInboxConversation": {
				MakeArg: func() interface{} {
					ret := make([]ChatInboxConversationArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatInboxConversationArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatInboxConversationArg)(nil), args)
						return
					}
					err = i.ChatInboxConversation(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatInboxFailed": {
				MakeArg: func() interface{} {
					ret := make([]ChatInboxFailedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatInboxFailedArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatInboxFailedArg)(nil), args)
						return
					}
					err = i.ChatInboxFailed(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatThreadCached": {
				MakeArg: func() interface{} {
					ret := make([]ChatThreadCachedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatThreadCachedArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatThreadCachedArg)(nil), args)
						return
					}
					err = i.ChatThreadCached(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"chatThreadFull": {
				MakeArg: func() interface{} {
					ret := make([]ChatThreadFullArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatThreadFullArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatThreadFullArg)(nil), args)
						return
					}
					err = i.ChatThreadFull(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodNotify,
			},
			"chatSearchHit": {
				MakeArg: func() interface{} {
					ret := make([]ChatSearchHitArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatSearchHitArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatSearchHitArg)(nil), args)
						return
					}
					err = i.ChatSearchHit(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatSearchDone": {
				MakeArg: func() interface{} {
					ret := make([]ChatSearchDoneArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatSearchDoneArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatSearchDoneArg)(nil), args)
						return
					}
					err = i.ChatSearchDone(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"chatConfirmChannelDelete": {
				MakeArg: func() interface{} {
					ret := make([]ChatConfirmChannelDeleteArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ChatConfirmChannelDeleteArg)
					if !ok {
						err = rpc.NewTypeError((*[]ChatConfirmChannelDeleteArg)(nil), args)
						return
					}
					ret, err = i.ChatConfirmChannelDelete(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type ChatUiClient struct {
	Cli rpc.GenericClient
}

func (c ChatUiClient) ChatAttachmentDownloadStart(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentDownloadStartArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentDownloadStart", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatAttachmentDownloadProgress(ctx context.Context, __arg ChatAttachmentDownloadProgressArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.chatUi.chatAttachmentDownloadProgress", []interface{}{__arg})
	return
}

func (c ChatUiClient) ChatAttachmentDownloadDone(ctx context.Context, sessionID int) (err error) {
	__arg := ChatAttachmentDownloadDoneArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatAttachmentDownloadDone", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatInboxUnverified(ctx context.Context, __arg ChatInboxUnverifiedArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatInboxUnverified", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatInboxConversation(ctx context.Context, __arg ChatInboxConversationArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatInboxConversation", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatInboxFailed(ctx context.Context, __arg ChatInboxFailedArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatInboxFailed", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatThreadCached(ctx context.Context, __arg ChatThreadCachedArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.chatUi.chatThreadCached", []interface{}{__arg})
	return
}

func (c ChatUiClient) ChatThreadFull(ctx context.Context, __arg ChatThreadFullArg) (err error) {
	err = c.Cli.Notify(ctx, "chat.1.chatUi.chatThreadFull", []interface{}{__arg})
	return
}

func (c ChatUiClient) ChatSearchHit(ctx context.Context, __arg ChatSearchHitArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatSearchHit", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatSearchDone(ctx context.Context, __arg ChatSearchDoneArg) (err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatSearchDone", []interface{}{__arg}, nil)
	return
}

func (c ChatUiClient) ChatConfirmChannelDelete(ctx context.Context, __arg ChatConfirmChannelDeleteArg) (res bool, err error) {
	err = c.Cli.Call(ctx, "chat.1.chatUi.chatConfirmChannelDelete", []interface{}{__arg}, &res)
	return
}
