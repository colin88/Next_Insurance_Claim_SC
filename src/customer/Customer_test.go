package main

//noinspection GoUnresolvedReference
import (
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	_ "fmt"
	"github.com/hyperledger/fabric/common/util"
	hosp "InsuranceClaim/Next_Insurance_Claim_SC/src/hospital"
	"fmt"
)

func TestCustomerChainCode_Init(t *testing.T) {//stub *shim.MockStub, args [][]byte
	scc := new(CustomerChainCode)
	stub := shim.NewMockStub("test1", scc)

	chaincodeToInvoke := "hosp01"
	cc2 := new(hosp.HospitalChainCode)
	stubExec2 := shim.NewMockStub(chaincodeToInvoke, cc2)

	stub.MockPeerChaincode(chaincodeToInvoke, stubExec2)

	f := "invoke"
	args := util.ToChaincodeArgs(f, chaincodeToInvoke, "", "3702821982")
	res := stub.MockInvoke("01", args)
	if res.Status != shim.OK {
		fmt.Println("------Invoker Hospital info failed", string(res.Message))
		t.FailNow()
	}

	q := "query"
	conId := "3702821982"

	args1 := util.ToChaincodeArgs(q, conId)
	queryRes := stub.MockInvoke("02", args1)
	if queryRes.Status != shim.OK {
		t.FailNow()
	}
	fmt.Printf("Query payload is %s", string(queryRes.Payload))

}