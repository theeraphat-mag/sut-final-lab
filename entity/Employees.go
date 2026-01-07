package entity

import (
	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	Name		string		`valid:"length(2|80)~Name is length 2-80"`
	Salary		float64		`valid:"range(15000|200000)~Salary must be between 15000 and 200000"`
	EmployeeCode string		`valid:"matches(^[A-Z][-][0-9]{4}$)~EmployeeCode must be 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"`
}