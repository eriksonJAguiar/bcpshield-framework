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

module.exports.info  = 'querying dicom values';


let bc, contx;
let dicom_array;

module.exports.init = function(blockchain, context, args) {
    const createAsset = require('./create-asset.js');
    bc       = blockchain;
    contx    = context;
    dicom_array = createAsset.dicom_array;

    return Promise.resolve();
};

module.exports.run = function() {
    const acc  = dicom_array[Math.floor(Math.random()*(dicom_array.length))];
        let args = {
            chaincodeFunction: 'getAsset',
            chaincodeArguments: [acc],
        };

    return bc.bcObj.querySmartContract(contx, 'dicom-caliper', '1', args, 10);
    
};

module.exports.end = function() {
    // do nothing
    return Promise.resolve();
};