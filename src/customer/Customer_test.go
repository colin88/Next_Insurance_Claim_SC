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

	jsonVal := `{"uid":"3702821982","expenseTime":"20001010010203","claimed":false,"medicines":[{"name":"med1000","id":"1000","number":10,"price":10},{"name":"med2000","id":"2000","number":10,"price":20},{"name":"med3000","id":"3000","number":10,"price":30}]}
`
	res := stubExec2.MockInvoke("1", [][]byte{[]byte("invoke"), []byte(jsonVal)})
	if res.Status != shim.OK {
		t.FailNow()
	}

	stub.MockPeerChaincode(chaincodeToInvoke, stubExec2)

	f := "invoke"
	args := util.ToChaincodeArgs(f, chaincodeToInvoke, "", "3702821982")
	res = stub.MockInvoke("01", args)
	if res.Status != shim.OK {
		fmt.Println("Invoker Hospital info failed", string(res.Message))
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