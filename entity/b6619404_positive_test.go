package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewWithT(t)

	t.Run("Employee test Positive", func(t *testing.T) {
		employees := Employees{
			Name:         "mag",
			Salary:       15000,
			EmployeeCode: "A-1234",
		}

		ok, err := govalidator.ValidateStruct(employees)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

