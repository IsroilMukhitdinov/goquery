package main

import (
	"fmt"
	"net/http"

	"github.com/IsroilMukhitdinov/goquery/pkg"
)

func main() {
	fmt.Println("Go Query Library Usage Demo")
	// title, err := pkg.GetTitle()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(title)

	// result, err := pkg.ReadLocalFile("/home/isroil/Desktop/goworkshop/code/goquery/static/index.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

	// result, err := pkg.ReadHTMLString()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, word := range result {
	// 	fmt.Println(word)
	// }

	// result, err := pkg.Filter()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, word := range result {
	// 	fmt.Println(word)
	// }

	// result, err := pkg.ReadAnyData("https://go-app.dev/", "h1")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result.Text())

	// result, err := pkg.FilterMore()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, link := range result {
	// 	fmt.Println(link)
	// }

	// result, err := pkg.UnionChildren()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, paragraph := range result {
	// 	fmt.Printf("-%s\n", paragraph)
	// }

	// result, err := pkg.GetRakuQuestions()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for i, title := range result {
	// 	fmt.Printf("[%d] %s\n", i+1, title)
	// }

	// log.Fatal(pkg.EarthquakeTable())

	http.HandleFunc("/quakes", pkg.EarthquakeTable)
	http.ListenAndServe(":9090", nil)

}
