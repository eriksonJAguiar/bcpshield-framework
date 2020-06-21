'use strict';

var utils = require('../utils/writer.js');
var Default = require('../service/DefaultService');
var { bcInitNetwork, bcAddAsset, bcAuditLog  } = require('../../blockchainApi');


module.exports.addAsset = function addAsset (req, res, next) {
  var body = req.swagger.params['body'].value;
  bcAddAsset(body, res);
  Default.addAsset(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.auditLog = function auditLog (req, res, next) {
  var user = req.swagger.params['user'].value;
  var tokenID = req.swagger.params['tokenID'].value;
  bcAuditLog(tokenID, user, res);
  Default.auditLog(user,tokenID)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.getDicomById = function getDicomById (req, res, next) {
  var user = req.swagger.params['user'].value;
  var dicomId = req.swagger.params['dicomId'].value;
  //BlockchainApi.getAsset(dicomId, user, res);
  Default.getDicomById(user,dicomId)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.getSharedAssetForResearcher = function getSharedAssetForResearcher (req, res, next) {
  var user = req.swagger.params['user'].value;
  var accessID = req.swagger.params['accessID'].value;
  //BlockchainApi.getSharedAssetForResearcher(accessID, user, res);
  Default.getSharedAssetForResearcher(user,accessID)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.getSharedAssetWithDoctor = function getSharedAssetWithDoctor (req, res, next) {
  var user = req.swagger.params['user'].value;
  var hashIPFS = req.swagger.params['hashIPFS'].value;
  //BlockchainApi.getSharedAssetWithDoctor(hashIPFS, user, res);
  Default.getSharedAssetWithDoctor(user,hashIPFS)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.initNetwork = function initNetwork (req, res, next) {
  bcInitNetwork(res);
  Default.initNetwork()
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.registerUser = function registerUser (req, res, next) {
  var body = req.swagger.params['body'].value;
  //BlockchainApi.registerUser(body, res);
  Default.registerUser(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.requestAssetForResearcher = function requestAssetForResearcher (req, res, next) {
  var body = req.swagger.params['body'].value;
  //BlockchainApi.requestAssetForResearcher(body, res);
  Default.requestAssetForResearcher(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.shareAssetForResearcher = function shareAssetForResearcher (req, res, next) {
  var body = req.swagger.params['body'].value;
  //BlockchainApi.shareAssetForResearcher(body, res);
  Default.shareAssetForResearcher(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.shareAssetWithDoctor = function shareAssetWithDoctor (req, res, next) {
  var body = req.swagger.params['body'].value;
  //BlockchainApi.shareAssetWithDoctor(body, res);
  Default.shareAssetWithDoctor(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
