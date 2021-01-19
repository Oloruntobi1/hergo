package main

import (
	// "errors"
	"gositelearn/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetMyGuyAgain(t *testing.T) {


	
	f := &mocks.MakeRequest{}
	f.On("GetMyGuy", "https://findert.herokuapp.com").Return("my man", nil).Once()
	// f.On("GetMyGuy").Return(nil, errors.New("Oops")).Once()

	res, err := ReturnMyGuy(f)

	assert.Nil(t, err)
	assert.Equal(t, "my man", res)
}
