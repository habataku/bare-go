package bare_go

import (
    "github.com/go-chi/chi"
    "github.com/ruby-network/bare-go/internal/routes"
)

func HandleBare(directory string, chiRouter *chi.Mux) *chi.Mux {
    router := routes.ChiRouter(directory, chiRouter)
    return router
}
