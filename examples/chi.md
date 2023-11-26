# Usage with Chi

main.go
```go
package main

import (
    "github.com/go-chi/chi/v5"
    "net/http"
    bare "github.com/ruby-network/bare-go/pkg"
)

func main() {
    initalRouter := chi.NewRouter()
    router := bare.HandleBareChi("/bare/", initalRouter)
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })
    http.ListenAndServe(":8080", router)
}
