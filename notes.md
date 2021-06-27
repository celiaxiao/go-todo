# notes

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>This is the homepage. Try /hello and /hello/Sammy\n</h1>")
    })

    r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Hello from Docker!\n</h1>")
    })

    r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["name"]

        fmt.Fprintf(w, "<h1>Hello, %s!\n</h1>", title)
    })

    http.ListenAndServe(":80", r)
}
```

first import net/http and gorilla/mux packages, which provide HTTP server functionality and routing.

The gorilla/mux package implements an easier and more powerful request router and dispatcher, while at the same time maintaining interface compatibility with the standard router. Here, you instantiate a new mux router and store it in variable r. Then, you define three routes: /, /hello, and /hello/{name}. The first (/) serves as the homepage and you include a message for the page. The second (/hello) returns a greeting to the visitor. For the third route (/hello/{name}) you specify that it should take a name as a parameter and show a greeting message with the name inserted.

At the end of your file, you start the HTTP server with http.ListenAndServe and instruct it to listen on port 80, using the router you configured.
