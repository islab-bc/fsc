const express = require("express");
const contractRouter = require("../../connection");
const config = require("../../config/config.json");
const { logger } = require("../../config/logConfig");
require('dotenv').config();
const router = express.Router();

contractRouter(config).then((connection) => {
    gateway = connection.gateway;
    contract = connection.contract;

    router.post("/create", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("CreateComment", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/update", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("UpdateComment", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/delete", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("DeleteComment", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/list/boardkey", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetCommentsByBoardKey", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/list/commenter", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetCommentsByCommenter", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/list/did", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetCommentsByDID", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });
});

module.exports.router = router;