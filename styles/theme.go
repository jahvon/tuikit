package styles

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

const (
	HeaderHeight = 3
	FooterHeight = 2
)

func (t Theme) Spinner() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(t.SecondaryColor)
}

func (t Theme) EntityView() lipgloss.Style {
	return lipgloss.NewStyle().MarginLeft(2)
}

func (t Theme) Collection() lipgloss.Style {
	return lipgloss.NewStyle().MarginLeft(2).Padding(0, 1)
}

func (t Theme) Box() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(t.BorderColor).
		Padding(0, 1).
		MarginLeft(2)
}

func (t Theme) RenderBold(text string) string {
	return lipgloss.NewStyle().Bold(true).Render(text)
}

func (t Theme) RenderInfo(text string) string {
	return lipgloss.NewStyle().Foreground(t.InfoColor).Render(text)
}

func (t Theme) RenderNotice(text string) string {
	return lipgloss.NewStyle().Foreground(t.EmphasisColor).Render(text)
}

func (t Theme) RenderSuccess(text string) string {
	return lipgloss.NewStyle().Foreground(t.SuccessColor).Render(text)
}

func (t Theme) RenderWarning(text string) string {
	return lipgloss.NewStyle().Foreground(t.WarningColor).Render(text)
}

func (t Theme) RenderError(text string) string {
	return lipgloss.NewStyle().Foreground(t.ErrorColor).Render(text)
}

func (t Theme) RenderEmphasis(text string) string {
	return lipgloss.NewStyle().Foreground(t.EmphasisColor).Render(text)
}

func (t Theme) RenderUnknown(text string) string {
	return lipgloss.NewStyle().Foreground(t.Gray).Render(text)
}

func (t Theme) RenderLevel(str string, lvl OutputLevel) string {
	if str == "" {
		return ""
	}

	switch lvl {
	case OutputLevelSuccess:
		return t.RenderSuccess(str)
	case OutputLevelNotice:
		return t.RenderNotice(str)
	case OutputLevelInfo:
		return t.RenderInfo(str)
	case OutputLevelWarning:
		return t.RenderWarning(str)
	case OutputLevelError:
		return t.RenderError(str)
	default:
		return t.RenderUnknown(str)
	}
}

func (t Theme) RenderHeader(appName, stateKey, stateVal string, width int) string {
	if width == 0 {
		return t.renderShortHeader(appName, stateKey, stateVal)
	}

	appNameStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(t.BorderColor).
		Foreground(t.PrimaryColor).
		AlignVertical(lipgloss.Center).
		Height(HeaderHeight - 2). // top and bottom borders
		Bold(true)
	ctxStyle := lipgloss.NewStyle().
		Foreground(t.SecondaryColor).
		Italic(true).
		AlignVertical(lipgloss.Center).
		Height(HeaderHeight)

	appNameContent := appNameStyle.Render(fmt.Sprintf(" %s ", appName))
	var stateContent string
	if stateKey != "" && stateVal != "" {
		stateContent = lipgloss.JoinHorizontal(
			0,
			ctxStyle.Render(t.RenderBold(fmt.Sprintf(" %s:", stateKey))),
			ctxStyle.Render(fmt.Sprintf("%s ", stateVal)),
		)
	}
	fullContent := lipgloss.JoinHorizontal(0, appNameContent, stateContent)
	borderWidth := width - lipgloss.Width(fullContent)
	if borderWidth < 0 {
		borderWidth = 0
	}
	borderStyle := lipgloss.NewStyle().
		Foreground(t.BorderColor).
		MaxWidth(borderWidth).
		Height(HeaderHeight).AlignVertical(lipgloss.Center)
	border := borderStyle.Render(strings.Repeat("─", borderWidth))
	return lipgloss.JoinHorizontal(0, fullContent, border, "\n")
}

func (t Theme) renderShortHeader(appName, ctxKey, ctxVal string) string {
	headerStyle := lipgloss.NewStyle().
		Foreground(t.PrimaryColor).
		Italic(true).Bold(true)
	return headerStyle.Render(fmt.Sprintf("<!-- %s %s(%s) --!>", appName, ctxKey, ctxVal))
}

func (t Theme) RenderFooter(text string, width int) string {
	footerStyle := lipgloss.NewStyle().
		Foreground(t.Gray).
		Border(lipgloss.NormalBorder()).
		BorderTop(true).BorderBottom(false).
		BorderLeft(false).BorderRight(false).
		BorderForeground(t.BorderColor).
		Height(FooterHeight - 1). // top border
		Width(width)
	return footerStyle.Render(text)
}

func (t Theme) RenderKeyAndValue(key, value string) string {
	keyStyle := lipgloss.NewStyle().Foreground(t.SecondaryColor).Bold(true)
	valueStyle := lipgloss.NewStyle().Foreground(t.Gray)
	return keyStyle.Render(key) + ": " + valueStyle.Render(value)
}

func (t Theme) RenderKeyAndValueWithBreak(key, value string) string {
	keyStyle := lipgloss.NewStyle().Foreground(t.SecondaryColor).Bold(true)
	valueStyle := lipgloss.NewStyle().Foreground(t.Gray)
	return keyStyle.Render(key) + "\n" + valueStyle.Render(value)
}

func (t Theme) RenderInputForm(text string) string {
	return lipgloss.NewStyle().PaddingLeft(2).Render(text)
}

func (t Theme) RenderInContainer(text string) string {
	return lipgloss.Style{}.MarginLeft(2).Render(text)
}

func (t Theme) ListStyles() list.Styles {
	styles := list.DefaultStyles()
	styles.StatusBar = styles.StatusBar.
		Padding(0, 0, 1, 0).
		Italic(true).
		Foreground(t.TertiaryColor)
	return styles
}

func (t Theme) ListItemStyles() list.DefaultItemStyles {
	styles := list.NewDefaultItemStyles()
	styles.NormalTitle = styles.NormalTitle.Foreground(t.SecondaryColor)
	styles.NormalDesc = styles.NormalDesc.Foreground(t.BodyColor)

	styles.SelectedTitle = styles.SelectedTitle.
		Border(lipgloss.DoubleBorder(), false, false, false, true).
		Foreground(t.PrimaryColor).
		BorderForeground(t.PrimaryColor).
		Bold(true)
	styles.SelectedDesc = styles.SelectedDesc.
		Border(lipgloss.HiddenBorder(), false, false, false, true).
		Foreground(t.Gray)
	return styles
}

func (t Theme) LoggerStyles() *log.Styles {
	baseStyles := log.DefaultStyles()
	baseStyles.Timestamp = baseStyles.Timestamp.Foreground(t.Gray)
	baseStyles.Levels = map[log.Level]lipgloss.Style{
		log.InfoLevel:  baseStyles.Levels[log.InfoLevel].Foreground(t.InfoColor).SetString("INF"),
		LogNoticeLevel: baseStyles.Levels[LogNoticeLevel].Foreground(t.WarningColor).SetString("NTC"),
		log.WarnLevel:  baseStyles.Levels[log.WarnLevel].Foreground(t.WarningColor).SetString("WRN"),
		log.ErrorLevel: baseStyles.Levels[log.ErrorLevel].Foreground(t.ErrorColor).SetString("ERR"),
		log.DebugLevel: baseStyles.Levels[log.DebugLevel].Foreground(t.EmphasisColor).SetString("DBG"),
		log.FatalLevel: baseStyles.Levels[log.FatalLevel].Foreground(t.ErrorColor).SetString("ERR"),
	}
	baseStyles.Message = baseStyles.Message.Foreground(t.BodyColor)
	baseStyles.Key = baseStyles.Key.Foreground(t.SecondaryColor)
	baseStyles.Value = baseStyles.Value.Foreground(t.Gray)

	return baseStyles
}

func (t Theme) MarkdownStyleJSON() (string, error) {
	tmpl, err := template.New("mdstyles").Parse(mdstylesTemplateJSON)
	if err != nil {
		return "", err
	}

	data := t.markdownTemplateData()
	builder := &strings.Builder{}
	err = tmpl.Execute(builder, data)
	if err != nil {
		return "", err
	}

	return builder.String(), nil
}

func (t Theme) HuhTheme() *huh.Theme {
	theme := huh.ThemeBase()
	theme.FieldSeparator = lipgloss.NewStyle().SetString("\n\n")

	theme.Focused.Base = theme.Focused.Base.BorderStyle(lipgloss.HiddenBorder())
	theme.Focused.Title = theme.Focused.Title.Foreground(t.PrimaryColor).Bold(true)
	theme.Focused.Description = theme.Focused.Description.Foreground(t.BodyColor)
	theme.Focused.ErrorMessage = theme.Focused.ErrorMessage.Foreground(t.ErrorColor)
	theme.Focused.FocusedButton = theme.Focused.FocusedButton.Foreground(t.PrimaryColor).Background(t.TertiaryColor)
	theme.Focused.BlurredButton = theme.Focused.BlurredButton.Foreground(t.SecondaryColor).Background(t.Gray)

	theme.Focused.TextInput.Placeholder = theme.Focused.TextInput.Placeholder.Foreground(t.BodyColor)
	theme.Focused.TextInput.Cursor = theme.Focused.TextInput.Cursor.Foreground(t.SecondaryColor)
	theme.Focused.TextInput.CursorText = theme.Focused.TextInput.Cursor.Foreground(t.SecondaryColor)
	theme.Focused.TextInput.Placeholder = theme.Focused.TextInput.Placeholder.Foreground(t.Gray)
	theme.Focused.TextInput.Text = theme.Focused.TextInput.Text.Foreground(t.BodyColor)
	theme.Focused.TextInput.Prompt = theme.Focused.TextInput.Prompt.Foreground(t.TertiaryColor)

	theme.Blurred.Title = theme.Blurred.Title.Foreground(t.White)
	theme.Blurred.Description = theme.Blurred.Description.Foreground(t.Gray)
	theme.Blurred.ErrorMessage = theme.Blurred.ErrorMessage.Foreground(t.EmphasisColor)
	theme.Blurred.FocusedButton = theme.Blurred.FocusedButton.Foreground(t.SecondaryColor).Background(t.Gray)
	theme.Blurred.BlurredButton = theme.Blurred.BlurredButton.Foreground(t.Gray).Background(t.Gray)

	theme.Blurred.TextInput.Placeholder = theme.Blurred.TextInput.Placeholder.Foreground(t.Gray)
	theme.Blurred.TextInput.Cursor = theme.Blurred.TextInput.Cursor.Foreground(t.Gray)
	theme.Blurred.TextInput.CursorText = theme.Blurred.TextInput.CursorText.Foreground(t.Gray)
	theme.Blurred.TextInput.Placeholder = theme.Blurred.TextInput.Placeholder.Foreground(t.Gray)
	theme.Blurred.TextInput.Text = theme.Blurred.TextInput.Text.Foreground(t.Gray)
	theme.Blurred.TextInput.Prompt = theme.Blurred.TextInput.Prompt.Foreground(t.Gray)

	return theme
}
