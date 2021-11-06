package main

import (

	// "fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func gallery() {
	// a := app.New()
	w := myApp.NewWindow("Gallery")
	w.Resize(fyne.NewSize(500,320))
	tabs := container.NewAppTabs()
	// var picsarr []string
	
	root_src:="D:\\VIT\\Micro\\LAB\\DA3\\Q1 Factorial\\images"
	files, err := ioutil.ReadDir(root_src)
    if err != nil {
        log.Fatal(err)
    }
	for _, file:=range files{
		// fmt.Println(file.Name(),file.IsDir())
		if !file.IsDir(){
			file_extension := strings.Split(file.Name(), ".")[1]
			// fmt.Println(file_extension)
			
			if file_extension=="png" || file_extension=="jpeg"{
				// file_location:=root_src+"\\"+file.Name()
				// fmt.Println(file_location)
				// picsarr = append(picsarr, file_location)
				image := canvas.NewImageFromFile(root_src+"\\"+file.Name())
				// image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(file.Name(),image))
			}
		}
	}

	// for i:=1;i<len(picsarr);i++{
	// 	fmt.Println(picsarr[i])
	// }
	tabs.SetTabLocation(container.TabLocationLeading)
	// tabs.Append(container.NewTabItem("Image",canvas.NewImageFromFile(picsarr[1])))
	w.SetContent(tabs)
	

	// w.ShowAndRun()
	w.Show()
}