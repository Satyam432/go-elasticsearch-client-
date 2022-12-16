package main

import (
	"errors"
	"fmt"
	"rest-api-golang/src/routes"
	z "rest-api-golang/src/utils"

	_ "github.com/lib/pq"
)

const (
	host     = "3.108.179.205"
	port     = 5432
	user     = "satyam"
	password = "evcesehBfiLY!eAr"
	dbname   = "satyam"
)

func main() {

	//Metricbeat approach
	// f, err := os.OpenFile("lagbanao.log", os.O_WRONLY|os.O_CREATE, 0755)
	// if err != nil {
	// 	fmt.Println("file not created")
	// }
	// logrus.SetOutput(f)
	// logrus.SetFormatter(&logrus.TextFormatter{
	// 	FullTimestamp:   true,
	// 	ForceColors:     true,
	// 	TimestampFormat: time.RFC822,
	// })

	//Create custom eror--------------------------------------------------------
	for num := 1; num <= 3; num++ {
		fmt.Printf("validating %d... ", num)
		err := validateValue(num)
		if err != nil {
			fmt.Println("there was an error:", err)
		} else {
			fmt.Println("valid!")
		}
	}
	//---------------------------------------------------------------------------

	//Postegres-Testing
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	// db, err := sql.Open("postgres", psqlconn)
	// CheckError(err)

	// close database
	// defer db.Close()

	// check db
	// err = db.Ping()
	// CheckError(err)

	// fmt.Println("Connected!")
	//-----------------------------------------------------------------------------------------

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	//Setup server
	routes.SetupServer("3000")

}

func validateValue(number int) error {
	if number == 1 {
		z.Errorlogger(errors.New("this is a sample error"))
		return fmt.Errorf("that's odd")
	} else if number == 2 {
		return fmt.Errorf("uh oh")
	}
	return nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
