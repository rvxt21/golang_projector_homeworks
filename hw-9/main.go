package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"students/collector"
)

// Створити вебсервер для перегляду інформації щодо класу школи.
// Користувач повинен мати можливість отримувати загальну інформацію про клас (список учнів, назва класу).
// Додаткові вимоги:
// • інформація про учнів має зберігатися в оперативній пам'яті та бути доступною під час кожного запиту;
// • отримання інформації про учня (наприклад, середній бал по предметах) має здійснюватись методом GET на адресі "/student/{id}", де {id} — унікальний ідентифікатор учня;
// • дані можна отримати, лише якщо користувач є вчителем у цьому класі.

func main() {
	mux := http.NewServeMux()

	students := StudentResource{
		s: collector.NewStorage(),
	}

	mux.HandleFunc("GET /allStudents", checkAuth(students.GetAll))
	mux.HandleFunc("POST /student", checkAuth(students.CreateOneStudent))
	mux.HandleFunc("GET /student/{id}", checkAuth(students.SearcByID))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Failed to listen and serve: %v\n", err)
	}

}

type StudentResource struct {
	s *collector.Storage
}

func (st *StudentResource) GetAll(w http.ResponseWriter, r *http.Request) {
	students := st.s.GetAllStudents()

	err := json.NewEncoder(w).Encode(students)
	if err != nil {
		fmt.Printf("Failed to encode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (st *StudentResource) CreateOneStudent(w http.ResponseWriter, r *http.Request) {
	var student collector.Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		fmt.Printf("Failed to decode: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	student.ID = st.s.CreateOneStudent(student)
	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		fmt.Printf("Failed to encode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (st *StudentResource) SearcByID(w http.ResponseWriter, r *http.Request) {
	idVal := r.PathValue("id")

	studentID, err := strconv.Atoi(idVal)
	if err != nil {
		fmt.Printf("Invalid id param: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	student, ok := st.s.FindStudentById(studentID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		fmt.Printf("Failed to encode: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

type User struct {
	Username string
	Password string
}

var singleUser = User{
	Username: "teacher",
	Password: "1234567",
}

func checkAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if username != singleUser.Username || password != singleUser.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
