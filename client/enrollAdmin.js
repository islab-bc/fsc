'use strict';

const FabricCAServices = require('fabric-ca-client');
const { Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const { logger } = require("./config/logConfig");

const config = require('./config/config.json');

async function main() {
    try {
        const ccpPath = path.resolve(__dirname, config.ccpPath);
        const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        const caInfo = ccp.certificateAuthorities[config.ca];
        const caTLSCACerts = caInfo.tlsCACerts.pem;
        const ca = new FabricCAServices(caInfo.url, { trustedRoots: caTLSCACerts, verify: false }, caInfo.caName);

        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        logger.info(`Wallet path: ${walletPath}`);

        const identity = await wallet.get('admin');
        if (identity) {
            logger.info('An identity for the admin user "admin" already exists in the wallet'+identity);
            return;
        }

        const enrollment = await ca.enroll({ enrollmentID: 'admin', enrollmentSecret: 'adminpw' });

        logger.info(enrollment.certificate)
        const x509Identity = {
            credentials: {
                certificate: enrollment.certificate,
                privateKey: enrollment.key.toBytes(),
            },
            mspId: config.msp,
            type: 'X.509',
        };
        await wallet.put('admin', x509Identity);
        logger.info('Successfully enrolled admin user "admin" and imported it into the wallet');

    } catch (error) {
        logger.error(`Failed to enroll admin user "admin": ${error}`);
        process.exit(1);
    }
}

main();
