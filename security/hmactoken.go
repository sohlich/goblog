package security

import(
	"errors"
	"bytes"
	"strings"
	"crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
	"encoding/json"
	"github.com/sohlich/goblog/repository"
)

const Principal = "Principal"

var appSecret = "secret"
var separator = "."

func ComputeHmac256(message string, secret string) []byte {
    key := []byte(secret)
    h := hmac.New(sha256.New, key)
    h.Write([]byte(message))
    return h.Sum(nil) 
}

func EncodeBase64(input []byte) string{
	return base64.StdEncoding.EncodeToString(input)
}

func DecodeBase64(input string) ([]byte,error){
	return base64.StdEncoding.DecodeString(input)
}


func CreateUserToken(user *repository.User) (string ,error) {
	userBytes,err := json.Marshal(user)
	if err != nil {return "", err}
	content := EncodeBase64(userBytes)
	hash := EncodeBase64(ComputeHmac256(content,appSecret))
	token := content + separator + hash
	return token , nil
}


func ParseUserToken(token string)(*repository.User,error){
	
	//Split token
	tokenPart := strings.Split(token,separator)
	
	if len(tokenPart) < 2 {return nil, errors.New("Malformed auth token")}
	
	//Decode token parts
	userBytes,err:= DecodeBase64(tokenPart[0]);
	if err != nil {return nil, err}
	hash,err := DecodeBase64(tokenPart[1]);
	if err != nil {return nil, err}

	//Compare both parts
	hashFromUserBytes := ComputeHmac256(EncodeBase64(userBytes),appSecret)
	
	validHash := bytes.Equal(hashFromUserBytes,hash)
	if validHash {
		user := &repository.User{}
		json.Unmarshal(userBytes,&user)
		return user, nil
	}
	
	return nil, err
	
}