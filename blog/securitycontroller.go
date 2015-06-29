package blog

import (
//	"time"
	"log"
	"net/http"
	"github.com/sohlich/goblog/repository"
	"github.com/sohlich/goblog/security"
	"golang.org/x/crypto/bcrypt"
)


//Route handlers
func RegisterFormProcess(w http.ResponseWriter, req *http.Request) {
	
	user := &repository.User{
		Password: req.FormValue("password"),
		Username: req.FormValue("username"),
		Email: req.FormValue("email"),
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password),6)
	
	if err != nil {http.Error(w, http.StatusText(405), 405)}
	
	user.Password = string(passwordBytes)	
	log.Println(user.Password)
	repository.UserRepository().Add(user)
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
		log.Fatal("Wrong password")
		notAuthenticatedRedirect(w,req)
	}else{
		token, err := security.CreateUserToken(databaseUser)
		if err != nil {http.Error(w, http.StatusText(405), 405)}
		w.Header().Set("X-AUTH",token)
		http.Redirect(w, req,"/", 302)
	}
}

//Security interceptor for pre-request and post-request operations
func HttpSecurityInterceptor(router http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		
        log.Print("Pre-request log")
        router.ServeHTTP(w, req)
        
		switch req.Method {
        case "GET":
		
        case "POST":
            // here we might use http.StatusCreated
        }

    })
}

//Default redirect if user not authenticated
func notAuthenticatedRedirect(w http.ResponseWriter, req *http.Request){
	http.Redirect(w, req,"/login", 302)
}




