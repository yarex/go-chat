package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App        fyne.App
	MainWindow fyne.Window
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
}

func main() {
	var myApp Config

	fyneApp := app.NewWithID("dev.yarex.go-chat")
	myApp.App = fyneApp

	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	myApp.MainWindow = fyneApp.NewWindow("Go-Chat")
	myApp.MainWindow.Resize(fyne.NewSize(600, 400))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUI()

	myApp.MainWindow.ShowAndRun()
}
