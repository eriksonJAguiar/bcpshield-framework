/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Contract } = require('fabric-contract-api');

class DicomContract extends Contract {

    async dicomExists(ctx, dicomId) {
        const buffer = await ctx.stub.getState(dicomId);
        return (!!buffer && buffer.length > 0);
    }

    async createDicom(ctx, dicomId, typeExam, owner) {
        const exists = await this.dicomExists(ctx, dicomId);
        if (exists) {
            throw new Error(`The dicom ${dicomId} already exists`);
        }
        const dicom = { 
            dicomId,
            typeExam, 
            owner
         };
        const buffer = Buffer.from(JSON.stringify(dicom));
        await ctx.stub.putState(dicomId, buffer);
    }   

     async shareDicom(ctx, tokenDicom, to, toOrganization, accessTime) {
        const exists = await this.dicomExists(ctx, tokenDicom);
        if (exists) {
            throw new Error(`The logs ${tokenDicom} already exists`);
        }
        const logsAccess = { 
            tokenDicom, 
            to, 
            toOrganization, 
            accessTime
        };
        const buffer = Buffer.from(JSON.stringify(logsAccess));
        await ctx.stub.putState(tokenDicom, buffer);
    }


    async readAccessLog(ctx, tokenDicom) {
        const exists = await this.dicomExists(ctx, tokenDicom);
        if (!exists) {
            throw new Error(`The logs ${tokenDicom} does not exist`);
        }
    
        const buffer = await ctx.stub.getState(tokenDicom);
        const logsAccess = JSON.parse(buffer.toString());
        return logsAccess;
    }
    
    async readDicom(ctx, dicomId) {
        const exists = await this.dicomExists(ctx, dicomId);
        if (!exists) {
            throw new Error(`The dicom ${dicomId} does not exist`);
        }
        const buffer = await ctx.stub.getState(dicomId);
        const dicom = JSON.parse(buffer.toString());
        return dicom;
    }

}

module.exports = DicomContract;
