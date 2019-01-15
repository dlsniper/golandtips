// Name: Florin Patan
// Work: Developer Advocate @ JetBrains (GoLand IDE)
// Blog: https://blog.jetbrains.com/go
// Twitter: @GoLandIDE  @JetBrainsJP
// Twitter: @DLSNIPER
// Mail: florin@jetbrains.com
// Gophers Slack: @DLSNIPER (https://invite.slack.golangbridge.org) #goland

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Page struct {
	Title string
	Users []User
}

type User struct {
	ID       int
	Username string
	Job      string
}

func main() {
	fileContents, err := ioutil.ReadFile("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	t, err := template.New("Demo").Parse(string(fileContents))
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sqlx.Open("postgres", "postgres://postgres:postgres@10.0.75.1:5432/goland?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	user := User{}
	err = db.Get(&user, "SELECT * FROM users WHERE id = $1", "Florin")
	if err != nil {
		log.Fatalln(err)
	}

	page := Page{
		Title: "Demo",

		Users: []User{user},
	}

	if err := t.Execute(os.Stdout, page); err != nil {
		log.Fatalln(err)
	}
}
