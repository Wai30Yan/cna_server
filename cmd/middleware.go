package main

import (
	"fmt"
	"net/http"
	"strings"
)

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authenticating middleware...")
        authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
        fmt.Println("authHeader", authHeader)
		next.ServeHTTP(w, r)
	})
}

// func preflightMiddleware(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         if r.Method == http.MethodOptions {
//             // Set CORS headers for preflight requests
//             w.Header().Set("Access-Control-Allow-Origin", "*")
//             w.Header().Set("Access-Control-Allow-Methods", "POST")
//             w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//             w.WriteHeader(http.StatusNoContent)
//             return
//         }

//         // Call the next handler in the chain
//         next.ServeHTTP(w, r)
//     })
// }

// func middleware1(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("middleware1")
// 		next.ServeHTTP(w, r)
// 	})
// }