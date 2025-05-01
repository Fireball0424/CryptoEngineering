const { Client } = require('pg');

const client = new Client({
  user: 'client',
  host: '127.0.0.1',
  database: 'clientdev',
  password: 'client',
  port: 5433
});

client.connect()
  .then(() => {
    console.log('Connected');
    return client.query('SELECT NOW()');
  })
  .then(res => {
    console.log(res.rows);
    return client.end();
  });