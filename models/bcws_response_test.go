package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestDecodeResponse(t *testing.T) {
	a := assert.New(t)

	body := []byte(`<?BrightCloud version=bcap/1.1?><bcap><response><status>200</status><statusmsg>OK</statusmsg><uri>www.whatismyclassification.com</uri><categories><cat><catid>0</catid></cat></categories><bcri>40</bcri><a1cat>0</a1cat></response></bcap>
`)
	r, err := DecodeURIResponse(body)
	a.NoError(err)
	a.Equal(200, r.StatusCode)
	a.Equal(40, r.Bcri)
	a.Equal(0, r.A1cat)
	a.Equal(0, r.Categories[0].CatID)
	fmt.Println(r)
}

func TestDecodeResponseError(t *testing.T) {
	a := assert.New(t)

	body := []byte(`<?BrightCloud version=bcap/1.1?><bcap><respnse><status>200</status><statusmsg>OK</statusmsg><uri>www.whatismyclassification.com</uri><categories><cat><catid>0</catid></cat></categories><bcri>40</bcri><a1cat>0</a1cat></response></bcap>
`)
	r, err := DecodeURIResponse(body)
	a.Error(err)
	a.Nil(r)
}

func TestDecodeResponseComplex(t *testing.T) {
	a := assert.New(t)
	xmlFile, err := os.Open("../resources/serialsWs.xml")
	defer xmlFile.Close()

	a.NoError(err)
	// read our opened xmlFile as a byte array.
	body, err := ioutil.ReadAll(xmlFile)
	a.NoError(err)

	r, err := DecodeURIResponse(body)
	a.NoError(err)
	a.Equal(200, r.StatusCode)
	a.Equal(92, r.Bcri)
	a.Equal(0, r.A1cat)
	a.Equal(33, r.Categories[0].CatID)
	a.Equal(85, r.Categories[0].Confidence)
	a.Equal(44, r.Categories[1].CatID)
}
