package render

import (
	"bookings/packs/config"
	"bookings/packs/models"

	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.App

func Config(a *config.App) {
	app = a
}

var functions = template.FuncMap{}

func addDefaultData(td *models.TemplateData) {

}

func Template(w http.ResponseWriter, tmplName string, td *models.TemplateData) {
	// fmt.Fprintln(w, "this is the line from the render")

	var err error
	if !app.UseCache {
		app.TemplateCache, err = CreateTmplCache()
		if err != nil {
			log.Fatal("error while parsing templates", err)
		}
	}

	tmplSet, ok := app.TemplateCache[tmplName]
	if !ok {
		log.Fatal("the given page doesn't exist")
	}

	// Create buffer to pass executed template
	buff := new(bytes.Buffer)
	addDefaultData(td)
	if err := tmplSet.Execute(buff, td); err != nil {
		log.Fatal("error while executing template set", err)
	}

	if _, err := buff.WriteTo(w); err != nil {
		log.Fatal("error while writing to the response writer", err)
	}

	// // Parse html page
	// parsedTemp, err := template.ParseFiles("packs/templates/" + tmplName)
	// if err != nil {
	// 	fmt.Println("error whilw parsing template")
	// 	fmt.Println(err)
	// 	return
	// }
	// // Execute template
	// err = parsedTemp.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error whilw executing template")
	// 	fmt.Println(err)
	// 	return
	// }

}

func CreateTmplCache() (map[string]*template.Template, error) {

	/*
		Steprs
		1. Get all the pages
		2. For each page
			a. Create a template set for that file. It will have
				1. name, same as the page name
				2. Set of functions which have defined to use indise the template
				3. Layouts which need to be parsed to render the complate page
			b. Get all the layouts
			c. check if no of layouts available is more than 1. If there are than parse them
	*/
	tmplCache := make(map[string]*template.Template)

	// Get all the page templates
	pages, err := filepath.Glob("packs/templates/*page.html")
	if err != nil {
		return nil, err
	}

	// For each template page
	for _, page := range pages {
		name := filepath.Base(page)
		// fmt.Println("name of the file is", name)

		// Create a new template with the same name as the page
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Get all the the layouts
		matches, err := filepath.Glob("packs/templates/*layout.html")
		if err != nil {
			return nil, err
		}

		// If there are any layouts then parse them
		if len(matches) > 0 {
			templateSet, err := templateSet.ParseGlob("packs/templates/*layout.html")
			if err != nil {
				return nil, err
			}
			tmplCache[name] = templateSet
		}

	}
	return tmplCache, nil
}
