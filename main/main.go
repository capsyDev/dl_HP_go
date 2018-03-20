package main

func main() {

	var  url string = "test"
	var path = "/home/capsy/telechargement/test/"
	//var  url2 string = "test"

	resp := request(url)
	defer resp.Body.Close()
	htmlCode := getHtmlCode(resp)

	name := getName(htmlCode)
	path = path + name

	createDir(path)
	nbPages := getNbPage(htmlCode)
	urlPage := getUrlPages(htmlCode)
	urls := generatedUrl(urlPage,nbPages)

	writeJpg(urls,path,nbPages)
}


