package string

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestByte2String(t *testing.T) {

	byte1 := []byte{168, 254, 224, 167, 15, 246, 126, 25, 212, 254, 125, 87, 243, 181, 137, 121, 50, 48, 50, 49, 48, 53}
	fmt.Println(string(byte1))

	err := bcrypt.CompareHashAndPassword([]byte("$2a$04$mtf7bYCPnXkWaqY3dTXkA.aa/B20WTQP4t8ELk6iAPjEIS0FeRE8W"), byte1)
	fmt.Println(err)

	fmt.Println([]byte("202105"))
	//str, _ := bcrypt.GenerateFromPassword([]byte("334921"), 4)
	//fmt.Println(string(str))
	//$2a$04$nEVdfOJa7UBbLeC4olCH5enruEyPAqSOMJhmEXJdTCkLU7JW4V40S
	str := "$2a$04$E4f4xu9XNtqsAlnl1NokR.lH9gvo7KECdbDaP8perVAQEJfGCMFq6"
	//str := "$2a$04$yAyqxUzFrZLt5W8hT4jEF.gm/gjAbqfy9XMAUYwfYOgqbuzPKShma"

	fmt.Println(bcrypt.CompareHashAndPassword([]byte(str), []byte("334921")))
}
