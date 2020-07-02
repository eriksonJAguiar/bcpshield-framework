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

// module.exports.info  = 'create Assets';

// let dicom_array = [];
// let txnPerBatch;
// let initAsset;
// let bc, contx;
// module.exports.init = async function(blockchain, context, args) {
//     // if(!args.hasOwnProperty('dicom')) {
//     //     return Promise.reject(new Error('create asset - \'dicom\' is missed in the arguments'));
//     // }

//     if(!args.hasOwnProperty('txnPerBatch')) {
//         args.txnPerBatch = 1;
//     }
//     //initMoney = args.money;
//     txnPerBatch = args.txnPerBatch;
//     bc = blockchain;
//     contx = context;

//     return Promise.resolve();
// };


/**
 * Generate string by picking characters from dic variable
 * @param {*} number character to select
 * @returns {String} string generated based on @param number
 */

/**
 * Generate unique dicom key for the transaction
 * @returns {String} dicom key
 */
const dic = 'abcdefghijklmnopqrstuvwxyz';
/**
 * Generate string by picking characters from dic variable
 * @param {*} number character to select
 * @returns {String} string generated based on @param number
 */
function get26Num(number){
    let result = '';
    while(number > 0) {
        result += dic.charAt(number % 26);
        number = parseInt(number/26);
    }
    return result;
}

let prefix;
/**
 * Generate unique account key for the transaction
 * @returns {String} account key
 */
function generateAccount() {
    // should be [a-z]{1,9}
    if(typeof prefix === 'undefined') {
        prefix = get26Num(process.pid);
    }
    return prefix + get26Num(dicom_array.length+1);
}

/**
 * Generates simple workload
 * @returns {Object} array of json objects
 */
function generateWorkload() {
    let workload = [];
    for(let i= 0; i < txnPerBatch; i++) {
        let acc_id = generateAccount();
        dicom_array.push(acc_id);

        workload.push({
            chaincodeFunction: 'addAsset',
            chaincodeArguments: [acc_id,  "2808886", "Jeff", "Slavech", "4221513", "310 South Crouse Avenue Syracuse NY", "54", "01-01-1966", "OHIP", "AAAAA", "None", "Male", "ASASSAS", "1780694000", "67.0", "1.77", "C3L-01285"],
        });
    }
    return workload;
}

module.exports.run = function() {
    let args = generateWorkload();
    return bc.invokeSmartContract(contx, 'dicom-caliper', '1', args, 100);
};

module.exports.end = function() {
    return Promise.resolve();
};

module.exports.dicom_array = dicom_array;

// 'use strict';

// const contractId = 'dicom-caliper';
// const version = '1';

// let bc, ctx, clientArgs, clientIdx;
// let dicom_array = [];

// module.exports.init = async function(blockchain, context, args) {
//     bc = blockchain;
//     ctx = context;
//     clientArgs = args;
//     clientIdx = context.clientIdx.toString();

//     return Promise.resolve();
// };

// const dic = 'abcdefghijklmnopqrstuvwxyz';
// function get26Num(number){
//     let result = '';
//     while(number > 0) {
//         result += dic.charAt(number % 26);
//         number = parseInt(number/26);
//     }
//     return result;
// }

// let prefix;
// /**
//  * Generate unique account key for the transaction
//  * @returns {String} account key
//  */
// function generateAccount() {
//     // should be [a-z]{1,9}
//     if(typeof prefix === 'undefined') {
//         prefix = get26Num(process.pid);
//     }
//     return prefix + get26Num(dicom_array.length+1);
// }

// module.exports.run = function() {
//     let acc_id = generateAccount();
//     dicom_array.push(acc_id);
//     let myArgs = {
//         chaincodeFunction: 'addAsset',
//         chaincodeArguments: [acc_id,  "2808886", "Jeff", "Slavech", "4221513", "310 South Crouse Avenue Syracuse NY", "54", "01-01-1966", "OHIP", "AAAAA", "None", "Male", "ASASSAS", "1780694000", "67.0", "1.77", "C3L-01285"]
//     };

//     return bc.invokeSmartContract(ctx, contractId, version, myArgs, 60);
// };

// module.exports.end = async function() {
//     return Promise.resolve();
// };

// module.exports.dicom_array = dicom_array;