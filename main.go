package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fyneApp := app.NewWithID("work-time-track")

	fyneWindow := fyneApp.NewWindow("Work Time Track")
	fyneWindow.Resize(fyne.NewSize(400, 600))

	var timer *Timer

	txtTimer := canvas.NewText("Error", theme.Color(theme.ColorNameForeground))
	txtTimer.TextSize = theme.TextHeadingSize()
	txtTimer.TextStyle.Bold = true
	// lblTimer.Alignment = fyne.TextAlignCenter

	btnWorkDayStart := widget.NewButton("Begin Work Day", func() { timer.WorkDayBegin() })
	btnOfficeLeave := widget.NewButton("Leave Office", func() { timer.OfficeLeave() })
	btnOfficeEnter := widget.NewButton("Go Back to Office", func() { timer.OfficeEnter() })
	btnWorkDayEnd := widget.NewButton("End Work Day", func() { timer.WorkDayEnd() })

	timer = NewTimer(txtTimer, btnWorkDayStart, btnOfficeLeave, btnOfficeEnter, btnWorkDayEnd)

	fyneWindow.SetContent(
		container.NewGridWithColumns(
			1,
			container.NewCenter(txtTimer),
			widget.NewLabel(""),
			btnWorkDayStart,
			widget.NewLabel(""),
			btnOfficeLeave,
			widget.NewLabel(""),
			btnOfficeEnter,
			widget.NewLabel(""),
			btnWorkDayEnd,
		),
	)

	fyneWindow.ShowAndRun()
}
