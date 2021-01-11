package chooseadventure

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var t *template.Template

func init() {
	var err error
	t, err = template.ParseFiles("coa.gohtml")
	if err != nil {
		panic(err)
	}
}

// StoryDecoder takes a io.Reader as input as returns Story
func StoryDecoder(r io.Reader) (Story, error) {
	fileDecoder := json.NewDecoder(r)
	var mainStory Story
	err := fileDecoder.Decode(&mainStory)
	if err != nil {
		return nil, err
	}
	return mainStory, err
}

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	// remove leading slash ; /intro -> intro
	path = path[1:]
	if path == "" || path == "/" {
		path = "intro"
	}
	if c, ok := h.s[path]; ok {
		err := t.Execute(w, c)
		if err != nil {
			log.Println(err)
			http.Error(w, "Something went wrong ...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusInternalServerError)
}

// Story is map of Chapters
type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
