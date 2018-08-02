package routers

import (
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("page not found"))

	//t,_:= template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/404.html")
	//data :=make(map[string]interface{})
	//data["content"] = "page not found"
	//t.Execute(w, data)
}