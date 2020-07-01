'use strict';

const { WorkloadModuleBase } = require('@hyperledger/caliper-core');

const dic = 'abcdefghijklmnopqrstuvwxyz';
let dicom_array = [];

/**
 * Workload module for the benchmark round.
 */
class Workload extends WorkloadModuleBase {
    /**
     * Initializes the workload module instance.
     */
    constructor() {
        super();
        this.txnPerBatch = 1;
        this.initDicom = 0;
        this.prefix = '';
    }

    /**
     * Initialize the workload module with the given parameters.
     * @param {number} workerIndex The 0-based index of the worker instantiating the workload module.
     * @param {number} totalWorkers The total number of workers participating in the round.
     * @param {number} roundIndex The 0-based index of the currently executing round.
     * @param {Object} roundArguments The user-provided arguments for the round from the benchmark configuration file.
     * @param {BlockchainInterface} sutAdapter The adapter of the underlying SUT.
     * @param {Object} sutContext The custom context object provided by the SUT adapter.
     * @async
     */
    async initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext) {
        await super.initializeWorkloadModule(workerIndex, totalWorkers, roundIndex, roundArguments, sutAdapter, sutContext);

        if(!this.roundArguments.hasOwnProperty('dicom')) {
            throw new Error('simple.open - \'ducom\' is missed in the arguments');
        }

        this.initDicom = this.roundArguments.dicom;
        this.txnPerBatch = this.roundArguments.txnPerBatch || 1;
        this.prefix = this.workerIndex.toString();
    }

    /**
     * Generate string by picking characters from dic variable
     * @param {*} number character to select
     * @returns {String} string generated based on @param number
     */
    _get26Num(number){
        let result = '';
        while(number > 0) {
            result += dic.charAt(number % 26);
            number = Math.floor(number/26);
        }
        return result;
    }

    /**
     * Generate unique dicom key for the transaction
     * @returns {String} dicom key
     */
    _generateDicom() {
        return this.prefix + this._get26Num(dicom_array.length+1);
    }

    /**
     * Generates simple workload
     * @returns {Object} array of json objects
     */
    _generateWorkload() {
        let workload = [];
        for(let i= 0; i < this.txnPerBatch; i++) {
            let acc_id = this._generateDicom();
            dicom_array.push(acc_id);

            
            workload.push({
                chaincodeFunction: 'addAsset',
                chaincodeArguments: [acc_id, "2808886", "Jeff", "Slavech", "4221513", "310 South Crouse Avenue Syracuse NY", "54", "01-01-1966", "OHIP", "AAAAA", "None", "Male", "ASASSAS", "1780694000", "67.0", "1.77", "C3L-01285"],
            });
        }
        
        return workload;
    }

    /**
     * Assemble TXs for the round.
     * @return {Promise<TxStatus[]>}
     */
    async submitTransaction() {
        let args = this._generateWorkload();
        return this.sutAdapter.invokeSmartContract(this.sutContext, 'dicom-caliper', '1', args, 100);
    }
}

/**
 * Create a new instance of the workload module.
 * @return {WorkloadModuleInterface}
 */
function createWorkloadModule() {
    return new Workload();
}

module.exports.createWorkloadModule = createWorkloadModule;
module.exports.dicom_array = dicom_array;