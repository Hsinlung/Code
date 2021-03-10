package main


import (
	"fmt"
	"github.com/satori/go.uuid"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())

	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	//u2, err := uuid.NewV4()
	//if err != nil {
	//	fmt.Printf("出现错误：%s", err)
	//}
	//fmt.Printf("UUIDv4: %s\n", u2)

	u2, err := uuid.FromString("6ba7b810")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)

}

