package Models

type User struct {
	Id         uint    `json:"id"`
	First_Name string  `json:"first_name"`
	Last_Name  string  `json:"last_name"`
	DOB        string  `json:"date_of_birth"`
	Address    string  `json:"address"`
	subject    Subject `json:"subject_marks"`
	//Marks   `json:"marks_obtained"`
}
type Subject struct{
	maths , physics, chemistry int
}
//Id
//Name
//Last Name
//DOB
//Address
//Subject
//Marks
func (b *User) TableName() string {
	return "user"
}
