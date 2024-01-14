package funcs

import (
	"crud/moduls"
	"database/sql"
)

// this func inserts Student to database
func CreateStudent(db *sql.DB, student *moduls.Student) (*moduls.Student, error) {
	result := db.QueryRow(`INSERT INTO student (f_name, l_name, age) VALUES ($1, $2, $3) RETURNING id, f_name, l_name, age `,
		student.FirstName,
		student.LastName,
		student.Age)

	var respostuden moduls.Student

	err := result.Scan(&respostuden.Id, &respostuden.FirstName, &respostuden.LastName, &respostuden.Age)

	if err != nil {
		return &moduls.Student{}, err
	}

	for _, id := range student.CourseId {
		_, err := db.Exec(`INSERT INTO it_center (course_id, student_id) VALUES ($1, $2)`, id, respostuden.Id)

		if err != nil {
			return &moduls.Student{}, err
		}
		respostuden.CourseId = student.CourseId
	}

	return &respostuden, nil

}

// this func insert Course to database
func CreateCourse(db *sql.DB, course *moduls.Course) (*moduls.Course, error) {

	query := `
  INSERT INTO course (name, teacher, price)
  VALUES ($1, $2, $3) RETURNING id, name, teacher, price`

  // run the query
  CourseRows := db.QueryRow(query, course.Name, course.Teacher, course.Price)

  // copy result of the query
  var respocourse moduls.Course
  err := CourseRows.Scan(
    &respocourse.Id,
    &respocourse.Name,
    &respocourse.Teacher,
    &respocourse.Price,
  )


	// result := db.QueryRow(`
	// INSERT INTO course (name, teacher, price) 
	// VALUES ($1, $2, $3) RETURNING id, name, teacher, price`,
	// 	course.Name,
	// 	course.Teacher,
	// 	course.Price)

	// var respocourse moduls.Course

	// err := result.Scan(
	// 	&respocourse.Id, 
	// 	&respocourse.Name, 
	// 	&respocourse.Teacher,
	// 	&respocourse.Price, 
	// )

	if err != nil {
		panic(err)
		// return &moduls.Course{}, err
	}

	for _, id := range course.StudentId {
		_, err := db.Exec(`INSERT INTO it_center(course_id, student_id) VALUES ($1, $2)`, respocourse.Id, id)

		if err != nil {
			panic(err)
			// return &moduls.Course{}, err
		}
		respocourse.StudentId = course.StudentId
	}
	return &respocourse, nil
}
