package main

import (
	"net/http"

	"github.com/lu97chi/gofirst/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
