package routes


import (
    "net/http"
    "github.com/rs/cors"
    //imports from project
    "github.com/ruby-network/bare-go/internal/utils"
    "github.com/ruby-network/bare-go/internal/v3"
)

func NetHttp(directory string, router *http.ServeMux) *http.ServeMux {
    cors := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
        ExposedHeaders: []string{"*"},
        MaxAge: 300,
    })

    router.HandleFunc(directory, func(w http.ResponseWriter, r *http.Request) {
        e := utils.GetJson()
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(e))
    })
    router.HandleFunc(directory + "health", func(w http.ResponseWriter, r *http.Request) {
        e := `{"status": "ok"}`
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(e))
    })

    router.HandleFunc(directory + "v3/", func(w http.ResponseWriter, r *http.Request) {
        method := string(r.Method)
        userAgent := string(r.Header.Get("User-Agent"))
        headers := r.Header
        v3.Handler(method, userAgent, w, r, headers)
    })

    return router
}
