package security

import(
	"log"
	"testing"
	"github.com/sohlich/goblog/repository"
)


//Test to generate HMAC token
func TestUserToken(*testing.T){
	user := repository.User{
		Username: "radek",
		Password: "password",
	}
	
	token,err := createUserToken(&user)
	log.Println(err)
	log.Println(token)
	
	
	parsedUser, err := parseUserToken(token)
	log.Println(err)
	log.Println(parsedUser)
	
}