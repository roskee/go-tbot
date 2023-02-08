package entity

import "strings"

// Message holds a message object the telegram server sends.
//
// TODO: incomplete
type Message struct {
	// MessageID is a unique message identifier inside this chat.
	//
	// it is a required field.
	MessageID int64 `json:"message_id,omitempty"`
	// MessageThreadID is a unique identifier of a message thread to which the message belongs.
	// it is for supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
	// From is the sender of the message.
	// It is empty for messages sent to channels.
	From *User `json:"from,omitempty"`
	// SenderChat is the sender of the message, sent on behalf of a chat.
	SenderChat *Chat `json:"sender_chat,omitempty"`
	// Date is the date the message was sent in Unix time.
	//
	// It is a required field
	Date int64 `json:"date,omitempty"`
	// Chat is the conversation this message belongs to.
	//
	// It is a required field
	Chat *Chat `json:"chat,omitempty"`
	// ForwardFrom is, for forwarded messages, sender of the original message.
	ForwardFrom *User `json:"forward_from,omitempty"`
	// ForwardFromChat is,  For messages forwarded from channels or from anonymous administrators,
	// information about the original sender chat.
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`
	// ForwardFromMessageID is, for messages forwarded from channels, identifier of the original message in the channel.
	ForwardFromMessageID int64 `json:"forward_from_message_id,omitempty"`
	// ForwardSignature is, for forwarded messages that were originally sent in channels or by an anonymous chat administrator,
	// signature of the message sender if present.
	ForwardSignature string `json:"forward_signature,omitempty"`
	// ForwardSenderName is sender's name for messages forwarded from users
	// who disallow adding a link to their account in forwarded messages.
	ForwardSenderName string `json:"forward_sender_name,omitempty"`
	// ForwardDate is, for forwarded messages, date the original message was sent in Unix time.
	ForwardDate int64 `json:"forward_date,omitempty"`
	// IsTopicMessage is true, if the message is sent to a forum topic.
	IsTopicMessage bool `json:"is_topic_message,omitempty"`
	// IsAutomaticForward is true, if the message is a channel post
	// that was automatically forwarded to the connected discussion group.
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`
	// ReplyToMessage is, for replies, the original message.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	// ViaBot is the bot through which the message was sent.
	ViaBot *User `json:"via_bot,omitempty"`
	// EditDate is the date the message was last edited in Unix time.
	EditDate int64 `json:"edit_date,omitempty"`
	// HasProtectedContent is true, if the message can't be forwarded.
	HasProtectedContent bool `json:"has_protected_content,omitempty"`
	// MediaGroupID is the unique identifier of a media message group this message belongs to.
	MediaGroupID string `json:"media_group_id,omitempty"`
	// AuthorSignature is signature of the post author for messages in channels,
	// or the custom title of an anonymous group administrator.
	AuthorSignature string `json:"author_signature,omitempty"`
	// Text is, for text messages, the actual UTF-8 text of the message.
	Text string `json:"text,omitempty"`
	// Entities is, for text messages, special entities
	// like usernames, URLs, bot commands, etc. that appear in the text.
	Entities []MessageEntity `json:"entities,omitempty"`
	// Animation is, when message is an animation, information about the animation.
	Animation *Animation `json:"animation,omitempty"`
	// Audio is, when message is an audio file, information about the file.
	Audio *Audio `json:"audio,omitempty"`
	// Document is, when essage is a general file, information about the file.
	Document *Document `json:"document,omitempty"`
	// Photo is, when essage is a photo, available sizes of the photo.
	Photo []PhotoSize `json:"photo,omitempty"`
	// Sticker is, when message is a sticker, information about the sticker.
	Sticker *Sticker `json:"sticker,omitempty"`
	// Video is, when message is a video, information about the video.
	Video *Video `json:"video,omitempty"`
	// VideoNote is, when message is a video note, information about the video message.
	VideoNote *VideoNote `json:"video_note,omitempty"`
	// Voice is, when message is a voice message, information about the file.
	Voice *Voice `json:"voice,omitempty"`
	// Caption is the caption for the animation, audio, document, photo, video or voice.
	Caption string `json:"caption,omitempty"`
	// CaptionEntities is, for messages with a caption,
	// special entities like usernames, URLs, bot commands, etc. that appear in the caption.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
}

// GetCommand checks if this message has Text starting with '/'.
// if so, it returns the immediate word attached to it
func (m *Message) GetCommand() string {
	if len(m.Text) == 0 || m.Text[0] != '/' {
		return ""
	}
	index := strings.Index(m.Text, " ")
	if index == -1 {
		index = len(m.Text)
	}
	return m.Text[1:index]
}

// MessageEntity  represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type is Type of the entity.
	// Currently, can be “mention” (@username), “hashtag” (#hashtag), “cashtag” ($USD),
	// “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org),
	// “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text),
	// “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler” (spoiler message),
	// “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable text URLs),
	// “text_mention” (for users without usernames), “custom_emoji” (for inline custom emoji stickers)
	//
	// It is a required field
	Type string `json:"type,omitempty"`
	// Offset is offset in UTF-16 code units to the start of the entity.
	//
	// It is a required field
	Offset int64 `json:"offset,omitempty"`
	// Length is length of the entity in UTF-16 code units.
	//
	// It is a required field
	Length int64 `json:"length,omitempty"`
	// URL is, for “text_link” only, URL that will be opened after user taps on the text.
	URL string `json:"url,omitempty"`
	// User is, for “text_mention” only, the mentioned user.
	User *User `json:"user,omitempty"`
	// Language is, for “pre” only, the programming language of the entity text.
	Language string `json:"language,omitempty"`
	// CustomEmojiID is, for “custom_emoji” only, unique identifier of the custom emoji.
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

type ReplyMarkup struct {
	*InlineKeyboardMarkup
	*ReplyKeyboardMarkup
	*ReplyKeyboardRemove
	*ForceReply
}
