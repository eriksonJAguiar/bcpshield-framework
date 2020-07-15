var express = require('express');
var bodyParser = require('body-parser');
var app = express();
const fabricNetwork = require('./fabricNetwork');
const enrollAdmin = require('../enrollAdmin');
const registerUser = require('../registerUser');
app.set('view engine', 'ejs');
app.use(bodyParser.json());
urlencoder = bodyParser.urlencoded({ extended: true });
const {v4:uuid} = require('uuid');


/**
 * This function starting blockchain network and add admin credentials
 * @public
 * @returns {json} status respose
 */
app.get('/api/initNetwork', urlencoder, async function (req, res) {
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

});

/**
 * This function adds one to its input.
 * @param {string} req.body.org organization {hprovider, research or patient}
 * @param {string} req.body.msp organization MSP {HProviderMSP, ResearchMSP or PatientMSP}
 * @returns {json} status json
 */
app.post('/api/registerUser', urlencoder, async function (req, res) {
  try {
    let userID = uuid();
    let result = registerUser.registerUser(req.body.org, userID, req.body.msp);
    res.json({
      status: 'True',
      result: result.toString(),
      userID: userID
    });
    console.log('OK - Transaction has been submitted');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error,
      status: 'False'
    });
  }
});

/**
 * This function adds asset on blockchain
 * @param {string} req.body.user user id on blockchain
 * @param {string} req.body.dicomID dicom ID
 * @param {string} req.body.patientID patient ID
 * @param {string} req.body.patientFirstname patient first name
 * @param {string} req.body.patientLastname patient last name
 * @param {string} req.body.patientTelephone patient phone number
 * @param {string} req.body.patientAddress patient address
 * @param {int} req.body.patientAge patient age
 * @param {string} req.body.patientOrganization org which patient to join
 * @param {string} req.body.patienRace  patient insuranceplan
 * @param {string} req.body.patientGender  patient gender
 * @param {string} req.body.patientInsuranceplan patient insuranceplan
 * @param {float} req.body.patientWeigth patient weigth
 * @param {float} req.body.patientHeigth patient heigth
 * @param {string} req.body.machineModel patient insuranceplan
 * @returns {number} that number, plus one.
 */
app.post('/api/addAsset', urlencoder, async function (req, res) {

  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../wallet/wallet-hprovider', req.body.user.toString());
    console.log(req.body);
    let response = await contract.submitTransaction('addAsset', req.body.dicomID.toString(), req.body.patientID.toString(), req.body.patientFirstname.toString(), req.body.patientLastname.toString(), 
                                              req.body.patientTelephone.toString(), req.body.patientAddress.toString(), req.body.patientAge.toString(), req.body.patientOrganization.toString(), 
                                              req.body.patientRace.toString(), req.body.patientGender.toString(), req.body.patientInsuranceplan.toString(), 
                                              req.body.patientWeigth.toString(), req.body.patientHeigth.toString(), req.body.machineModel.toString());
    res.json({
      status: 'OK - Transaction has been submitted',
      result: response
    });
    console.log('OK - Transaction has been submitted');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }

});

//addAssetDiff


/**
 * This function add asset on blockchain apply differential privacy
 * @param {string} req.body.user user id on blockchain
 * @param {string} req.body.dicomID dicom private will add on blockchain
 * @param {string} req.body.asset string for convert to byte
 * @returns {json} updated shared log
 */
app.post('/api/addAssetDiff', urlencoder, async function (req, res) {

  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../wallet/wallet-hprovider', req.body.user.toString());
    console.log(req.body);
    let response = await contract.submitTransaction('addAssetDiff', req.body.dicomID.toString(), req.body.asset);
    res.json({
      status: 'OK - Transaction has been submitted',
      result: response
    });
    console.log('OK - Transaction has been submitted');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }

});

/**
 * This function notify a requester and add struct on blockchain
 * @param {string} req.body.user user id on blockchain
 * @param {string} req.body.requestID ID for request
 * @param {string} req.body.assets assets shared
 * @returns {json} updated shared log
 */
app.post('/api/notify', urlencoder, async function (req, res) {

  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../wallet/wallet-hprovider', req.body.user.toString());
    console.log(req.body);
    let response = await contract.submitTransaction('notifyRequester', req.body.requestID.toString(), req.body.assets.toString());
    res.json({
      status: 'OK - Transaction has been submitted',
      result: response
    });
    console.log('OK - Transaction has been submitted');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }

});

/**
 * This function add private assets on blockchain
 * @param {string} req.body.user user id on blockchain
 * @param {string} req.body.dicomID dicom ID
 * @param {string} req.body.patientID patient ID
 * @param {string} req.body.patientFirstname patient first name
 * @param {string} req.body.patientLastname patient last name
 * @param {string} req.body.patientTelephone patient phone number
 * @param {string} req.body.patientAddress patient address
 * @param {int} req.body.patientAge patient age
 * @param {string} req.body.patientOrganization org which patient to join
 * @param {string} req.body.patienRace  patient insuranceplan
 * @param {string} req.body.patientGender  patient gender
 * @param {string} req.body.patientInsuranceplan patient insuranceplan
 * @param {float} req.body.patientWeigth patient weigth
 * @param {float} req.body.patientHeigth patient heigth
 * @param {string} req.body.machineModel patient insuranceplan
 * @returns {number} that number, plus one.
 */
app.post('/api/addAssetPriv', urlencoder, async function (req, res) {

  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../wallet/wallet-hprovider', req.body.user);
    console.log(req.body);
    let response = await contract.submitTransaction('addAssetPriv', req.body.dicomID, req.body.patientID, req.body.patientFirstname, req.body.patientLastname, 
                                              req.body.patientTelephone, req.body.patientAddress, req.body.patientAge.toString(), " ", req.body.patientOrganization, 
                                              " ", req.body.patientRace, " ", req.body.patientGender, req.body.patientInsuranceplan, 
                                              req.body.patientWeigth.toString(), req.body.patientHeigth.toString(), req.body.machineModel);
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

});

/**
 * This get asset on blockchain
 * @param {string} req.body.user user id on blockchain
 * @param {string} req.body.dicomID dicom ID
 * @returns {json:Dicom} asset request
 */
app.get('/api/getAsset', async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../wallet/wallet-hprovider', req.body.user);
    const result = await contract.evaluateTransaction('getAsset', req.body.dicomId);
    let response = JSON.parse(result);
    res.json({ 
      status: 'OK - Transaction has been submitted',
      result: response
     });
    console.log('OK - Query Successful');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }
});

/**
 * This get private asset on blockchain
 * @param {number} input any number
 * @returns {number} that number, plus one.
 */
app.get('/api/getAssetPriv', async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../wallet/wallet-hprovider', req.body.user);
    const result = await contract.evaluateTransaction('getAssetPriv', req.body.dicomId);
    let response = JSON.parse(result.toString());
    res.json({ 
      status: 'OK - Transaction has been submitted',
      result: response
     });
    console.log('OK - Query Successful');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }
});

/**
 * This function a doctor request asset from patient
 * @param {string} req.body.user user id on blockchain
 * @param {string} req.body.patientID patient that own the imaging
 * @returns {json:Dicom} asset request
 */
app.post('/api/requestAssetDoctor', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../wallet/wallet-hprovider', req.body.user.toString());
    console.log(req.body);
    let response = await contract.submitTransaction('requestAsset', req.body.user.toString(), req.body.patientID.toString(), "Doctor");
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
});

/**
 * This function a researcher request asset to study it 
 *  @param {string}  req.body.user  researcher ID
 * @param {string}  req.body.patientID the patient that asset owns
 * @returns {json:string} A string with request ID 
 */
app.post('/api/requestAssetResearcher', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-research.json', '../wallet/wallet-research', req.body.user.toString());
    console.log(req.body);
    let response = await contract.submitTransaction('requestAsset', req.body.user.toString(), req.body.patientID.toString(), "Researcher");
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
});

/**
/**
 * This function to get a asset shared with doctor
 *  @param {string}  req.body.user  doctor ID
 * @param {string}  req.body.accessID patient that asset owns
 * @returns {json:string} A string with request ID 
 */
app.get('/api/getRequestedDoctor', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../wallet/wallet-hprovider', req.body.user);
    let result = await contract.submitTransaction('getRequested', req.body.accessID.toString());
    const response = JSON.parse(result.toString());
    res.json({ result: response.toString() });
    console.log('OK - Transaction has been submitted');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }

});

/**
/**
 * This function to get a asset shared with doctor
 *  @param {string}  req.body.user  doctor ID
 * @param {string}  req.body.accessID patient that asset owns
 * @returns {json:string} A string with request ID 
 */
app.get('/api/getRequestedResearcher', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-research.json', '../wallet/wallet-research', req.body.user);
    let result = await contract.submitTransaction('getRequested', req.body.accessID.toString());
    const response = JSON.parse(result.toString());
    res.json({ result: response.toString() });
    console.log('OK - Transaction has been submitted');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }

});

/**
 * This function observer imaging requested on blockchain
 * @param {string} req.body.user id hprovider
 * @param {string} req.body.timestamp last time observing
 * @returns {json} json with last requests
 */
app.get('/api/observerRequests', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../wallet/wallet-hprovider', req.body.user);
    let result = await contract.submitTransaction('observerRequests',  req.body.user.toString(),req.body.timestamp.toString());
    const response = JSON.parse(result.toString());
    res.json({ result: response.toString() });
    console.log('OK - Transaction has been submitted');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }

});

/**
/**
 * This function to get a asset shared with doctor
 *  @param {string}  req.body.user  doctor ID
 * @param {string}  req.body.tokenID token ID to audit leakage on blockchain
 * @returns {json} Logs from token ID
 */
app.get('/api/auditLog', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../wallet/wallet-hprovider', req.body.user);
    const result = await contract.submitTransaction('auditLog', req.body.tokenID.toString());
    let response = JSON.parse(result.toString());
    res.json({ result: response.toString() });
    console.log('OK - Transaction has been submitted');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }

});



app.listen(3000, () => {
  console.log("***********************************");
  console.log("API server listening at localhost:3000");
  console.log("***********************************");
});
