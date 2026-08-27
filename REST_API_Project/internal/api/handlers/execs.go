package handlers

import (
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"restapi/pkg/utils"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GetExecsHandler(w http.ResponseWriter, r *http.Request) {

	var execs []models.Exec
	execs, err := sqlconnect.GetExecsDbHandler(execs, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "success",
		Count:  len(execs),
		Data:   execs,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetOneExecHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	exec, err := sqlconnect.GetExecByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exec)
}

func AddExecsHandler(w http.ResponseWriter, r *http.Request) {

	var newExecs []models.Exec
	// Step: 8)
	var rawExecs []map[string]interface{}

	// Step: 15)
	// request body become empty once we use it
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Step: 9)
	// err := json.NewDecoder(r.Body).Decode(&rawExecs)
	// Step: 16)
	err = json.Unmarshal(body, &rawExecs)
	// Step: 9)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Step: 13)
	fmt.Println(rawExecs)

	// Step: 4)
	// fields := GetFieldNames()
	// Step: 19)
	fields := GetFieldNames(models.Exec{})

	// Step: 6)
	allowedFields := make(map[string]struct{})
	for _, field := range fields {
		allowedFields[field] = struct{}{}
	}

	// Step: 12) // comment these codes in Step: 14
	// fmt.Println(fields)
	// fmt.Println(allowedFields)

	// Step: 7)
	// for _, exec := range newExecs {
	// Step: 10)
	for _, exec := range rawExecs {
		for key := range exec { // this exec will show an error so we need Step: 8
			_, ok := allowedFields[key]
			if !ok {
				http.Error(w, "Unacceptable field found in request. Only use allowed fields.", http.StatusBadRequest)
				return
			}
		}
	}

	// Step: 11)
	// err = json.NewDecoder(r.Body).Decode(&newExecs)
	// Step: 17)
	err = json.Unmarshal(body, &newExecs)
	// Step: 11)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		// Step: 14)
		fmt.Println("New Execs: ", newExecs)
		return
	}

	// Step: 1)
	for _, exec := range newExecs {
		// Step: 2)
		// if exec.FirstName == "" || exec.LastName == "" || exec.Email == "" || exec.Class == "" || exec.Subject == "" {
		// 	http.Error(w, "All Fields are required", http.StatusBadRequest)
		// 	return
		// }

		// Step: 3)
		err := CheckBlankFields(exec)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedExecs, err := sqlconnect.AddExecsDBHandler(newExecs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string        `json:"status"`
		Count  int           `json:"count"`
		Data   []models.Exec `json:"data"`
	}{
		Status: "success",
		Count:  len(addedExecs),
		Data:   addedExecs,
	}

	json.NewEncoder(w).Encode(response)
}

func PatchExecsHandler(w http.ResponseWriter, r *http.Request) {

	var updates []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchExecs(updates)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func PatchOneExecsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Exec Id", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	updatedExec, err := sqlconnect.PatchOneExec(id, updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedExec)
}

func DeleteOneExecsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Exec Id", http.StatusBadRequest)
		return
	}

	err = sqlconnect.DeleteOneExec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ---- Alternate approach
	// w.WriteHeader(http.StatusNoContent)

	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Exec successfully deleted",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Exec // store request body in this req varibale

	// Data Validation
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and Password are required", http.StatusBadRequest)
		return
	}

	// Search for user if user actually exists
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		utils.ErrorHandler(err, "error updating data")
		http.Error(w, "error connecting to database", http.StatusBadRequest)
		return
	}
	defer db.Close()

	user := &models.Exec{} // create a blank instance of models.exec

	err = db.QueryRow(`SELECT id, first_name, last_name, email, username, password, inactive_status, role FROM execs WHERE username = ?`, req.Username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Username, &user.Password, &user.InactiveStatus, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.ErrorHandler(err, "user not found")
			http.Error(w, "user not found", http.StatusBadRequest)
			return
		}
		http.Error(w, "database query error", http.StatusBadRequest)
		return
	}

	// -------------- Start: 085 ----------------
	// is user active
	if user.InactiveStatus {
		http.Error(w, "Account is inactive", http.StatusForbidden)
		return
	}

	// verify password
	parts := strings.Split(user.Password, ".")
	if len(parts) != 2 {
		utils.ErrorHandler(errors.New("invalid encoded hash format"), "invalid encoded hash format")
		http.Error(w, "invalid encoded hash format", http.StatusForbidden)
		return
	}

	saltBse64 := parts[0]
	hashedPasswordBase64 := parts[1]

	salt, err := base64.StdEncoding.DecodeString(saltBse64)
	if err != nil {
		utils.ErrorHandler(err, "failed to decode the salt")
		http.Error(w, "failed to decode the salt", http.StatusForbidden)
		return
	}

	hashedPassword, err := base64.StdEncoding.DecodeString(hashedPasswordBase64)
	if err != nil {
		utils.ErrorHandler(err, "failed to decode the hashed password")
		http.Error(w, "failed to decode the hashed password", http.StatusForbidden)
		return
	}

	hash := argon2.IDKey([]byte(req.Password), salt, 1, 64*1024, 4, 32)

	if len(hash) != len(hashedPassword) {
		utils.ErrorHandler(errors.New("incorrect password"), "incorrect password")
		http.Error(w, "incorrect password", http.StatusForbidden)
		return
	}

	if subtle.ConstantTimeCompare(hash, hashedPassword) == 1 {
		// do nothing
	} else {
		utils.ErrorHandler(errors.New("incorrect password"), "incorrect password")
		http.Error(w, "incorrect password", http.StatusForbidden)
		return
	}
	// -------------- End: 085 ----------------
	// Generate Token

	// Send token as a response or as a cookie
}
