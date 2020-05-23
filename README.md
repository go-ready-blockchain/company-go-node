# COMPANY NODE

## Blockchain Implementation in GoLang For Placement System

## The Consensus Algorithm implemented in Blockchain System is a combination of Proof Of Work and Proof Of Elapsed Time


### Run `go run src/main.go` to Start the Server and listen on localhost:8082

### Usage :

#### To Print Usage
####    Make POST request to `/usage`

#### Advance Pipeline - 

#### To Add a new Company    
####    Make POST request to `/company` with body -
```json
{
    "company": "GE"
}
```

#### To Send Request to Placement Dept for Eligible Students based on Eligibility Criteria
####    Make POST request to `/request` with body -
```json
{
	"company" : "JPMC",
	"backlog" : "",
	"starOffer" : "",
	"branch" : ["CSE","ISE"],
	"gender" : "",
	"cgpaCond" : "GreaterThan",
	"cgpa" : "2",
	"perc10thCond" : "GreaterThan",
	"perc10th" : "10",
	"perc12thCond" : "GreaterThan",
	"perc12th" : "10"
}
```
#### Part of the Pipeline - 

#### To Retrieve the data for the Company
####    Make POST request to `/companyRetrieveData` with body -
```json
{
	"name":"1MS16CS034",
    "company": "JPMC"
  
}
```




