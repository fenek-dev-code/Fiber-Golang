package vacancy

import (
	"fmt"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type VacancyCreateForm struct {
	Email    string `form:"email"`
	Name     string `form:"name"`
	Role     string `form:"role"`
	Type     string `form:"type"`
	Salary   int    `form:"salary"`
	Location string `form:"location"`
}

type FormParser interface {
	FormValue(key string, defaultValue ...string) string
}

func CreateVacancyForm(form FormParser) *VacancyCreateForm {
	return &VacancyCreateForm{
		Email: form.FormValue("email"),
		Name:  form.FormValue("name"),
		Role:  form.FormValue("role"),
		Type:  form.FormValue("type"),
		Salary: func() int {
			salaryStr := form.FormValue("salary")
			var salary int
			_, err := fmt.Sscanf(salaryStr, "%d", &salary)
			if err != nil {
				return 0
			}
			return salary
		}(),
		Location: form.FormValue("location"),
	}
}

func (f *VacancyCreateForm) IsValid() *validate.Errors {
	return validate.Validate(&validators.EmailIsPresent{
		Name: "Email", Field: f.Email, Message: "Введите корректный email",
	}, &validators.StringIsPresent{
		Name: "Name", Field: f.Name, Message: "Введите название компании",
	}, &validators.StringIsPresent{
		Name: "Role", Field: f.Role, Message: "Введите должность",
	}, &validators.StringIsPresent{
		Name: "Type", Field: f.Type, Message: "Введите тип занятости",
	}, &validators.IntIsPresent{
		Name: "Salary", Field: f.Salary, Message: "Введите зарплату",
	}, &validators.StringIsPresent{
		Name: "Location", Field: f.Location, Message: "Введите локацию",
	})
}
