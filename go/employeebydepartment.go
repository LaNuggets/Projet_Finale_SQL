package Projet_Final_SQL

import (
    "database/sql"
    "net/http"
)


func GetEmployeesByDepartment(w http.ResponseWriter, r *http.Request, departmentID string) []CompletEmployee {
    db, err := sql.Open("sqlite3", "bdd.db?_foreign_keys=on")
    CheckErr(err, w, r)
    defer db.Close()
    
    

    rows, err := db.Query(`SELECT DISTINCT e.Id, e.LastName, e.FirstName, e.Birthday, e.Phone, e.Address, p.Title AS PostTitle, d.Title AS Department, s.Wage, CONCAT(rf.FirstName, ' ', rf.LastName) AS ReferentName
	FROM employees e 
	LEFT JOIN posts p ON e.PostId = p.Id 
	LEFT JOIN department d ON p.DepartmentId = d.Id 
	LEFT JOIN salary s ON p.SalaryId = s.Id 
	LEFT JOIN referent r ON e.referentId = r.ReferentId 
	LEFT JOIN employees rf ON r.ReferentId = rf.Id
    WHERE d.Id = ?`, departmentID)
    CheckErr(err, w, r)
    defer rows.Close()

    var employeesList []CompletEmployee
    for rows.Next() {
        var employee CompletEmployee
		err = rows.Scan(&employee.Id, &employee.LastName, &employee.FirstName, &employee.Birthday, &employee.Phone, &employee.Address, &employee.PostTitle, &employee.Department, &employee.Wage, &employee.ReferentName)
        CheckErr(err, w, r)
        employeesList = append(employeesList, employee)
    }

    return employeesList
}