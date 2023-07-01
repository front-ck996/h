package csy

import "github.com/parnurzeal/gorequest"

func NewRequest() *gorequest.SuperAgent {
	return gorequest.New()
}
