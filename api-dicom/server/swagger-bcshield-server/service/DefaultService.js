'use strict';


/**
 * Add a new DICOM image on blockchain
 * 
 *
 * body Dicom 
 * returns ApiResponse
 **/
exports.addAsset = function(body) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "result" : "result",
  "status" : "status"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Audit logs from token hash
 * Returns Dicom reference
 *
 * user String ID of user
 * tokenID String Reference token to audit image leaked
 * returns ApiResponseDicom
 **/
exports.auditLog = function(user,tokenID) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "result" : "result"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Find Dicom imaging by ID
 * Returns a single Dicom
 *
 * user String ID of user
 * dicomId Integer ID of Dicom to return
 * returns ApiResponseDicom
 **/
exports.getDicomById = function(user,dicomId) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "result" : "result"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Find shared imaging ID
 * Returns Imaging reference shared with reserchers
 *
 * user String ID of user
 * accessID String ID of imaging request
 * returns ApiResponseDicom
 **/
exports.getSharedAssetForResearcher = function(user,accessID) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "result" : "result"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Find shared imaging ID
 * Returns Imaging reference with a doctor
 *
 * user String ID of user
 * hashIPFS String Hash string of Dicom to get
 * returns ApiResponseDicom
 **/
exports.getSharedAssetWithDoctor = function(user,hashIPFS) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "result" : "result"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Start wallets on network
 * 
 *
 * returns ApiResponse
 **/
exports.initNetwork = function() {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "result" : "result",
  "status" : "status"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Register a user on blockchain
 * 
 *
 * body User 
 * returns ApiResponse
 **/
exports.registerUser = function(body) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "result" : "result",
  "status" : "status"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * A researcher request to patient Dicom imaging
 * 
 *
 * body ResearcherRequest 
 * returns ApiResponse
 **/
exports.requestAssetForResearcher = function(body) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "result" : "result",
  "status" : "status"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * A patient share imging with a researcher
 * 
 *
 * body PatientShareResearcher 
 * returns ApiResponse
 **/
exports.shareAssetForResearcher = function(body) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "result" : "result",
  "status" : "status"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Sharing an imaging with a doctor
 * 
 *
 * body SharedDicom 
 * returns ApiResponse
 **/
exports.shareAssetWithDoctor = function(body) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "result" : "result",
  "status" : "status"
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}

