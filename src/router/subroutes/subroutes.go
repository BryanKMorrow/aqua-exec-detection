package subroutes

import (
	"github.com/BryanKMorrow/aqua-exec-detection/src/handlers"
	"github.com/BryanKMorrow/aqua-exec-detection/src/routes"
	"net/http"
)

// SubRoutePackage routes the middleware
type SubRoutePackage struct {
	Routes     routes.Routes
	Middleware func(next http.Handler) http.Handler
}

// Middleware - Handler
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// GetRoutes - Returns the list of Sub Routes
func GetRoutes() (SubRoute map[string]SubRoutePackage) {
	/* SUB ROUTES */
	SubRoute = map[string]SubRoutePackage{
		"/api/v1": {
			Routes: routes.Routes{
				routes.Route{Name: "Status", Method: "POST", Pattern: "/webhook", HandlerFunc: handlers.WebhookHandler},
			},
			Middleware: Middleware,
		},
	}
	return
}
