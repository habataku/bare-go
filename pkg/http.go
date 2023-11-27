package bare_go

import (
    "net/http"
    "github.com/ruby-network/bare-go/internal/routes"
)

func HandleBare(directory string, router *http.ServeMux) http.Handler {
    router = routes.NetHttp(directory, router)
    return router
}
