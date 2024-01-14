package funcs

import (
	"crud/moduls"
	"database/sql"
)

// this func deleted Student from table on database
func DeleteStudent(db *sql.DB, student *moduls.Student) error {

	_, err := db.Exec(`DELETE FROM it_center WHERE student_id = $1`, student.Id)

	if err != nil {
		return err
	}

	_, err = db.Exec(`DELETE FROM student WHERE id = $1`, student.Id)

	if err != nil {
		return err
	}

	return nil
}

// this func deleted Course from table on database
func DeleteCourse(db *sql.DB, course *moduls.Course) error {
	_, err := db.Exec(`DELETE FROM it_center WHERE course_id = $1`, course.Id)

	if err != nil {
		return err
	}

	_, err = db.Exec(`DELETE FROM course WHERE id = $1`, course.Id)

	if err != nil {
		return err
	}
	return nil
}
