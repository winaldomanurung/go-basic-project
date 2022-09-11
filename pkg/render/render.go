package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string){
	// kita parse salah satu template
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, nil)
	if err !=nil {
		fmt.Println("error parsing template:", err)
	}
}

// the key will look things yang merupakan string dan dengan apa yang di store di dalamnya merupakan apa yang di return template.ParseFiles (kalau kita hover dia mereturn template.Template)
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string){
	// we need to store parsed template in variable
	// instead of reading from the disk every single time, we'll have some kind of data structure we can store a parsed template into, selanjutnya we'll check to see if the template exists in that data structure. If it does, we'll use it, and if it doesnt, we'll read it from disk, parse it and then store the resulting template in that data structure.
	// the best data structure to use is map
	var tmpl *template.Template
	var err error

	// check if we already have the template in our cache
	// kita lihat di map tc, adakah key t
	// _ merupakan apa isinya, inMap merupakan apakah dia true
	_, inMap := tc[t]
	if !inMap{
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}

	} else {
		// we have the template in the cache
		log.Println(("using cached template"))
	}

	// kita ambil templatenya
	tmpl = tc[t]

	// lalu kita execute
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

// kita buat sama seperti function ini:
// 	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
func createTemplateCache(t string) error{
	// kita create templates dengan value of a slice of string. In that slice, we'll put one entry for each of the things required to render a template to the web browser
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		// kita define juga layout
		"./templates/base.layout.tmpl",
	}

	// ... => sama seperti spread
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)
	tc[t] = tmpl
	return nil
}