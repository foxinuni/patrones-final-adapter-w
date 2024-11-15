package main

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Employee struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Position  string  `json:"position"`
	Salary    float64 `json:"salary"`
}

type EmployeeCreateDTO struct {
	FirstName string  `json:"first_name" validate:"required"`
	LastName  string  `json:"last_name" validate:"required"`
	Position  string  `json:"position" validate:"required"`
	Salary    float64 `json:"salary" validate:"required"`
}

type UpdateEmployeeDTO struct {
	FirstName *string  `json:"first_name"`
	LastName  *string  `json:"last_name"`
	Position  *string  `json:"position"`
	Salary    *float64 `json:"salary"`
}

var employees []Employee

func CreateEmployee(employee *Employee) {
	employee.ID = len(employees) + 1
	employees = append(employees, *employee)
}

func GetEmployees() []Employee {
	return employees
}

func GetEmployeeById(id int) (*Employee, error) {
	for _, employee := range employees {
		if employee.ID == id {
			return &employee, nil
		}
	}
	return nil, errors.New("Employee not found")
}

func UpdateEmployee(id int, updateEmployee *UpdateEmployeeDTO) (*Employee, error) {
	for index, employee := range employees {
		if employee.ID == id {
			if updateEmployee.FirstName != nil {
				employees[index].FirstName = *updateEmployee.FirstName
			}
			if updateEmployee.LastName != nil {
				employees[index].LastName = *updateEmployee.LastName
			}
			if updateEmployee.Position != nil {
				employees[index].Position = *updateEmployee.Position
			}
			if updateEmployee.Salary != nil {
				employees[index].Salary = *updateEmployee.Salary
			}

			return &employees[index], nil
		}
	}
	return nil, errors.New("employee not found")
}

func DeleteEmployee(id int) error {
	for index, employee := range employees {
		if employee.ID == id {
			employees = append(employees[:index], employees[index+1:]...)
			return nil
		}
	}
	return errors.New("Employee not found")
}

func init() {
	employees = []Employee{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Position:  "Software Engineer",
			Salary:    3500.00,
		},
		{
			ID:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Position:  "Project Manager",
			Salary:    4500.00,
		},
		{
			ID:        3,
			FirstName: "Michael",
			LastName:  "Smith",
			Position:  "Data Analyst",
			Salary:    3000.00,
		},
		{
			ID:        4,
			FirstName: "Maria",
			LastName:  "Garcia",
			Position:  "Software Engineer",
			Salary:    3500.00,
		},
	}
}

func main() {
	app := echo.New()

	app.GET("/employees", func(c echo.Context) error {
		return c.JSON(200, GetEmployees())
	})

	app.GET("/employees/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, "Invalid ID")
		}

		employee, err := GetEmployeeById(id)
		if err != nil {
			return c.JSON(404, "Employee not found")
		}

		return c.JSON(200, employee)
	})

	app.POST("/employees", func(c echo.Context) error {
		dto := new(EmployeeCreateDTO)
		if err := c.Bind(dto); err != nil {
			return c.JSON(400, "Invalid request")
		}

		employee := &Employee{
			FirstName: dto.FirstName,
			LastName:  dto.LastName,
			Position:  dto.Position,
			Salary:    dto.Salary,
		}

		CreateEmployee(employee)
		return c.JSON(201, employee)
	})

	app.PUT("/employees/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, "Invalid ID")
		}

		dto := new(UpdateEmployeeDTO)
		if err := c.Bind(dto); err != nil {
			return c.JSON(400, "Invalid request")
		}

		employee, err := UpdateEmployee(id, dto)
		if err != nil {
			return c.JSON(404, "Employee not found")
		}

		return c.JSON(200, employee)
	})

	app.DELETE("/employees/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(400, "Invalid ID")
		}

		err = DeleteEmployee(id)
		if err != nil {
			return c.JSON(404, "Employee not found")
		}

		return c.NoContent(204)
	})

	app.Logger.Fatal(app.Start(":8080"))
}
