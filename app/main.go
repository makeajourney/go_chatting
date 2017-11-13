package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
)

var renderer *render.Render

func init() {
	// create renderer
	renderer = render.New()
}

func main() {
	// create router
	router := httprouter.New()

	// define handler
	router.GET("/", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		// 렌더러를 사용하여 템플릿 렌더링
		renderer.HTML(w, http.StatusOK, "index", map[string]string{"title": "Simple Chat!"})
	})

	// negroni 미들웨어 생성
	n := negroni.Classic()

	// negroni에 router를 핸들러로 등록
	n.UseHandler(router)

	// execute web server
	n.Run(":3000")
}
