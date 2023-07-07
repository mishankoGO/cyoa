package cli

import (
	"fmt"
	"github.com/mishankoGO/cyoa/internal/storyteller"
	"os"
	temp "text/template"
)

var templates map[string]*temp.Template

type Cli struct {
	storyTeller map[string]storyteller.Arc
}

func NewCli(storyTeller map[string]storyteller.Arc) *Cli {
	return &Cli{storyTeller: storyTeller}
}

func (c *Cli) Game() error {
	options, err := c.ArcHandler("intro")
	if err != nil {
		return err
	}

	var choice int
	for {
		fmt.Scan(&choice)
		currOption := options[choice].Next
		options, err = c.ArcHandler(currOption)
		if err != nil {
			if err.Error() == "the end" {
				return nil
			}
			return err
		}
	}
}

func (c *Cli) ArcHandler(arc string) ([]storyteller.Option, error) {
	var viewModel storyteller.Arc

	if story, ok := c.storyTeller[arc]; ok {
		viewModel.Title = story.Title
		viewModel.Story = story.Story
		viewModel.Options = story.Options
		if len(viewModel.Options) == 0 {
			err := renderTemplate("end", viewModel)
			if err != nil {
				return nil, err
			}
			return nil, fmt.Errorf("the end")
		}
		err := renderTemplate("index", viewModel)
		if err != nil {
			return nil, err
		}
	}
	return c.storyTeller[arc].Options, nil
}

//Render templates for the given name, template definition and data object
func renderTemplate(name string, viewModel storyteller.Arc) error {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		fmt.Println("no such template")
		return fmt.Errorf("no such template")
	}
	err := tmpl.ExecuteTemplate(os.Stdout, name, viewModel)
	if err != nil {
		fmt.Println("error executing template")
		return err
	}
	return nil
}

//Compile view templates
func init() {
	const tmplIndex = "Title: {{ .Title }}\n\nStory:\n{{ range $val := .Story }}{{ $val }}\n{{ end }}\nВарианты продолжения\n{{ range $index, $val := .Options }}{{ $val.Text }}\nНажмите {{ $index }}\n{{ end }}"
	const tmplEnd = "Title: {{ .Title }}\n\nStory:\n{{ range $val := .Story }}{{ $val }}\n{{ end }}\nКонец!"
	if templates == nil {
		templates = make(map[string]*temp.Template)
	}
	t := temp.New("index")
	templates["index"] = temp.Must(t.Parse(tmplIndex))

	t1 := temp.New("end")
	templates["end"] = temp.Must(t1.Parse(tmplEnd))
}
