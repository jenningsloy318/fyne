package widget_test

import (
	"testing"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/stretchr/testify/assert"
)

func TestButton_MinSize(t *testing.T) {
	button := widget.NewButton("Hi", nil)
	min := button.MinSize()

	assert.Greater(t, min.Width, theme.InnerPadding())
	assert.Greater(t, min.Height, theme.InnerPadding())
}

func TestButton_SetText(t *testing.T) {
	button := widget.NewButton("Hi", nil)
	min1 := button.MinSize()

	button.SetText("Longer")
	min2 := button.MinSize()

	assert.Greater(t, min2.Width, min1.Width)
	assert.Equal(t, min2.Height, min1.Height)
}

func TestButton_MinSize_Icon(t *testing.T) {
	button := widget.NewButton("Hi", nil)
	min1 := button.MinSize()

	button.SetIcon(theme.CancelIcon())
	min2 := button.MinSize()

	assert.Greater(t, min2.Width, min1.Width)
	assert.Equal(t, min2.Height, min1.Height)
}

func TestButton_Cursor(t *testing.T) {
	button := widget.NewButton("Test", nil)
	assert.Equal(t, desktop.DefaultCursor, button.Cursor())
}

func TestButton_Tapped(t *testing.T) {
	tapped := make(chan bool)
	button := widget.NewButton("Hi", func() {
		tapped <- true
	})

	go test.Tap(button)
	func() {
		select {
		case <-tapped:
		case <-time.After(1 * time.Second):
			assert.Fail(t, "Timed out waiting for button tap")
		}
	}()
}

func TestButton_Disable(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	tapped := false
	button := widget.NewButtonWithIcon("Test", theme.HomeIcon(), func() {
		tapped = true
	})
	button.Disable()
	w := test.NewWindow(button)
	defer w.Close()

	test.Tap(button)
	assert.False(t, tapped, "Button should have been disabled")
	test.AssertImageMatches(t, "button/disabled.png", w.Canvas().Capture())

	button.Enable()
	test.AssertImageMatches(t, "button/initial.png", w.Canvas().Capture())
}

func TestButton_Enable(t *testing.T) {
	tapped := false
	button := widget.NewButton("Test", func() {
		tapped = true
	})

	button.Disable()
	test.Tap(button)
	assert.False(t, tapped, "Button should be disabled")

	button.Enable()
	test.Tap(button)
	assert.True(t, tapped, "Button should have been re-enabled")
}

func TestButton_Disabled(t *testing.T) {
	button := widget.NewButton("Test", func() {})
	assert.False(t, button.Disabled())
	button.Disable()
	assert.True(t, button.Disabled())
	button.Enable()
	assert.False(t, button.Disabled())
}

func TestButton_LowImportance(t *testing.T) {
	test.NewTempApp(t)

	b := &widget.Button{Text: "Text", Icon: theme.HomeIcon(), Importance: widget.LowImportance}
	w := test.NewWindow(b)
	defer w.Close()

	test.AssertImageMatches(t, "button/low_importance.png", w.Canvas().Capture())

	b.Disable()
	test.AssertImageMatches(t, "button/low_importance_disabled.png", w.Canvas().Capture())
}

func TestButton_Hover(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	b := widget.NewButtonWithIcon("Test", theme.HomeIcon(), func() {})
	w := test.NewWindow(b)
	defer w.Close()

	if !fyne.CurrentDevice().IsMobile() {
		test.MoveMouse(w.Canvas(), fyne.NewPos(5, 5))
		test.AssertImageMatches(t, "button/hovered.png", w.Canvas().Capture())
	}

	b.Importance = widget.HighImportance
	b.Refresh()
	if !fyne.CurrentDevice().IsMobile() {
		test.AssertImageMatches(t, "button/high_importance_hovered.png", w.Canvas().Capture())

		test.MoveMouse(w.Canvas(), fyne.NewPos(0, 0))
		b.Refresh()
	}
	test.AssertImageMatches(t, "button/high_importance.png", w.Canvas().Capture())

	b.Importance = widget.MediumImportance
	b.Refresh()
	test.AssertImageMatches(t, "button/initial.png", w.Canvas().Capture())
}

func TestButton_Layout(t *testing.T) {
	test.NewTempApp(t)

	for name, tt := range map[string]struct {
		text      string
		icon      fyne.Resource
		alignment widget.ButtonAlign
		placement widget.ButtonIconPlacement
	}{
		"text_only_center_leading": {
			text:      "Test",
			alignment: widget.ButtonAlignCenter,
			placement: widget.ButtonIconLeadingText,
		},
		"text_only_center_trailing": {
			text:      "Test",
			alignment: widget.ButtonAlignCenter,
			placement: widget.ButtonIconTrailingText,
		},
		"text_only_leading_leading": {
			text:      "Test",
			alignment: widget.ButtonAlignLeading,
			placement: widget.ButtonIconLeadingText,
		},
		"text_only_leading_trailing": {
			text:      "Test",
			alignment: widget.ButtonAlignLeading,
			placement: widget.ButtonIconTrailingText,
		},
		"text_only_trailing_leading": {
			text:      "Test",
			alignment: widget.ButtonAlignTrailing,
			placement: widget.ButtonIconLeadingText,
		},
		"text_only_trailing_trailing": {
			text:      "Test",
			alignment: widget.ButtonAlignTrailing,
			placement: widget.ButtonIconTrailingText,
		},
		"text_only_multiline": {
			text:      "Test\nLine2",
			alignment: widget.ButtonAlignCenter,
		},
		"icon_only_center_leading": {
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignCenter,
			placement: widget.ButtonIconLeadingText,
		},
		"icon_only_center_trailing": {
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignCenter,
			placement: widget.ButtonIconTrailingText,
		},
		"icon_only_leading_leading": {
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignLeading,
			placement: widget.ButtonIconLeadingText,
		},
		"icon_only_leading_trailing": {
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignLeading,
			placement: widget.ButtonIconTrailingText,
		},
		"icon_only_trailing_leading": {
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignTrailing,
			placement: widget.ButtonIconLeadingText,
		},
		"icon_only_trailing_trailing": {
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignTrailing,
			placement: widget.ButtonIconTrailingText,
		},
		"text_icon_center_leading": {
			text:      "Test",
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignCenter,
			placement: widget.ButtonIconLeadingText,
		},
		"text_icon_center_trailing": {
			text:      "Test",
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignCenter,
			placement: widget.ButtonIconTrailingText,
		},
		"text_icon_leading_leading": {
			text:      "Test",
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignLeading,
			placement: widget.ButtonIconLeadingText,
		},
		"text_icon_leading_trailing": {
			text:      "Test",
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignLeading,
			placement: widget.ButtonIconTrailingText,
		},
		"text_icon_trailing_leading": {
			text:      "Test",
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignTrailing,
			placement: widget.ButtonIconLeadingText,
		},
		"text_icon_trailing_trailing": {
			text:      "Test",
			icon:      theme.CancelIcon(),
			alignment: widget.ButtonAlignTrailing,
			placement: widget.ButtonIconTrailingText,
		},
	} {
		t.Run(name, func(t *testing.T) {
			button := &widget.Button{
				Text:          tt.text,
				Icon:          tt.icon,
				Alignment:     tt.alignment,
				IconPlacement: tt.placement,
			}

			window := test.NewWindow(button)
			defer window.Close()
			window.Resize(button.MinSize().Max(fyne.NewSize(150, 200)))

			test.AssertRendersToMarkup(t, "button/layout_"+name+".xml", window.Canvas())
		})
	}
}

func TestButton_ChangeTheme(t *testing.T) {
	test.NewTempApp(t)

	b := widget.NewButton("Test", func() {})
	w := test.NewWindow(b)
	defer w.Close()
	w.Resize(fyne.NewSize(200, 200))
	b.Resize(b.MinSize())
	b.Move(fyne.NewPos(10, 10))

	test.AssertImageMatches(t, "button/theme_initial.png", w.Canvas().Capture())

	test.WithTestTheme(t, func() {
		b.Resize(b.MinSize())
		b.Refresh()
		time.Sleep(100 * time.Millisecond)
		test.AssertImageMatches(t, "button/theme_changed.png", w.Canvas().Capture())
	})
}

func TestButtonCompatImportance(t *testing.T) {
	// Test backward compatibility of widget.Importance
	var imp widget.ButtonImportance = widget.HighImportance

	btn := widget.NewButton("test", func() {})
	btn.Importance = imp
}

func TestButtonSuccess(t *testing.T) {
	test.NewTempApp(t)
	test.ApplyTheme(t, test.Theme())

	b := widget.NewButtonWithIcon("Test", theme.HomeIcon(), func() {})
	w := test.NewWindow(b)
	defer w.Close()

	b.Importance = widget.SuccessImportance
	b.Refresh()

	test.AssertImageMatches(t, "button/success_importance.png", w.Canvas().Capture())
}
