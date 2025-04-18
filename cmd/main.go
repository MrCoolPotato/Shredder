package main

//MARK: Imports
import (
	"strconv"
	"time"

	"github.com/MrCoolPotato/Shredder/logic"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// MARK: Utility Functions
func isChannelClosed(ch <-chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

// MARK: Main Function
func main() {

	app := tview.NewApplication()
	app.EnableMouse(true)

	// MARK: Input Fields
	fileField := tview.NewInputField()
	fileField.SetLabel("File Path: ")
	fileField.SetPlaceholder("Awaiting selection...")
	fileField.SetPlaceholderTextColor(tcell.ColorWhite)
	fileField.SetDisabled(true)
	fileField.SetFieldTextColor(tcell.ColorWhite)

	passesField := tview.NewInputField()
	passesField.SetLabel("Passes: ")
	passesField.SetText("1")
	passesField.SetFieldTextColor(tcell.ColorWhite)
	passesField.SetAcceptanceFunc(func(textToCheck string, lastChar rune) bool {
		return lastChar >= '0' && lastChar <= '9'
	})

	// MARK: Log View
	logView := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	logView.SetBorder(true).SetTitle(" Status ")
	logView.Write([]byte("Awaiting file selection...\n"))

	// MARK: Form Setup
	form := tview.NewForm().
		AddFormItem(fileField).
		AddFormItem(passesField).
		AddButton("Start Shredding", func() {

			filePath := fileField.GetText()
			if filePath == "" {
				logView.Write([]byte("[red]No file selected.\n"))
				return
			}

			passesStr := passesField.GetText()
			passes, err := strconv.Atoi(passesStr)
			if err != nil || passes <= 0 {
				logView.Write([]byte("[red]Invalid number of passes: " + passesStr + "\n"))
				return
			}

			logView.Write([]byte("[green]Starting file shredding...\n"))
			err = logic.CorruptFile(filePath, passes)
			if err != nil {
				logView.Write([]byte("[red]Error shredding file: " + err.Error() + "\n"))
			} else {
				logView.Write([]byte("[green]File shredded successfully.\n"))
			}
		}).
		AddButton("Select a File", func() {
			done := make(chan struct{})
			var selected string
			var err error
			go func() {
				selected, err = logic.PickFile()
				close(done)
			}()

			go func() {
				time.Sleep(5 * time.Second)
				if !isChannelClosed(done) {
					logView.Write([]byte("[yellow]File dialog active...\n"))
				}
			}()

			<-done
			if err != nil {
				logView.Write([]byte("[red]Error selecting file: " + err.Error() + "\n"))
				return
			}
			if selected != "" {
				fileField.SetText(selected)
			}
		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	// MARK: Layout Setup
	mainLayout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true).
		AddItem(logView, 0, 2, false)

	// MARK: Run Application
	if err := app.SetRoot(mainLayout, true).Run(); err != nil {
		panic(err)
	}
}
