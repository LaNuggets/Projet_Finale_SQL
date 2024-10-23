package Projet_Final_SQL

import (
    "database/sql"
    "net/http"
)


func GetEmployeesByDepartment(w http.ResponseWriter, r *http.Request, departmentID string) []CompletEmployee {
    db, err := sql.Open("sqlite3", "bdd.db?_foreign_keys=on")
    CheckErr(err, w, r)
    defer db.Close()
    
    

    rows, err := db.Query("SELECT e.Id, e.LastName, e.FirstName, e.Birthday, e.Phone, e.Address, e.PostId, e.referentId, d.Title, s.Wage FROM employees e LEFT JOIN posts p ON e.PostId = p.Id LEFT JOIN department d ON p.DepartmentId = d.Id LEFT JOIN salary s ON p.SalaryId = s.Id WHERE d.Id = ?", departmentID)
    CheckErr(err, w, r)
    defer rows.Close()

    var employeesList []CompletEmployee
    for rows.Next() {
        var employee CompletEmployee
        err = rows.Scan(&employee.Id, &employee.LastName, &employee.FirstName, &employee.Birthday, &employee.Phone, &employee.Address, &employee.PostId, &employee.ReferentId, &employee.Department, &employee.Wage)
        CheckErr(err, w, r)
        employeesList = append(employeesList, employee)
    }

    return employeesList
}