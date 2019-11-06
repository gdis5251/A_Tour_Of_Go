# Go 语言操作 MySQL 数据库

## 起步

- 首先需要引入一个驱动，在命令行使用 go get 引入即可：

  ```
  go get github.com/go-sql-driver/mysql
  ```

- 然后在 import 的时候，只需要该驱动的 init() 函数，所以可以在该驱动前加 _

- 最后使用 database/sql 来操作数据库

  ```go
  import (
  	"database/sql"
  	"fmt"
  	_ "github.com/go-sql-driver/mysql"
  )
  ```

## 操作

完整代码如下，其中需要一些函数，理解记住就好。

```go
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

type student struct {
	id   int
	name string
	sex  string
	age  int
}

// 插入数据
func insertInfo(DB *sql.DB) {
	stmt, prepareErr := DB.Prepare("insert into students(name, sex, age) values(?, ?, ?)")
	if prepareErr != nil {
		fmt.Println("prepare insert query error : ", prepareErr)
		return
	}

	// 将信息写入数据库
	result, execErr := stmt.Exec("apple", "female", 18)
	if execErr != nil {
		fmt.Println("exec error : ", execErr)
		return
	}

	// 返回值 id 为，上面的操作，影响了几行
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
	stmt, prepareErr := DB.Prepare("update students set name=?, sex=?, age=? where id=?")
	if prepareErr != nil {
		fmt.Println("prepare update error : ", prepareErr)
		return
	}

	// 再实际更新
	result, execErr := stmt.Exec("lili", "female", 17, 3)
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

// 删除数据
func DeleteInfo(DB *sql.DB) {
	stmt, prepareErr := DB.Prepare("delete from students where id=? or id=?")
	if prepareErr != nil {
		fmt.Println("prepare delete failed, error : ", prepareErr)
		return
	}

	result, execErr := stmt.Exec(2, 3)
	if execErr != nil {
		fmt.Println("delete failed, error : ", execErr)
		return
	}

	id, affectedErr := result.RowsAffected()
	if affectedErr != nil {
		fmt.Println("get affected row failed, error : ", affectedErr)
		return
	}

	fmt.Println("delete affect row is ", id)

}

// 单行查询
func QueryOne(DB *sql.DB) {
	student := new(student)
	row := DB.QueryRow("select * from students where id=?", 1)
	if scanErr := row.Scan(&student.id, &student.name, &student.sex, &student.age); scanErr != nil {
		fmt.Println("scan error : ", scanErr)
		return
	}

	fmt.Println(*student)
}

// 多行查询
func QueryMutil(DB *sql.DB) {
	student := new(student)
	rows, queryErr := DB.Query("select * from students where id>?", 1)
	// 写一个 defer 函数，如果最后 rows 没有显示完，就关闭 row
	// 主要是 如果发生错误，就把已经加载进来的关闭掉，防止内存泄漏
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	// 错误检查
	if queryErr != nil {
		fmt.Println("query mutil failed, error : ", queryErr)
		return
	}

	for rows.Next() {
		if scanErr := rows.Scan(&student.id, &student.name, &student.sex, &student.age); scanErr != nil {
			fmt.Println("scan failed, error : ", scanErr)
			return
		}

		fmt.Println(*student)
	}
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
		CHARSET  string = "charset=utf8"
	)
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE, CHARSET)

	// 连接数据库
	DB, openErr := sql.Open("mysql", dsn)
	if openErr != nil {
		fmt.Println("open mysql err : ", openErr)
		return
	}

	// insertInfo(DB)
	// UpdateInfo(DB)
	// DeleteInfo(DB)
	// QueryOne(DB)
	QueryMutil(DB)
}

```

