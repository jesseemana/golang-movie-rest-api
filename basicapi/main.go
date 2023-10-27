package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// install gorilla mux	

// Models setup
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"full_name"`
	Website  string `json:"website"`
}

// middleware for checking if course id and course name are empty
func (c *Course) IsEMpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// slice working as a mockup db to hold all the courses
var courses []Course

func main() {
	fmt.Println("API - Get Courses Online")
	r := mux.NewRouter()

	// seeding the courses slice
	courses = append(courses, Course{CourseId: "1", CourseName: "Node Masterclass", CoursePrice: 32, Author: &Author{Fullname: "Zero To Mastery", Website: "zerotomastery.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Advanced React Recipes", CoursePrice: 28, Author: &Author{Fullname: "Kent C. Dodds", Website: "kentdodds.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

// CONTROLLERS
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Courses Online</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Response) {
	fmt.Println("Get Single Courses")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course with given id has been found")
	return
}

func createCourse(w http.ResponseWriter, r *http.Response) {
	fmt.Println("Create a single course")
	w.Header().Set("Content-Type", "application/json")

	// check if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please fill the fields")
	}

	var course Course
	// decode the course reference
	_ = json.NewDecoder(r.Body).Decode(&course)
	// check if body is {}
	if course.IsEMpty() {
		json.NewEncoder(w).Encode("Please provide all the data")
		return
	}

	// TODO: check duplicate course title

	// generate random course id
	// add course into courses slice
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Response) {
	fmt.Println("Update a single course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			if course.IsEMpty() {
				json.NewEncoder(w).Encode("Please proivde necessary fields for the course")
				return
			}
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Course with given id is not found")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Response) {
	fmt.Println("Delet a single course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send json response for confirming deletion
			break
		}
	}
}
