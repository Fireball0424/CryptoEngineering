const tunnel = require('tunnel-ssh');
const { Client } = require('pg');

const tunnelConfig = {
  username: 'ubuntu',
  host: '13.55.95.185',
  port: 22,
  privateKey: require('fs').readFileSync('SSH_KMS.pem'), // 或使用 password
  dstHost: '127.0.0.1', // PostgreSQL server 在遠端的位址（通常是 localhost）
  dstPort: 5432,        // PostgreSQL port
  localHost: '127.0.0.1',
  localPort: 5433       // 在本機開的轉接 port（你接下來會連這個）
};

const pgConfig = {
  user: 'client',
  host: '127.0.0.1',
  database: 'clientdev',
  password: 'client',
  port: 5433, // 要對應 localPort
};

tunnel(tunnelConfig, (error, server) => {
  if (error) {
    console.error('SSH tunnel error:', error);
    return;
  }

  const client = new Client(pgConfig);

  client.connect(err => {
    if (err) {
      console.error('PG connection error:', err.stack);
    } else {
      console.log('Connected to PostgreSQL via SSH tunnel');
      client.query('SELECT NOW()', (err, res) => {
        console.log(err ? err.stack : res.rows);
        client.end();
        server.close();
      });
    }
  });
});
