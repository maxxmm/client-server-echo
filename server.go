package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
)

type Note struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Note    string `json:"note"`
}

var sliceOfNotes = []Note{}

func main() {
	e := echo.New()
	e.POST("/save_note", SaveNote)
	e.GET("/watch_notes", WatchNotes)
	log.Fatalln(e.Start("127.0.0.1:4000"))
}

func SaveNote(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Println(err)
		return c.NoContent(500)
	}
	n := Note{}
	err = json.Unmarshal(body, &n)
	if err != nil {
		log.Println(err)
		return c.NoContent(500)
	}
	//newNote := Note{n.Name, n.Surname, n.Note}
	fmt.Println("Имя:", n.Name)
	fmt.Println("Фамилия", n.Surname)
	fmt.Println("Текст заметки", n.Note)
	sliceOfNotes = append(sliceOfNotes, n)
	fmt.Println("Все заметки\n", sliceOfNotes)
	return c.NoContent(200)
}

func WatchNotes(c echo.Context) error {
	//fmt.Println("Got get req")
	return c.JSON(http.StatusOK, &sliceOfNotes)
}
