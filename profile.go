package profile

import (
	"fmt"
)

// PrintHeader renders header of personal profile
func PrintHeader(c string) {
	fmt.Printf("Personal Profile of Oliver Koeth")
}

// PrintConstants renders constants in Markdown
func PrintConstants(c string) {
	fmt.Printf("")
}

// BuildAgenda inserts agenda items
func BuildAgenda(agenda chan string, topic string, items ...string) {

}

// OnMyBookshelf build a list of books with links to Amazon in Markdown
func OnMyBookshelf(items ...string) {

}

// OnMyPlaylist build a list of artists with links to Spotify in Markdown
func OnMyPlaylist(items ...string) {

}

// OnMyWatchlist build a list of TV shows with links to Netflix in Markdown
func OnMyWatchlist(items ...string) {

}

// OnMyTimeline builds a list of Twitter users in Markdown
func OnMyTimeline(items ...string) {

}

// CatOrDog reveals some trivia in Markdown
func CatOrDog() {

}

// EmacsOrVi reveals some more trivia in Markdown
func EmacsOrVi() {

}
