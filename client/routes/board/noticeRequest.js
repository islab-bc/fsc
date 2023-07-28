const express = require("express");
const contractRouter = require("../../connection");
const config = require("../../config/config.json");
const { logger } = require("../../config/logConfig");
require('dotenv').config();
const router = express.Router();

contractRouter(config).then((connection) => {
    gateway = connection.gateway;
    contract = connection.contract;

    router.post("/query/one", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetOneNoticeDetail", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        console.log(jsonData);
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/fixed/create", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("CreateFixedNotice", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/fixed/delete", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("DeleteFixedNotice", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/fixed/update", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("UpdateFixedNotice", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });


    router.get("/fixed/query/list/all", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetFixedNoticeList");
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/notfixed/create", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("CreateNotFixedNotice", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/notfixed/delete", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("DeleteNotFixedNotice", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/notfixed/update", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("UpdateNotFixedNotice", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/notfixed/query/list/all", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetNotFixedNoticeList", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });


    router.post("/query/list/writer", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetNoticeListByWriter", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/list/did", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetNoticeListByDID", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });
});

module.exports.router = router;