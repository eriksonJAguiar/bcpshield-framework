var express = require('express');
var bodyParser = require('body-parser');
const ipfsClient = require('ipfs-http-client');
const { globSource } = ipfsClient
var app = express();
const ipfs = ipfsClient('');
const fabricNetwork = require('./fabricNetwork')
const enrollAdmin = require('../enrollAdmin')
const registerUser = require('../registerUser')
app.set('view engine', 'ejs');
app.use(bodyParser.json());
urlencoder = bodyParser.urlencoded({ extended: true });



app.get('/api/initNetwork', urlencoder, async function (req, res) {
  try {
      enrollAdmin.enrollAdmin('hprovider', 'HProviderMSP');
      enrollAdmin.enrollAdmin('research', 'ResearchMSP');
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
    let result = registerUser.registerUser(req.body.org, req.body.user, req.body.msp);
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

app.post('/api/addAsset', urlencoder, async function (req, res) {

  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../../wallet/wallet-hprovider', req.body.user);
    console.log(req.body);
    var currentdate = new Date();
    let response = await contract.submitTransaction('addAsset', req.body.dicomID, req.body.patientID, req.body.patientFirstname, req.body.patientLastname, 
                                              req.body.patientTelephone, req.body.patientAddress, req.body.patientAge, req.body.patientBirth, req.body.patientOrganization, 
                                              req.body.patientMothername, req.body.patientReligion, req.body.patientSex, req.body.patientGender, req.body.patientInsuranceplan, 
                                              req.body.patientWeigth, req.body.patientHeigth, req.body.machineModel, currentdate.getDate().toString());
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

app.get('/api/getAsset/:dicomId', async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../../wallet/wallet-hprovider', req.params.user.toString());
    const result = await contract.evaluateTransaction('getAsset', req.params.dicomId.toString());
    let response = JSON.parse(result.toString());
    res.json({ result: response });
    console.log('OK - Query Successful');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }
});

app.post('/api/shareAssetWithDoctor', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../../wallet/wallet-hprovider', req.body.user);
    console.log(req.body);
    let response = await contract.submitTransaction('shareAssetWithDoctor', req.body.patientID, req.body.doctorID, req.body.hashIPFS, req.body.dicomID);
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

app.get('/api/getSharedAssetWithDoctor/:hashIPFS', async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-research.json', '../../wallet/wallet-research', req.params.user.toString());
    const result = await contract.evaluateTransaction('getSharedAssetWithDoctor', req.params.hashIPFS.toString());
    let response = JSON.parse(result.toString());
    res.json({ result: response });
    console.log('OK - Query Successful');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }
})

app.post('/api/requestAssetForResearcher', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-research.json', '../../wallet/wallet-research', req.body.user);
    console.log(req.body);
    let response = await contract.submitTransaction('requestAssetForResearcher', req.body.amount.toString(), req.body.researchID, req.body.patientID);
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

app.post('/api/shareAssetForResearcher', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../../wallet/wallet-hprovider', req.body.user);
    console.log(req.body);
    let response = await contract.submitTransaction('shareAssetForResearcher', req.body.requestID.toString());
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

app.get('/api/getSharedAssetForResearcher', urlencoder, async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-research.json', '../../wallet/wallet-research', req.body.user);
    let result = await contract.submitTransaction('getSharedAssetForResearcher', req.body.accessID.toString());
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
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../../wallet/wallet-hprovider', req.body.user);
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


app.post('/api/addIPFS', urlencoder, async function (req, res) {
  try {
    const response = await ipfs.add(globSource(req.body.path.toString(), { recursive: true }))
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

// For getting file we should use a local client



app.listen(3000, () => {
  console.log("***********************************");
  console.log("API server listening at localhost:3000");
  console.log("***********************************");
});
