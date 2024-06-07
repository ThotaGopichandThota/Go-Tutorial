package main

import (
    "fmt"
)

// Define struct for Student
type Student struct {
    Name     string
    Age      int
    Grade    float64
    Subjects []string
}

// Global slice to store student records
var students []Student

// Function to add a new student
func addStudent(name string, age int, grade float64, subjects []string) {
    student := Student{Name: name, Age: age, Grade: grade, Subjects: subjects}
    students = append(students, student)
    fmt.Println("Student added successfully.")
}

// Function to update student details by name
func updateStudentByName(name string, age int, grade float64, subjects []string) {
    for i, student := range students {
        if student.Name == name {
            students[i].Age = age
            students[i].Grade = grade
            students[i].Subjects = subjects
            fmt.Println("Student details updated successfully.")
            return
        }
    }
    fmt.Println("Student not found.")
}

// Function to delete a student by name
func deleteStudentByName(name string) {
    for i, student := range students {
        if student.Name == name {
            students = append(students[:i], students[i+1:]...)
            fmt.Println("Student deleted successfully.")
            return
        }
    }
    fmt.Println("Student not found.")
}

// Function to display all student records
func displayStudents() {
    if len(students) == 0 {
        fmt.Println("No student records found.")
        return
    }
    fmt.Println("Student Records:")
    for _, student := range students {
        fmt.Printf("Name: %s, Age: %d, Grade: %.2f, Subjects: %v\n", student.Name, student.Age, student.Grade, student.Subjects)
    }
}

// Function to calculate the average grade of all students
func calculateAverageGrade() float64 {
    if len(students) == 0 {
        return 0.0
    }
    totalGrade := 0.0
    for _, student := range students {
        totalGrade += student.Grade
    }
    return totalGrade / float64(len(students))
}

func main() {
    // Add some initial students
    addStudent("Alice", 18, 85.5, []string{"Math", "Science"})
    addStudent("Bob", 17, 90.0, []string{"History", "English"})
    addStudent("Charlie", 19, 75.0, []string{"Physics", "Chemistry"})

    // Display all student records
    displayStudents()

    // Update student details
    updateStudentByName("Alice", 19, 88.0, []string{"Math", "Science", "Computer Science"})

    // Display updated student records
    displayStudents()

    // Delete a student
    deleteStudentByName("Bob")

    // Display student records after deletion
    displayStudents()

    // Calculate average grade of all students
    avgGrade := calculateAverageGrade()
    fmt.Printf("Average grade of all students: %.2f\n", avgGrade)
}
