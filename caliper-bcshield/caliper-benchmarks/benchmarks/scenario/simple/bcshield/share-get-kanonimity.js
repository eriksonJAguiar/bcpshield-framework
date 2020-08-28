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

module.exports.info  = 'get Asset from K-Anonymity values';


let bc, contx;
let request_array;

module.exports.init = function(blockchain, context, args) {
    const sharedAsset = require('./share-kanonimity.js');
    bc       = blockchain;
    contx    = context;
    request_array = sharedAsset.request_array;

    return Promise.resolve();
};

module.exports.run = function() {
        const id  = request_array[Math.floor(Math.random()*(request_array.length))];
        let args = {
            chaincodeFunction: 'getSharedAssetWithDoctor',
            chaincodeArguments: [id],
        };

    return bc.querySmartContract(contx, 'dicom-contract', '1', args, 1000);
    
};

module.exports.end = function() {
    // do nothing
    return Promise.resolve();
};