package main

import (
"fmt"
_ "strconv"
	_ "encoding/json"
"github.com/hyperledger/fabric/core/chaincode/shim"
pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"encoding/binary"
)

type MedicineDetail struct {
	Id string
	Name string
	Price int
	Number int
}

type ExpenseDetail struct {
	Uid string
	//yyyyMMddHHmmss
	ExpenseTime string
	Claimed bool
	Medicines []MedicineDetail
}


type HospitalChainCode struct {
}

func (t *HospitalChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *HospitalChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		return t.invoke(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	}
	return shim.Error(`invalid invoke function name: "invoke" "query"`)
}

func (t *HospitalChainCode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	jsonVal := args[0]

	var jsonObj ExpenseDetail
	err := json.Unmarshal([]byte(jsonVal), &jsonObj)
	if err != nil {
		return shim.Error("Fail to unmarshal json data!")
	}

	// usrMapdataBytes is []byte
	usrMapdataBytes, err := stub.GetState(jsonObj.Uid)


	// map is not found
	if err != nil {
	//	usrMapdata := map[string][]ExpenseDetail{}
}
	usrMapdata[jsonObj.ExpenseTime] := []byte(jsonObj)



	return nil
}

func (t *HospitalChainCode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return nil
}

func main() {
	err := shim.Start(new(HospitalChainCode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
