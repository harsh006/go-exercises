package main

import (
	"fmt"
)

type Employee interface {
	Salary_Calculator()
}

type FT struct{
	worktime int
	Pay int
}
type Contractor struct{
	worktime int
	Pay int
}
type Freelancer struct{
	worktime int
	Pay int
}
func (f *FT)Salary_Calculator(){
	fmt.Println(f.worktime*f.Pay)
}
func (f *Contractor)Salary_Calculator(){
	fmt.Println(f.worktime*f.Pay)
}
func (f *Freelancer)Salary_Calculator(){
	fmt.Println(f.worktime*f.Pay)
}

func main() {
	employees := []Employee{&FT{30,500},&Contractor{30,100},&Freelancer{20,10}}
	for _, e := range employees{
		fmt.Printf("The salary of %T is ",e)
		e.Salary_Calculator()
	}
}
