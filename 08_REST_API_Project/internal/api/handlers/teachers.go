package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"restapi/pkg/utils"
	"strconv"
)

func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {

	var teachers []models.Teacher
	teachers, err := sqlconnect.GetTeachersDbHandler(teachers, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(teachers),
		Data:   teachers,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetOneTeacherHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	teacher, err := sqlconnect.GetTeacherByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}

func AddTeacherHandler(w http.ResponseWriter, r *http.Request) {

	var newTeachers []models.Teacher
	// Step: 8)
	var rawTeachers []map[string]interface{}

	// Step: 15)
	// request body become empty once we use it
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Step: 9)
	// err := json.NewDecoder(r.Body).Decode(&rawTeachers)
	// Step: 16)
	err = json.Unmarshal(body, &rawTeachers)
	// Step: 9)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Step: 13)
	fmt.Println(rawTeachers)

	// Step: 4)
	// fields := GetFieldNames()
	// Step: 19)
	fields := GetFieldNames(models.Teacher{})

	// Step: 6)
	allowedFields := make(map[string]struct{})
	for _, field := range fields {
		allowedFields[field] = struct{}{}
	}

	// Step: 12) // comment these codes in Step: 14
	// fmt.Println(fields)
	// fmt.Println(allowedFields)

	// Step: 7)
	// for _, teacher := range newTeachers {
	// Step: 10)
	for _, teacher := range rawTeachers {
		for key := range teacher { // this teacher will show an error so we need Step: 8
			_, ok := allowedFields[key]
			if !ok {
				http.Error(w, "Unacceptable field found in request. Only use allowed fields.", http.StatusBadRequest)
				return
			}
		}
	}

	// Step: 11)
	// err = json.NewDecoder(r.Body).Decode(&newTeachers)
	// Step: 17)
	err = json.Unmarshal(body, &newTeachers)
	// Step: 11)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		// Step: 14)
		fmt.Println("New Teachers: ", newTeachers)
		return
	}

	// Step: 1)
	for _, teacher := range newTeachers {
		// Step: 2)
		// if teacher.FirstName == "" || teacher.LastName == "" || teacher.Email == "" || teacher.Class == "" || teacher.Subject == "" {
		// 	http.Error(w, "All Fields are required", http.StatusBadRequest)
		// 	return
		// }

		// Step: 3)
		err := CheckBlankFields(teacher)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	addedTeachers, err := sqlconnect.AddTeachersDBHandler(newTeachers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Teacher Id", http.StatusBadRequest)
		return
	}

	var updatedTeacher models.Teacher
	err = json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	updatedTeacherFromDB, err := sqlconnect.UpdateTeacher(id, updatedTeacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTeacherFromDB)
}

func PatchTeachersHandler(w http.ResponseWriter, r *http.Request) {

	var updates []map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = sqlconnect.PatchTeachers(updates)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func PatchOneTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Teacher Id", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	updatedTeacher, err := sqlconnect.PatchOneTeacher(id, updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTeacher)
}

func DeleteOneTeacherHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Teacher Id", http.StatusBadRequest)
		return
	}

	err = sqlconnect.DeleteOneTeacher(id)
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
		Status: "Teacher successfully deleted",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteTeachersHandler(w http.ResponseWriter, r *http.Request) {

	var ids []int
	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	deletedIds, err := sqlconnect.DeleteTeachers(ids)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status     string `json:"status"`
		DeletedIDs []int  `json:"deleted_ids"`
	}{
		Status:     "Teachers successfully deleted",
		DeletedIDs: deletedIds,
	}
	json.NewEncoder(w).Encode(response)

}

func GetStudentsByTeacherId(w http.ResponseWriter, r *http.Request) {
	teacherId := r.PathValue("id")

	var students []models.Student

	students, err := sqlconnect.GetStudentsByTeacherIdFromDb(teacherId, students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Student `json:"data"`
	}{
		Status: "success",
		Count:  len(students),
		Data:   students,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetStudentsCountByTeacherId(w http.ResponseWriter, r *http.Request) {
	// --------------- Start: 098 ----------------
	// admin, manager, exec
	_, err := utils.AuthorizeUser(r.Context().Value(utils.ContextKey("role")).(string), "admin", "manager", "exec")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// --------------- End: 098 ----------------

	teacherId := r.PathValue("id")

	var studentCount int

	studentCount, err = sqlconnect.GetStudentsCountByTeacherIdFromDb(teacherId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Status string `json:"status"`
		Count  int    `json:"count"`
	}{
		Status: "success",
		Count:  studentCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
