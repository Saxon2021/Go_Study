package main

import (
	assets "Go_Study/Asset_Admin/interface/asset"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/interface/assets/add_assets_name", assets.Add_assets_name)
	// http.HandleFunc("/interface/assets/add_design_name", Assets.Add_design_name)

	err := http.ListenAndServe(":1248", nil)

	if err != nil {
		fmt.Println("http server failed!, err: ", err)
	}

}
