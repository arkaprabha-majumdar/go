package main

import ("fmt"
	"net/http")


func index_handler(w http.ResponseWriter , r *http.Request){ //*request refers to read here, not memory
	fmt.Fprintf(w,`<h1>Web Dev with Go</h1>
	<p>Here is my medium link for those who like to read:</p>
	<p>https://medium.com/@pravaa8/golang-a-titular-introduction-dd59eef4334d?sk=490f301d53297e48e8a8100b841a1927</p>`)
}
//we could make a decent website with this method, but using templates is much easier.

func about_handler(w http.ResponseWriter , r *http.Request){ //*request refers to read here, not memory
	fmt.Fprintf(w,"This is the about page")
}

func main(){
	http.HandleFunc("/",index_handler) //homepage wih some function - index handler
	http.HandleFunc("/about/",about_handler)
	http.ListenAndServe(":8000",nil) //nil ~ None (Python)
}


//You'll find the output in your browser on 127.0.0.1:8000
//we can add as many pages as we want, using the same format above... find them at 127.0.0.1:8000/page_name
