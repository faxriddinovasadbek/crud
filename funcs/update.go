package funcs

import (
	"crud/moduls"
	"database/sql"
)

// this func updates Student from table on database
func UpdateStudent(db *sql.DB, student *moduls.Student) (*moduls.Student, error) {

	shablon :=
		`UPDATE
		student 
	SET age=age+1 
	WHERE 
		id = $1 
	RETURNING 
		id,
		f_name,
		l_name,
		age`

	result := db.QueryRow(shablon, student.Id)

	var respostudent moduls.Student
	err := result.Scan(&respostudent.Id, &respostudent.FirstName, &respostudent.LastName, &respostudent.Age)

	if err != nil {
		return &respostudent, err
	}

	return &respostudent, nil
}

// this func updates Course from table on database
func UpdateCourse(db *sql.DB, course *moduls.Course) (*moduls.Course, error) {

	shablon :=
		`UPDATE 
		course 
	SET 
		price = price + price * 0.1 
	WHERE 
		id = $1 
	RETURNING 
		id,
		name,
		price,
		teacher`

	result := db.QueryRow(shablon, course.Id)

	var respocourse moduls.Course

	err := result.Scan(&respocourse.Id, &respocourse.Name, &respocourse.Price, &respocourse.Teacher)

	if err != nil {
		return &respocourse, err
	}

	return &respocourse, nil
}
