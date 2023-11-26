package bare_go

import (
    "github.com/go-chi/chi/v5"
    "github.com/ruby-network/bare-go/internal/routes"
)

func HandleBareChi(directory string, chiRouter *chi.Mux) *chi.Mux {
    router := routes.ChiRouter(directory, chiRouter)
    return router
}
