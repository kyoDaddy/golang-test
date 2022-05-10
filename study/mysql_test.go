package study

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "log"
	"math/rand"
	_ "math/rand"
	_ "strconv"
	"testing"
	_ "time"
)

func TestMysql(t *testing.T) {
	// sql.DB 객체 생성ß
	db, err := sql.Open("mysql", "root:mariadb@tcp(127.0.0.1:3306)/sql_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 하나의 Row를 갖는 SQL 쿼리
	var usersId int
	err = db.QueryRow("SELECT users_id as usersId from Users WHERE users_id = 1").Scan(&usersId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usersId)

	// INSERT 문 실행
	usersId = rand.Intn(100000)
	result, err := db.Exec("INSERT INTO Users (users_id, banned, `role`) VALUES(?, ?, ?)", usersId, "No", "client")
	if err != nil {
		log.Fatal(err)
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}

	// Prepared Statement 생성
	stmt, err := db.Prepare("UPDATE Users SET `role`=? WHERE users_id = ?")
	checkError(err)
	defer stmt.Close()

	// Prepared Statement 실행
	var updateCount int
	updateCount = 0
	result, err = stmt.Exec("partner", 1) // Placeholder 파라미터 순서대로 전달
	checkError(err)
	updateCount = updateCounting(result, err, updateCount)

	result, err = stmt.Exec("driver", usersId)
	checkError(err)
	updateCount = updateCounting(result, err, updateCount)
	fmt.Println("update count ", updateCount)

	// 복수 Row를 갖는 SQL 쿼리
	var id int
	var role string
	rows, err := db.Query("SELECT users_id as id, role from Users WHERE users_id > ?", 0)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // 반드시 닫는다 (지연하여 닫기)

	for rows.Next() {
		err := rows.Scan(&id, &role)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, role)
	}

}

func updateCounting(result sql.Result, err error, count int) int {
	// RowsAffected() 함수를 통해 update 한 갯수를 확인한다.
	nRow, err := result.RowsAffected()
	if nRow == 1 {
		count++
	}
	return count
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
