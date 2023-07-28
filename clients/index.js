const express = require("express");
const bodyParser = require("body-parser");
const router = require('./routes/index');
const { logger } = require("./config/logConfig");
require('dotenv').config();

const app = express();

app.use(bodyParser.json());
app.use(
    bodyParser.urlencoded({
        extended: true,
    })
);

app.use('/chaincode/test-api', router.app);

app.listen(process.env.PORT, () => {
    logger.info(`-api listening on port ${process.env.PORT}!`);
});

process.on("SIGINT", async() => {
    logger.error("Caught interrupt signal - start disconnect from the gateway");
    gateway.disconnect();
    contract.removeContractListener();
    process.exit();
});