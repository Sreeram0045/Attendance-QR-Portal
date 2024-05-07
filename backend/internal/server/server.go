package ServerOpen

import (
	"fmt"
	"log"
	"net/http"
)

func ServerStart() {
	http.Handle("/", http.FileServer(http.Dir("../")))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server Hosting.....")
}
