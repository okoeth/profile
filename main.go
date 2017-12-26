package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type (
	// QandA carries a question with an answer
	QandA struct {
		Q, A string
	}

	// Insights carry some personal information
	Insights struct {
		Pic, Constants, SelectedTrivia string
		QandAs                         []QandA
		Trivia                         []string
	}
)

// Jeder von uns hat besondere Interessen, besondere Eigenschaften und etwas, das ihn auszeichnet.
// Als technikbegeisterter Mensch möchte Oliver mit dem kleinen Programm unten zeigen, daß
// auch nicht Techniker die wesentlichen Informationen aus übersichtlich geschriebenen
// Code ziehen können.
// Gerne hätte sich Oliver auf ein paar Fragen nach Lieblingsautor und  -musiker beschränkt. Seine
// Lektorin war allerdings der Meinung, dass die folgenden Fragen für die Leser viel interessanter
// sind.

func main() {
	key := os.Getenv("KEY")
	http.HandleFunc("/personalinsights/"+key, func(w http.ResponseWriter, r *http.Request) {
		i := Insights{
			Pic:       "https://avatars1.githubusercontent.com/u/4511670?s=460&v=4",
			Constants: "Familie, Reisen, Unabhängigkeit",
			QandAs: []QandA{
				{Q: "Wie schaffe ich mein hohes Arbeitspensum?",
					A: "Zum einen steht auch bei mir der Input in Korrelation zum Output, d.h. lange Arbeitstage. " +
						"Aber bei den vielen Themen, die ich bearbeiten darf, gibt es auch immer die Chance Teilergebnisse " +
						"wiederzuverwenden, so daß der Output manchmal nach mehr aussieht als es eigentlich ist."},
				{Q: "Woher nehme ich die Energie?",
					A: "Zum Großteil aus meiner Begeisterung für Technik, aber auch aus einem (angeborenen?) " +
						"Qualitätsanspruch. 7-8 Stunden Schlaf und regelmäßige Tagesabläufe helfen ebenfalls."},
				{Q: "Wie bekommst Du Job und Familie unter einen Hut?",
					A: "Das wichtigste ist, mit der Familie einen gemeinsame Sicht auf die Prioritäten zu haben. " +
						"Es hilft mir auch, die Tage oder Tagesabschnitte klar zu trennen: Job oder Familie. " +
						"Und nicht ein Mischmasch, bei dem man keiner Seite wirklich gerecht wird."}},
			Trivia: []string{"Katze oder Hund: Katze", "Flugzeug oder Bahn: Bahn", "Emacs oder Vi: Vi"},
		}
		c := make(chan string)
		for _, t := range i.Trivia {
			go func(c chan string, t string) {
				fmt.Printf("Trivia: %s", t)
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				c <- t
			}(c, t)
		}
		i.SelectedTrivia = <-c
		t := template.New("i")
		t.ParseFiles("insights.html")
		t.ExecuteTemplate(w, "insights.html", i)
	})
	http.ListenAndServe("0.0.0.0:8017", nil)
}
