package main

import (
	"fmt"
	"testing"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


func TestInsuranceChainCode_InitheckInit(t *testing.T) {//stub *shim.MockStub, args [][]byte
	scc := new(InsuranceChainCode)
	stub := shim.NewMockStub("test1", scc)
	f := "init"
	jsonStr := `{"userID":"3702821982","policies":[{"id":"new20170012","expenseRate":"0.8", "amount":0}]}`
	args := util.ToChaincodeArgs(f, jsonStr)
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}

	q := "query"
	conId := "3702821982"

	args1 := util.ToChaincodeArgs(q, conId)
	queryRes := stub.MockInvoke("2", args1)
	if queryRes.Status != shim.OK {
		t.FailNow()
	}
	fmt.Printf("Query payload is %s", string(queryRes.Payload))

	//stub.MockInvoke("1", [][]byte{[]byte("query"), []byte("123456")})
	//fmt.Printf("test = %s\n", "Init test is Ok!");
}

//func TestInsuranceChainCode_Invoke(t *testing.T) {
//	scc := new(InsuranceChainCode)
//	stub := shim.NewMockStub("test2", scc)
//	f := "query"
//	jsonStr := "3702821982"
//
//	args := util.ToChaincodeArgs(f, jsonStr)
//	queryRes := stub.MockInvoke("2", args)
//	if queryRes.Status != shim.OK {
//		t.FailNow()
//	}
//	fmt.Printf("Query payload is %s", string(queryRes.Payload))
//}


