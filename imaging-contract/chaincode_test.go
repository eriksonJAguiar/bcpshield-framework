/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"strings"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func Init(t *testing.T) *shim.MockStub {
	cc := new(HealthcareChaincode)
	stub := shim.NewMockStub("dicom-v10@contract", cc)
	res := stub.MockInit("1", [][]byte{[]byte("Init")})
	if res.Status != shim.OK {
		t.Error("Init failed", res.Status, res.Message)
	}

	return stub
}

// func TestInvoke(t *testing.T) {
// 	cc := new(HealthcareChaincode)
// 	stub := shim.NewMockStub("dicom-v9@contract", cc)
// 	res := stub.MockInit("1", [][]byte{[]byte("initFunc")})
// 	if res.Status != shim.OK {
// 		t.Error("Init failed", res.Status, res.Message)
// 	}
// 	res = stub.MockInvoke("1", [][]byte{[]byte("invokeFunc")})
// 	if res.Status != shim.OK {
// 		t.Error("Invoke failed", res.Status, res.Message)
// 	}
// }

func Invoke(test *testing.T, stub *shim.MockStub, function string, args ...string) {

	ccArgs := make([][]byte, 1+len(args))
	ccArgs[0] = []byte(function)
	for i, arg := range args {
		ccArgs[i+1] = []byte(arg)
	}
	result := stub.MockInvoke("1", ccArgs)
	// fmt.Println("Call:    ", function, "(", strings.Join(args, ","), ")")
	// fmt.Println("RetCode: ", result.Status)
	// fmt.Println("RetMsg:  ", result.Message)
	// fmt.Println("Payload: ", string(result.Payload))

	test.Log("Call:    ", function, "(", strings.Join(args, ","), ")")
	test.Log("RetCode: ", result.Status)
	test.Log("RetMsg:  ", result.Message)
	test.Log("Payload: ", string(result.Payload))

	if result.Status != shim.OK {
		test.Error("Invoke failed", result.Status, result.Message)
	}
}

func TestAddImaging(t *testing.T) {
	stub := Init(t)
	Invoke(t, stub, "addImaging", "10005", "Bob", "Truth", "(43) 0000-0000", "São Carlos - SP", "23", "1996-31-08", "USP", "AAAAA", "None", "Male", "ASASSAS", "Plan X", "75.5", "1.80", "ASDFG")

}

func TestGetImaging(t *testing.T) {
	stub := Init(t)
	dicomID := "10005"
	Invoke(t, stub, "getImaging", dicomID)
}

//Three args Patient ID ,Doctor or research Id and for sharing Several assets ID for sharing
func TestSharingImagingWithDoctor(t *testing.T) {
	stub := Init(t)
	patientID := "3001"
	doctorID := "2001"
	Invoke(t, stub, "sharingImagingWithDoctor", patientID, doctorID, "10005")
}

// // One param batch ID send to who request
func TestGetSharedImagingWithDoctor(t *testing.T) {
	cc := new(HealthcareChaincode)
	stub := shim.NewMockStub("dicom-v10@contract", cc)
	batchID := "c1ff0a4174f6db2a57e92cca34e877b40602b78b"
	Invoke(t, stub, "getSharedImagingWithDoctor", batchID)
}

// //One param amount image
// func TestRequestImagingForResearchers(t *testing.T) {
// 	cc := new(HealthcareChaincode)
// 	amount := string(2)
// 	stub := shim.NewMockStub("dicom-v9@contract", cc)
// 	Invoke(t, stub, "requestImagingForResearchers", amount)
// }

// //Two atributes Research ID and Batch Shared Dicom ID
// func TestSharingImagingForResearchers(t *testing.T) {
// 	cc := new(HealthcareChaincode)
// 	stub := shim.NewMockStub("dicom-v9@contract", cc)
// 	researchID := "10001"
// 	batchID := "3001"
// 	Invoke(t, stub, "sharingImagingForResearchers", researchID, batchID)
// }

// // One param Log ID
// func TestAuditLogs(t *testing.T) {
// 	cc := new(HealthcareChaincode)
// 	stub := shim.NewMockStub("dicom-v9@contract", cc)
// 	logID := "10001"
// 	Invoke(t, stub, "auditLogs", logID)
// }
