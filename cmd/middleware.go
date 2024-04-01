package main

import (
	"net/http"
)

type User struct {
	Authenticated bool
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// session, err := handlers.Store.Get(r, "user-id")
		// if err != nil {
		// 	http.Error(w, "Forbidden", http.StatusForbidden)
		// 	return
		// }
		// val := session.Values["user"]
		// _, ok := val.(User)

		// if !ok {
		// 	http.Redirect(w, r, "user/login", http.StatusSeeOther)
		// 	return
		// }
		next.ServeHTTP(w, r)
	})
}
