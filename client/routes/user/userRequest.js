const express = require("express");
const contractRouter = require("../../connection");
const config = require("../../config/config.json");
const { logger } = require("../../config/logConfig");
require('dotenv').config();
const router = express.Router();

function getTimeStamp() {
    let today = new Date();
    let year = today.getFullYear(); // 년도
    let month = today.getMonth() + 1; // 월
    let date = today.getDate(); // 날짜
    let hour = today.getHours(); //시간
    let minute = today.getMinutes(); //분
    let seconds = today.getSeconds(); //초
    let milli = today.getMilliseconds(); // 밀리세컨즈
    let timeStamp = year.toString() + '.' + month.toString().padStart(2, "0") + '.' + date.toString().padStart(2, "0") +
                    ' ' + hour.toString().padStart(2, "0") + ":" + minute.toString().padStart(2, "0") + ":"
                    + seconds.toString().padStart(2, "0") + ":" + milli.toString().padStart(2, "0");
    return timeStamp
}

contractRouter(config).then((connection) => {
    gateway = connection.gateway;
    contract = connection.contract;

    router.get("/query/list", async(request, response) => {
        const bufferedData = await contract.submitTransaction("GetUserList");
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.get("/query/list/approved", async(request, response) => {
        const bufferedData = await contract.submitTransaction("GetApprovedUserList");
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });
    
    router.get("/query/list/pending", async(request, response) => {
        const bufferedData = await contract.submitTransaction("GetPendingUserList");
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.get("/query/list/reject", async(request, response) => {
        const bufferedData = await contract.submitTransaction("GetRejectedUserList");
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.get("/query/list/expired", async(request, response) => {
        const bufferedData = await contract.submitTransaction("GetExpiredUserList");
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.get("/query/list/paging", async(request, response) => {
        const bufferedData = await contract.submitTransaction("GetUserList");
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/create", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        let did = {"did": request.body.did};
        let didJsonData = JSON.stringify(did);
        const userData = await contract.submitTransaction("GetUserProfile", didJsonData);
        const userJson = JSON.parse(String(userData));
        let result = {"data": requestData, "message": "이미 등록된 사용자", "status_code": 500, "level":"info", "timestamp": getTimeStamp()}
        let resultData = JSON.parse(JSON.stringify(result))
        if(userJson.status_code == 500) {
            const bufferedData = await contract.submitTransaction("EnrollUser", requestData);
            const jsonData = JSON.parse(String(bufferedData));
            logger.info(jsonData);
            response.send(jsonData);
        }
        else if(userJson.status_code == 200){
            logger.info(resultData);
            response.send(resultData);
        }
    });

    router.post("/update/did", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const userData = await contract.submitTransaction("GetUserByCI", requestData);
        const userJson = JSON.parse(String(userData));
        const userDataJson = JSON.parse(String(userJson.data));
        let beforeUserData = userDataJson.user;
        beforeUserData['before_did'] = userDataJson.key;
        beforeUserData['did'] = request.body.after_did;
        let JsonToString = JSON.stringify(beforeUserData);
        const bufferedData = await contract.submitTransaction("UpdateDID", JsonToString);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/update/profile", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const userData = await contract.submitTransaction("GetUserProfile", requestData);
        const userJson = JSON.parse(String(userData));
        const userDataJson = JSON.parse(String(userJson.data));
        let beforeUserData = userDataJson.user;
        beforeUserData['before_did'] = userDataJson.key;
        beforeUserData['did'] = request.body.did;
        beforeUserData['company_name'] = request.body.company_name;
        beforeUserData['position'] = request.body.position;
        let JsonToString = JSON.stringify(beforeUserData);
        const bufferedData = await contract.submitTransaction("UpdateUser", JsonToString);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/update/state", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const userData = await contract.submitTransaction("GetUserProfile", requestData);
        const userJson = JSON.parse(String(userData));
        const userDataJson = JSON.parse(String(userJson.data));
        let beforeUserData = userDataJson.user;
        beforeUserData['before_did'] = userDataJson.key;
        beforeUserData['did'] = request.body.did;
        beforeUserData['state'] = request.body.state;
        let JsonToString = JSON.stringify(beforeUserData);
        const bufferedData = await contract.submitTransaction("UpdateUserState", JsonToString);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/delete", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("DeleteUser", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/list/sql", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetUserListByQuery", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/profile/byci", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetUserByCI", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/profile/image", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetUserProfileImage", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/profile", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetUserProfile", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });
    
    router.post("/query/profile/devicetoken", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetUserDeviceToken", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/profile/image", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetUserProfileImage", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/profile/ci", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetUserCI", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });

    router.post("/query/transaction", async(request, response) => {
        const requestData = JSON.stringify(request.body);
        const bufferedData = await contract.submitTransaction("GetTransactionList", requestData);
        const jsonData = JSON.parse(String(bufferedData));
        logger.info(jsonData);
        response.send(jsonData);
    });
});

module.exports.router = router;