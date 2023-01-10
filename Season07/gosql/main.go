package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "192.168.0.77"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "latihandb"
)

var (
	err error
	db  *sql.DB
)

/*
create table employees (

	id SERIAL primary key ,
	full_name VARCHAR(50) not null,
	email VARCHAR(50) not null,
	age INT not null,
	division varchar(20) not null

);
*/
type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func CreateEmployee() {
	empl := Employee{}
	sql := `INSERT INTO employees (full_name,email,age,division) 
			VALUES ($1, $2, $3, $4) 
			Returning *
			`
	err = db.QueryRow(sql, "IKetut ", "iketutg@gmail.com", 45, "Engineer1").
		Scan(&empl.ID,
			&empl.Full_name,
			&empl.Email,
			&empl.Age,
			&empl.Division)
	if err != nil {
		fmt.Println("Create Employee Error : ", err.Error())
		return
	}
	fmt.Println("Create Emplyee Success : ", empl)
}
func GetEmployee() {
	var empl = []Employee{}
	sql := `SELECT * FROM employees`
	rows, err2 := db.Query(sql)
	if err2 != nil {
		fmt.Println("Error Query : ", err2.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var emp = Employee{}
		err2 = rows.Scan(&emp.ID,
			&emp.Full_name,
			&emp.Email,
			&emp.Age,
			&emp.Division)
		if err2 != nil {
			fmt.Println("Error rows next : ", err.Error())
		} else {
			empl = append(empl, emp)
		}
	}
	fmt.Println("Data Emplyee : ", empl)
}

func UpdateEmployee() {
	sql := `UPDATE employees 
			SET full_name= $1,
			 division = $2, 
			 email=$3,  
             age=$4 
			WHERE id= $5`

	result, err2 := db.Exec(sql, "Joni", "HRD", "Joni@joni.com", 30, 2)

	if err2 != nil {
		fmt.Println("Error Execute Query : ", err2)
	}
	totoalrecord, err3 := result.RowsAffected()

	if err3 != nil {
		fmt.Println("Error Update : ", err3.Error())
	}

	fmt.Println("Total Record : ", totoalrecord)
}
func DeleteRecord() {
	sql := "DELETE FROM employees WHERE id=$1"

	result, err2 := db.Exec(sql, 4)

	if err2 != nil {
		fmt.Println("Error Execute Query : ", err2)
	}
	totoalrecord, err3 := result.RowsAffected()

	if err3 != nil {
		fmt.Println("Error Update : ", err3.Error())
	}

	fmt.Println("Total Record : ", totoalrecord)
}
func main() {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable connect_timeout=5",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", dataSourceName)
	defer db.Close()
	if err != nil {
		fmt.Println("Error : ", err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error : ", err)
		panic(err)
	}
	fmt.Println("Connection Success")
	CreateEmployee()
	GetEmployee()
	UpdateEmployee()
	DeleteRecord()
}
