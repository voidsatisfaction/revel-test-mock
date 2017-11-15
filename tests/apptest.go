package tests

import (
	"fmt"
	"mock-test-revel/app/configs"

	"github.com/jarcoal/httpmock"
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
}

// 애초에 서버에서 seastallion에 api를 보내는 거니까 의미없는거 아닌가....
// 여기는 클라이언트 아닌가...
func tempRequest(t *AppTest) []byte {
	c := configs.Initialize()
	// httpmock.DeactivateAndReset()
	ch := make(chan []byte)
	go func() {
		fmt.Println("first")
		t.Get("/")
		ch <- t.ResponseBody
	}()
	fmt.Println("second")
	httpmock.RegisterResponder("GET", c.SeastallionOwnd,
		httpmock.NewStringResponder(200, `[{"id": 1, "name": hello}]`),
	)
	httpmock.Activate()
	return <-ch
}

func (t *AppTest) Before() {
	// httpmock.RegisterResponder("GET", t.BaseUrl()+"/",
	// 	httpmock.NewStringResponder(200, `[{"id": 1, "name": hello}]`),
	// )
	// httpmock.Activate()
}

func (t *AppTest) TestIndex() {

	fmt.Printf("base: %s", t.BaseUrl())

	resp := tempRequest(t)
	results := string(resp)
	fmt.Printf("%+v", results)
	t.AssertOk()
}

func (t *AppTest) After() {
	// httpmock.DeactivateAndReset()
}
