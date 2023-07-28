const express = require("express");
const contractRouter = require("../../connection");
const config = require("../../config/config.json");
const { logger } = require("../../config/logConfig");
require('dotenv').config();
const router = express.Router();

// function getTimeStamp() {
//     let today = new Date();
//     let year = today.getFullYear(); // 년도
//     let month = today.getMonth() + 1; // 월
//     let date = today.getDate(); // 날짜
//     let hour = today.getHours(); //시간
//     let minute = today.getMinutes(); //분
//     let seconds = today.getSeconds(); //초
//     let milli = today.getMilliseconds(); // 밀리세컨즈
//     let timeStamp = year.toString() + '.' + month.toString().padStart(2, "0") + '.' + date.toString().padStart(2, "0") +
//                     ' ' + hour.toString().padStart(2, "0") + ":" + minute.toString().padStart(2, "0") + ":"
//                     + seconds.toString().padStart(2, "0") + ":" + milli.toString().padStart(2, "0");
//     return timeStamp
// }

contractRouter(config).then((connection) => {
    gateway = connection.gateway;
    contract = connection.contract;

    router.get("/query/list/all", async(request, response) => {
        const bufferedData = await contract.submitTransaction("GetAdmissionList");
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/create", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("EnrollAdmission", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/list/did", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetAdmissionListByDID", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });
});

module.exports.router = router;