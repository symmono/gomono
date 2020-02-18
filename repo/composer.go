/*
 * This file is part of the GoMono application
 *
 * (c) Bertin van den Ham <bertinvandenham@gmail.com>
 *
 * Compser struct that contains the composer mono repo's data
 */

package repo

type Composer struct {
	Name        string
	Type		string
	Repository	map[string] string
	Extra		map[string] map[string] string
}