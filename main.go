package main

import (
	//"github.com/arl/assertgo"
	"blackfriday"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	//"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"time"
)

//Firsttime 基础时间，每次更新后也会更新
var Firsttime = "2018-07-02 18:13:29"

//GetFileModTime 获取文件修改时间 返回unix时间戳
func GetFileModTime(path string) {
	fileInfo, _ := os.Stat(path)
	modTime := fileInfo.ModTime().Format("2006-01-02 15:04:05") //坑，必须是这个时间，其他的会出错
	fmt.Println(modTime, reflect.TypeOf(modTime).String())      //打印字符类型
	if modTime > Firsttime {
		fmt.Println(modTime, path)
		file := strings.Replace(path, "blog/posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1) //如果n<0会替换所有old子串
		fileread, _ := ioutil.ReadFile(path)
		lines := strings.Split(string(fileread), "\n")
		titleGet := string(lines[0])
		dateGet := modTime
		summaryGet := string(lines[2])
		bodyGet := strings.Join(lines[3:len(lines)], "\n")
		bodyGet = string(blackfriday.MarkdownCommon([]byte(bodyGet)))
		//fmt.Println(file, titleGet, dateGet, summaryGet, bodyGet)

		//sql.Open并不会立即建立一个数据库的网络连接, 也不会对数据库链接参数的合法性做检验
		//它仅仅是初始化一个sql.DB对象. 当真正进行第一次数据库查询操作时, 此时才会真正建立网络连接
		Odb, err := sql.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/my_blog")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer Odb.Close() //defer语句延迟执行一个函数，该函数被推迟到当包含它的程序返回时

		var row *sql.Row
		row = Odb.QueryRow("select * from article where filename = ?", file)
		if err != nil {
			fmt.Println(err)
		}
		var filename, title, summary, body, date string
		var uid int
		var result sql.Result
		err = row.Scan(&uid, &title, &summary, &body, &date, &filename)
		if err != nil {
			fmt.Println(err)
			fmt.Println("结果为空")
			result, err = Odb.Exec("insert into article(title, summary, body, date, filename) values(?,?,?,?,?)", titleGet, summaryGet, bodyGet, dateGet, file)
			if err != nil {
				fmt.Println(err)
				return
			}
			lastID, _ := result.LastInsertId()
			fmt.Println("新插入记录的ID为", lastID)
		}else{
		    result ,err = Odb.Exec("update article set title= ? , summary= ?, body = ? , date = ? where filename = ?",titleGet,summaryGet,bodyGet,dateGet,file)
		    if err != nil {
			    fmt.Println(err)
			    return 
		    }
		    lastID, _ := result.LastInsertId()
		    fmt.Println("新插入记录的ID为", lastID)
		}
		//fmt.Println(uid,title, summary, body, date, filename)
	}
}

//readBlogPost 读取blog列表
func readBlogPost() {
	var tempTime string = time.Now().Format("2006-01-02 15:04:05")
	files, _ := filepath.Glob("blog/posts/*.md")
	fmt.Println(files)
	for _, f := range files {
		GetFileModTime(f)
	}
	Firsttime = tempTime
}

//gitPullBlog git pull
func gitPullBlog(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("/bin/bash", "-c", "cd /home/youmi/git/go_web/blog; git pull")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		fmt.Fprintf(w, "Execute Command failed:"+err.Error())
		return
	}
	readBlogPost()
	fmt.Println("git pull ok!")
	fmt.Fprintf(w, "git pull ok!") //这个写入到w的是输出到网页
}

//GitPullBlogiTest pull
func GitPullBlogiTest() string {
	cmd := exec.Command("/bin/bash", "-c", "cd /home/youmi/git/go_web/blog; git pull")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return "git pull error!"
	}
	readBlogPost()
	fmt.Println("git pull ok!")
	return "git pull ok!"
}

func main() {
	//端口设置
	ports := []string{":25000", ":25001"}
	for _, v := range ports {
		go func(port string) { //每个端口都扔进一个goroutine中去监听
			mux := http.NewServeMux()
			if port == ":25000" {
				mux.HandleFunc("/", handlerequest)
			} else if port == ":25001" {
				mux.HandleFunc("/", gitPullBlog)
			}
			http.ListenAndServe(port, mux)
		}(v)
	}
	select {}
}
