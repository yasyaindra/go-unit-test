package helper

import "testing"
func TestHelloWorld(t *testing.T){
	result := HelloWorld("Indra")
	if result != "Hello Indra" {
		panic("result it not indra")
	}
}

func TestHelloWorldMaulana(t *testing.T){
	result := HelloWorld("Maulana")
	if result != "Hello Maulana" {
		panic("result it not Maulana")
	}
}