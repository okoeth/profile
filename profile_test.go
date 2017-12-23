package main

import (
	"html/template"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestTemplate(t *testing.T) {
	profile := struct {
		Pic, Constants                    string
		Agenda, Authors, Timeline, Trivia []string
	}{
		Pic:       "https://avatars1.githubusercontent.com/u/4511670?s=460&v=4",
		Constants: "Family, Travel, Independence",
		Agenda:    []string{},
		Authors:   []string{"Haruki Murakami", "William Gibson", "Siegfried Lenz"},
		Timeline:  []string{"@stewartbrand", "@jerryweinberg", "@therealbanksy"},
		Trivia:    []string{"Cat or Dog: Cat", " Plane or Train: Train", "Emacs or Vi: Vi"},
	}
	pt := template.New("profile")
	pt, err := pt.ParseFiles("profile.html")
	if err != nil {
		t.Errorf("Error parsing template: %v", err)
	}
	err = pt.ExecuteTemplate(os.Stdout, "profile.html", profile)
	if err != nil {
		t.Errorf("Error parsing template: %v", err)
	}
}
func TestConcurrency(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	c := make(chan string)
	go buildAgenda(c, "Topic1", "Item 1.1", "Item 1.2", "Item 1.3")
	go buildAgenda(c, "Topic2", "Item 2.1", "Item 2.2")
	go buildAgenda(c, "Topic3", "Item 3.1", "Item 3.2", "Item 3.3", "Item 3.4")
	agenda := []string{}
	collectAgenda(c, 3, &agenda)
	if len(agenda) != 9 {
		t.Errorf("Expected 9 agenda items, got %d", len(agenda))
	}
}
