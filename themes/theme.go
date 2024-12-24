package themes

import (
	"github.com/charmbracelet/bubbles/v2/list"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/charmbracelet/log"
)

type Theme interface {
	RenderBold(text string) string
	RenderInfo(text string) string
	RenderNotice(text string) string
	RenderSuccess(text string) string
	RenderWarning(text string) string
	RenderError(text string) string
	RenderEmphasis(text string) string
	RenderUnknown(text string) string
	RenderLevel(str string, lvl OutputLevel) string
	RenderHeader(appName, stateKey, stateVal string, width int) string
	RenderFooter(text string, width int) string
	RenderKeyAndValue(key, value string) string
	RenderKeyAndValueWithBreak(key, value string) string
	RenderInputForm(text string) string
	RenderInContainer(text string) string

	SpinnerStyle() lipgloss.Style
	EntityViewStyle() lipgloss.Style
	CollectionStyle() lipgloss.Style
	BoxStyle() lipgloss.Style
	ListStyles() list.Styles
	ListItemStyles() list.DefaultItemStyles
	
	LoggerStyles() *log.Styles
	GlamourMarkdownStyleJSON() (string, error)
	HuhTheme() *huh.Theme
}
