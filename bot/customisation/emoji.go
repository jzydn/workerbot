package customisation

import (
	"fmt"
	"github.com/rxdn/gdl/objects"
	"github.com/rxdn/gdl/objects/guild/emoji"
)

type CustomEmoji struct {
	Name     string
	Id       uint64
	Animated bool
}

func NewCustomEmoji(name string, id uint64, animated bool) CustomEmoji {
	return CustomEmoji{
		Name: name,
		Id:   id,
	}
}

func (e CustomEmoji) String() string {
	if e.Animated {
		return fmt.Sprintf("<a:%s:%d>", e.Name, e.Id)
	} else {
		return fmt.Sprintf("<:%s:%d>", e.Name, e.Id)
	}
}

func (e CustomEmoji) BuildEmoji() *emoji.Emoji {
	return &emoji.Emoji{
		Id:       objects.NewNullableSnowflake(e.Id),
		Name:     e.Name,
		Animated: e.Animated,
	}
}

var (
	EmojiId         = NewCustomEmoji("id", 1370547041964392579, false)
	EmojiOpen       = NewCustomEmoji("open", 1370547510623600781, false)
	EmojiOpenTime   = NewCustomEmoji("opentime", 1370547830871298078, false)
	EmojiClose      = NewCustomEmoji("close", 1370548144273625189, false)
	EmojiCloseTime  = NewCustomEmoji("closetime", 1339008285390405713, false)
	EmojiReason     = NewCustomEmoji("reason", 1370548565608366182, false)
	EmojiSubject    = NewCustomEmoji("subject", 1013527369832738907, false)
	EmojiTranscript = NewCustomEmoji("transcript", 1339007964018901103, false)
	EmojiClaim      = NewCustomEmoji("claim", 1370549020401078456, false)
	EmojiPanel      = NewCustomEmoji("panel", 1339004212222230640, false)
	EmojiRating     = NewCustomEmoji("rating", 1339004985224200222, false)
	EmojiStaff      = NewCustomEmoji("staff", 1339003554383527997, false)
	EmojiThread     = NewCustomEmoji("thread", 1339003444098498561, false)
	EmojiBulletLine = NewCustomEmoji("bulletline", 1339003248329228441, false)
	EmojiPatreon    = NewCustomEmoji("patreon", 1339007402044817470, false)
	EmojiDiscord    = NewCustomEmoji("discord", 1339002870657454091, false)
	//EmojiTime       = NewCustomEmoji("time", 974006684622159952, false)
)

// PrefixWithEmoji Useful for whitelabel bots
func PrefixWithEmoji(s string, emoji CustomEmoji, includeEmoji bool) string {
	if includeEmoji {
		return fmt.Sprintf("%s %s", emoji, s)
	} else {
		return s
	}
}