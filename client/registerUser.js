"use strict";

const { Wallets } = require("fabric-network");
const FabricCAServices = require("fabric-ca-client");
const fs = require("fs");
const path = require("path");
const { logger } = require("./config/logConfig");

const config = require("./config/config.json");

async function main() {
  try {
    const ccpPath = path.resolve(__dirname, config.ccpPath);
    const ccp = JSON.parse(fs.readFileSync(ccpPath, "utf8"));

    const caURL = ccp.certificateAuthorities[config.ca].url;
    const ca = new FabricCAServices(caURL);

    const walletPath = path.join(process.cwd(), "wallet");
    const wallet = await Wallets.newFileSystemWallet(walletPath);
    logger.info(`Wallet path: ${walletPath}`);

    const userIdentity = await wallet.get(config.identityLabel);
    if (userIdentity) {
      logger.info(
        `An identity for the user ${config.identityLabel} already exists in the wallet`
      );
      return;
    }

    const adminIdentity = await wallet.get("admin");
    if (!adminIdentity) {
      logger.info(
        'An identity for the admin user "admin" does not exist in the wallet'
      );
      logger.info("Run the enrollAdmin.js application before retrying");
      return;
    }

    const provider = wallet
      .getProviderRegistry()
      .getProvider(adminIdentity.type);
    const adminUser = await provider.getUserContext(adminIdentity, "admin");

    const secret = await ca.register(
      {
        affiliation: "btp.example.com.user",
        enrollmentID: config.identityLabel,
        role: "client",
      },
      adminUser
    );
    const enrollment = await ca.enroll({
      enrollmentID: config.identityLabel,
      enrollmentSecret: secret,
    });
    const x509Identity = {
      credentials: {
        certificate: enrollment.certificate,
        privateKey: enrollment.key.toBytes(),
      },
      mspId: config.msp,
      type: "X.509",
    };
    await wallet.put(config.identityLabel, x509Identity);
    logger.info("identity: ", config.identityLabel, x509Identity);
    logger.info(
      `Successfully registered and enrolled admin user ${config.identityLabel} and imported it into the wallet`
    );
  } catch (error) {
    logger.error(`Failed to register user ${config.identityLabel}: ${error}`);
    process.exit(1);
  }
}

main();
