package quotes

import (
	"github.com/hackebrot/turtle"
)

func GetEmoji(name string) string {
	emoji, ok := turtle.Emojis[name]
	if !ok {
		return ""
	}
	return emoji.Char
}
