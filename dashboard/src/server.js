const express = require('express');
const path = require('path');

const app = express();
const cors = require('cors');

app.use(cors()); // Used to cheat 'Same Origem Policy' from browsers

app.use('/talks', (req, res) => {
  res.sendFile(path.join(__dirname, './talks.html'));
});

app.use('/', (req, res) => {
  res.sendFile(path.join(__dirname, './index.html'));
});

app.listen(8080, () => {
  console.log(`Dashboard up and running at localhost:8080 :)`);
});
