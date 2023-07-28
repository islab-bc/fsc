const axios = require("axios");
const {
    logger
} = require("../config/logConfig");
require('dotenv').config();

function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

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


async function main() {
    let promises = [];
    //TODO ...
}

main()