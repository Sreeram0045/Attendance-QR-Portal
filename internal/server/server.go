package ServerOpen

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	

	LogInHandle "github.com/Sreeram0045/backend/internal/loginhandle"
	SignUpHandle "github.com/Sreeram0045/backend/internal/signuphandle"
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

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	path := checkingPathForServer()
	http.ServeFile(w, r, path+"login.html")
}

// Custom handler for assets
func assetsHandler(w http.ResponseWriter, r *http.Request) {
	assetPath := checkingPathForServer() + "assets"
	requestedPath := filepath.Join(assetPath, r.URL.Path[len("/assets/"):])

	// If the request is for the /assets directory itself, return an error
	if r.URL.Path == "/assets" || r.URL.Path == "/assets/" {
		http.Error(w, "Access to this resource is restricted", http.StatusForbidden)
		return
	}

	// If requestedPath is a directory, return an error
	fileInfo, err := os.Stat(requestedPath)
	if err == nil && fileInfo.IsDir() {
		http.Error(w, "Access to this resource is restricted", http.StatusForbidden)
		return
	}

	// Prevent directory traversal
	if strings.Contains(requestedPath, "..") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	// Serve the file if it exists
	http.ServeFile(w, r, requestedPath)
}

func ServerStart() {
	path := checkingPathForServer()
	http.Handle("/", http.FileServer(http.Dir(path)))
	http.HandleFunc("/login", LoginPageHandler)
	http.HandleFunc("/sign-up", SignUpHandle.UserSignUp)
	http.HandleFunc("/log-in", LogInHandle.UserLogIn)
	// Handler for serving assets while preventing directory traversal
	http.HandleFunc("/assets/", assetsHandler)
	logger := log.New(os.Stdout, "httpServer: ", log.Lmsgprefix)
	logger.Println("server listening on http://localhost:3001")
	err := http.ListenAndServe(":3001", nil)
	log.Fatal(err)
	
	
}
