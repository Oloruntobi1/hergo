package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main() {

	g := &getter{}
	fmt.Println(ReturnMyGuy(g))
}


type MakeRequest interface {
	GetMyGuy(string) (string, error)
}

type getter struct {

}

func (g *getter) GetMyGuy(url string) (string, error) {
	

	res, err := http.Get(url)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil

}

func ReturnMyGuy(m MakeRequest) (string, error ){

	var url ="https://findert.herokuapp.com"
	
	return m.GetMyGuy(url)

	

}







// type MakeRequest struct {
// 	mock.Mock
// }

// func (m *MakeRequest) GetMyGuy(url string) (string, error) {

// 	args := m.Called(url)

// 	return args.String(0), args.Error(1)
// }