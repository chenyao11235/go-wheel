package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID   int64          `db:"id"`
	Name sql.NullString `db:"name"`
	Age  int            `db:"age"`
}

const (
	USERNAME = "root"
	PASSWORD = "cy.89757"
	NETWORK  = "tcp"
	SERVER   = "10.211.55.5"
	PORT     = 3306
	DATABASE = "dispatcher"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second) // 设置最大链接周期，超时自动关闭链接
	DB.SetMaxOpenConns(2)                    // 设置最大链接数，链接池中的数量
	DB.SetMaxIdleConns(16)                   // 设置闲置链接数

	queryOne(DB)
	queryMulti(DB)
	insertData(DB)
	updateData(DB)
	deleteData(DB)

}

// 注意 row必须scan，否则会导致链接不释放

//查询单行
func queryOne(DB *sql.DB) {
	user := new(User)
	row := DB.QueryRow("select * from users where id=?", 1)
	if err := row.Scan(&user.ID, &user.Name, &user.Age); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}

	fmt.Println(*user)
}

//查询多行
// 查询多行的时候rows如果不遍历或者没有遍历完(查看Next源码可知)，该查询操作会一直占用链接不释放，所有可以用defer手工释放
func queryMulti(DB *sql.DB) {
	user := new(User)
	rows, err := DB.Query("select * from users where id > ?", 1)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}
		fmt.Print(*user)

		time.Sleep(5 * time.Second)
	}
}

//插入数据
func insertData(DB *sql.DB) {
	result, err := DB.Exec("insert INTO users(name,age) values(?,?)", "YDZ", 23)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return
	}
	fmt.Println("LastInsertID:", lastInsertID)
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

//更新数据
func updateData(DB *sql.DB) {
	result, err := DB.Exec("UPDATE users set age=? where id=?", "30", 3)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}

//删除数据
func deleteData(DB *sql.DB) {
	result, err := DB.Exec("delete from users where id=?", 1)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Println("RowsAffected:", rowsaffected)
}
