package commons

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func GetPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func GetEnv() string {
	return os.Getenv("MYENV")
}

// PrintErr function
func PrintErr(err error, errorCode string) {
	fmt.Printf("%v [%s]\n", err, errorCode)
}

// Error function
func Error(err error, errorCode string) error {
	return fmt.Errorf("%v [%s]", err, errorCode)
}

// GetTemplatePath returns format of "templates/%s.html"
func GetTemplatePath(templateName string) string {
	return fmt.Sprintf("templates/%s.html", templateName)
}

func GetTemplate() *template.Template {
	return template.Must(template.New("").ParseGlob("templates/*"))
}

func Middleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		if GetEnv() == "production" {
			r.URL.Scheme = "https"
		}

		next(w, r, ps)
	}
}
