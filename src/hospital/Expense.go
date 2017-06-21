package main

import (
"fmt"
_ "strconv"

"github.com/hyperledger/fabric/core/chaincode/shim"
pb "github.com/hyperledger/fabric/protos/peer"
)

type HospitalChainCode struct {
}

func (t *HospitalChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return nil
}

func (t *HospitalChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
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
