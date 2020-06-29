package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
   
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// User ...
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
   
var db *sql.DB

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Println("Error readin .env: ", err)
		os.Exit(1)
	}
   
	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbNAME := os.Getenv("DB_NAME")
   
	dsn := dbUserName + ":" + dbPassword + "@/" + dbNAME   
	db, err = sql.Open("mysql", dsn)   
	if err != nil {
	    log.Println(err)
	    os.Exit(1)
	}
}   

func main() {
	r := mux.NewRouter()
	r.Use(middleware)
   
	// curl http://127.0.0.1:8080/
	r.HandleFunc("/", rootHandler)
	// curl -X POST http://127.0.0.1:8080/user -d '{"email":"", "password":""}' -H "content-type: application/json"
	// 无效的 json 数据会导致程序 panic, 但 supervisor 会重启程序
	// curl -X POST http://127.0.0.1:8080/user -d '{il":"", "password":""}' -H "content-type: application/json"
	r.HandleFunc("/user", createUserHandler).Methods("POST")
   
	fmt.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		// 退出程序
		log.Println("Failed starting server ", err)
		os.Exit(1)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is root handler")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}   
	err := json.NewDecoder(r.Body).Decode(user)
	// 对请求响应 JSON 数据
	// 在实际应用中你可能想要创建一个进行错误处理的函数
	if err != nil {
		// 我们也可以这么做
		// errREsp := `"error": "Invalid input", "status": 400`
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(400)
		// w.Write([]byte(errREsp))
   
		// 然而我们会让服务器崩溃
	 	log.Fatal(err)
	 	return
	}
   
	// 在实际应用中必须对密码进行哈希, 可以使用 bcrypt 算法
	_, err = db.Exec("INSERT INTO users(email, password) VALUES(?,?)", user.Email, user.Password)
   
	if err != nil {
		log.Println(err)
		// 简单起见, 发送明文字符串
		// 创建一个有效的 JSON 响应
	 	errREsp := `"error": "Internal error", "status": 500` // 返回 500 状态码，因为这是我们而非用户的问题
	 	w.Header().Set("Content-Type", "application/json")
	 	w.WriteHeader(500)
	 	w.Write([]byte(errREsp))   
		return
	}
}

// 一个简单的中间件, 只用来记录请求的 URI
var middleware = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {   
		requestPath := r.URL.Path
		log.Println(requestPath)
		next.ServeHTTP(w, r) // 在中间件调用链中进行处理
	})
}
