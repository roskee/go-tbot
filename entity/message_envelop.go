package entity

// MessageEnvelop holds the object that is used to send a new message
type MessageEnvelop struct {
	// ChatID is the unique identifier for the target chat
	// or username of the target channel
	//
	// It is a required field.
	ChatID string `json:"chat_id,omitempty"`
	// MessageThreadID is the unique identifier for the target message thread (topic) of the forum;
	// for forum supergroups only
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
	// Text is text of the message to be sent, 1-4096 characters after entities parsing.
	//
	// It is a required field
	Text string `json:"text,omitempty"`
	// ParseMode is the mode for parsing entities in the message text.
	ParseMode string `json:"parse_mode,omitempty"`
	// Entities are A JSON-serialized list of special entities that appear in message text,
	// which can be specified instead of parse_mode.
	Entities []MessageEntity `json:"entities,omitempty"`
	// DisableWebPagePreview disables link previews for links in this message.
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	// DisableNotification sends the message silently.
	// Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// ProtectContent protects the contents of the sent message from forwarding and saving.
	ProtectContent bool `json:"protect_content,omitempty"`
	// ReplyToMessageID is, if the message is a reply, ID of the original message.
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`
	// AllowSendingWithoutReply can be true if the message should be sent
	// even if the specified replied-to message is not found.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup contains additional interface options.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
	// Photo is the photo to send.
	// Pass a file_id as String to send a photo that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a photo from the Internet,
	// or upload a new photo using multipart/form-data.
	//
	// It is a required field for sending a photo.
	Photo any `json:"photo,omitempty"`
	// Caption is the media caption
	Caption string `json:"caption,omitempty"`
	// CaptionEntities is list of special entities that appear in the caption,
	// which can be specified instead of parse_mode.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// Audio is the audio file to send.
	// Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get an audio file from the Internet,
	// or upload a new one using multipart/form-data.
	//
	// It is a required field for sending an audio.
	Audio any `json:"audio,omitempty"`
	// Duration of the media in seconds.
	Duration int64 `json:"duration,omitempty"`
	// Performer of the media
	Performer string `json:"performer,omitempty"`
	// Title is the track name
	Title string `json:"title,omitempty"`
	// Thumb is the thumbnail of the file sent.
	// It should be in JPEG format and less than 200 kB in size.
	// It's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	Thumb any `json:"thumb,omitempty"`
	// Document is the document to send.
	// Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a file from the Internet,
	// or upload a new one using multipart/form-data.
	//
	// It is a required field for sending a document.
	Document any `json:"document,omitempty"`
	// DisableContentTypeDetection disables automatic server-side content type detection
	// for files uploaded using multipart/form-data.
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
	// Video is the Video to send.
	// Pass a file_id as String to send a video that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a video from the Internet,
	// or upload a new video using multipart/form-data.
	//
	// It is a required field for sending video.
	Video any `json:"video,omitempty"`
	// Width of the media
	Width int64 `json:"width,omitempty"`
	// Height of the media
	Height int64 `json:"height,omitempty"`
	// SupportsStreaming can be true if the uploaded media is suitable for streaming.
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
	// Animation is the animation to send.
	// Pass a file_id as String to send an animation that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get an animation from the Internet,
	// or upload a new animation using multipart/form-data.
	//
	// It is a required field for sending an animation.
	Animation any `json:"animation,omitempty"`
	// Voice is the audio file to send.
	// Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a file from the Internet,
	// or upload a new one using multipart/form-data.
	//
	// It is a required field for sending a voice.
	Voice any `json:"voice,omitempty"`
	// VideoNote is the video note to send.
	// Pass a file_id as String to send a video note that exists on the Telegram servers (recommended)
	// or upload a new video using multipart/form-data.
	// Sending video notes by a URL is currently unsupported.
	//
	// It is a required field for sending a video note.
	VideoNote any `json:"video_note,omitempty"`
	// Length is the media width and height, i.e. diameter of the video message.
	Length int64 `json:"length,omitempty"`
	// Media is an array of media messages to be sent, must include 2-10 items.
	Media []InputMedia `json:"media,omitempty"`

	// All Parameters for sendLocation Method
	*Location

	// All Parameters for sendContact Method
	*Contact

	// All Parameters for sendPoll Method
	*Poll

	// All Parameters for sendDice Method
	*Dice

	// Action is type of action to broadcast. Choose one, depending on what the user is about to receive.
	// typing for text messages, upload_photo for photos, record_video or upload_video for videos,
	// record_voice or upload_voice for voice notes, upload_document for general files, choose_sticker for stickers,
	// find_location for location data, record_video_note or upload_video_note for video notes.
	//
	// It is a required field for sending a chat action.
	Action string `json:"action"`
}

// FileEnvelop represents a file to be uploaded to the telegram server.
type FileEnvelop struct {
	// Name is the field name for this file for the upload request.
	Name string
	// Path is the location of this file on this device.
	Path string
}
