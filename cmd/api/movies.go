package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "create a new movie")

	// params := httprouter.ParamsFromContext(r.Context())

	// id, err := strconv.ParseInt(params.ByName("id"), 10, 64)

	// if err != nil || id < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }

	// fmt.Fprintf(w, " show the details of the moview %d\n", id)

	id, err := app.readIDParam(r) // this is in helpers.go, as it is being used often
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, " show the details of the moviez %d\n ", id)
}

// func main() {

// }
