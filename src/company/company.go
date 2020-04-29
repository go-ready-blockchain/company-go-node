package company

import (
	"fmt"

	"github.com/jugalw13/company-go-node/blockchain"
	"github.com/jugalw13/company-go-node/security"
	"github.com/jugalw13/company-go-node/student"
	"github.com/jugalw13/company-go-node/utils"
)

//Supports only 1 company for now
//Retrieve Data from Buffer itself

func RetrieveData(name string, company string) bool {

	block := blockchain.GetBlockFromBuffer(name, company)

	studentdata, dflag := security.DecryptMessage(block.StudentData, security.GetUserFromDB(company).PrivateKey)
	if dflag == false {
		fmt.Println("Decrytion of Message Failed")
		utils.DeleteBlockFromBuffer(name, company)
		return false
	}

	sflag := security.VerifyPSSSignature(security.GetPublicKeyFromDB("PlacementDept"), block.Signature, studentdata)
	if sflag == false {
		fmt.Println("Signature Verification Failed, Authentication Failed")
		utils.DeleteBlockFromBuffer(name, company)
		return false
	}

	v := blockchain.DecodeToStruct(block.Verification)
	vflag := blockchain.CheckIfVerifiedByAll(v)
	if vflag == false {
		fmt.Println("Verification Not Yet Done. Company not allowed to retrieve Data")
		return false
	}
	fmt.Println("Student Data:\n")
	studentstruct := student.DecodeToStruct(studentdata)
	student.PrintStudentData(studentstruct)
	fmt.Println("\nRetrieved Data Successfully")

	return true
}
