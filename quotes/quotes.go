// This is the quotes package

package quotes

import (
	"github.com/hackebrot/turtle"
)

// Returns an emoji based on supplied name
func GetEmoji(name string) string {
	emoji, ok := turtle.Emojis[name]
	if !ok {
		return ""
	}
	return emoji.Char
}
