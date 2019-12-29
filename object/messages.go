package object // import "github.com/SevereCloud/vksdk/object"

import (
	"encoding/json"
	"fmt"
)

// MessagesAudioMessage struct
type MessagesAudioMessage struct {
	AccessKey string `json:"access_key"` // Access key for the document
	ID        int    `json:"id"`         // Document ID
	OwnerID   int    `json:"owner_id"`   // Document owner ID
	Duration  int    `json:"duration"`   // Audio message duration in seconds
	LinkMp3   string `json:"link_mp3"`   // MP3 file URL
	LinkOgg   string `json:"link_ogg"`   // OGG file URL
	Waveform  []int  `json:"waveform"`   // Sound visualisation
}

// ToAttachment return attachment format
func (doc MessagesAudioMessage) ToAttachment() string {
	return fmt.Sprintf("doc%d_%d", doc.OwnerID, doc.ID)
}

// MessagesGraffiti struct
type MessagesGraffiti struct {
	AccessKey string `json:"access_key"` // Access key for the document
	ID        int    `json:"id"`         // Document ID
	OwnerID   int    `json:"owner_id"`   // Document owner ID
	URL       string `json:"url"`        // Graffiti URL
	Width     int    `json:"width"`      // Graffiti width
	Height    int    `json:"height"`     // Graffiti height
}

// ToAttachment return attachment format
func (doc MessagesGraffiti) ToAttachment() string {
	return fmt.Sprintf("doc%d_%d", doc.OwnerID, doc.ID)
}

// MessagesMessage struct
type MessagesMessage struct {
	AdminAuthorID         int                         `json:"admin_author_id"` // Only for messages from community. Contains user ID of community admin, who sent this message.
	Action                MessagesMessageAction       `json:"action"`
	Attachments           []MessagesMessageAttachment `json:"attachments"`
	ConversationMessageID int                         `json:"conversation_message_id"` // Unique auto-incremented number for all messages with this peer
	Date                  int                         `json:"date"`                    // Date when the message has been sent in Unixtime
	Deleted               int                         `json:"deleted"`                 // Is it an deleted message
	FromID                int                         `json:"from_id"`                 // Message author's ID
	FwdMessages           []MessagesMessage           `json:"fwd_Messages"`            // Forwarded messages
	ReplyMessage          *MessagesMessage            `json:"reply_message"`
	Geo                   BaseGeo                     `json:"geo"`
	ID                    int                         `json:"id"`        // Message ID
	Important             bool                        `json:"important"` // Is it an important message
	IsHidden              bool                        `json:"is_hidden"`
	IsCropped             bool                        `json:"is_cropped"`
	Keyboard              MessagesKeyboard            `json:"keyboard"`
	Out                   int                         `json:"out"` // Information whether the message is outcoming
	Payload               string                      `json:"payload"`
	PeerID                int                         `json:"peer_id"`   // Peer ID
	RandomID              int                         `json:"random_id"` // ID used for sending messages. It returned only for outgoing messages
	Ref                   string                      `json:"ref"`
	RefSource             string                      `json:"ref_source"`
	Text                  string                      `json:"text"`          // Message text
	UpdateTime            int                         `json:"update_time"`   // Date when the message has been updated in Unixtime
	MembersСount          int                         `json:"members_count"` // Members number
}

// MessagesKeyboard struct
type MessagesKeyboard struct {
	AuthorID int                        `json:"author_id,omitempty"` // Community or bot, which set this keyboard
	Buttons  [][]MessagesKeyboardButton `json:"buttons"`
	OneTime  bool                       `json:"one_time"` // Should this keyboard disappear on first use
	Inline   bool                       `json:"inline,omitempty"`
}

func NewMessagesKeyboard(oneTime bool, inline bool) MessagesKeyboard {
	return MessagesKeyboard{
		Buttons: [][]MessagesKeyboardButton{},
		OneTime: oneTime,
		Inline:  inline,
	}
}

// AddRow add row in MessagesKeyboard
func (keyboard *MessagesKeyboard) AddRow() {
	if len(keyboard.Buttons) == 0 {
		keyboard.Buttons = make([][]MessagesKeyboardButton, 1)
	} else {
		row := make([]MessagesKeyboardButton, 0)
		keyboard.Buttons = append(keyboard.Buttons, row)
	}
}

// AddTextButton add Text button in last row
func (keyboard *MessagesKeyboard) AddTextButton(label string, payload string, color string) {
	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Type:    ButtonText,
			Label:   label,
			Payload: payload,
		},
		Color: color,
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)
}

// AddOpenLinkButton add Open Link button in last row
func (keyboard *MessagesKeyboard) AddOpenLinkButton(link, label, payload string) {
	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Type:    ButtonOpenLink,
			Payload: payload,
			Label:   label,
			Link:    link,
		},
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)
}

// AddLocationButton add Location button in last row
func (keyboard *MessagesKeyboard) AddLocationButton(payload string) {
	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Type:    ButtonLocation,
			Payload: payload,
		},
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)
}

// AddVKPayButton add VK Pay button in last row
func (keyboard *MessagesKeyboard) AddVKPayButton(payload string, hash string) {
	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Type:    ButtonVKPay,
			Payload: payload,
			Hash:    hash,
		},
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)
}

// AddVKAppsButton add VK Apps button in last row
func (keyboard *MessagesKeyboard) AddVKAppsButton(appID, ownerID int, payload, label, hash string) {
	button := MessagesKeyboardButton{
		Action: MessagesKeyboardButtonAction{
			Type:    ButtonVKApp,
			AppID:   appID,
			OwnerID: ownerID,
			Payload: payload,
			Label:   label,
			Hash:    hash,
		},
	}

	lastRow := len(keyboard.Buttons) - 1
	keyboard.Buttons[lastRow] = append(keyboard.Buttons[lastRow], button)
}

// ToJSON returns the JSON encoding of MessagesKeyboard
func (keyboard MessagesKeyboard) ToJSON() string {
	b, _ := json.Marshal(keyboard)
	return string(b)
}

// MessagesKeyboardButton struct
type MessagesKeyboardButton struct {
	Action MessagesKeyboardButtonAction `json:"action"`
	Color  string                       `json:"color"` // Button color
}

// MessagesKeyboardButtonAction struct
type MessagesKeyboardButtonAction struct {
	AppID   int    `json:"app_id,omitempty"`   // Fragment value in app link like vk.com/app{app_id}_-654321#hash
	Hash    string `json:"hash,omitempty"`     // Fragment value in app link like vk.com/app123456_-654321#{hash}
	Label   string `json:"label,omitempty"`    // Label for button
	OwnerID int    `json:"owner_id,omitempty"` // Fragment value in app link like vk.com/app123456_{owner_id}#hash
	Payload string `json:"payload,omitempty"`  // Additional data sent along with message for developer convenience
	Type    string `json:"type"`               // Button type
	Link    string `json:"link"`               // Link URL
}

// MessagesChat struct
type MessagesChat struct {
	AdminID      int                       `json:"admin_id"`  // Chat creator ID
	ID           int                       `json:"id"`        // Chat ID
	Kicked       int                       `json:"kicked"`    // Shows that user has been kicked from the chat
	Left         int                       `json:"left"`      // Shows that user has been left the chat
	Photo100     string                    `json:"photo_100"` // URL of the preview image with 100 px in width
	Photo200     string                    `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo50      string                    `json:"photo_50"`  // URL of the preview image with 50 px in width
	PushSettings MessagesChatPushSettings  `json:"push_settings"`
	Title        string                    `json:"title"` // Chat title
	Type         string                    `json:"type"`  // Chat type
	Users        []int                     `json:"users"`
	MembersCount int                       `json:"members_count"`
	Members      []int                     `json:"members"`
	Photo        MessagesChatSettingsPhoto `json:"photo"`
	Joined       bool                      `json:"joined"`
	LocalID      int                       `json:"local_id"`
}

// MessagesChatFull struct
type MessagesChatFull struct {
	AdminID      int                        `json:"admin_id"`  // Chat creator ID
	ID           int                        `json:"id"`        // Chat ID
	Kicked       int                        `json:"kicked"`    // Shows that user has been kicked from the chat
	Left         int                        `json:"left"`      // Shows that user has been left the chat
	Photo100     string                     `json:"photo_100"` // URL of the preview image with 100 px in width
	Photo200     string                     `json:"photo_200"` // URL of the preview image with 200 px in width
	Photo50      string                     `json:"photo_50"`  // URL of the preview image with 50 px in width
	PushSettings MessagesChatPushSettings   `json:"push_settings"`
	Title        string                     `json:"title"` // Chat title
	Type         string                     `json:"type"`  // Chat type
	Users        []MessagesUserXtrInvitedBy `json:"users"`
}

// MessagesChatPushSettings struct
type MessagesChatPushSettings struct {
	DisabledUntil int `json:"disabled_until"` // Time until that notifications are disabled
	Sound         int `json:"sound"`          // Information whether the sound is on
}

// MessagesChatSettingsPhoto struct
type MessagesChatSettingsPhoto struct {
	Photo100 string `json:"photo_100"`
	Photo200 string `json:"photo_200"`
	Photo50  string `json:"photo_50"`
}

// MessagesConversation struct
type MessagesConversation struct {
	CanWrite        MessagesConversationCanWrite     `json:"can_write"`
	ChatSettings    MessagesConversationChatSettings `json:"chat_settings"`
	InRead          int                              `json:"in_read"`         // Last message user have read
	LastMessageID   int                              `json:"last_message_id"` // ID of the last message in conversation
	Mentions        []int                            `json:"mentions"`        // IDs of messages with mentions
	MessageRequest  string                           `json:"message_request"`
	OutRead         int                              `json:"out_read"` // Last outcoming message have been read by the opponent
	Peer            MessagesConversationPeer         `json:"peer"`
	PushSettings    MessagesConversationPushSettings `json:"push_settings"`
	Important       bool                             `json:"important"`
	Unanswered      bool                             `json:"unanswered"`
	UnreadCount     int                              `json:"unread_count"` // Unread messages number
	CurrentKeyboard MessagesKeyboard                 `json:"current_keyboard"`
}

// MessagesConversationCanWrite struct
type MessagesConversationCanWrite struct {
	Allowed bool `json:"allowed"`
	Reason  int  `json:"reason"`
}

// MessagesConversationChatSettings struct
type MessagesConversationChatSettings struct {
	MembersCount  int                       `json:"members_count"`
	Photo         MessagesChatSettingsPhoto `json:"photo"`
	PinnedMessage MessagesPinnedMessage     `json:"pinned_message"`
	State         string                    `json:"state"`
	Title         string                    `json:"title"`
	ActiveIDS     []int                     `json:"active_ids"`
	ACL           struct {
		CanInvite           bool `json:"can_invite"`
		CanChangeInfo       bool `json:"can_change_info"`
		CanChangePin        bool `json:"can_change_pin"`
		CanPromoteUsers     bool `json:"can_promote_users"`
		CanSeeInviteLink    bool `json:"can_see_invite_link"`
		CanChangeInviteLink bool `json:"can_change_invite_link"`
		CanCopyChat         bool `json:"can_copy_chat"`
		CanModerate         bool `json:"can_moderate"`
	} `json:"acl"`
	IsGroupChannel bool `json:"is_group_channel"`
	OwnerID        int  `json:"owner_id"`
}

// MessagesConversationPeer struct
type MessagesConversationPeer struct {
	ID      int    `json:"id"`
	LocalID int    `json:"local_id"`
	Type    string `json:"type"`
}

// MessagesConversationPushSettings struct
type MessagesConversationPushSettings struct {
	DisabledUntil   int  `json:"disabled_until"`
	DisabledForever bool `json:"disabled_forever"`
	NoSound         bool `json:"no_sound"`
}

// MessagesConversationWithMessage struct
type MessagesConversationWithMessage struct {
	Conversation MessagesConversation `json:"conversation"`
	LastMessage  MessagesMessage      `json:"last_message"`
}

// MessagesDialog struct
type MessagesDialog struct {
	Important  int             `json:"important"`
	InRead     int             `json:"in_read"`
	Message    MessagesMessage `json:"message"`
	OutRead    int             `json:"out_read"`
	Unanswered int             `json:"unanswered"`
	Unread     int             `json:"unread"`
}

// MessagesHistoryAttachment struct
type MessagesHistoryAttachment struct {
	Attachment MessagesHistoryMessageAttachment `json:"attachment"`
	MessageID  int                              `json:"message_id"` // Message ID
}

// MessagesHistoryMessageAttachment struct
type MessagesHistoryMessageAttachment struct {
	Audio  AudioAudioFull `json:"audio"`
	Doc    DocsDoc        `json:"doc"`
	Link   BaseLink       `json:"link"`
	Market BaseLink       `json:"market"`
	Photo  PhotosPhoto    `json:"photo"`
	Share  BaseLink       `json:"share"`
	Type   string         `json:"type"`
	Video  VideoVideo     `json:"video"`
	Wall   BaseLink       `json:"wall"`
}

// MessagesLastActivity struct
type MessagesLastActivity struct {
	Online int `json:"online"` // Information whether user is online
	Time   int `json:"time"`   // Time when user was online in Unixtime
}

// MessagesLongpollParams struct
type MessagesLongpollParams struct {
	Key    string `json:"key"`    // Key
	Pts    int    `json:"pts"`    // Persistent timestamp
	Server string `json:"server"` // Server URL
	Ts     int    `json:"ts"`     // Timestamp
}

// MessagesMessageAction struct
type MessagesMessageAction struct {
	ConversationMessageID int                        `json:"conversation_message_id"` // Message ID
	Email                 string                     `json:"email"`                   // Email address for chat_invite_user or chat_kick_user actions
	MemberID              int                        `json:"member_id"`               // User or email peer ID
	Message               string                     `json:"message"`                 // Message body of related message
	Photo                 MessagesMessageActionPhoto `json:"photo"`
	Text                  string                     `json:"text"` // New chat title for chat_create and chat_title_update actions
	Type                  string                     `json:"type"`
}

// MessagesMessageActionPhoto struct
type MessagesMessageActionPhoto struct {
	Photo100 string `json:"photo_100"` // URL of the preview image with 100px in width
	Photo200 string `json:"photo_200"` // URL of the preview image with 200px in width
	Photo50  string `json:"photo_50"`  // URL of the preview image with 50px in width
}

// MessagesMessageAttachment struct
type MessagesMessageAttachment struct {
	Audio             AudioAudioFull       `json:"audio"`
	Doc               DocsDoc              `json:"doc"`
	Gift              GiftsLayout          `json:"gift"`
	Link              BaseLink             `json:"link"`
	Market            MarketMarketItem     `json:"market"`
	MarketMarketAlbum MarketMarketAlbum    `json:"market_market_album"`
	Photo             PhotosPhoto          `json:"photo"`
	Sticker           BaseSticker          `json:"sticker"`
	Type              string               `json:"type"`
	Video             VideoVideo           `json:"video"`
	Wall              WallWallpostAttached `json:"wall"`
	WallReply         WallWallComment      `json:"wall_reply"`
	AudioMessage      DocsDoc              `json:"audio_message"`
	Graffiti          DocsDoc              `json:"graffiti"`
	Poll              PollsPoll            `json:"poll"`
	Call              MessageCall          `json:"call"`
}

// MessageCall struct
type MessageCall struct {
	InitiatorID int    `json:"initiator_id"`
	ReceiverID  int    `json:"receiver_id"`
	State       string `json:"state"`
	Time        int    `json:"time"`
	Duration    int    `json:"duration"`
	Video       bool   `json:"video"`
}

// MessagesPinnedMessage struct
type MessagesPinnedMessage struct {
	Attachments           []MessagesMessageAttachment `json:"attachments"`
	ConversationMessageID int                         `json:"conversation_message_id"` // Unique auto-incremented number for all Messages with this peer
	Date                  int                         `json:"date"`                    // Date when the message has been sent in Unixtime
	FromID                int                         `json:"from_id"`                 // Message author's ID
	FwdMessages           []*MessagesMessage          `json:"fwd_Messages"`
	Geo                   BaseGeo                     `json:"geo"`
	ID                    int                         `json:"id"`      // Message ID
	PeerID                int                         `json:"peer_id"` // Peer ID
	ReplyMessage          *MessagesMessage            `json:"reply_message"`
	Text                  string                      `json:"text"` // Message text
}

// MessagesUserXtrInvitedBy struct
type MessagesUserXtrInvitedBy struct {
}
