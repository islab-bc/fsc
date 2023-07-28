const axios = require("axios");
const {
    logger
} = require("../config/logConfig");
require('dotenv').config();

function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

const boardEnrolledTime = ["2022.02.22 01:39:312", "2022.02.22 23:10:111", "2022.02.21 11:40:000"]
const boardWriter = ["자크", "헤카림", "리신"]

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

function createComment(requestURL, comment) {
    return new Promise(function (resolve, reject) {
        axios.post(requestURL, comment).then(result => {
            resolve(result.data)
        });
    });
}

async function main() {
    let promises = [];
    for (let i = 0; i < 10; ++i) {
        let comment = {}
        const timeStamp =  getTimeStamp();
        comment.board_company = `boardCompany${i}`;
        comment.board_writer = boardWriter[i%3];
        comment.board_enrolled_time = boardEnrolledTime[i%3];
        comment.parent_commenter = "";
        comment.parent_enrolled_time = "";
        comment.commenter = `댓글작성자${i}`;
        comment.comment_content = `${timeStamp}에 쓴 댓글입니다.`
        comment.comment_enrolled_time = timeStamp
        comment.comment_updated_time = timeStamp
        console.log(comment)
        await sleep(300)
        const prom = createComment(`http://${process.env.SERVER_HOST}:${process.env.PORT}/chaincode/test-api/comment/create`, comment);
        promises.push(prom)
    }
    Promise.all(promises).then(() => {
        logger.info("Init Complete");
    })
    await sleep(5000)

}

main()