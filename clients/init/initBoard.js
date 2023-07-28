const axios = require("axios");
const {
    logger
} = require("../config/logConfig");
require('dotenv').config();


const boardEnrolledTime = ["2022.02.22 01:39:312", "2022.02.22 23:10:111", "2022.02.21 11:40:000"]
const boardWriter = ["자크", "헤카림", "리신"]
const doctype = ["community", "infoshare", "notice", "suggestion"]
const company = ["company1", "company2", "company3", "company4", "company5", "company6", "company7"];
const writer = ["user1", "user2", "user3", "user4", "user5", "user6", "user7"];
const title = ["제목입니다1", "제목입니다2", "제목입니다3", "제목입니다4", "제목입니다5"];
const content = ["본문입니다1", "본문입니다2"];
const files = [
    {
        file_name: "파일이름1",
        file_hash: "2873413f1c0b757400be1e65f4edb2b1f7b354cf497a6811d4e1ff92b4e0d3f0",
        file_path: "file/file1.png"
    },
    {
        file_name: "파일이름2",
        file_hash: "9873413f1c0b757400be1e65f4edb2b1f7b354cf497a6811d4e1ff92b4e0d3f0",
        file_path: "file/file2.png"
    },
    {
        file_name: "파일이름3",
        file_hash: "6543413f1c0b757400be1e65f4edb2b1f7b354cf497a6811d4e1ff92b4e0d3f0",
        file_path: "file/file3.png"
    },
    {
        file_name: "파일이름4",
        file_hash: "1533413f1c0b757400be1e65f4edb2b1f7b354cf497a6811d4e1ff92b4e0d3f0",
        file_path: "file/file4.png"
    },
]

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

function createBoard(requestURL, board) {
    return new Promise(function (resolve, reject) {
        axios.post(requestURL, board).then(result => {
            resolve(result.data)
        });
    });
}

async function main() {
    let promises = [];

    for (let i = 0; i < 30; ++i) {
        let board = {}
        board.doctype = doctype[i%4];
        board.company = company[i%3];
        board.did = `did${i}`;
        board.writer = boardWriter[i%3];
        board.category = `category${i}`;
        board.title = title[i%5];
        board.content = content[i%2];
        board.files = files;
        board.enrolled_time = getTimeStamp()
        console.log(board);
        await sleep(300)

        const prom = createBoard(`http://${process.env.SERVER_HOST}:${process.env.PORT}/chaincode/test-api/board/create`, board);

        promises.push(prom)
    }
    Promise.all(promises).then(() => {
        logger.info("Init Complete");
    })

}

main()