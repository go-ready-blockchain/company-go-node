package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	

	"github.com/go-ready-blockchain/blockchain-go-core/Init"
	"github.com/go-ready-blockchain/blockchain-go-core/company"
	
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("Make POST request to /company \tTo Add Company")
	fmt.Println("Make POST request to /request \tTo Send Request to Placement Dept for Eligible Students based on Eligibility Criteria")
	fmt.Println("Make POST request to /companyRetrieveData \tCompany retrieves Student's data")

}

func addCompany(company string) {
	Init.InitCompanyNode(company)
	fmt.Println("Company Added!")

}
func request(w http.ResponseWriter, r *http.Request) {
	//name := time.Now().String()
	//logger.FileName = "Company Request" + name + ".log"
	//logger.NodeName = "Company Node"
	//logger.CreateFile()

	fmt.Println("Sending Request to Placement Dept!")
	Apiurl := os.Getenv("PLACEMENT_URL")
	Apiurl = Apiurl + "/send"
	fmt.Println(Apiurl)

	resp, err := http.Post(Apiurl,
		"application/json", r.Body)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

	//logger.UploadToS3Bucket(//logger.NodeName)

	//logger.DeleteFile()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Request Sent!"))
}

func companyRetrieveData(name string, companyname string) bool {
	flag := company.RetrieveData(name, companyname)
	if flag == true {
		fmt.Println("Company retrieved the data!")
		return true
	} else {
		fmt.Println("Company failed to retrieve the data!")
		return false
	}
}

func calladdCompany(w http.ResponseWriter, r *http.Request) {
	//name := time.Now().String()
	//logger.FileName = "Add Company" + name + ".log"
	//logger.NodeName = "Company Node"
	//logger.CreateFile()

	type jsonBody struct {
		Company string `json:"company"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}
	addCompany(b.Company)

	//logger.UploadToS3Bucket(//logger.NodeName)

	//logger.DeleteFile()

	message := "Company Added!"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func callcompanyRetrieveData(w http.ResponseWriter, r *http.Request) {
	//name := time.Now().String()
	//logger.FileName = "Company Retrieve Data" + name + ".log"
	//logger.NodeName = "Company Node"
	//logger.CreateFile()

	type jsonBody struct {
		Name    string `json:"name"`
		Company string `json:"company"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}

	message := ""
	flag := companyRetrieveData(b.Name, b.Company)
	if flag == true {
		message = "Company retrieved the data!"
	} else {
		message = "Company failed to retrieve the data!"
	}

	//logger.UploadToS3Bucket(//logger.NodeName)

	//logger.DeleteFile()

	w.Write([]byte(message))
}

func callprintUsage(w http.ResponseWriter, r *http.Request) {

	printUsage()

	w.Header().Set("Content-Type", "application/json")
	message := "Printed Usage!!"
	w.Write([]byte(message))
}

func main() {
	port := "8080"
	http.HandleFunc("/company", calladdCompany)
	http.HandleFunc("/request", request)
	http.HandleFunc("/companyRetrieveData", callcompanyRetrieveData)
	http.HandleFunc("/usage", callprintUsage)
	fmt.Printf("Server listening on localhost:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
