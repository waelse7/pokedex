package main

type currentPage struct {
	Next *string
	Prev *string
}

var page currentPage

func SetNextPage(url *string){
	page.Next = url
}

func GetNextPage() *string {
	return page.Next
}

func SetPrevPage(url *string){
	page.Prev = url
}

func GetPrevPage() *string {
	return page.Prev
}
