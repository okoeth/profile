package profile

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// PrintHeader renders header of personal profile
func PrintHeader(c string) {
	fmt.Printf("Personal Profile of Oliver Koeth")
}

// PrintConstants renders constants in Markdown
func PrintConstants(c string) {
	fmt.Printf(`
		In hectic times, it's important to have some constants in your life.
		Mine can be summarised as follows: %s`, c)
}

// BuildAgenda inserts agenda items
func BuildAgenda(agenda chan string, topic string, items ...string) {
	for _, item := range items {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		agenda <- "[" + topic + "] " + item
	}
	agenda <- "[" + "] !END"
}

// PrintAgenda inserts agenda items
func PrintAgenda(agenda chan string, topics int) {
	i := 0
	for {
		item := <-agenda
		if strings.HasSuffix(item, "!END") {
			i = i + 1
			if i == topics {
				break
			}
		} else {
			fmt.Printf("* %s\n", item)
		}
	}
}

// OnMyBookshelf build a list of books with links to Amazon in Markdown
func OnMyBookshelf(items ...string) {
	fmt.Println("TODO: Bookshelf")
}

// OnMyPlaylist build a list of artists with links to Spotify in Markdown
func OnMyPlaylist(items ...string) {
	fmt.Println("TODO: Bookshelf")
}

// OnMyWatchlist build a list of TV shows with links to Netflix in Markdown
func OnMyWatchlist(items ...string) {
	fmt.Printf("TODO: Bookshelf")
}

// OnMyTimeline builds a list of Twitter users in Markdown
func OnMyTimeline(items ...string) {
	fmt.Printf("TODO: Bookshelf")
}

// CatOrDog reveals some trivia in Markdown
func CatOrDog() {
	fmt.Printf("### Cat or Dog?")
	fmt.Printf("Cat")
}

// EmacsOrVi reveals some more trivia in Markdown
func EmacsOrVi() {
	fmt.Printf("### Emacs or Vi?")
	fmt.Printf("Vi")
}
