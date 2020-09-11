package commons

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

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
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		if GetEnv() == "production" && ctx.Value("is-https") != "1" {
			target := "https://" + r.Host + r.URL.Path
			if len(r.URL.RawQuery) > 0 {
				target += "?" + r.URL.RawQuery
			}
			log.Printf("redirect to: %s", target)

			ctx = context.WithValue(ctx, "is-https", "1")
			r = r.WithContext(ctx)

			http.Redirect(w, r, target, http.StatusTemporaryRedirect)
			return
		} else {
			next(w, r, ps)
		}

	}
}
