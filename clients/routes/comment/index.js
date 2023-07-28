const express = require("express");
const comment = require('./commentRequest');
const router = express.Router();

router.use('/', comment.router);

module.exports.router = router;