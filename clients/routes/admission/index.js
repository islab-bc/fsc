const express = require("express");
const admission = require('./admissionRequest');
const router = express.Router();

router.use('/', admission.router);

module.exports.router = router;