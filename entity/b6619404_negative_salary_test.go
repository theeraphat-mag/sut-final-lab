package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositiveSalary(t *testing.T) {
	g := NewWithT(t)

	t.Run("Employee test negative salary", func(t *testing.T) {
		employees := Employees{
			Name:         "mag",
			Salary:       150.0,
			EmployeeCode: "A-1234",
		}

		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
	})

	t.Run("Employee test negative salary > 200000", func(t *testing.T) {
		employees := Employees{
			Name:         "mag",
			Salary:       1000000.0,
			EmployeeCode: "A-1234",
		}

		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).ToNot(BeTrue())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
	})
}
