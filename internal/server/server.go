package server

import (
	"html/template"
	"io"
	"net/http"

	"github.com/dxtym/zfetch/internal/specs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type FetchInfo struct {
	specs.HostInfo
	specs.CpuInfo
	specs.MemInfo
}

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = &Templates{
		templates: template.Must(template.ParseGlob("web/views/*.html")),
	}

	e.GET("/", func(c echo.Context) error {
		var res FetchInfo

		host, err := specs.GetHostInfo()
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error", err)
		}
		res.HostInfo = host

		cpu, err := specs.GetCpuInfo()
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error", err)
		}
		res.CpuInfo = cpu

		mem, err := specs.GetMemInfo()
		if err != nil {
			return c.Render(http.StatusInternalServerError, "error", err)
		}
		res.MemInfo = mem

		return c.Render(http.StatusOK, "index", res)
	})

	e.Logger.Fatal(e.Start(":6969"))
}
