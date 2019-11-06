/*
@Time : 2019-11-6 下午 3:04
@Author : Gerald
@File : CURD.go
@Software: GoLand
*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 插入数据
func insertInfo(DB *sql.DB) {
	stmt, prepareErr := DB.Prepare("insert into students(name, sex, age) values(?, ?, ?)")
	if prepareErr != nil {
		fmt.Println("prepare insert query error : ", prepareErr)
		return
	}

	// 将信息写入数据库
	result, execErr := stmt.Exec("Gerald", "male", 22)
	if execErr != nil {
		fmt.Println("exec error : ", execErr)
		return
	}

	id, rowsAffectedErr := result.RowsAffected()
	if rowsAffectedErr != nil {
		fmt.Println("insert faild : ", rowsAffectedErr)
		return
	}

	fmt.Println("affected row is ", id)
}

// 更新数据
func UpdateInfo(DB *sql.DB) {
	// 先做准备
	stmt, prepareErr := DB.Prepare("update students set name=? sex=? age=? where id=?")
	if prepareErr != nil {
		fmt.Println("prepare update error : ", prepareErr)
		return
	}

	// 再实际更新
	result, execErr := stmt.Exec("二丫", "female", 16, 2)
	if execErr != nil {
		fmt.Println("exec update error : ", execErr)
		return
	}

	// 再打印影响行数
	id, affectedErr := result.RowsAffected()
	if affectedErr != nil {
		fmt.Println("affected row error : ", affectedErr)
		return
	}

	fmt.Println("affect row is ", id)
}

func main() {
	// 定义一组 const 变量，准备连接数据库
	const (
		USERNAME string = "root"
		PASSWORD string = "peiban493."
		NETWORK  string = "tcp"
		SERVER   string = "localhost"
		PORT     int    = 3306
		DATABASE string = "test"
	)
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	// 连接数据库
	DB, openErr := sql.Open("mysql", dsn)
	if openErr != nil {
		fmt.Println("open mysql err : ", openErr)
		return
	}

	// insertInfo(DB)
	UpdateInfo(DB)
}
