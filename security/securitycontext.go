package security

import(
	"github.com/gorilla/context"
	"net/http"
	"github.com/sohlich/goblog/repository"
)

//Sets the user to context for authorization purpouses.
func SetSecurityContext(request *http.Request,user *repository.User){
	context.Set(request,Principal,user)
}

func GetSecurityContext(request *http.Request) *repository.User{
	return context.Get(request,Principal).(*repository.User)
}