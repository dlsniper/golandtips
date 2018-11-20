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
	"os"

	// you must import this package manually, when you type _ " then the completion should show up for the rest of the packages
	_ "github.com/lib/pq"
)

func main() {

	// start from typing ReadFile and show how awesome our completion is
	// - we can show here that we don't need the package to be imported
	// - show that we can show the quick documentation and quick definition from here
	// - show that we can use the Example link in the documentation to create a new scratch file
	// - show that we can create a private gist with the file contents, get the link and share it back to people
	ioutil.ReadFile("tpl.gohtml")

	// file is the first return value of the ReadFile call above. show that we can convert string to []byte  using the quickfix
	template.New("Demo").Parse(file)

	// use the sqlxconnect live template

	// to get this variable created, first go to the page := and then create the struct, assign the User{user} part
	// then create the User{} and then create the local variable user
	user := User{}
	// use .rrv and then replace return with Fatalln(err) - choose log.Fatalln()

	// say that you don't have a database running and you'd like to have one before trying the code
	// it's a good opportunity to introduce the Docker and (failing) Database plugins
	// show we support DB query completion with the tables from above
	db.Get(&user, "SELECT * FROM users WHERE ID = $1", "Florin")

	// first type this type all things below
	// then use the quickfix to create the missing Page type
	page := Page{
		// then create the missing field Title
		Title: "Demo",

		// Then create the missing type User
		// Then create the missing field Users
		Users: []User{user},
	}

	// use .rrv then same Fatalln
	t.Execute(os.Stdout, page)
}

// after we cover all contents in main(), ask about k8s plugin users
// show support for quick documentation, completion, and navigate to support
// show kdep live template in a yaml file
// show that and helm support (right click on k8s/helm/postgresql/templates/svc.yaml and use Kubernetes | Helm Template)

// then show terraform support