package main

import (
	"blackfriday"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	//"io/ioutil"
	"net/http"
	//"path/filepath"
	"reflect"
	//"strings"
)

// Post for index html
type Post struct {
	Title   string
	Auth    string
	Summary string
	Body    string
	File    string
}

//Postone for test html
type Postone struct {
	Title string
	Body  string
}

func handlerequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query()["name"])
	articleFilename := r.URL.Query()["name"]
	//fmt.Println(articleFilename,reflect.TypeOf(articleFilename))
	for _, s := range articleFilename {
		fmt.Println(reflect.TypeOf(s))
		if "" != s {
			//posts := getonePosts(s)
			posts := getDataMysql(s)
			//t := template.New("test.html")
			//t, _ = t.ParseFiles("test.html")
			//t.Execute(w, "")
			//t.Execute(w, posts)
			fmt.Fprintf(w, posts[0].Body)
			return
		}
	}
	//posts := getPosts()[]Postone
	posts := getDataMysql("all")
	t := template.New("index.html")
	t, _ = t.ParseFiles("index.html")
	t.Execute(w, posts)
}

//Mytest test for handlerequest
func Mytest(s string) string {
	if "test" != s {
		getDataMysql(s)
		t := template.New("test.html")
		t, _ = t.ParseFiles("test.html")
		return "handle request ok"
	}
	//posts := getPosts()[]Postone
	getDataMysql("all")
	t := template.New("index.html")
	t, _ = t.ParseFiles("index.html")
	return "handle request ok"
}

//getDataMysql mysql search
func getDataMysql(nameSearch string) []Post {
	a := []Post{}
	Odb, err := sql.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/my_blog")
	if err != nil {
		fmt.Println(err)
	}
	defer Odb.Close() //defer语句延迟执行一个函数，该函数被推迟到当包含它的程序返回时
	//var row  *sql.Row
	var uid int
	var filename, title, summary, body, date string

	if "all" == nameSearch {
		row, err := Odb.Query("select uid ,filename,title, summary, body, date from article order by date desc")
		if err != nil {
			fmt.Println(err)
		}
		//err = row.Scan(&uid, &title, &summary, &body, &date, &filename)

		for row.Next() {
			row.Scan(&uid, &filename, &title, &summary, &body, &date)
			fmt.Println(filename, title, summary, body, date)
			body = string(blackfriday.MarkdownCommon([]byte(body)))
			a = append(a, Post{title, date, summary, body, filename})
		}
		//fmt.Println(row)
		return a
	} else if "all" != nameSearch {
		//fmt.Println("search one", nameSearch)
		row, err := Odb.Query("select uid ,filename,title, summary, body, date from article where filename = ?", nameSearch)
		if err != nil {
			fmt.Println(err)
		}
		for row.Next() {
			row.Scan(&uid, &filename, &title, &summary, &body, &date)
			//fmt.Println("vvvvvvvvvvvv")
			//fmt.Println(uid, filename, title, summary, body, date)
			body = string(blackfriday.MarkdownCommon([]byte(body)))
			a = append(a, Post{title, date, summary, body, filename})
		}
		return a
	}
	return a
}

func sdd() int {
	return 3
}

func mytest(s string) string {
	return "handle request ok"
}

/*
func getonePosts(name string) []Postone {
	a := []Postone{}
	a = append(a, Postone{title, body})
	return a
}

func getPosts() []Post {
	a := []Post{}
	files, _ := filepath.Glob("posts/*.md")
	//fmt.Println(files)
	for _, f := range files {
		file := strings.Replace(f, "posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fileread, _ := ioutil.ReadFile(f)
		lines := strings.Split(string(fileread), "\n")
		title := string(lines[0])
		date := string(lines[1])
		summary := string(lines[2])
		body := strings.Join(lines[3:len(lines)], "\n")
		body = string(blackfriday.MarkdownCommon([]byte(body)))
		a = append(a, Post{title, date, summary, body, file})
	}
	return a
}
*/
