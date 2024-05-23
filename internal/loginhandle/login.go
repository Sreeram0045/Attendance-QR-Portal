package LogInHandle

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func LogInValidateUserName(username []string) bool {
	expression := regexp.MustCompile(`^[a-zA-Z0-9_@]{1,50}$`)
	var concat_username string = strings.Join(username, "")
	return expression.MatchString(concat_username)
}

func LogInValidatePassword(password []string) bool {
	// Join the slice of passwords into a single string
	concatPassword := strings.Join(password, "")

	// Regular expression to match valid passwords
	expression := regexp.MustCompile(`^[A-Za-z\d!@#$%^&*]{8,}$`)

	// Check if the concatenated password matches the regular expression
	if !expression.MatchString(concatPassword) {
		return false
	}

	// Check for at least one lowercase letter
	if !strings.ContainsAny(concatPassword, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}

	// Check for at least one uppercase letter
	if !strings.ContainsAny(concatPassword, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}

	// Check for at least one digit
	if !strings.ContainsAny(concatPassword, "0123456789") {
		return false
	}

	// Check for at least one special character
	if !strings.ContainsAny(concatPassword, "!@#$%^&*") {
		return false
	}

	return true
}

func UserLogIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm() error", http.StatusInternalServerError)
		return
	}

	if LogInValidateUserName(r.Form["u_name"]) && LogInValidatePassword(r.Form["pass_key"]) {
		fmt.Printf("Username: %s\n Password: %s", r.Form["u_name"], r.Form["pass_key"])
	} else {
		fmt.Printf("Username: %s\n Password: %s", r.Form["u_name"], r.Form["pass_key"])
		http.Redirect(w, r, "/login.html", http.StatusFound)
		fmt.Println("Can't Get the data")
	}
}
