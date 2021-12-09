package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ChechAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking authentication")
			if flag {
				f(w, r)
			} else {
				return
			}

		}
	}
}

func Loggin() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func(w http.ResponseWriter, r *http.Request) {
				log.Println(r.URL.Path, time.Since(start))
			}(w, r)
		}
	}
}
