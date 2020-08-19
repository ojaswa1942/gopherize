package adventure

import (
	"fmt"
	"encoding/json"
	"io"
	"html/template"
	"os"
	"net/http"
	"strings"
)

var defaultHTMLTemplate = `<!DOCTYPE html>
	<html>
	<head>
		<title>Choose your own story</title>
	</head>
		<body>
			<section class="page">
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
			</section>
		</body>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 4px 4px 6px 2px #77777740;
        border-radius: 10px;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #166fa7;
      }
      a:active,
      a:hover {
        color: #00395b;
      }
      p {
        text-indent: 1em;
      }
    </style>
	</html>
`

func init() {
	// initiate default template
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

// Functional Options design paradigm
type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFn(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = fn
	}
}

// uses internal mux instead wanting the user to handle
func WithURLPrefix(urlPrefix string) HandlerOption {
	return func(h *handler) {
		h.urlPrefix = urlPrefix
	}
}

// Not efficient for conditional arguments
// type HandlerOptions struct {
// 	t *template.Template
// 	ParseFunc func(r *http.Request) string
// }

func DecodeJSONStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}


type handler struct {
	s Story
	t *template.Template
	pathFn func(r *http.Request) string
	urlPrefix string
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl, defaultPathFn, "/"}
	for _, opt := range opts {
		opt(&h)
	}
	mux := http.NewServeMux()
	mux.Handle(h.urlPrefix, h)
	return mux
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)

	if chapter, okay := h.s[path]; okay {
		err := h.t.Execute(w, chapter)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			exit(fmt.Sprintln("cannot execute template: ", err), true)
		}
	} else {
		http.Error(w, "Chapter not found", http.StatusNotFound)
	}
}

func defaultPathFn(r *http.Request) (string) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]
	return path
}

func exit(msg string, error bool) {
	fmt.Println(msg)
	if error {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}