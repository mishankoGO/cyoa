package controllers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mishankoGO/cyoa/internal/storyteller"
	temp "html/template"
	"net/http"
	"os"
)

var templates map[string]*temp.Template

type Controller struct {
	storyTeller map[string]storyteller.Arc
}

func NewController(storyTeller map[string]storyteller.Arc) *Controller {
	return &Controller{storyTeller: storyTeller}
}

func (c *Controller) Route() *chi.Mux {
	// init chi router
	router := chi.NewRouter()

	// init middleware
	router.Use(middleware.Recoverer)

	// group request /api/user
	router.Route("/", func(r chi.Router) {
		r.Get("/", c.IntroHandler().ServeHTTP)
		r.Get("/{arc}", c.ArcHandler().ServeHTTP)
	})
	return router
}

func (c *Controller) IntroHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var viewModel storyteller.Arc
		//change: moved render template call inside if block.
		//Read value from route variable
		//arc := chi.URLParam(r, "arc")

		if story, ok := c.storyTeller["intro"]; ok {
			viewModel.Title = story.Title
			viewModel.Story = story.Story
			viewModel.Options = story.Options
			renderTemplate(w, "index", "base", viewModel)
		} else {
			http.Error(w, "Could not find the resource to edit.", http.StatusBadRequest)
		}
	})
}

func (c *Controller) ArcHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var viewModel storyteller.Arc
		//change: moved render template call inside if block.
		//Read value from route variable
		arc := chi.URLParam(r, "arc")

		if story, ok := c.storyTeller[arc]; ok {
			viewModel.Title = story.Title
			viewModel.Story = story.Story
			viewModel.Options = story.Options
			if len(viewModel.Options) == 0 {
				renderTemplate(w, "end", "base", viewModel)
				return
			}
			renderTemplate(w, "index", "base", viewModel)
		} else {
			http.Error(w, "Could not find the resource to edit.", http.StatusBadRequest)
		}
	})
}

//Render templates for the given name, template definition and data object
func renderTemplate(w http.ResponseWriter, name string, template string, viewModel storyteller.Arc) {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	// add style to html
	style, err := os.ReadFile("templates/css/style.css")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmplData := struct {
		Style     temp.CSS
		ViewModel storyteller.Arc
	}{Style: temp.CSS(style), ViewModel: viewModel}
	err = tmpl.ExecuteTemplate(w, template, tmplData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//Compile view templates
func init() {
	if templates == nil {
		templates = make(map[string]*temp.Template)
	}
	templates["index"] = temp.Must(temp.ParseFiles("templates/index.html", "templates/base.html"))
	templates["end"] = temp.Must(temp.ParseFiles("templates/end.html", "templates/base.html"))

}
