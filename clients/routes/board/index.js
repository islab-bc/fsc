const express = require("express");
const notice = require('./noticeRequest');
const community = require('./communityRequest');
const infoshare = require('./infoshareRequest');
const suggestion = require('./suggestionRequest');
const router = express.Router();

router.use('/notice', notice.router);
router.use('/community', community.router);
router.use('/infoshare', infoshare.router);
router.use('/suggestion', suggestion.router);

module.exports.router = router;