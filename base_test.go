package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	const (
		ErrUserNotFound   = 100610
		ErrUserIsNotValid = 100720
	)

	tests:= []struct{
		code int
		http int
		ext string
	}{
		{100610,500,"user is not found"},
		{100611,500,"user is not validate"},
		{ErrDatabase,500,"database error"},
	}

	for _,v := range tests {
		assert.Equal(t,v.code,1)
	}

}
