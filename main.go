package main

import (
	"fyne.io/fyne/v2"
	// "fyne.io/fyne"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)
var myApp fyne.App = app.New()

var myWindow fyne.Window=myApp.NewWindow("My OS")

var weatherBtn fyne.Widget 
var calBtn fyne.Widget 
var galleryBtn fyne.Widget 
var editorBtn fyne.Widget 

var HomeBtn fyne.Widget

var img fyne.CanvasObject

var panel *fyne.Container

func main(){

	myApp.Settings().SetTheme(theme.LightTheme())
	img= canvas.NewImageFromFile("D:\\DEV-20211013T102902Z-001\\Hackathon\\wallpaper.jpg")

	weatherBtn=widget.NewButtonWithIcon("Weather",theme.InfoIcon(), func() {
		weatherAPP(myWindow)
	})

	calBtn=widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		calc()
	})

	galleryBtn=widget.NewButtonWithIcon("Gallery", theme.StorageIcon(), func() {
		gallery()
	})

	editorBtn=widget.NewButtonWithIcon("Editor", theme.DocumentIcon(), func() {
		editor()
	})

	HomeBtn =widget.NewButtonWithIcon("Home",theme.HomeIcon(),func() {
		myWindow.SetContent(container.NewBorder(panel,nil,nil,nil,img))
	})

	panel = container.NewVBox(container.NewGridWithColumns(5,HomeBtn,weatherBtn,calBtn,galleryBtn,editorBtn))
	// HomeBtn = widget.NewButtonWithIcon()

	myWindow.SetContent(
		container.NewBorder(panel,nil,nil,nil,img),
	)

	myWindow.CenterOnScreen()

	myWindow.Resize(fyne.NewSize(1280,720))

	myWindow.ShowAndRun()
}


// package main

// import (
// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/theme"
// 	"fyne.io/fyne/v2/widget"
// )

//     var myApp fyne.App = app.New()
//     var myWindow fyne.Window = myApp.NewWindow("Virtual Operating System")
//     var btn1 fyne.Widget
// 	var btn2 fyne.Widget
// 	var btn3 fyne.Widget
// 	var btn4 fyne.Widget
// 	var img fyne.CanvasObject
// 	var DeskBtn fyne.Widget 
// 	var panelContent *fyne.Container

	 
// func main() {
// 	//myApp.Settings().SetTheme(theme.LightTheme())
// 	// img := canvas.NewImageFromFile("C:\\Users\\HP\\OneDrive\\Desktop\\VOS\\Wallpaper.jpg")

//     img= canvas.NewImageFromFile("wallpaper.jpg")
// 	btn1=widget.NewButtonWithIcon("Calculator",theme.ContentAddIcon(),func(){
// 		calc()
// 	}) 


// 	btn2=widget.NewButtonWithIcon("GalleryApp",theme.StorageIcon(),func(){
// 		gallery()
// 	}) 

// 	btn3=widget.NewButtonWithIcon("Text Editor",theme.DocumentIcon(),func(){
// 		editor()
// 	})	
// 	btn4=widget.NewButtonWithIcon("Weather App",theme.InfoIcon(),func(){
// 		weatherAPP(myWindow)
// 	}) 

// 	DeskBtn = widget.NewButtonWithIcon("Home",theme.HomeIcon(),func(){
// 		myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
// 	})

// 	panelContent=container.NewVBox((container.NewGridWithColumns(5,DeskBtn,btn1,btn2,btn3,btn4)))

// 	myWindow.Resize(fyne.NewSize(1280,720))
// 	myWindow.CenterOnScreen();

// 	myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img),)
	

// 	 myWindow.ShowAndRun()
	
// }