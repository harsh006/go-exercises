package main

import "fmt"

type Matrix struct{
	rows int
	columns int
	elements [][]int
}

func (m Matrix)No_Rows() int{
	return m.rows
}
func (m Matrix)No_Cols() int{
	return m.columns
}
func (m *Matrix)Set_Element(i,j,value int){
	m.elements[i][j] = value
}

func (m *Matrix) add(m2 Matrix){ // cases of unequal dimensions not covered.
	// This adds the other matrix directly to main called matrix
	for i:=0;i<m.No_Rows();i++{
		for j:=0;j<m.No_Cols();j++{
			m.elements[i][j] += m2.elements[i][j]
		}
	}
}
func main() {
	var mat Matrix = Matrix{2,2,[][]int{{1,1},{1,1}}}
	mat.add(Matrix{2,2,[][]int{{1,1},{1,1}}})
	fmt.Println(mat)
}
