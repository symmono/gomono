// GOMONO - monorepo tool written in GOLang

package main

import (
	"fmt"
	"github.com/symmono/gomono/repo"
)

func init() {
	fmt.Println("Init gomono application")
}

func main() {
	fmt.Println("time to buld the parmas")

	repo.SearchComposer("composer.json")
}