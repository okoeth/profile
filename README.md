# Einleitung 

Dies ist eine kurze Beschreibung zu meiner in [Golang](http://golang.org) entwickelten
*Personal Insights* Seite.


# Interessante Aspekte von Golang

Golang ist inzwischen auch schon [8 Jahre](https://blog.golang.org/8years) alt. Ursprünglich 
von Google entwickelt, hat sich die Sprache zunächst als Nachfolger von C und C++ für die 
Systemprogrammierung durchgesetzt. Bekannte Beispiele für Systeme, die in Golang entwickelt 
wurden sind [Docker](http://www.docker.com) und [Kubernetes](http://kubernetes.io).

Inzwischen setzt sich Golang aber auch immer weiter im Bereich der Anwendungsentwicklung durch
und Firmen wie [Dropbox nutzen Golang](https://www.youtube.com/watch?v=5doOcaMXx08), aber auch 
Anwedungen wie [Mattermost](https://github.com/mattermost/mattermost-server)

Was zeichnet also Golang im Vergleich zu anderen Programmiersprachen aus? Aus meiner Sicht kann 
das mit drei Begriffen "einfach", "vollständig" und "modern" gut beschrieben werden:


## Einfach
Die Syntax von Golang ist frei von unnötigen Deklarationen und Steuerzeichen, so daß Prgramme
schnell, einfach und übersichtlich geschrieben werden können. Dennoch ist die Sprache typsicher.
In den Beispiel unten wird eine String Variable implizit aus dem Ergebnis der Funktion
initialisiert:
```
key := os.Getenv("KEY")
```

Anstatt auf komplexe Objekthierarchien setzt Golang auf Interfaces. Es ist typisch für
Golang, daß Interfaces nur sehr wenige Methoden haben (kein Vergleich mit Java) und eines
der [Golang Sprichworte](http://go-proverbs.github.io/) ist auch: Je größer das Interface, 
desto schwächer die Abstraktion. Wie wahr.

Ebenfalls zur Übersichtlichkeit tragen anonyme Funktionen bei. In meinem Beispiel, kann dadurch
sogar auf die Verwendung von Interfaces bei der Implementerung der Serverfunktionalität
verzichtet werden.
```
http.HandleFunc("/personalinsights/"+key, func(w http.ResponseWriter, r *http.Request) {
    // Implementierung...
})
```

Das spannendste Element sind für mich aber Go Routinen, mit denen nebenläufige Funktionen
entwickelt werden können, ohne daß man sich mit Betriebssystemaspekten wie Threads
herumschlagen muß. Die nebenläufigen Funktionen kommunizieren dabei über Channels und
vermeiden so konkurrierende Zugriff auf gemeinsame Variablen.
Auch wenn man nicht mit Golang programmiert, kann man [hier](https://blog.golang.org/concurrency-is-not-parallelism) 
viel über Nebenläufigkeit lernen.
In meinem Beispiel werden Go Routinen bemüht um einen Trivia
Aspekt zufällig auszuwählen (ja, etwas sehr bemüht, das ginge auf anders ;-)
```
c := make(chan string)
for _, t := range i.Trivia {
	go func(c chan string, t string) {
		fmt.Printf("Trivia: %s", t)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		c <- t
	}(c, t)
}
i.SelectedTrivia = <-c
```


## Vollständig
Neben der gelungenen Konzeption der Sprache selbst, bietet die Standardbibliothek von
Golang aber auch (fast) alle Funktionen, die man für die Programmierung von Backend
Services benötigt. Das hier vorgestellte Beispiel benötigt keine einzige externe
Library für Test, Templating, oder Netzwerkkommunikation.
Zudem steht die Standardbibliothek als [Quelltext](https://github.com/golang/go) 
zur Verfügung.
```
t := template.New("i")
t.ParseFiles("insights.html")
t.ExecuteTemplate(w, "insights.html", i)
```


## Modern (aber nicht immer neu!)
Was Golang -- zum Beispiel im Vergleich zu Java -- neben vielen extenen Bibliothekten 
ebenfalls nicht benötigt ist eine Laufzeitumgebung. Damit stehen einem Java Server 
Prozess, der ohne Anwendung (nur für die virtuelle Maschine und den Application Server)
schon 150MB Speicher benötigt, ein in Golang entwickelt Service mit einem minimalen 
Speicherbedarf von 4MB entgegen.

In Zeiten von Microservices und Container Orchestrierung ein nicht zu vernachlässigender
und kommerziell durchaus relevanter Aspekt. 

Während Golang also einerseits bestens in das moderne Cloud Zeitalter paßt, ist aber bei weitem nicht 
alles neu (i.S.v. es sind viele Weisheiten unserer Industrie in Golang erhalten). Das fängt zum einen
schon bei dem Team an, das Golang entwickelt hat: Robert Griesemer, Rob Pike und Ken Thompson (ja, *der* 
Ken Thompson ;-)

Viele Aspekte wie zum Beispiel der [Golang Assembler](https://golang.org/doc/asm), der viel zur 
Portabilität von Golang ganz ohne Virtuelle Maschine beiträgt, sind direkt auf 
[Plan9](https://en.wikipedia.org/wiki/Plan_9_from_Bell_Labs) zurückzuführen.

Der elegente Ansatz zur Erstellung von nebenfäufigen Programmen geht sogar bis auf einen 
[Artikel von C.A.R. Hoare](http://spinroot.com/courses/summer/Papers/hoare_1978.pdf) aus dem Jahr 
1978 zurück.

Alles in allem für mich eine faszinierende Kombination aus Tradition und Moderne.


# Das Programm ausprobieren

Wer jetzt Lust auf etwas Go bekommen hat, kann entweder ein [Tour durch Go](https://tour.golang.org/welcome/1) 
direkt im Browser mit dem starten.


# Lokal
Oder Ihr [installiert Golang](https://golang.org/dl/) auf Eurem Rechner und laßt das Beispiel hier lokal laufen.
Es wird im weiteren angenommen, daß in einer Unix Shell (z.B Git Bash unter Windows) gearbeitet wird.
```
git clone git@github.com:okoeth/profile.git ~/go/src/profile
cd ~/go/src/profile
go test 
go run main.go
curl http://localhost:8017/personalinsights/
```


# In Docker
Alternativ kann man das Programm auch in Docker bauen und starten. Dazu ist nicht einmal eine Installation von
Golang notwendig.
```
git clone git@github.com:okoeth/profile.git /tmp/profile
cd /tmp/profile
docker build -t profile .
docker run -p 8017:8017 -t -i --rm profile
curl http://localhost:8017/personalinsights/
```
