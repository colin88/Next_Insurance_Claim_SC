package main

import (
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
)

func TestHospitalChainCode_Invoke(t *testing.T) {

	scc := new(HospitalChainCode)
	stub := shim.NewMockStub("test1", scc)

	jsonVal := `{"uid":"3702821982","expenseTime":"20001010010203","claimed":false,"medicines":[{"name":"med1000","id":"1000","number":10,"price":10},{"name":"med2000","id":"2000","number":10,"price":20},{"name":"med3000","id":"3000","number":10,"price":30}]}`
	res := stub.MockInvoke("1", [][]byte{[]byte("invoke"), []byte(jsonVal)})
	if res.Status != shim.OK {
		t.FailNow()
	}

	jsonVal2 := `{"uid":"3702821982","expenseTime":"20001010010204","claimed":false,"medicines":[{"name":"med1000","id":"1000","number":10,"price":10},{"name":"med2000","id":"2000","number":10,"price":20},{"name":"med3000","id":"3000","number":10,"price":30}]}`
	res2 := stub.MockInvoke("1", [][]byte{[]byte("invoke"), []byte(jsonVal2)})
	if res2.Status != shim.OK {
		t.FailNow()
	}


	queryRes := stub.MockInvoke("2", [][]byte{[]byte("query"), []byte("3702821982")})
	if queryRes.Status != shim.OK {
		t.FailNow()
	}

	fmt.Printf("Query payload is %s", string(queryRes.Payload))


}