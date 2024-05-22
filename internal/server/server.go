package ServerOpen

import (
	"fmt"
	SignUpHandle "github.com/Sreeram0045/backend/internal/signuphandle"
	"log"
	"net/http"
	"os"
	"strings"
)

func checkingPathForServer() string {
	// created two variables to check/compare the path
	comparison := "/Attendance-QR-Portal"
	lastindex := "/pages/"
	// created two variables one will get the current working directory and other will recieve the error
	rootDir, err := os.Getwd()
	// error handling
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return ""
	}
	// checks if the current working directory contains the compared path
	if strings.Contains(rootDir, comparison) {
		// The current working directory is split into before and after the comparison variable
		slices := strings.SplitAfter(rootDir, comparison)
		// The first part of the string is joined with the lastindex varible in-order to get the correct path to serve
		return (strings.Join([]string{slices[0], lastindex}, ""))
	}
	// if the variable is not present as a part of the working directory an empty string is returned which will obviously result in error
	return ""
}

func ServerStart() {
	path := checkingPathForServer()
	http.Handle("/", http.FileServer(http.Dir(path)))
	http.HandleFunc("/sign-up", SignUpHandle.UserSignUp)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
