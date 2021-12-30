package main

import (
	"fmt"
	"net/http"
)

func Add_assets_name(w http.ResponseWriter, r *http.Request) {

	// 注册一个关闭连接的事件
	defer r.Body.Close()

	// 解析客户端传来的参数
	r.ParseForm()
	fmt.Println(r.PostForm)
	pmName := r.PostForm.Get("pmName")
	// designName := r.PostForm.Get("designName")
	// finalName := r.PostForm.Get("finalName")
	// sort := r.PostForm.Get("sort")
	// recording := r.PostForm.Get("recording")
	// skeleton := r.PostForm.Get("skeleton")
	// designer := r.PostForm.Get("designer")

	fmt.Println(pmName)

	answer := `{"status": pmName}`
	w.Write([]byte(answer))
}

func main() {

	http.HandleFunc("/interface/assets/add_assets_name", Add_assets_name)

	err := http.ListenAndServe(":1248", nil)

	if err != nil {
		fmt.Println("hhtp server failed!, err: ", err)
	}

}
