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
	stub := shim.NewMockStub("dicom-bug-@contract", cc)

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

func TestAddAsset(t *testing.T) {
	stub := Init(t)
	Invoke(t, stub, "addAsset", "10005", "11110", "Bob", "Singer", "(43) 9900 0000", "São Paulo SP", "28", "1992-15-08", "USP", "AAAAA", "None", "Male", "ASASSAS", "Plan X", "75.5", "1.89", "ASAEDF")
	Invoke(t, stub, "addAsset", "10006", "11110", "Alice", "Truth", "(43) 8100 0000", "Londrina PR", "23", "1996-05-10", "IBM", "BBBB", "None", "Male", "ASASSAS", "Plan X", "80.1", "1.75", "ASIFA")
	Invoke(t, stub, "addAsset", "10007", "11110", "Bob", "Winshester", "(43) 4200 0000", "São Carlos SP", "41", "1979-31-08", "Microsoft", "EEEEE", "None", "Female", "ASASSAS", "Plan X", "60", "1.60", "QOASXA")
	Invoke(t, stub, "addAsset", "10008", "11110", "Jonh", "Truth", "(43) 9300 0000", "Rib Preto SP", "30", "1990-31-08", "Apple", "DDDDD", "None", "Male", "ASASSAS", "Plan X", "75.5", "1.80", "OASKZA")

}

func TestGetAsset(t *testing.T) {
	stub := Init(t)
	dicomID := "10005"
	Invoke(t, stub, "getAsset", dicomID)
}

//Three args Patient ID ,Doctor or research Id and for sharing Several assets ID for sharing
func TestShareAssetWithDoctor(t *testing.T) {
	stub := Init(t)
	patientID := "3001"
	doctorID := "2001"
	Invoke(t, stub, "shareAssetWithDoctor", patientID, doctorID, "10005")
}

// // One param batch ID send to who request
func TestGetSharedImagingWithDoctor(t *testing.T) {
	cc := new(HealthcareChaincode)
	stub := shim.NewMockStub("dicom-v12@contract", cc)
	batchID := "09ff4216d04ba6e204b0ad27c1fd2e40bcd070fd"
	Invoke(t, stub, "getSharedImagingWithDoctor", batchID)
}

// //One param amount image
func TestRequestImagingForResearchers(t *testing.T) {
	cc := new(HealthcareChaincode)
	amount := string(2)
	stub := shim.NewMockStub("dicom-v9@contract", cc)
	Invoke(t, stub, "requestImagingForResearchers", amount)
}

// //Two atributes Research ID and Batch Shared Dicom ID
func TestSharingImagingForResearchers(t *testing.T) {
	cc := new(HealthcareChaincode)
	stub := shim.NewMockStub("dicom-v9@contract", cc)
	researchID := "10001"
	batchID := "3001"
	Invoke(t, stub, "sharingImagingForResearchers", researchID, batchID)
}

// // One param Log ID
func TestAuditLogs(t *testing.T) {
	cc := new(HealthcareChaincode)
	stub := shim.NewMockStub("dicom-v9@contract", cc)
	logID := "10001"
	Invoke(t, stub, "auditLogs", logID)
}
