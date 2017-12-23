package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	key := os.Getenv("KEY")
	http.HandleFunc("/profile/"+key, func(w http.ResponseWriter, r *http.Request) {
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
		c := make(chan string)
		go buildAgenda(c, "Altemista", "Steering Board Call", "Triple Call", "Sales Activities")
		go buildAgenda(c, "Ensō", "Workshop", "Short Tour")
		go buildAgenda(c, "Innovation Hub", "Project Status", "Weekly Meeting")
		go buildAgenda(c, "This and That", "Hitoe Opportunity Call", "Analyst Briefing")
		collectAgenda(c, 4, &profile.Agenda)
		t := template.New("profile")
		t.ParseFiles("profile.html")
		t.ExecuteTemplate(w, "profile.html", profile)
	})
	http.ListenAndServe("0.0.0.0:8017", nil)
}

func buildAgenda(agenda chan string, topic string, items ...string) {
	for _, item := range items {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		agenda <- "[" + topic + "] " + item
	}
	agenda <- "[" + "] !END"
}

func collectAgenda(c chan string, topics int, agenda *[]string) {
	for i := 0; i < topics; {
		item := <-c
		if strings.HasSuffix(item, "!END") {
			i = i + 1
		} else {
			*agenda = append(*agenda, item)
		}
	}
}
