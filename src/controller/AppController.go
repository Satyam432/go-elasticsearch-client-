package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"rest-api-golang/src/models"
	"time"

	u "rest-api-golang/src/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	buf = new(bytes.Buffer)
	w   = multipart.NewWriter(buf)
)

var bububu map[string]interface{}
var dududu map[string]interface{}

var Login = func(w http.ResponseWriter, r *http.Request) {

	//Generate UUID
	id := uuid.New()
	start := time.Now()
	//Parsing Request body
	reqbody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqbody, &bububu)

	loggingreqobj := make(map[string]interface{})
	loggingreqobj["UID"] = id.String()
	loggingreqobj["time"] = start
	loggingreqobj["method"] = r.Method
	loggingreqobj["host"] = r.Host
	loggingreqobj["contentLength"] = r.ContentLength
	loggingreqobj["body"] = bububu
	loggingreqobj["uRI"] = r.RequestURI
	loggingreqobj["response"] = r.Response

	u.Loggerforrequest(loggingreqobj)
	//-------------------------------------

	//--------------------------------------

	user := &models.User{}

	user.Password = ""
	tk := &models.Token{UserId: user.Email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	elapsed := time.Since(start).Seconds()
	resp := u.Message(true, "Successful")
	resp["response_time"] = elapsed
	resp["token"] = tokenString
	resp["uid"] = id.String()
	resp["time"] = time.Now()
	resp["status_code"] = http.StatusBadGateway
	if resp["status_code"] != 400 {
		u.Erloggerobj(resp)
	}
	u.Respond(w, resp)
	u.Loggerforresponse(resp)

}

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	user1 := &models.User{}
	user2 := &models.User{}
	user1.Email = "cao.trung.thu@mail.com"
	user2.Email = "cao.trung.thu@hot.com"

	var users [2]*models.User
	users[0] = user1
	users[1] = user2
	resp := u.Message(true, "Successful")
	resp["data"] = users
	u.Respond(w, resp)
}
