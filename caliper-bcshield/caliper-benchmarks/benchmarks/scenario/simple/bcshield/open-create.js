
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

// 'use strict';

// const { WorkloadModuleBase } = require('@hyperledger/caliper-core');
// module.exports.info = 'create dicom imaging';


// /**
//  * Workload module for the benchmark round.
//  */
// class CreateCarWorkload extends WorkloadModuleBase {
//     /**
//      * Initializes the workload module instance.
//      */
//     constructor() {
//         super();
//         this.txIndex = 0;
//     }

//     /**
//      * Assemble TXs for the round.
//      * @return {Promise<TxStatus[]>}
//      */
//     async submitTransaction() {
//         this.txIndex++;
//         // let carNumber = 'Client' + this.workerIndex + '_CAR' + this.txIndex.toString();
//         // let carColor = colors[Math.floor(Math.random() * colors.length)];
//         // let carMake = makes[Math.floor(Math.random() * makes.length)];
//         // let carModel = models[Math.floor(Math.random() * models.length)];
//         // let carOwner = owners[Math.floor(Math.random() * owners.length)];
//         var dcm_id = 'Patient_' + this.workerIndex + '_Dicom' + this.txIndex.toString();
        
//         let args = {
//             chaincodeFunction: 'addAsset',
//             chaincodeArguments: [dcm_id, "2808886", "Jeff", "Slavech", "4221513", "310 South Crouse Avenue Syracuse NY", "54", "01-01-1966", "OHIP", "AAAAA", "None", "Male", "ASASSAS", "1780694000", "67.0", "1.77", "C3L-01285"],
//         };

//         return this.sutAdapter.invokeSmartContract(this.sutContext, 'dicom-cliper', '1', args, 30);
//     }
// }

// /**
//  * Create a new instance of the workload module.
//  * @return {WorkloadModuleInterface}
//  */
// function createWorkloadModule() {
//     return new CreateCarWorkload();
// }

// module.exports.createWorkloadModule = createWorkloadModule;


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