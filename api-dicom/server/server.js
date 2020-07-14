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

app.post('/api/notify', urlencoder, async function (req, res) {

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

app.post('/api/requestAssetDoctor', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../wallet/wallet-hprovider', req.body.user.toString());
    console.log(req.body);
    let response = await contract.submitTransaction('requestAsset', req.body.patientID, req.body.doctorID, "Doctor");
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

app.post('/api/requestAssetResearcher', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-research.json', '../wallet/wallet-research', req.body.user.toString());
    console.log(req.body);
    let response = await contract.submitTransaction('requestAsset', req.body.patientID, req.body.researcherID, "Researcher");
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
