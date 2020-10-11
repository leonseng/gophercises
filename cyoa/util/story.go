package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JsonStory(f *os.File) (Story, error) {
	d := json.NewDecoder(f)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

type StoryHandler struct {
	S Story
}

func (sh StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	htmlTemplate := `
	<h1>{{.Title}}</h1>

	{{range .Paragraphs}}
		<p>{{.}}</p>
	{{end}}

	<ul>
		{{range .Options}}
			<li><a href="/{{.Arc}}">{{.Text}}</a></li>
		{{end}}
	</ul>
	`

	var t = template.Must(template.New("").Parse(htmlTemplate))
	path := r.URL.Path

	if path == "/" {
		// default to intro if path undefined
		path = "/intro"
	}

	if c, ok := sh.S[path[1:]]; ok {
		err := t.Execute(w, c)
		if err != nil {
			fmt.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Chapter not found", http.StatusNotFound)
	}
}

// func StoryHandleFunc(s Story) func(http.ResponseWriter, *http.Request) {
// 	htmlTemplate := `
// 	<h1>{{.Title}}</h1>

// 	{{range .Paragraphs}}
// 		<p>{{.}}</p>
// 	{{end}}

// 	<ul>
// 		{{range .Options}}
// 			<li><a href="/{{.Arc}}">{{.Text}}</a></li>
// 		{{end}}
// 	</ul>
// 	`

// 	var t = template.Must(template.New("").Parse(htmlTemplate))

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		cName := r.URL.Path[1:]
// 		fmt.Println(cName)

// 		if c, ok := s[cName]; ok {
// 			fmt.Printf("%v", c)
// 			t.Execute(w, c)
// 		} else {
// 			fmt.Fprintf(w, "Chapter %s not found", cName)
// 		}
// 	}
// }
