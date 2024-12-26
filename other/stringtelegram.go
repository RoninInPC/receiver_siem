package other

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"slices"
	"strings"
)

var markdownSymbols = []string{"_", "`", "[", "*"}

const TelegramMessageLimit = 4096

type MessagesConfig []tgbotapi.MessageConfig

func ToBold(str string) string {
	return "*" + str + "*"
}

func ToItalic(str string) string {
	return "_" + str + "_"
}

func NewMarkdownMessage(chatID int64, text string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	return msg
}

func MakeNewMarkdownMessages(chatID int64, text string) []tgbotapi.MessageConfig {
	textes := SplitByTelegramLimits(text, TelegramMessageLimit-30)
	ans := make(MessagesConfig, 0)
	for _, t := range textes {
		ans = append(ans, NewMarkdownMessage(chatID, t))
	}
	return ans
}

func FixString(str string) string {
	str = strings.Replace(str, "_", "\\_", strings.Count(str, "_"))
	str = strings.Replace(str, "`", "\\`", strings.Count(str, "`"))
	str = strings.Replace(str, "[", "\\[", strings.Count(str, "["))
	return strings.Replace(str, "*", "\\*", strings.Count(str, "*"))
}

func GetCountMarkdownSymbol(str string, symbol string) int {
	if !slices.Contains(markdownSymbols, symbol) {
		return 0
	}
	return strings.Count(str, symbol) - strings.Count(str, "\\"+symbol)
}

func CheckMarkdownSymbols(str string) (int, string, bool) {
	minimum := len(str)
	symbolMin := ""
	for _, symbol := range markdownSymbols {
		if GetCountMarkdownSymbol(str, symbol)%2 == 1 {
			for j := len(str) - 1; j >= 0; j-- {
				if str[j] == symbol[0] {
					if j >= 1 {

					}
				}
			}
			minimum = min(minimum, strings.LastIndex(str, symbol))
			symbolMin = symbol
		}
	}
	if minimum == len(str) {
		return -1, "", true
	}
	return minimum, symbolMin, false
}

func SplitByTelegramLimits(text string, limit int) []string {
	var lines []string
	for i := 0; i < len(text); i += limit {
		end := i + limit
		if end > len(text) {
			end = len(text)
		} else {
			if text[end] == '\\' {
				end++
			}
		}
		lines = append(lines, text[i:end])
	}

	if len(lines) == 1 {
		return lines
	}

	for i := 0; i < len(lines)-1; i++ {
		for {
			_, symbol, ok := CheckMarkdownSymbols(lines[i])
			if ok {
				break
			} else {
				lines[i] += symbol
				lines[i+1] = symbol + lines[i+1]
			}
		}
	}

	lines[0] += " ..."
	lines[len(lines)-1] = "... " + lines[len(lines)-1]
	for i := 1; i < len(lines)-1; i++ {
		lines[i] = "... " + lines[i] + " ..."
	}

	return lines
}

func NopeMarkdown(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	for _, symbol := range markdownSymbols {
		msg.Text = strings.Replace(msg.Text, symbol, "", strings.Count(msg.Text, symbol))
	}
	return msg
}
