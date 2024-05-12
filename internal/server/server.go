package ServerOpen

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func checkingPathForServer() string {
	comparison := "/Attendance-QR-Portal"
	lastindex := "/pages/"
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return ""
	}
	if strings.Contains(rootDir, comparison) {
		slices := strings.SplitAfter(rootDir, comparison)
		return (strings.Join([]string{slices[0], lastindex}, ""))
	}
	return ""
}

func ServerStart() {
	path := checkingPathForServer()
	http.Handle("/", http.FileServer(http.Dir(path)))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server Hosting.....")
}
