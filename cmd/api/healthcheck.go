//Filename: cmd/api/healthcheck.go

package main

import (
    "fmt"
    "net/http"
)

// Declare a handler which writes a plain-text response with information about the 
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "status: available")
    fmt.Fprintf(w, "environment: %s\n", app.config.env)
    fmt.Fprintf(w, "version: %s\n", version)
}


// import "crypto/rand"

// type Tools struct{}

// const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_#$-!~"

// func (t *Tools) generateRandomString(length int) string {
// 	s := make([]rune, length)
// 	r := []rune(randomStringSource)

// 	for i := range s {
// 		p, _ := rand.Prime(rand.Reader, len(r))
// 		x := p.Uint64()
// 		y := uint64(len(r))
// 		s[i] = r[x%y]
// 	}

// 	return string(s)

// }