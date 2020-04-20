/*
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { FileSystemWallet, Gateway, X509WalletMixin } = require('fabric-network');
const path = require('path');

async function registerUser(org,user,MSP) {
    try {

        // const args = process.argv.slice(2);
        // const org = args[0];
        // const user = args[1];
        const ccpPath = path.resolve(__dirname, '..', 'hyperledger-network','connections', `connection-${org}.json`);
        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), `../../wallet/wallet-${org}`);
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const userExists = await wallet.exists(user);
        if (userExists) {
            console.log(`An identity for the user "${user}" already exists in the wallet`);
            return false;
        }

        // Check to see if we've already enrolled the admin user.
        const adminExists = await wallet.exists('admin');
        if (!adminExists) {
            console.log('An identity for the admin user "admin" does not exist in the wallet');
            console.log('Run the enrollAdmin.js application before retrying');
            return false;
        }
        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'admin', discovery: { enabled: true, asLocalhost: true } });
        
        // Get the CA client object from the gateway for interacting with the CA.
        const ca = gateway.getClient().getCertificateAuthority();
        const adminIdentity = gateway.getCurrentIdentity();

        const secret = await ca.register({enrollmentID: user, role: 'client' }, adminIdentity);
        const enrollment = await ca.enroll({ enrollmentID: user, enrollmentSecret: secret });
        const userIdentity = X509WalletMixin.createIdentity(MSP, enrollment.certificate, enrollment.key.toBytes());
        await wallet.import(user, userIdentity);
        console.log(`Successfully registered and enrolled admin user "${user}" and imported it into the wallet`);

    } catch (error) {
        console.error(`Failed to register user "${user}": ${error}`);
        process.exit(1);
    }
}

module.exports = {registerUser}