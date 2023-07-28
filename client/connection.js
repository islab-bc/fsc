'use strict';
const { logger } = require("./config/logConfig");

module.exports = (config) => {
  const {Wallets, Gateway} = require('fabric-network');
  const fs = require('fs');

  const path = require('path');

  return new Promise (async (resolve, reject) => {
    const ccpPath = path.resolve(__dirname, config.ccpPath);
    const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));
    const walletPath = path.join(__dirname, '.', 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);

    const gateway = new Gateway();
    logger.info("connect시 identityLabel: " + config.identityLabel);
    await gateway.connect(ccp, {
      wallet,
      identity: config.identityLabel,
      discovery: {enabled: true, asLocalhost: true}
    });

    const network = await gateway.getNetwork(config.channel);
    const contract = network.getContract(config.cc);

    let listener;
    let checkTransaction;

    try {
      listener = async (event) => {
        const asset = JSON.parse(event.payload.toString());
        const eventTransaction = event.getTransactionEvent();
        const transaction = eventTransaction.transactionData

        try {
          const checkChaincode = transaction.actions[0].payload.chaincode_proposal_payload.input.chaincode_spec;
          logger.info("listening....")
          logger.info("function :", checkChaincode.input.args[0].toString())
        } catch (error) {
          logger.error(`Failed to submit transaction: ${error}`);
          process.exit(1);
        }

      };
      contract.addContractListener(listener);
    } catch (eventError) {
      logger.info("error");
    }

    const result = {
      gateway: gateway,
      contract: contract,
      listener: listener,
      checkTransaction: checkTransaction
    }
    resolve(result);
  })
}