// 예제 2-6. util.go 파일에 선언한 세션 유틸리티 함수

func index(w http.ResponseWriter, r * http.Request) {
	threads, err := data.Threads(); if err == nil {
		_, err := session(w, r)
		public_tmpl_files := [] string {"templates/layout.html",
			                            "templates/public.navbar.html",
			                            "templates/index.html"}
		private_tmpl_files := [] string {"templates/layout.html",
			                             "templates/private.navbar.html",
			                             "templates/index.html"}
		var templates * template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(private_tmpl_files...))
		} else {
			templates = template.Must(template.ParseFiles(public_tmpl_files...))
		}
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
