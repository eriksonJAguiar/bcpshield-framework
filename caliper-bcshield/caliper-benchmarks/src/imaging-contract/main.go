/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import "github.com/hyperledger/fabric/core/chaincode/shim"

func main() {
	err := shim.Start(new(HealthcareChaincode))
	if err != nil {
		panic(err)
	}
}
