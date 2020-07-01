'use strict';

module.exports.info = 'create dicom imaging';
const { v1: uuidv4 } = require('uuid')

let dicom_array = [];

let bc, contx;
var txnPerBatch = 1
module.exports.init = function (blockchain, context, args) {
    if (!args.hasOwnProperty('txnPerBatch')) {
        args.txnPerBatch = 1;
    }
    txnPerBatch = args.txnPerBatch;
    bc = blockchain;
    contx = context;

    return Promise.resolve();
};


function generateWorkload() {
    let workload = [];
    for (let i = 0; i < txnPerBatch; i++) {
        workload.push({
            chaincodeFunction: 'addAsset',
            chaincodeArguments: [uuidv4(), "2808886", "Jeff", "Slavech", "4221513", "310 South Crouse Avenue Syracuse NY", "54", "01-01-1966", "OHIP", "AAAAA", "None", "Male", "ASASSAS", "1780694000", "67.0", "1.77", "C3L-01285"],
        });
    }
    return workload;
}

module.exports.run = function () {
    let args = generateWorkload();
    return bc.invokeSmartContract(contx, 'dicom-caliper', '1', args);
};

module.exports.end = function () {
    return Promise.resolve();
};

module.exports.dicom_array = dicom_array;
