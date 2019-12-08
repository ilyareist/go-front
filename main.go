package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)
type Template struct {
	templates *template.Template
}

type Data struct {
	Name string
}
var d = Data{"Collin"}
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/index.gohtml")),
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/hello", Hello)
	e.GET("/getJWT", getJWT)
	e.Logger.Fatal(e.Start(":8081"))
}


func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", d)
}


func  getJWT(c echo.Context) error {
	return c.JSON(http.StatusOK, "Fsd")
}