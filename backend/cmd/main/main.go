package main

import (
	"builders/forweb/blueprints"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error al obtener el directorio de trabajo: %v", err)
	}
	dir := root + "/output/web"

	fileServer := http.FileServer(http.Dir(dir))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpPath := path.Join(dir, "index_template.html")

		tmpl, err := template.ParseFiles(tmpPath)
		if err != nil {
			log.Fatalf("Error al cargar la plantilla: %v", err)
		}

		data := blueprints.IndexBlueprint{
			Title: "Bienvenido a mi sitio web",
			Body:  "Este es un ejemplo de una página web servida por un servidor Go.",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error al procesar la plantilla", http.StatusInternalServerError)
			return
		}
		log.Println("Página servida correctamente")
	})

	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
