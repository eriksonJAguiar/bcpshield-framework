var express = require('express');
var bodyParser = require('body-parser');
var app = express();
const fabricNetwork = require('./fabricNetwork')
const enrollAdmin = require('../enrollAdmin')
const registerUser = require('../registerUser')
app.set('view engine', 'ejs');
app.use(bodyParser.json());
urlencoder = bodyParser.urlencoded({ extended: true });

try {
  enrollAdmin.enrollAdmin('hprovider', 'HProviderMSP');
  enrollAdmin.enrollAdmin('research', 'ResearchMSP'); 
} catch (error) {
  console.log(error);
}


app.post('/api/registerUser', urlencoder, async function (req, res) {
  try {
    let result = registerUser.registerUser(req.body.org, req.body.user,req.body.msp);
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

app.post('/api/createDicom', urlencoder, async function (req, res) {

  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../../wallet/wallet-hprovider', req.body.user);
    console.log(req.body);
    let tx = await contract.submitTransaction('createDicom', req.body.dicomId, req.body.typeExam, req.body.owner);
    res.json({
      status: 'OK - Transaction has been submitted',
      txid: tx.toString()
    });
    console.log('OK - Transaction has been submitted');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }

});

app.get('/api/readDicom/:dicomId', async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-hprovider.json', '../../wallet/wallet-hprovider', req.params.user.toString());
    const result = await contract.evaluateTransaction('readDicom', req.params.dicomId.toString());
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

app.post('/api/shareDicom', urlencoder, async function (req, res) {

  try {
    const contract = await fabricNetwork.connectNetwork('connection-research.json', '../../wallet/wallet-research', req.body.user);
    console.log(req.body);
    let tx = await contract.submitTransaction('shareDicom', req.body.tokenDicom, req.body.to, req.body.toOrganization, Date.now().toString());
    res.json({
      status: 'OK - Transaction has been submitted',
      txid: tx.toString()
    });
    console.log('OK - Transaction has been submitted');
  } catch (error) {
    console.error(`Failed to evaluate transaction: ${error}`);
    res.status(500).json({
      error: error
    });
  }

});

app.get('/api/readAccessLog/:tokenDicom', async function (req, res) {
  try {
    const contract = await fabricNetwork.connectNetwork('connection-research.json', '../../wallet/wallet-research', req.params.user.toString());
    const result = await contract.evaluateTransaction('readAccessLog', req.params.tokenDicom.toString());
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





app.listen(3000, () => {
  console.log("***********************************");
  console.log("API server listening at localhost:3000");
  console.log("***********************************");
});
