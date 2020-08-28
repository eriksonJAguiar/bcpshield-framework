'use strict';

module.exports.info  = 'Share asset for k-anoninity';
const { v4: uuidv4 } = require('uuid');

let bc, contx;
let request_array = []
module.exports.init = async function(blockchain, context, args) {
    bc = blockchain;
    contx = context;

    return Promise.resolve();
};


/**
 * Generates simple workload
 * @returns {Object} array of json objects
 */
module.exports.run = function() {
    
    // let args = {
    //     chaincodeFunction: 'shareAssetWithDoctor',
    //     chaincodeArguments: ["2808886", "2002020","hash IPFS", "haewaf"],
    // };
    let settings = {
        chaincodeFunction: 'shareAssetWithDoctor',
        chaincodeArguments: [uuidv4(),"2808886", "2002020","hash IPFS", "00eeccd4-623e-4504-bf16-deee0d2ee01c"],
    };
    var contracValue =  bc.invokeSmartContract(contx, 'dicom-contract', '1', settings, 100);
    contracValue.
    then((result) =>{
        var objResp = JSON.stringify(result);
        var jsonResp = JSON.parse(objResp);
        request_array.push(jsonResp[0].status.id)
    });
    return contracValue;
};

module.exports.end = function() {
    return Promise.resolve();
};

module.exports.request_array = request_array;