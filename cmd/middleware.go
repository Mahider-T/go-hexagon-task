package main

import (
	"fmt"
	"go-hexagon-task/internal/adapters/handlers"
	"log"
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware works 0")
		session, err := handlers.Store.Get(r, "user-id")
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		val := session.Values["user"]
		var usr = handlers.User{}
		usr, ok := val.(handlers.User)

		if !ok || !usr.Authenticated {
			// http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			// fmt.Fprintf(w, "You are not authorized for this operation")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		fmt.Println("Middleware works 2")
		next.ServeHTTP(w, r)
	})
}
