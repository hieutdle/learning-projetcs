package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	// New command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.

	addr := flag.String("addr", ":4000", "HTTP network address")

	// Parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000".

	flag.Parse()

	// Use the http.NewServeMux() function to initialize a new servemux
	// servermux is a routerstores a mapping between the URL patterns
	// for your application and the corresponding handlers.

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// register the home function as the handler for the "/" URL pattern.
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet/view", snippetView)

	mux.HandleFunc("/snippet/create", snippetCreate)

	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e.
	// prefix it with the * symbol) before using it. Note that we're using the
	// log.Printf() function to interpolate the address with the log message.

	log.Printf("Starting server on %s", *addr)

	// Start a new web server
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created
	err := http.ListenAndServe(*addr, mux)

	// If http.ListenAndServe() returns an error
	// log.Fatal() function to log the error message and exit
	log.Fatal(err)

}
