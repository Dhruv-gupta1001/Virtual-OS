package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func calc() {
	// a := app.New()
	w := myApp.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500,320))
	historyString:=""
	output:= ""
	input := widget.NewLabel(output)
	var historyarr []string
	history:=widget.NewLabel(historyString)
	var historyflag bool= true
	w.SetContent(container.NewVBox(
		HomeBtn,
		// DeskBtn,
		input,
		history,
		// widget.NewButton("Hi!", func() {
		// 	input.SetText("Welcome :)")
		// }),
		container.NewGridWithColumns(2,
			widget.NewButton("History",func() {
				if historyflag{
					for i:=len(historyarr)-1;i>=0;i--{
						historyString=historyString+ historyarr[i];
						historyString=historyString+"\n"
					}
					
					historyflag=!historyflag
				}else{
					historyString=""
					historyflag=!historyflag
				}
				history.SetText(historyString)
			}),
			widget.NewButton("Back",func() {
				if len(output)>0{
					output=output[:len(output)-1]
					input.SetText(output)
				}
			}),
		),
		container.NewGridWithColumns(4,
			widget.NewButton("Clear",func() {
				output=""
				input.SetText(output)
			}),
			widget.NewButton("(",func() {
				output=output+"("
				input.SetText(output)
			}),
			widget.NewButton(")",func() {
				output=output+")"
				input.SetText(output)
			}),
			widget.NewButton("/",func() {
				output=output+"/"
				input.SetText(output)
			}),
		),
		container.NewGridWithColumns(4,
			widget.NewButton("7",func() {
				output=output+"7"
				input.SetText(output)
			}),
			widget.NewButton("8",func() {
				output=output+"8"
				input.SetText(output)
			}),
			widget.NewButton("9",func() {
				output=output+"9"
				input.SetText(output)
			}),
			widget.NewButton("*",func() {
				output=output+"*"
				input.SetText(output)
			}),
		),
		container.NewGridWithColumns(4,
			widget.NewButton("4",func() {
				output=output+"4"
				input.SetText(output)
			}),
			widget.NewButton("5",func() {
				output=output+"5"
				input.SetText(output)
			}),
			widget.NewButton("6",func() {
				output=output+"6"
				input.SetText(output)
			}),
			widget.NewButton("-",func() {
				output=output+"-"
				input.SetText(output)
			}),
		),
		container.NewGridWithColumns(4,
			widget.NewButton("1",func() {
				output=output+"1"
				input.SetText(output)
			}),
			widget.NewButton("2",func() {
				output=output+"7"
				input.SetText(output)
			}),
			widget.NewButton("3",func() {
				output=output+"3"
				input.SetText(output)
			}),
			widget.NewButton("+",func() {
				output=output+"+"
				input.SetText(output)
			}),
		),
		container.NewGridWithColumns(2,
			container.NewGridWithColumns(2,
				widget.NewButton("0",func() {
					output=output+"0"
					input.SetText(output)
				}),
				widget.NewButton(".",func() {
					output=output+"."
					input.SetText(output)
				}),
			),
			widget.NewButton("=",func() {
				expression, err := govaluate.NewEvaluableExpression(output);
				
				if err==nil{
					result, err := expression.Evaluate(nil);
					if err==nil{
						ans := strconv.FormatFloat(result.(float64),'f',-1,64)
						stringToAppend := output +"="+ ans
						historyarr = append(historyarr, stringToAppend)
						output= ans

					}else{
						output="error"
					}
				}else{
					output="error"
				}
				// historyString=output + "="+ans

				// output=ans
				input.SetText(output)	
			}),
		),

	))

	// w.ShowAndRun()
	w.Show()
}