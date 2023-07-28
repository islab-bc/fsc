const express = require("express");
const user = require('./userRequest');
const router = express.Router();

router.use('/', user.router);

module.exports.router = router;