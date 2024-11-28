package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"main/src/middlewares"
	"main/src/utils"
)

type User struct {
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data, _ := json.Marshal(map[string]string{"message": "Hello World"})

		w.Write(data)
	})

	router.Handle("GET /users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal([]User{
			{"John Doe", "g3QzW@example.com", nil},
		})

		w.Write(data)
	}))

	router.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(map[string]string{"message": "Hello World"})

		if err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	middlewares := []func(http.Handler) http.Handler{
		middlewares.Header,
		// middlewares.Auth,
		// middlewares.Logging,
	}

	routerWithMiddleware := utils.Middlewares(router, middlewares...)

	fmt.Println("Server started running on port 8080")

	if err := http.ListenAndServe(":8080", routerWithMiddleware); err != nil {
		fmt.Println("Error", err)
	}
}
