/*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* 'di'stributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

'use strict';

module.exports.info  = 'querying asset';


let bc, contx;
let dicom_array;

module.exports.init = function(blockchain, context, args) {
    const open = require('./open.js');
    bc       = blockchain;
    contx    = context;
    dicom_array = open.dicom_array;

    return Promise.resolve();
};

module.exports.run = function() {
    const acc  = dicom_array[Math.floor(Math.random()*(dicom_array.length))];

    if (bc.getType() === 'fabric') {
        let args = {
            chaincodeFunction: 'getAsset',
            chaincodeArguments: [acc],
        };

        return bc.bcObj.querySmartContract(contx, 'dicom-caliper', '1', args, 10);
    } else {
        // NOTE: the query API is not consistent with the invoke API
        return bc.queryState(contx, 'dicom-caliper', '1', acc);
    }
};

module.exports.end = function() {
    // do nothing
    return Promise.resolve();
};
