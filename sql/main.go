package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

type Emp struct {
	empNo   int
	empName string
	job     string
}

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:19990414@tcp(127.0.0.1:3306)/dljd?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select empno, ename, job from emp where empno = %s", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var e Emp
	err := db.QueryRow(sqlStr).Scan(&e.empNo, &e.empName, &e.job)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	fmt.Printf("user:%#v\n", e)
}

func transactionDemo() {
	// 创建事务 transaction： 先删除，再创建，后更新
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			// 事务已经创建，先回滚
			tx.Rollback()
		}
		fmt.Printf("begin transaction failed, err: %v", err)
		return
	}
	fmt.Println("transaction begin...")
	// 删除
	sqlStr := "delete from emp_tmp where empno = ?"
	res, err := db.Exec(sqlStr, 9900)
	if err != nil {
		tx.Rollback()
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	affRow0, err := res.RowsAffected() // 操作影响的行数
	if err != nil {
		tx.Rollback()
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", affRow0)

	// 插入新数据
	sqlStr = "insert into emp_tmp(empno, ename, job, sal) values (?,?,?,?)"
	res, err = db.Exec(sqlStr, 9900, "cj", "boss", 9999)
	if err != nil {
		tx.Rollback()
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	affRow1, err := res.RowsAffected() // 操作影响的行数
	if err != nil {
		tx.Rollback()
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, affected rows:%d\n", affRow1)

	// 更新插入的新数据
	sqlStr = "update emp_tmp set DEPTNO = 10 where empno = ?"
	res, err = db.Exec(sqlStr, 9900)
	if err != nil {
		tx.Rollback()
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	affRow2, err := res.RowsAffected() // 操作影响的行数
	if err != nil {
		tx.Rollback()
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", affRow2)

	// 根据插入和更新时的影响的行数，决定是否提交事务
	if affRow1 == 1 && affRow2 == 1 {
		tx.Commit()
		fmt.Println("transaction commit...")
	} else {
		tx.Rollback()
		fmt.Println("transaction rollback...")
	}

	// 查询多行数据，此查询不包含在事务内
	sqlStr = "select empno, ename, job from emp_tmp where sal > ?"
	rows, err := db.Query(sqlStr, 5000)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var e Emp
		err := rows.Scan(&e.empNo, &e.empName, &e.job)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("Query- empno:  %d ename: %s  job: %s\n", e.empNo, e.empName, e.job)
	}

	// 查询一行数据
	sqlStr = "select empno, ename, job from emp_tmp where empno = ?" // 可能返回多行数据，但只取第一行
	var e Emp
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err = db.QueryRow(sqlStr, 9900).Scan(&e.empNo, &e.empName, &e.job)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Printf("QueryRow- empno: %d  ename: %s  job: %s\n", e.empNo, e.empName, e.job)
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	transactionDemo()
}
