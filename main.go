/*
 * GoMono
 *
 * (c) Bertin van den Ham <bertinvandenham@gmail.com>
 *
 * Main file to startup to gomono application
 */

package main

import (
	"fmt"
	"github.com/symmono/gomono/repo"
)

func init() {
	fmt.Println("Init GoMono application")
}

func main() {
	var composer repo.Composer
	composer.Name = "symmono/page-bundle"
	composer.Type = "symfony-bundle"

	fmt.Println("name: ", composer.Name)
	fmt.Println("repo: ", composer.Type)

	c := repo.Composer{
		Name:       "symmono/page-bundle",
		Type:       "symfony-bundle",
		Repository: nil,
		Extra:      nil,
	}

	fmt.Println(c)
}