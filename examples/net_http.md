# Usage with basic net/http

## Setup
```bash
go mod init your-module-name
```
---

> main.go
```go
package main

import (
    "net/http"
    "fmt"
    bare "github.com/ruby-network/bare-go/pkg"
)

func main() {
    router := bare.HandleBareHttp("/bare/")
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    })
    fmt.Println("Listening on http://localhost:8080")
    http.ListenAndServe(":8080", router)
}
```
---
## Run
```bash
go get && go run main.go
```
And navigate to http://localhost:8080/bare/
