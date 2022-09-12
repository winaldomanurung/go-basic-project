package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string){
	// get the template cache from the app config
	// how?

	// =================================================

	// 1. create a template cache
	tc, err := CreateTemplateCache()
	// jika ada error maka kita panggil Fatal
	if err != nil {
		log.Fatal(err)
	}

	// jika lewat error checking maka kita akan punya template cache dan lanjut ke langkah 2
	// 2. get requested template from cache
	// t merupakan index, ok adalah booelan true or false. dia mengecek apakah ada t di dalam tmpl (yang merupakan argument function RenderTemplate)
	// t ini antara si about atau home
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// langkah selanjutnya adalah optional, yaitu membuat buff
	// create buffer dengan menggunakan bytes.Buffer
	// disini what we are going to do is try to execute the value that we got from the map. Namun, rather than doing it directly (we can do it if we want), we are going to execute buffer directly and then write it out. And the only reason we are doing this is for finer grained error checking, karena ketika kita create this buffer, kita bisa ignore the result. 
	buf := new(bytes.Buffer)

	// kita bisa execute buff dan nil. Ini memberi clear indication that the value we got from that map, there is something wrong with it. It parsed it, but we cant execute it and we dont know what situation that might be.
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// 3. render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error){
	// myCache := make(map[string]*template.Template)
	// atau bisa juga ditulis:
	myCache := map[string]*template.Template{}

  // we want to make our cache, but we want to create entire cache at once and populate it with every available templates
	// when you are rendering a template that uses layout, you typically must have as the first thing you try to parse, the template you want to render and then the associated layout and partials and so forth.
	// artinya, when we start parsing our templates and adding them to myCache, we want to do everything that ends in .page.tmpl first
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// if we get past the error checking it means we now have a slice of string with all the files that end in .page.tmpl from the templates folder 
	// selanjutnya kita range through all the files ending with *.page.tmpl
	// setiap iterasi, page akan berisi value whatever we get from the slice of strings
	for _, page := range pages {
		// page akan mengambil bagian belakang dari nama file
		// Base berfungsi untuk mereturn last element of the path
		name := filepath.Base(page)

		// kita parse file dengan nama 'page' dan store that in a template called 'name', lalu kita masukkan ke variable ts
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// lalu kita loop through layout that exist in that directory
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// ParseGlob parses the template definitions in the files identified by the pattern and associates the resulting templates with t
		// all this is doing is dia menerangkan bahwa some of the file di line:
		//  ts, err := template.New(name).ParseFiles(page)
		//...might require the file layout down here
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// setelah semua selesai maka kita simpan myCache dengan key name ke dalam ts (template set)
		myCache[name] = ts
	}

	return myCache, nil
}