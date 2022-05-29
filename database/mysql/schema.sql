CREATE DATABASE IF NOT EXISTS Todolist;
use Todolist;

CREATE TABLE Todo (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE Student (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    blocked BOOL NOT NULL,
    role VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE TaskList (
    todoId INT NOT NULL,
    studentId INT NOT NULL,
    isFinished BOOL NOT NULL,
    FOREIGN KEY (todoId) REFERENCES Todo(id) ON DELETE SET NULL,
    FOREIGN KEY (studentId) REFERENCES Student(id) ON DELETE SET NULL
)