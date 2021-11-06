package main

import (
	"encoding/json"
	"fmt"

	// "image"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// func main(){
// 	a:=app.New()
// 	w:=a.NewWindow("Weather")

// 	w.Resize(fyne.NewSize(500,320))

// 	//API part
// }

func weatherAPP (w fyne.Window){
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=mumbai&appid=b944d88d96fc530f5fbe2b202b7abe10")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err)
		log.Fatal(err)
	}

	weather, err := UnmarshalWeather(responseData)
	if err!=nil{
		fmt.Print(err)
		log.Fatal(err)
	}

	image := canvas.NewImageFromFile("weather.png")
	image.FillMode = canvas.ImageFillOriginal
	text1 := canvas.NewText(fmt.Sprintf("Country %s",weather.Sys.Country), color.Black)
	text2 := canvas.NewText(fmt.Sprintf("Temp %.2f", weather.Main.Temp), color.Black)
	text3 := canvas.NewText(fmt.Sprintf("Humidity %d", weather.Main.Humidity), color.Black)
	hello := widget.NewLabel("Hello Fyne!")
	// w.SetContent(
	// 	container.NewVBox(
	// 		hello,
	// 		image,
	// 		text1,
	// 		text2,
	// 		text3,
	// 	),
	// )

	weatherContent:= container.NewVBox(
			hello,
			image,
			text1,
			text2,
			text3,
	)

	w.SetContent(container.NewBorder(panel,nil,nil,nil,weatherContent))
	w.Show()
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()


func UnmarshalWeather(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`     
	Weather    []Weather `json:"weather"`   
	Base       string    `json:"base"`      
	Main       Main      `json:"main"`      
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`      
	Clouds     Clouds    `json:"clouds"`    
	Dt         int64     `json:"dt"`        
	Sys        Sys       `json:"sys"`       
	Timezone   int64     `json:"timezone"`  
	ID         int64     `json:"id"`        
	Name       string    `json:"name"`      
	Cod        int64     `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
}

type Sys struct {
	Type    int64  `json:"type"`   
	ID      int64  `json:"id"`     
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type Weather struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
}
