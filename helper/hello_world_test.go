package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestHelloWorld(t *testing.T){
	result := HelloWorld("Indra")
	if result != "Hello Indra" {
		t.FailNow()
	}
}

func TestHelloWorldMaulana(t *testing.T){
	result := HelloWorld("Maulana")
	assert.Equal(t, "Hello Maulana",result)
}
func TestGoodbye(t *testing.T){
	result := SayGoodbye()
	if result != "Good Bye" {
		t.Fail()
	}
}