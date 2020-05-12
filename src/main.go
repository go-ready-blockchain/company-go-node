package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-ready-blockchain/blockchain-go-core/Init"
	"github.com/go-ready-blockchain/blockchain-go-core/company"
)

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("company -name NAME \tAddCompany")
	fmt.Println("request -company COMPANY -student USN \tCompany requests for Student's Data")
	fmt.Println("companyRetrieveData -student USN \tCompany retrieves Student's data")

}

func addCompany(company string) {
	Init.InitCompanyNode(company)
	fmt.Println("Company Added!")

}
func request(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Sending Request to Placement Dept!")
	resp, err := http.Post("http://localhost:8084/send",
		"application/json", r.Body)
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()

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
	type jsonBody struct {
		Company string `json:"company"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}
	addCompany(b.Company)

	message := "Company Added!"
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(message))
}

func callrequestBlock(w http.ResponseWriter, r *http.Request) {
	type jsonBody struct {
		Name    string `json:"name"`
		Company string `json:"company"`
	}
	decoder := json.NewDecoder(r.Body)
	var b jsonBody
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}
	//TODO: SendEmail(name,company)
	message := "\n\nSent Email to Student: " + b.Name + " for Requested Data for Company: " + b.Company + "\n\n"
	fmt.Println(message)
	w.Write([]byte(message))
	//callStudentRequestBlock(b.Name, b.Company)

}

func callStudentRequestBlock(name string, company string) {
	reqBody, err := json.Marshal(map[string]string{
		"name":    name,
		"company": company,
	})
	if err != nil {
		print(err)
	}
	resp, err := http.Post("http://localhost:8081/handlerequest",
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		print(err)
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	print(err)
	// }
	// fmt.Println(string(body))
}

func callcompanyRetrieveData(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte(message))
}

func callprintUsage(w http.ResponseWriter, r *http.Request) {

	printUsage()

	w.Header().Set("Content-Type", "application/json")
	message := "Printed Usage!!"
	w.Write([]byte(message))
}

func main() {
	port := "8082"
	http.HandleFunc("/company", calladdCompany)
	http.HandleFunc("/request", request)
	http.HandleFunc("/requestStudent", callrequestBlock)
	http.HandleFunc("/companyRetrieveData", callcompanyRetrieveData)
	http.HandleFunc("/usage", callprintUsage)
	fmt.Printf("Server listening on localhost:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
