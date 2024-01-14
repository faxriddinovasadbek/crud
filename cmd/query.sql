CREATE DATABASE najot_talimdb;

\c najot_talimdb

CREATE TABLE course (
    id SERIAL PRIMARY KEY,
    name VARCHAR(32),
    price INT,
    teacher VARCHAR(32)
);

INSERT INTO course (name, price, teacher) VALUES
    ('Go', 75, 'Abbos Azimov'),
    ('Java', 90, 'Javlon Mirzayev'),
    ('QA', 80, 'Raxim Hoshimiv'),
    ('Foundation', 100, 'Erkin Tolipov'),
    ('Frontend', 85, 'Doston Jabborov'),
    ('English For It', 95, 'Azim Azizov'),
    ('Python', 70, 'Diyor Sunnatov');


CREATE TABLE student (
    id SERIAL PRIMARY KEY,
    f_name VARCHAR(32),
    l_name VARCHAR(32),
    age INT
);

INSERT INTO student (f_name, l_name, age) VALUES
    ('Anvar', 'Ortiqov', 20),
    ('Dilshod', 'Ismoilov', 22),
    ('Malika', 'Nizomova', 18),
    ('Zilola', 'Xudaybergenova', 21),
    ('Sardor', 'Xayrullayev', 19),
    ('Shaxzoda', 'Xamidova', 23),
    ('Javlon', 'Usmonov', 20),
    ('Nodir', 'Baxramov', 21),
    ('Nigora', 'Murodova', 22),
    ('Ravshan', 'Qosimov', 19),
    ('Gulnoza', 'Abdullayeva', 20),
    ('Oybek', 'Rahmonov', 21),
    ('Shirin', 'Xudaybergenova', 18),
    ('Sunnat', 'Ibrohimov', 22),
    ('Lobar', 'Hamroeva', 20);


CREATE TABLE it_center (
    course_id INT, FOREIGN KEY(course_id) REFERENCES course(id),
    student_id INT, FOREIGN KEY(student_id) REFERENCES student(id)
);


INSERT INTO it_center (course_id, student_id) VALUES
    (1, 1),
    (2, 2),
    (3, 3),
    (4, 4),
    (5, 5),
    (6, 6),
    (7, 7),
    (1, 8),
    (2, 9),
    (3, 10),
    (4, 11),
    (4, 12),
    (4, 13),
    (4, 14),
    (4, 1),
    (5, 12);
