package main

import (
	"os"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const _TimeFormat = time.RFC3339
const _WorkDayStartFile = "deleteme" // TODO: this needs to be relative to the app storage

type Timer struct {
	working  bool
	inOffice bool

	workDayStart time.Time

	txtTimer        *canvas.Text
	btnWorkDayStart *widget.Button
	btnOfficeLeave  *widget.Button
	btnOfficeEnter  *widget.Button
	btnWorkDayEnd   *widget.Button
}

func NewTimer(
	txtTimer *canvas.Text,
	btnWorkDayStart *widget.Button,
	btnOfficeLeave *widget.Button,
	btnOfficeEnter *widget.Button,
	btnWorkDayEnd *widget.Button,
) *Timer {
	txtTimer.Text = "00:00:00"
	txtTimer.Refresh()

	t := &Timer{
		txtTimer:        txtTimer,
		btnWorkDayStart: btnWorkDayStart,
		btnOfficeLeave:  btnOfficeLeave,
		btnOfficeEnter:  btnOfficeEnter,
		btnWorkDayEnd:   btnWorkDayEnd,
	}

	t.WorkDayEnd()

	return t
}

// TODO: I don't like that this and the button updates are in different functions
// another solution would be to simplify `UpdateButtonsBasedOnState`
func (t *Timer) WorkDayBegin() {
	t.working = true

	t.workDayStart = time.Now()
	timeStr := t.workDayStart.Format(_TimeFormat)
	err := os.WriteFile(_WorkDayStartFile, []byte(timeStr), 0644)
	if err != nil {
		panic(err) // TODO:
	}

	t.OfficeEnter()
}

func (t *Timer) WorkDayEnd() {
	t.working = false
	t.inOffice = false

	t.btnWorkDayStart.Enable()
	t.btnOfficeLeave.Disable()
	t.btnOfficeEnter.Disable()
	t.btnWorkDayEnd.Disable()
}

func (t *Timer) OfficeLeave() {
	t.inOffice = false

	t.btnWorkDayStart.Disable()
	t.btnOfficeLeave.Disable()
	t.btnOfficeEnter.Enable()
	t.btnWorkDayEnd.Disable()
}

func (t *Timer) OfficeEnter() {
	t.inOffice = true

	t.btnWorkDayStart.Disable()
	t.btnOfficeLeave.Enable()
	t.btnOfficeEnter.Disable()
	t.btnWorkDayEnd.Enable()
}
