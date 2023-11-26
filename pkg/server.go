package bare_go

import (
    "github.com/go-chi/chi/v5"
    "github.com/Ruby-Network/bare-go/internal/routes"
)

func HandleBareChi(directory string, chiRouter *chi.Mux) *chi.Mux {
    router := routes.Router(directory, chiRouter)
    return router
}
