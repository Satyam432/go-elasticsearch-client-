# Building simple restful apis with Go language 
This is an essential example to build server with restful api using Golang.

Step to run
1. Clone the [repo](https://github.com/diegothucao/rest-api-golang)
2. `go run main.go`

Define route 

```go 
var SetupServer = func(appPort string) {
	var router = mux.NewRouter()

	router.HandleFunc("/api/login", controller.Login).Methods("POST")
	router.HandleFunc("/api/getUser", controller.GetUser).Methods("GET")

	router.Use(JwtAuthentication) 

	err := http.ListenAndServe(":"+appPort, router)
	if err != nil {
		fmt.Print(err)
	}
}
```

Process a request 

```go 
var Login = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) 
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	user.Password = ""
	tk := &models.Token{UserId: user.Email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	resp := u.Message(true, "Successful")
	resp["data"] = user
	resp["token"] = tokenString
	u.Respond(w, resp)
}
```


Then you can request 
```
Post http://localhost:3000/api/login

{
"email": "drife@gmail.com",
"password": "drife"
}

Get http://localhost:3000/api/getUser

Token:
"Authorization": `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiJjYW8udHJ1bmcudGh1QGdtYWlsLmNvbSJ9.7N7vWh73ELZmqG0AxRtuzGVlB8JaAVSncmCQowP6cWQ`

```


---
Logging  Functions are in -
src\utils\logme.go
---

Funtions used-

```
-Loggerforrequest
	Services can import the function by package
	Pass the Object to function
	Pass env -
	Use function to log HTTP request
```
```
-Loggerforresponse
	Services can import the function by package
	Pass the Object to function
	Pass env -
	Use function to log HTTP response
```
```
-GetESClient() (*elastic.Client, error)
	Services can import the function by package
	Pass the Object to function
	Pass env -
	Inbuilt function for initializing index for ES
```
```
-Errorlogger(errobj error)
	Services can import the function by package
	Pass the Object to function
	Pass env -
	Use function for logging errors while handling the code 
```
```
-Erloggerobj(erbjecto map[string]interface{})
Services can import the function by package
	Pass the Object to function
	Pass env -
	Inbuilt function to log objects within the Loggerforresponse for unusual responses.
```


