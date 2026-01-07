package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositiveEmployeeCode(t *testing.T) {
	g := NewWithT(t)

	t.Run("Employee test negative EmployeeCode", func(t *testing.T) {
		employees := Employees{
			Name:         "mag",
			Salary:       15000.0,
			EmployeeCode: "a-1234",
		}

		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"))
	})

	t.Run("Employee test negative EmployeeCode", func(t *testing.T) {
		employees := Employees{
			Name:         "mag",
			Salary:       15000.0,
			EmployeeCode: "A1234",
		}

		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"))
	})
}
