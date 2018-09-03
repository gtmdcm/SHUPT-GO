package main

import (
	"SHUPT-GO/graphQL"
	_ "SHUPT-GO/models"
	_ "SHUPT-GO/routers"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"io/ioutil"
	"log"
	"net/http"
)

//func main() {
//	beego.Run()
//
//}

func main() {
	b, _ := ioutil.ReadFile("./graphQL/schema/schema.graphql")

	schema := graphql.MustParseSchema(string(b), &graphQL.Resolver{})

	//http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.Write(page)
	//}))

	http.Handle("/graphql", &relay.Handler{Schema: schema})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
