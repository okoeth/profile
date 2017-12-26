package main

import (
	"html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	i := Insights{
		Pic:       "A",
		Constants: "1, 2, 3",
		QandAs:    []QandA{{Q: "Q1", A: "A1"}, {Q: "Q1", A: "A1"}, {Q: "Q1", A: "A1"}},
		Trivia:    []string{"x", "y", "z"},
	}
	pt := template.New("i")
	pt, err := pt.ParseFiles("insights.html")
	if err != nil {
		t.Errorf("Error parsing template: %v", err)
	}
	err = pt.ExecuteTemplate(os.Stdout, "insights.html", i)
	if err != nil {
		t.Errorf("Error parsing template: %v", err)
	}
}
