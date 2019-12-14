
package main

import (
    "fmt"
   "io"
    "log"
	"net/http"
	"encoding/json"
    "time"
   "database/sql"
 _ "github.com/go-sql-driver/mysql"
	// "strings"
)

func add(w http.ResponseWriter, r *http.Request){
	
	db, err := sql.Open("mysql", "root:xfz123456@tcp(127.0.0.1:3306)/work?charset=utf8")
	if err != nil{
		return  
	}
	defer db.Close()
	r.ParseForm()
	title  := r.Form.Get("title")        //标题
	messages := r.Form.Get("messages")   //信息文本
	date := r.Form.Get("time")			//时间
	id := time.Now().Unix()             //id
	fmt.Println(id,title,messages,date)
	result, err := db.Exec("insert into work (id,title,messages,date) values (?,?,?,?)",id,title,messages,date)
	if err != nil{
		fmt.Println(err)
		return 
	}
	fmt.Println(result)
	fmt.Println("success！")
	io.WriteString(w, "发布成功！")

}
func cha(w http.ResponseWriter, r *http.Request){
 	db, err := sql.Open("mysql", "root:xfz123456@tcp(127.0.0.1:3306)/work?charset=utf8")
	if err != nil{
        return 
	}
	defer db.Close()
	rows,err:=db.Query("SELECT * FROM work")
	data :=map[int]interface{}{}
	i :=0
	for rows.Next(){
		var id,title,messages,date string
		rows.Scan(&id,&title,&messages,&date)
		fmt.Println(title)
		item := map[string]string{"title":title,"messages":messages,"date":date}
		data[i] = item
		i++
	}	
	data_json,_ := json.Marshal(data)
	io.WriteString(w,string(data_json))

}
func shan(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "root:xfz123456@tcp(127.0.0.1:3306)/work?charset=utf8")
	if err != nil{
        return 
	}
	defer db.Close()
	query := r.URL.Query()
	uid := query["shan"][0]
	rows,err := db.Exec("delete from work where id=?",uid)
	fmt.Println(rows)
	if err != nil{
		return
	}
	io.WriteString(w,"删除成功 ！！！")
}
func gai(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "root:xfz123456@tcp(127.0.0.1:3306)/work?charset=utf8")
	if err != nil{
		return  
	}
	defer db.Close()
	r.ParseForm()
	title  := r.Form.Get("title")        //标题
	messages := r.Form.Get("messages")   //信息文本
	date := r.Form.Get("time")			//时间
	query := r.URL.Query()
	id := query["gai"][0]  				//id
	fmt.Println(id,title,messages,date)
	result, err := db.Exec("update work set title=?,messages=?,date=? where id=?",title,messages,date,id)
	if err != nil{
		fmt.Println(err)
		return 
	}
	fmt.Println(result)
	fmt.Println("success！")
	io.WriteString(w, "修改成功！")
}

func main() {
	http.HandleFunc("/add", add)              //添加路由
	http.HandleFunc("/cha", cha)              //查询路由
	http.HandleFunc("/shan", shan)            //删除路由
	http.HandleFunc("/gai", gai)              //修改路由

        err := http.ListenAndServe(":8000", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

