package blog

import (
//	"time"
	"log"
	"net/http"
	"github.com/sohlich/goblog/repository"
	"github.com/sohlich/goblog/security"
	"golang.org/x/crypto/bcrypt"
	"strings"
)


//Route handlers
func RegisterFormProcess(w http.ResponseWriter, req *http.Request) {
	user := &repository.User{
		Password: req.FormValue("password"),
		Username: req.FormValue("username"),
		Email: req.FormValue("email"),
	}
	bCryptPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password),6)
	if err != nil {http.Error(w, http.StatusText(405), 405)}
	user.Password = string(bCryptPasswordBytes)	
	user, err = repository.UserRepository().Add(user)
	
	if err != nil {
		http.Error(w, http.StatusText(405), 405)
	}	
}


//Route handlers
func RegisterForm(w http.ResponseWriter, req *http.Request) {
	generatedTemplate := LoadTemplate("register")
	log.Print("Accessing register page")
	generatedTemplate.Execute(w,nil)
}


//Route handlers
func LoginForm(w http.ResponseWriter, req *http.Request) {
	generatedTemplate := LoadTemplate("login")
	generatedTemplate.Execute(w,nil)
}

//Route handlers
func LoginFormProcess(w http.ResponseWriter, req *http.Request) {
	formPassword := req.FormValue("password")
	formUsername := req.FormValue("username")
	databaseUser,err := repository.UserRepository().FindByUsername(formUsername)
	if err != nil {http.Error(w, http.StatusText(405), 405)}
	err = bcrypt.CompareHashAndPassword([]byte (databaseUser.Password),[]byte (formPassword))
	if err != nil {
		log.Println("Wrong password")
		notAuthenticatedRedirect(w,req)
	}else{
		token, err := security.CreateUserToken(databaseUser)
		if err != nil {http.Error(w, http.StatusText(405), 405)}
		w.Header().Set("X-AUTH",token)
		authCookie := &http.Cookie{Name:"X-AUTH",
								  Value:token,
								  Path:"/"}
		http.SetCookie(w,authCookie)
		http.Redirect(w, req,"/", 302)
	}
}

func Logout(w http.ResponseWriter, req *http.Request){
	authCookie := &http.Cookie{Name:"X-AUTH",
								  Value:"",
								  Path:"/"}
	http.SetCookie(w,authCookie)
	http.Redirect(w, req,"/login", 302)
}


//Security interceptor for pre-request and post-request operations
func HttpSecurityInterceptor(router http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
		path := req.URL.Path
		
        log.Print("Pre-request log")
		log.Print(path)
		log.Println(req.Header.Get("X-AUTH"))
		
		user, err := getAuthenticatedUser(req)
			
//		needAuthentication := path!="/login" && path !="/register" && 
//		path !="/" && !strings.Contains(path,"/css") && 
//		!strings.Contains(path,"/js")

		needAuthentication := strings.Contains(path,"/new") || 
				strings.Contains(path,"/admin")
			
		if err != nil && needAuthentication {
			http.Redirect(w,req,"/login",302)
			return
		}
		security.SetSecurityContext(req,user)
        router.ServeHTTP(w, req)
    })
}

//Check if request is authenticated
func getAuthenticatedUser(request *http.Request) (* repository.User,error){
	authCookie, err := request.Cookie("X-AUTH");
	var user *repository.User
	if err == nil{
		user, err = security.ParseUserToken(authCookie.Value)
	}
	return user,err
}





//Default redirect if user not authenticated
func notAuthenticatedRedirect(w http.ResponseWriter, req *http.Request){
	http.Redirect(w, req,"/login", 302)
}




