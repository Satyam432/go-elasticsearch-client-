package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// "github.com/olivere/elastic"
	elastic "github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

//-----------------------------------------------------------------------------------------

func externalapilogger(apiname string, uid string, h http.Handler) http.Handler {
	fmt.Println("External API logger")
	fn := func(w http.ResponseWriter, r *http.Request) {

		// call the original http.Handler we're wrapping
		h.ServeHTTP(w, r)

		// gather information about request and log it
		uri := r.URL.String()
		method := r.Method
		ip := r.RemoteAddr
		req_payload := r.GetBody

		//Add additional fields

		//Send the data to ELK
		logrus.Info(uri, method, ip, req_payload)
	}

	// http.HandlerFunc wraps a function so that it
	// implements http.Handler interface
	return http.HandlerFunc(fn)
}

//-----------------------------------------------------------------------------------------------------------

func infologger(functionname string, uid string, logobj map[string]interface{}) {
	fmt.Println("General logger")

	// logobj["function_name"] = functionname
	// logobj["uid"] = functionname
	// logrus.Info(&logobj)

	//MAP to JSON
	jsonstr, err := json.Marshal(logobj)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	js := string(jsonstr)

	//----------------------------
	//JSON to ELK

	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	ind, err := esclient.Index().
		Index("infologgerservicename").
		BodyJson(js).
		Do(ctx)
	fmt.Print(ind)
	if err != nil {
		fmt.Println(err)
	}
}

//------------------------------------------------------------------------------------------------------------

func Errorlogger(errobj error) {

	values, err := json.Marshal(errobj.Error())
	if err != nil {
		fmt.Println("error converting to json")
	}

	js := string(values)

	fmt.Print(js)

	dudududud := make(map[string]interface{})
	dudududud["error"] = js
	dudududud["time"] = time.Now()
	jsonstr, err := json.Marshal(dudududud)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	js1 := string(jsonstr)

	//JSON to ELK
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	ind, err := esclient.Index().
		Index("errservicename").
		BodyJson(js1).
		Do(ctx)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(ind)
}

func Erloggerobj(erbjecto map[string]interface{}) {
	//MAP to JSON
	jsonstr, err := json.Marshal(erbjecto)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	js := string(jsonstr)

	//----------------------------
	//JSON to ELK

	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	ind, err := esclient.Index().
		Index("someerrorofapiservicename").
		BodyJson(js).
		Do(ctx)
	fmt.Print(ind)
	if err != nil {
		fmt.Println(err)
	}
}

//------------------------------------------------------------------------------------------------

func dblogger(dbname string, uid string, dbobj map[string]interface{}) {
	fmt.Println("DB related log")
	dbobj["db_name"] = dbname
	dbobj["uid"] = uid
	logrus.Info(&dbobj)
}

// -----------------------------------------------------------------
// Logger for HTTP request

func Loggerforrequest(loggingreqobj map[string]interface{}) {
	//----------------------------
	//MAP to JSON
	jsonstr, err := json.Marshal(loggingreqobj)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	js := string(jsonstr)

	//----------------------------
	//JSON to ELK

	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	ind, err := esclient.Index().
		Index("reqservicename").
		BodyJson(js).
		Do(ctx)
	fmt.Print(ind)
	if err != nil {
		fmt.Println(err)
	}
}

// --------------------------------------------------------------------------------------------------
// Logger for HTTP response
func Loggerforresponse(loggingresobj map[string]interface{}) {
	//MAP to JSON
	jsonstr, err := json.Marshal(loggingresobj)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	js := string(jsonstr)

	//---------------------------
	//JSON to ELK
	ctx := context.Background()
	esclient, err := GetESClient()
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	ind, err := esclient.Index().
		Index("resservicename").
		BodyJson(js).
		Do(ctx)
	fmt.Print(ind)
	if err != nil {
		fmt.Println(err)
	}
}

// ------------------------------------------------------------------
// ESClient setup

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("https://KibanDrife:d0jpG20l1232wwZ@elasticsearch.drife.io"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err
}
