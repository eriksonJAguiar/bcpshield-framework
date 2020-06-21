const fabricNetwork = require('./fabricNetwork');
const enrollAdmin = require('../enrollAdmin');
const registerUserMod = require('../registerUser');

module.exports.initNetwork = async function (res) {
    try {
        enrollAdmin.enrollAdmin('hprovider', 'HProviderMSP');
        enrollAdmin.enrollAdmin('research', 'ResearchMSP');
        enrollAdmin.enrollAdmin('patient', 'PatientMSP');
        res.json({
            status: 'True'
        });
        console.log('OK - Transaction has been submitted');
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
            error: error,
            status: 'False'
        });
    }
}

module.exports.registerUser = async function (body, res) {
    try {
        let result = registerUserMod.registerUser(body.org, body.user, body.msp);
        res.json({
            status: 'True'
        });
        console.log('OK - Transaction has been submitted');
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
            error: error,
            status: 'False'
        });
    }
}

module.exports.addAsset = async function(body, res) {
    try {
        const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../../wallet/wallet-hprovider', body.user);
        console.log(body);
        var currentdate = new Date();
        let response = await contract.submitTransaction('addAsset', body.dicomID, body.patientID, body.patientFirstname, body.patientLastname,
            body.patientTelephone, body.patientAddress, body.patientAge, "", body.patientOrganization,
            "", body.patientRace, "", body.patientGender, body.patientInsuranceplan,
            body.patientWeigth, body.patientHeigth, body.machineModel, currentdate.getDate().toString());
        res.json({
            status: 'OK - Transaction has been submitted',
            result: response.toString()
        });
        console.log('OK - Transaction has been submitted');
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
            error: error
        });
    }
}

module.exports.getAsset = async function(dicomId, user, res) {
    try {
        const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../../wallet/wallet-hprovider', user.toString());
        const result = await contract.evaluateTransaction('getAsset', dicomId.toString());
        let response = JSON.parse(result.toString());
        res.json({ result: response });
        console.log('OK - Query Successful');
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
            error: error
        });
    }
}

module.exports.shareAssetWithDoctor = async function(body, res) {
    try {
        const contract = await fabricNetwork.connectNetwork('connection-patient.json', '../../wallet/wallet-patient', body.user.toString());
        console.log(body);
        let response = await contract.submitTransaction('shareAssetWithDoctor', body.patientID, body.doctorID, body.hashIPFS, body.dicomID);
        res.json({
            status: 'OK - Transaction has been submitted',
            result: response.toString()
        });
        console.log('OK - Transaction has been submitted');
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
            error: error
        });
    }
}

module.exports.getSharedAssetWithDoctor = async function(hashIPFS, user, res) {
    try {
        const contract = await fabricNetwork.connectNetwork('connection-research.json', '../../wallet/wallet-research', user.toString());
        const result = await contract.evaluateTransaction('getSharedAssetWithDoctor', hashIPFS.toString());
        let response = JSON.parse(result.toString());
        res.json({ result: response });
        console.log('OK - Query Successful');
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
            error: error
        });
    }
}

module.exports.requestAssetForResearcher = async function(body, res) {
    try {
        const contract = await fabricNetwork.connectNetwork('connection-research.json', '../../wallet/wallet-research', body.user);
        console.log(body);
        let response = await contract.submitTransaction('requestAssetForResearcher', body.amount.toString(), body.researchID, body.patientID);
        res.json({
            status: 'OK - Transaction has been submitted',
            result: response.toString()
        });
        console.log('OK - Transaction has been submitted');
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
            error: error
        });
    }
}

module.exports.shareAssetForResearcher = async function(body, res) {
    try {
        const contract = await fabricNetwork.connectNetwork('connection-patient.json', '../../wallet/wallet-patient', body.user);
        console.log(req.body);
        let response = await contract.submitTransaction('shareAssetForResearcher', body.requestID.toString());
        res.json({
            status: 'OK - Transaction has been submitted',
            result: response.toString()
        });
        console.log('OK - Transaction has been submitted');
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
            error: error
        });
    }
}

module.exports.getSharedAssetForResearcher = async function(accessID, user, res) {
    try {
        const contract = await fabricNetwork.connectNetwork('connection-research.json', '../../wallet/wallet-research', user);
        let result = await contract.submitTransaction('getSharedAssetForResearcher', accessID.toString());
        const response = JSON.parse(result.toString());
        res.json({ result: response.toString() });
        console.log('OK - Transaction has been submitted');
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
            error: error
        });
    }
}

module.exports.auditLog = async function(tokenID, user, res) {
    try {
        const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../../wallet/wallet-hprovider', user);
        const result = await contract.submitTransaction('auditLog', tokenID.toString());
        let response = JSON.parse(result.toString());
        res.json({ result: response.toString() });
        console.log('OK - Transaction has been submitted');
    } catch (error) {
        console.error(`Failed to evaluate transaction: ${error}`);
        res.status(500).json({
            error: error
        });
    }
}