package funcs

import (
	"crud/moduls"
	"database/sql"
)

// this func gets all Students with all Courses
func GetAllStudent(db *sql.DB) ([]*moduls.Student, error) {
	var all_student []*moduls.Student

	result, err := db.Query(`SELECT id, f_name, l_name, age FROM student`)

	if err != nil {
		return []*moduls.Student{}, err
	}

	for result.Next() {
		var respostudent moduls.Student
		err := result.Scan(
			&respostudent.Id,
			&respostudent.FirstName,
			&respostudent.LastName,
			&respostudent.Age)
		if err != nil {
			return []*moduls.Student{}, err
		}

		it_result, err := db.Query(`SELECT course_id FROM it_center WHERE student_id = $1`, respostudent.Id)

		if err != nil {
			return []*moduls.Student{}, err
		}

		var arr []int
		for it_result.Next() {
			var id int
			err := it_result.Scan(&id)

			if err != nil {
				return []*moduls.Student{}, err
			}

			arr = append(arr, id)
		}

		respostudent.CourseId = arr

		all_student = append(all_student, &respostudent)

	}

	return all_student, nil
}

// this func gets all Courses with all Students
func GetAllCourse(db *sql.DB) ([]*moduls.Course, error) {

	var all_course []*moduls.Course

	result, err := db.Query(`SELECT id, name, price, teacher FROM course`)

	if err != nil {
		return []*moduls.Course{}, err
	}

	for result.Next() {
		var respocourse moduls.Course
		err := result.Scan(
			&respocourse.Id,
			&respocourse.Name,
			&respocourse.Price,
			&respocourse.Teacher)
		if err != nil {
			return []*moduls.Course{}, err
		}

		it_result, err := db.Query(`SELECT student_id FROM it_center WHERE course_id = $1`, respocourse.Id)

		if err != nil {
			return []*moduls.Course{}, err
		}

		var arr []int
		for it_result.Next() {
			var id int
			err := it_result.Scan(&id)

			if err != nil {
				return []*moduls.Course{}, err
			}

			arr = append(arr, id)
		}

		respocourse.StudentId = arr

		all_course = append(all_course, &respocourse)

	}

	return all_course, nil
}
