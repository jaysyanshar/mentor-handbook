package main

func main() {

}

// Student, Teacher, Study, Contact
// Student -> Study: Student do/take Study
// Teacher -> Study: Teacher teach Study
// Student -> Contact: Student have Contact
// Teacher -> Contact: Teacher have Contact
// Contact <- Student -> Study <- Teacher -> Contact

// Relasi DB: 1-1, 1-N, N-N

type OwnerType string

const (
	OwnerType_Student OwnerType = "student"
	OwnerType_Teacher OwnerType = "teacher"
)

type Student struct {
	ID    int // primary key
	Name  string
	Grade string
}

type Study struct {
	ID     int // primary key
	Name   string
	Credit int
}

type Teacher struct {
	ID      int // primary key
	StudyID int // foreign key
	Name    string
}

type Contact struct {
	ID          int // primary key
	ReferenceID int // foreign key
	OwnerType   OwnerType
	Phone       string
	Email       string
}

type StudentStudy struct {
	ID        int // primary key
	StudentID int // foreign key
	StudyID   int // foreign key
}

// StudentID=1, StudyID=1
// StudentID=1, StudyID=2
