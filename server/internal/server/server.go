package server

import (
	"net/http"

	"benton.codes/anonmsg/internal/core"
	"benton.codes/anonmsg/internal/inbox"
	"benton.codes/anonmsg/internal/messages"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/rs/cors"
)

func Run(config *core.Context) {
	router := http.NewServeMux()

	humaCfg := huma.DefaultConfig("App", "1.0.0")
	humaCfg.DocsPath = ""

	api := humago.New(router, humaCfg)
	v1 := huma.NewGroup(api, "/v1")

	inbox.Register(v1, config)
	messages.Register(v1, config)

	c := cors.New(cors.Options{
		AllowedOrigins: config.AllowedOrigins,
		AllowedMethods: config.AllowedMethods,
	})

	if config.EnableDocs {
		router.HandleFunc("GET /docs", serveDocs)
	}

	http.ListenAndServe(":"+config.Port, c.Handler(router))
}

func serveDocs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<!doctype html>
<html>
 <head>
   <title>API Reference</title>
   <meta charset="utf-8" />
   <meta
     name="viewport"
     content="width=device-width, initial-scale=1" />
 </head>
 <body>
   <script
     id="api-reference"
     data-url="/openapi.json"></script>
   <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
 </body>
</html>`))
}
