const { Client } = require('pg');
const {logger} = require('../logger.js')

const client = new Client({
  user: 'client',
  host: '127.0.0.1',
  database: 'clientdev',
  password: 'client',
  port: 5433,
});

client.connect();

async function handleLogin(event, { username, password }) {
  logger.info(`Login Attempt Username: ${username} Password: ${password}`);
  try {
    const res = await client.query(
      'SELECT * FROM users WHERE username = $1 AND password = $2',
      [username, password]
    );
    return res.rows.length > 0
      ? { success: true }
      : { success: false, error: 'Invalid credentials' };
  } catch (err) {
    return { success: false, error: err.message };
  }
}

module.exports = { handleLogin };
