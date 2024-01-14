package funcs

import (
	"crud/moduls"
	"database/sql"
)

// this func gets Student with all Courses
func GetStudent(db *sql.DB, std_id int) (*moduls.Student, error) {
	result := db.QueryRow(`SELECT id, f_name, l_name, age FROM student WHERE id = $1`, std_id)

	var respostuden moduls.Student

	err := result.Scan(&respostuden.Id, &respostuden.FirstName, &respostuden.LastName, &respostuden.Age)

	if err != nil {
		return &moduls.Student{}, err
	}

	it_result, err := db.Query(`SELECT course_id FROM it_center WHERE student_id = $1`, std_id)

	if err != nil {
		return &moduls.Student{}, err
	}

	for it_result.Next() {
		var newId int
		err := it_result.Scan(&newId)

		if err != nil {
			return &moduls.Student{}, err
		}
		respostuden.CourseId = append(respostuden.CourseId, newId)
	}

	return &respostuden, nil
}

// this func gets Course with all Students
func GetCourse(db *sql.DB, crs_id int) (*moduls.Course, error) {
	result := db.QueryRow(`SELECT id, name, price, teacher FROM course WHERE id = $1`, crs_id)

	var respocourse moduls.Course

	err := result.Scan(&respocourse.Id, &respocourse.Name, &respocourse.Price, &respocourse.Teacher)

	if err != nil {
		return &moduls.Course{}, err
	}

	it_result, err := db.Query(`SELECT student_id FROM it_center WHERE course_id = $1`, crs_id)

	for it_result.Next() {
		var newId int
		err := it_result.Scan(&newId)

		if err != nil {
			return &moduls.Course{}, err
		}

		respocourse.StudentId = append(respocourse.StudentId, newId)

	}

	if err != nil {
		return &moduls.Course{}, err
	}

	return &respocourse, nil
}
