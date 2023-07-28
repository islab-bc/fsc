const express = require('express');
const user = require('./user');
const admission = require('./admission');
const board = require('./board');
const comment = require('./comment');
const app = express();

app.use('/user', user.router);
app.use('/admission', admission.router);
app.use('/board', board.router);
app.use('/comment', comment.router);

module.exports.app = app;
