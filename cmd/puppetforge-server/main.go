package main

import (
	"github.com/syberalexis/puppetforge-server/pkg/server"
)

func main() {
	server.TestModel()
	/*http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello Mars!")
	}))
	log.Println("Now server is running on port 3000")
	http.ListenAndServe(":3000", nil)*/
}
