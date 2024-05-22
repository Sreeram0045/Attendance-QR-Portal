package SignUpHandle

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func validateEmail(mail []string) bool {
	expression := regexp.MustCompile("/^(?:[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-zA-Z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:a-z0-9?\\.)+a-z0-9?|\\[(?:\\[(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])$/")
	var concat_mail string = strings.Join(mail, "")
	return expression.MatchString(concat_mail)
}

func validateName(name []string) bool {
	// Regular expression to match names with letters (including accented), spaces, apostrophes, and hyphens
	expression := regexp.MustCompile(`^[a-zA-ZÀ-ÖØ-öø-ÿ\s'-]+$`)
	var concat_name string = strings.Join(name, "")
	return expression.MatchString(concat_name)
}

func validateAge(age []string) bool {
	expression := regexp.MustCompile(`^(18|19|[2-9]\d)$`)
	var concat_age string = strings.Join(age, "")
	return expression.MatchString(concat_age)
}

func validatePhoneNumber(num []string) bool {
	expression := regexp.MustCompile(`^(?:\+91|0)?[6-9]\d{9}$`)
	var concat_phone_num string = strings.Join(num, "")
	return expression.MatchString(concat_phone_num)
}

func validateUserName(username []string) bool {
	expression := regexp.MustCompile(`^[a-zA-Z0-9_@]{1,50}$`)
	var concat_username string = strings.Join(username, "")
	return expression.MatchString(concat_username)
}

func validatePassword(password []string) bool {
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

func validateOrganisation(org []string) bool {
	var concat_org string = strings.Join(org, "")
	expression := regexp.MustCompile(`^[A-Za-z]+$`)
	return expression.MatchString(concat_org)
}

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm() error", http.StatusInternalServerError)
		return
	}
	if r.Method == "POST" {
		if validateName(r.Form["f_name"]) &&
			validateAge(r.Form["u_age"]) &&
			validateEmail(r.Form["u_mail"]) &&
			validatePhoneNumber(r.Form["u_num"]) &&
			validateUserName(r.Form["u_username"]) &&
			validatePassword(r.Form["u_password"]) &&
			validateOrganisation(r.Form["u_org"]) {
			full_name := r.Form["f_name"]
			user_age := r.Form["u_age"]
			user_dob := r.Form["u_dob"]
			user_e_mail := r.Form["u_mail"]
			user_ph_num := r.Form["u_num"]
			user_org := r.Form["u_org"]
			user_username := r.Form["u_username"]
			user_password := r.Form["u_password"]
			fmt.Println(full_name, "\n", user_age, "\n", user_dob, "\n", user_e_mail, "\n", user_ph_num, "\n", user_org, "\n", user_username, "\n", user_password)
		} else {
			// path := checkingPathForServer() + "index.html"
			// http.ServeFile(w, r, path)
			http.Redirect(w, r, "/", http.StatusFound)
			fmt.Println("Can't Get the data")
		}
	} else {
		log.Fatal("Invalid Submission")
	}
}
