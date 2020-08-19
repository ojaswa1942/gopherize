package adventure

import (
	"fmt"
	"encoding/json"
	"io"
	"html/template"
	"os"
	"net/http"
)

var defaultHTMLTemplate = `<!DOCTYPE html>
	<html>
	<head>
		<title>Choose your own story</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
			<p>{{.}}</p>
		{{end}}
		<ul>
			{{range .Options}}
			<li>
				<a href="/{{.Chapter}}">{{.Text}}</a>
			</li>
			{{end}}
		</ul>
	</body>
	</html>
`

func init() {
	tpl = template.Must(template.New("").Parse(defaultHTMLTemplate))
}

var tpl *template.Template

type Story map[string]Chapter

type Chapter struct {
	Title   string    `json:"title"`
	Paragraphs   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Chapter  string `json:"arc"`
}

func DecodeJSONStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		exit(fmt.Sprintln("cannot execute template: ", err), true)
	}
}


func main() {
	fmt.Println("yep")
}

func exit(msg string, error bool) {
	fmt.Println(msg)
	if error {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}