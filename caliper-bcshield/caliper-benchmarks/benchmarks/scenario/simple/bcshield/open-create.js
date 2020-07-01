/*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

'use strict';

const logger = require('@hyperledger/caliper-core').CaliperUtils.getLogger('my-module');
const { v1: uuidv4 } = require('uuid')
// save the objects during init
let bc, contx;

/**
* Initializes the workload module before the start of the round.
* @param {BlockchainInterface} blockchain The SUT adapter instance.
* @param {object} context The SUT-specific context for the round.
* @param {object} args The user-provided arguments for the workload module.
*/
module.exports.init = async (blockchain, context, args) => {
    bc = blockchain;
    contx = context;
    logger.debug('Initialized workload module');
};

module.exports.run = async () => {
    let txArgs = {
        chaincodeFunction: 'addAsset',
        chaincodeArguments: [uuidv4(), "2808886", "Jeff", "Slavech", "4221513", "310 South Crouse Avenue Syracuse NY", "54", "01-01-1966", "OHIP", "AAAAA", "None", "Male", "ASASSAS", "1780694000", "67.0", "1.77", "C3L-01285"],
    };
    
    return bc.invokeSmartContract(contx, 'dicom-caliper', '1', txArgs, 30);
};

module.exports.end = async () => {
    // Noop
    logger.debug('Disposed of workload module');
};