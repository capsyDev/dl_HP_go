package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strconv"
	"os"
	s "strings"
)

func request (url string)*http.Response{
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error request")
	}

	return resp
}

func getHtmlCode (response *http.Response)string {
	body,err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error getHtmlCode")
	}
	htmlCode := string(body)
	return htmlCode
}


func getNbPage(htmlCode string) int{

	index := s.Index(htmlCode, "'lastImg'>") +10
	var s2 = ""
	for x:=0;x<10 ;x++  {
		if string(htmlCode[index+x]) == "<" {
			break
		}
		s2 = s2 + string(htmlCode[index+x])
	}
	nbPages ,err := strconv.Atoi(s2)
	if err != nil {fmt.Printf("error")}
	return nbPages
}

func getUrlPages(htmlCode string)string{
	index := s.Index(htmlCode,"data-url='1' data-uri='")+23
	var url = ""
	for x:=0;x<300 ;x++  {
		if string(htmlCode[index+x]) == "'" {
			break
		}
		url = url + string(htmlCode[index+x])
	}
	return url
}

func generatedUrl(url string,nbPage int)[]string{
	var urls []string
	for x:=1 ;x<=nbPage ;x++  {
		var ur = url
		switch {
		case x < 10:
			ur = ur + "00"
		case x < 100:
			ur = ur + "0"
		default:
		}
		ur = ur + strconv.Itoa(x) + ".jpg"
		urls = append(urls, ur)
	}
	return urls
}

func getName(htmlCode string) string {
	index := s.Index(htmlCode,"<span class='nameBook'>")+23
	index2 := index
	for x:=0;x<300 ;x++ {
		if string(htmlCode[index+x]) == ">" {
			index2++
			break
		}
		index2++
	}
	var name = ""
	for x:=0;x<300 ;x++ {
		if string(htmlCode[index2+x]) == "<" {
			break
		}
		if string(htmlCode[index2+x]) == " "{
			name = name + "-"
		} else {
			name = name + string(htmlCode[index2+x])
		}

	}
	return name
}

func createDir(path string){
	os.MkdirAll(path, os.ModePerm);
}

func writeJpg(urls []string , path string,nbPage int){

	for x:=0;x<nbPage ;x++  {
		resp := request(urls[x])
		body,err := ioutil.ReadAll(resp.Body)
		if err != nil {fmt.Printf("error jpg")}
		pathJpg := path + "/0" + strconv.Itoa(x) + ".jpg"
		fmt.Println(pathJpg)
		err2 := ioutil.WriteFile(pathJpg , body , 0644)
		if err2 != nil {fmt.Printf("error write")}
	}

}
