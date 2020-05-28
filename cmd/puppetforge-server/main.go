package main

import (
	"github.com/syberalexis/puppetforge-server/pkg/server"
)

func main() {
	s := server.PuppetForge{
		Host:     "127.0.0.1",
		Port:     8080,
		ForgeURI: "https://forgeapi.puppet.com/",
	}
	s.Run()
	/*http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello Mars!")
	}))
	log.Println("Now server is running on port 3000")
	http.ListenAndServe(":3000", nil)*/
}
