package i18n

import (
	"github.com/nicksnyder/go-i18n/src/pkg/msg"
)

const defaultLocale = ""

var currentLocale = defaultLocale
var translations = make(map[string]map[string]string)

// SetLocale sets the locale to use for translated messages.
func SetLocale(locale string) {
	currentLocale = locale
}

// AddTranslation adds a translated string to the dictionary and returns
// an id for the message.
//
// This function is used by the Go code generated by the goi18n command line tool.
func AddTranslation(locale, context, content, translation string) string {
	if translations[locale] == nil {
		translations[locale] = make(map[string]string)
	}
	id := msg.Id(context, content)
	translations[locale][id] = translation
	return id
}

// NewMessage returns a Message that may be translated into multiple langauges.
//
// The content of the message is the string to translate.
//
// The context of the message can be used to provide context to translators
// and to diambiguate homonyms. Messages with the same content and different contexts
// may have different translations.
//
// See the goi18n command line tool for documentation on how to extract messages for translation.
// http://github.com/nicksnyder/go-i18n
func NewMessage(content, context string) *Message {
	id := AddTranslation(defaultLocale, context, content, content)
	return &Message{id: id}
}

// Message is a string that is translated into multiple langauges.
type Message struct {
	id string
}

// String returns the message translated in the current locale.
// If there is no translation for the current locale, the message
// is returned untranslated
func (m *Message) String() string {
	t, found := translations[currentLocale][m.id]
	if !found {
		t = translations[defaultLocale][m.id]
	}
	return t
}
