package uiparser

import (
	"foldest-go/utils"

	"github.com/andlabs/ui"
)

var (
	// fontButton *ui.FontButton
	// alignment *ui.Combobox
	font *ui.FontDescriptor

	attrstr *ui.AttributedString
)

func appendWithAttributes(what string, attrs ...ui.Attribute) {
	start := len(attrstr.String())
	end := start + len(what)
	attrstr.AppendUnattributed(what)
	for _, a := range attrs {
		attrstr.SetAttribute(a, start, end)
	}
}

func makeAttributedString() {
	attrstr = ui.NewAttributedString(
		"Drawing strings with package ui is done with the ui.AttributedString and ui.DrawTextLayout objects.\n" +
			"ui.AttributedString lets you have a variety of attributes: ")

	appendWithAttributes("font family", ui.TextFamily("Courier New"))
	attrstr.AppendUnattributed(", ")

	appendWithAttributes("font size", ui.TextSize(18))
	attrstr.AppendUnattributed(", ")

	appendWithAttributes("font weight", ui.TextWeightBold)
	attrstr.AppendUnattributed(", ")

	appendWithAttributes("font italicness", ui.TextItalicItalic)
	attrstr.AppendUnattributed(", ")

	appendWithAttributes("font stretch", ui.TextStretchCondensed)
	attrstr.AppendUnattributed(", ")

	appendWithAttributes("text color", ui.TextColor{0.75, 0.25, 0.5, 0.75})
	attrstr.AppendUnattributed(", ")

	appendWithAttributes("text background color", ui.TextBackground{0.5, 0.5, 0.25, 0.5})
	attrstr.AppendUnattributed(", ")

	appendWithAttributes("underline style", ui.UnderlineSingle)
	attrstr.AppendUnattributed(", ")

	attrstr.AppendUnattributed("and ")
	appendWithAttributes("underline color",
		ui.UnderlineDouble,
		ui.UnderlineColorCustom{1.0, 0.0, 0.5, 1.0})
	attrstr.AppendUnattributed(". ")

	attrstr.AppendUnattributed("Furthermore, there are attributes allowing for ")
	appendWithAttributes("special underlines for indicating spelling errors",
		ui.UnderlineSuggestion,
		ui.UnderlineColorSpelling)
	attrstr.AppendUnattributed(" (and other types of errors) ")

	attrstr.AppendUnattributed("and control over OpenType features such as ligatures (for instance, ")
	appendWithAttributes("afford", ui.OpenTypeFeatures{
		ui.ToOpenTypeTag('l', 'i', 'g', 'a'): 0,
	})
	attrstr.AppendUnattributed(" vs. ")
	appendWithAttributes("afford", ui.OpenTypeFeatures{
		ui.ToOpenTypeTag('l', 'i', 'g', 'a'): 1,
	})
	attrstr.AppendUnattributed(").\n")

	attrstr.AppendUnattributed("Use the controls opposite to the text to control properties of the text.")
}

type areaHandler struct{}

func (areaHandler) Draw(a *ui.Area, p *ui.AreaDrawParams) {
	tl := ui.DrawNewTextLayout(&ui.DrawTextLayoutParams{
		String:      attrstr,
		DefaultFont: font,
		Width:       p.AreaWidth,
		Align:       ui.DrawTextAlignLeft,
	})
	defer tl.Free()
	p.Context.Text(tl, 0, 0)
}

func (areaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	// do nothing
}

func (areaHandler) MouseCrossed(a *ui.Area, left bool) {
	// do nothing
}

func (areaHandler) DragBroken(a *ui.Area) {
	// do nothing
}

func (areaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	// reject all keys
	return false
}

func SetupUI() {
	font = &ui.FontDescriptor{
		Family:  ui.TextFamily("Courier New"),
		Size:    ui.TextSize(12),
		Weight:  ui.TextWeightNormal,
		Italic:  ui.TextItalicNormal,
		Stretch: ui.TextStretchNormal,
	}

	makeAttributedString()
	window := ui.NewWindow("Foldest", 400, 400, true)
	window.SetMargined(true)
	window.OnClosing(func(*ui.Window) bool {
		window.Destroy()
		ui.Quit()
		return false
	})
	ui.OnShouldQuit(func() bool {
		window.Destroy()
		return true
	})

	canvas := ui.NewVerticalBox()
	canvas.SetPadded(true)
	window.SetChild(canvas)

	startBtn := ui.NewButton("Start")
	startBtn.OnClicked(func(*ui.Button) {
		utils.Start()
	})
	canvas.Append(startBtn, false)

	// outputLabel := ui.NewLabel("")
	// outputArea := ui.NewScrollingArea(AreaHandlerInstinc{}, 300, 200)
	outputArea := ui.NewArea(areaHandler{})
	canvas.Append(outputArea, true)

	window.Show()
}
