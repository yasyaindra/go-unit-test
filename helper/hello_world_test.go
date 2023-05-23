package helper

import "testing"
func TestHelloWorld(t *testing.T){
	result := HelloWorld("Indra")
	if result != "Hello Yasya" {
		t.FailNow()
	}
}

func TestHelloWorldMaulana(t *testing.T){
	result := HelloWorld("Maulana")
	if result != "Hello Maulana" {
		t.Fatal("result it not Maulana")
	}
}
func TestGoodbye(t *testing.T){
	result := SayGoodbye()
	if result != "Alright" {
		t.Fail()
	}
}