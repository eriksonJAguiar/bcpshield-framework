'use strict';

module.exports.info = 'create dicom imaging';
const { v1: uuidv4 } = require('uuid')

let account_array = [];

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
            chaincodeArguments: [uuidv4(), "10005", "11110", "Bob", "Singer", "(43) 9900 0000", "São Paulo SP", "28", "1992-15-08", "USP", "AAAAA", "None", "Male", "ASASSAS", "Plan X", "75.5", "1.89", "ASAEDF"],
        });
    }
    return workload;
}

module.exports.run = function () {
    let args = generateWorkload();
    return bc.invokeSmartContract(contx, 'dicom', '1', args);
};

module.exports.end = function () {
    return Promise.resolve();
};

module.exports.account_array = account_array;
