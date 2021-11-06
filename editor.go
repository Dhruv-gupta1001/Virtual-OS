package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	// "log"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func editor(){
	// a:= app.New()
	w:=myApp.NewWindow("Editor")

	w.Resize(fyne.NewSize(500,320))

	count := 0

	content := container.NewVBox(
		widget.NewLabel("GoLang Editor"),
	)

	content.Add(widget.NewButton("Add new file", func() {
		content.Add(widget.NewLabel(fmt.Sprintf("New File %d",count)))
		count++
	}))

	// content.Add(widget.NewButton("Add more items", func() {
	// 	content.Add(widget.NewLabel("Added"))
	// }))


	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	input.Resize(fyne.NewSize(800,200))

	saveBtn:= widget.NewButton("Save file", func() {
		saveFileDialog:= dialog.NewFileSave(
			 func(uc fyne.URIWriteCloser, _ error){
				textData:=[]byte(input.Text)
				uc.Write(textData)
			},w)

			saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1))
			saveFileDialog.Show()
	})

	openBtn:= widget.NewButton("Open File",func() {
		openFileDialog:=dialog.NewFileOpen(
				func (r fyne.URIReadCloser,_ error )  {
					readData ,_ := ioutil.ReadAll(r)
					output:= fyne.NewStaticResource("New File",readData)
					viewData:=widget.NewMultiLineEntry()

					viewData.SetText(string(output.StaticContent))

					w:=fyne.CurrentApp().NewWindow(
						string(output.StaticName),
					)
					saveBtn1:=  widget.NewButton("Save",func() {

				})
					w.SetContent(container.NewVBox(container.NewScroll(viewData),
					saveBtn1))

					w.Resize(fyne.NewSize(400,400))

					w.Show()
				},w)
				
				openFileDialog.SetFilter(
					storage.NewExtensionFileFilter([]string{".txt"}))
				openFileDialog.Show()

				
				// /w.SetContent(saveBtn1)

	})

	w.SetContent(container.NewVBox(
		content,
		input,
		container.NewHBox(
			saveBtn,
			openBtn,
		),
	))

	// w.ShowAndRun()
	w.Show()
}