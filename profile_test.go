package profile

import (
	"math/rand"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	agenda := make(chan string)
	go BuildAgenda(agenda, "Topic1", "Item 1.1", "Item 1.2", "Item 1.3")
	go BuildAgenda(agenda, "Topic2", "Item 2.1", "Item 2.2")
	go BuildAgenda(agenda, "Topic3", "Item 3.1", "Item 3.2", "Item 3.3", "Item 3.4")
	PrintAgenda(agenda, 3)
}
