package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strconv"
)

func GetStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var students []models.Student

	// --------------- Start: 096 ----------------
	page, limit := GetPaginationParams(r)

	students, totalStudents, err := sqlconnect.GetStudentsDbHandler(students, r, limit, page)
	// --------------- End: 096 ----------------
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// --------------- Start: 096 ----------------
	response := struct {
		Status   string           `json:"status"`
		Count    int              `json:"count"`
		Page     int              `json:"page"`
		PageSize int              `json:"page_size"`
		Data     []models.Student `json:"data"`
	}{
		Status:   "success",
		Count:    totalStudents,
		Page:     page,
		PageSize: limit,
		Data:     students,
	}
	// --------------- End: 096 ----------------

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// --------------- Start: 096 ----------------
func GetPaginationParams(r *http.Request) (int, int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}
	return page, limit
}

// --------------- End: 096 ----------------

func GetOneStudentHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	student, err := sqlconnect.GetStudentByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func AddStudentHandler(w http.ResponseWriter, r *http.Request) {

	var newStudents []models.Student
	// Step: 8)
	var rawStudents []map[string]interface{}

	// Step: 15)
	// request body become empty once we use it
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Step: 9)
	// err := json.NewDecoder(r.Body).Decode(&rawStudents)
	// Step: 16)
	err = json.Unmarshal(body, &rawStudents)
	// Step: 9)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Step: 13)
	fmt.Println(rawStudents)

	// Step: 4)
	// fields := GetFieldNames()
	// Step: 19)
	fields := GetFieldNames(models.Student{})

	// Step: 6)
	allowedFields := make(map[string]struct{})
	for _, field := range fields {
		allowedFields[field] = struct{}{}
	}

	// Step: 12) // comment these codes in Step: 14
	// fmt.Println(fields)
	// fmt.Println(allowedFields)

	// Step: 7)
	// for _, student := range newStudents {
	// Step: 10)
	for _, student := range rawStudents {
		for key := range student { // this student will show an error so we need Step: 8
			_, ok := allowedFields[key]
			if !ok {
				http.Error(w, "Unacceptable field found in request. Only use allowed fields.", http.StatusBadRequest)
				return
			}
		}
	}

	// Step: 11)
	// err = json.NewDecoder(r.Body).Decode(&newStudents)
	// Step: 17)
	err = json.Unmarshal(body, &newStudents)
	// Step: 11)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		// Step: 14)
		fmt.Println("New Students: ", newStudents)
		return
	}

	// Step: 1)
	for _, student := range newStudents {
		// Step: 2)
		// if student.FirstName == "" || student.LastName == "" || student.Email == "" || student.Class == "" || student.Subject == "" {
		// 	http.Error(w, "All Fields are required", http.StatusBadRequest)
		// 	return
		// }

		// Step: 3)
		err := CheckBlankFields(student)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedStudents, err := sqlconnect.AddStudentsDBHandler(newStudents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Student `json:"data"`
	}{
		Status: "success",
		Count:  len(addedStudents),
		Data:   addedStudents,
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Student Id", http.StatusBadRequest)
		return
	}

	var updatedStudent models.Student
	err = json.NewDecoder(r.Body).Decode(&updatedStudent)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	updatedStudentFromDB, err := sqlconnect.UpdateStudent(id, updatedStudent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedStudentFromDB)
}

func PatchStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var updates []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchStudents(updates)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func PatchOneStudentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Student Id", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	updatedStudent, err := sqlconnect.PatchOneStudent(id, updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedStudent)
}

func DeleteOneStudentHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Student Id", http.StatusBadRequest)
		return
	}

	err = sqlconnect.DeleteOneStudent(id)
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
		Status: "Student successfully deleted",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteStudentsHandler(w http.ResponseWriter, r *http.Request) {

	var ids []int
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	deletedIds, err := sqlconnect.DeleteStudents(ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status     string `json:"status"`
		DeletedIDs []int  `json:"deleted_ids"`
	}{
		Status:     "Students successfully deleted",
		DeletedIDs: deletedIds,
	}
	json.NewEncoder(w).Encode(response)

}
