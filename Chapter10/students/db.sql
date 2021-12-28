CREATE TABLE students (
id integer NOT NULL,
firstname char(256) NOT NULL,
middlename char(256) NOT NULL,
lastname char(256) NOT NULL,
class char(128) NOT NULL,
course char(128) NOT NULL);

INSERT INTO students (id, firstname, middlename, lastname, class, course) VALUES
('10149', 'Frank', 'Vincent', 'Zappa', '3A', 'Composition');

ALTER TABLE students ADD PRIMARY KEY (id);


