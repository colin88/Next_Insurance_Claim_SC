package main

import (
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
)

func TestHospitalChainCode_Invoke(t *testing.T) {

	scc := new(HospitalChainCode)
	stub := shim.NewMockStub("test1", scc)

	jsonVal := `{"uid":"3702821982","expenseTime":"20001010010203","claimed":false,"medicines":[{"name":"med1000","id":"1000","number":10,"price":10},{"name":"med2000","id":"2000","number":10,"price":20},{"name":"med3000","id":"3000","number":10,"price":30}]}
`
	res := stub.MockInvoke("1", [][]byte{[]byte("invoke"), []byte(jsonVal)})
	if res.Status != shim.OK {
		fmt.Printf("Failed!!")
	}



}