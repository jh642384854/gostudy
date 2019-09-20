package main

import (
	"dev/day5/db/util"
	"fmt"
	"log"
)

//Exec()方法，可以执行任意SQL语句。当然也就包括简单的增删改查
func dbdemo1()  {
	db := util.GetDbInstance()
	insertSql := "INSERT INTO `jh_articles` (`title`, `author`) VALUES ( 'title2', 'zhaoliu')"
	res,err := db.Exec(insertSql)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(res.RowsAffected())  //获得受影响的行数
	fmt.Println(res.LastInsertId())  //获取最后的记录ID
}

//Prepare()方法
/*
	一般用Prepared Statements和Exec()完成INSERT, UPDATE, DELETE操作。
*/
func dbdemo2()  {
	db := util.GetDbInstance()
	insertSql := "INSERT INTO `jh_articles` (`title`, `author`) VALUES ( ?, ?)"
	stmt,err := db.Prepare(insertSql)
	if err != nil{
		log.Fatal(err)
	}
	res ,err := stmt.Exec("title3","zhaoliu");
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(res.RowsAffected())  //获得受影响的行数
	fmt.Println(res.LastInsertId())  //获取最后的记录ID
}

//查询单条操作
func dbdemo3()  {
	var article Article
	db := util.GetDbInstance()
	selectSql := "SELECT * FROM `jh_articles` WHERE id = ?"
	row := db.QueryRow(selectSql,3)
	if err := row.Scan(&article.Id,&article.Title,&article.Author); err != nil{
		log.Fatal(err)
	}
	fmt.Println(article)
}

//查询多条记录
/**
	上面代码的过程为：db.Query()表示向数据库发送一个query，defer rows.Close()非常重要，遍历rows使用rows.Next()，
	把遍历到的数据存入变量使用rows.Scan(), 遍历完成后检查error。有几点需要注意：
	①、检查遍历是否有error
	②、结果集(rows)未关闭前，底层的连接处于繁忙状态。当遍历读到最后一条记录时，会发生一个内部EOF错误，自动调用rows.Close()，但是如果提前退出循环，
		rows不会关闭，连接不会回到连接池中，连接也不会关闭。所以手动关闭非常重要。rows.Close()可以多次调用，是无害操作。
 */
func dbdemo4()  {
	db := util.GetDbInstance()
	selectSql := "SELECT * FROM `jh_articles` WHERE id <?"
	rows,err := db.Query(selectSql,10)
	defer rows.Close()
	defer db.Close()
	if err != nil{
		log.Fatal(err)
	}
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.Id,&article.Title,&article.Author); err != nil{
			log.Fatal(err)
		}
		fmt.Println(article)
	}
}

type Article struct {
	Id int
	Title string
	Author string
}

func main() {
	/**
		可以参考：https://www.jianshu.com/p/bc8120bec94e   https://segmentfault.com/a/1190000003036452
		在事务处理的时候，不能使用db的查询方法，虽然后者可以获取数据，可是这不属于同一个事务处理，将不会接受commit和rollback的改变
		需要注意，Begin和Commit方法，与sql语句中的BEGIN或COMMIT语句没有关系。
		tx事务环境中，只有一个数据库连接，事务内的Eexc都是依次执行的，事务中也可以使用db进行查询，但是db查询的过程会新建连接，这个连接的操作不属于该事务。

		db.Begin()开始事务，Commit() 或 Rollback()关闭事务。Tx从连接池中取出一个连接，在关闭之前都是使用这个连接。Tx不能和DB层的BEGIN, COMMIT混合使用。

		Tx和statement不能分离，在DB中创建的statement也不能在Tx中使用，因为他们必定不是使用同一个连接使用Tx必须十分小心，例如下面的代码：

		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		defer tx.Rollback()
		stmt, err := tx.Prepare("INSERT INTO foo VALUES (?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close() // danger!
		for i := 0; i < 10; i++ {
			_, err = stmt.Exec(i)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
		// stmt.Close() runs here!
	 */

	db := util.GetDbInstance()
	//事务处理。注意，要测试这个事务功能，针对MySQL数据库，需要把表的引擎设置为InnoDB才行。
	tx,err := db.Begin()
	if err != nil{
		log.Fatal(err)
	}

	insertSql := "INSERT INTO `jh_articles` (`title`, `author`) VALUES ( ?, ?)"  // 这个用预处理方式
	updateSql := "UPDATE `jh_articles` SET title = ?,authord = ? WHERE id = ?"    // 这个没有用预处理方式，并且这里还故意写错一个字段(把author字段写成了authord，就是为了验证事务功能)

	stmt,err := db.Prepare(insertSql)
	defer stmt.Close()
	defer db.Close()
	defer tx.Rollback()
	if err != nil{
		 log.Fatal(err)
	}
	if _,err2 := db.Exec(updateSql,"title2","autor2",10); err2 != nil{
		err = tx.Rollback()
		if err != nil {
			log.Println("tx.Rollback1() Error:" + err.Error())
			return
		}
	}
	if _,err1 := stmt.Exec("title8","zhangdaxian"); err1 != nil{
		err = tx.Rollback()
		if err != nil {
			log.Println("tx.Rollback2() Error:" + err.Error())
			return
		}
	}
	if err := tx.Commit(); err != nil{
		err = tx.Rollback()
		if err != nil {
			log.Println("tx.Rollback3() Error:" + err.Error())
			return
		}
	}
	/*if err1 != nil || err2 !=nil{
		fmt.Println("操作有异常，不提交")
		if err := tx.Rollback(); err != nil{
			log.Fatal(err)
		}
	}else{

	}*/
	fmt.Println("success")
}
