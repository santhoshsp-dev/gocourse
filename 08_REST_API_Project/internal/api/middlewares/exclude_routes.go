package middlewares

import (
	"fmt"
	"net/http"
	"strings"
)

func MiddlewaresExcludePaths(middleware func(http.Handler) http.Handler, excludedPaths ...string) func(http.Handler) http.Handler {
	// Start ----------- 099 ------------
	fmt.Println("Middleware Exclude Paths Initialized")
	return func(next http.Handler) http.Handler {
		fmt.Println("=================== Middleware Exclude Paths Ran")
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, path := range excludedPaths {
				if strings.HasPrefix(r.URL.Path, path) {
					next.ServeHTTP(w, r)
					return
				}
			}
			fmt.Println("Send Response from Middleware Exclude Paths")
			middleware(next).ServeHTTP(w, r)
			// End ----------- 099 ------------
		})
	}
}
