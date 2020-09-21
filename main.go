package main

import (
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
  "fmt"
)

func sendPlayerInput(playerInput string) string {
  var solvedSudoku = "4,7,2,1,6,8,3,4,2,9,1,2,3,4,8,5,7,5,9,2,6,1,3,7,1,5,9,4,4,2,7,9,8,5,9,1,2,8,7,3,4,5,5,3,7,2,6,4,6,3,1"
    if playerInput != solvedSudoku {
    fmt.Printf("Tienes fallos, vuelve a intentarlo.")
    return "Tienes fallos, vuelve a intentarlo."
    }
    return "¡Eres un fenómeno como Diego y Javi!"
    }

func main() {

  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "sudoku",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(sendPlayerInput)
  app.Run()
}
