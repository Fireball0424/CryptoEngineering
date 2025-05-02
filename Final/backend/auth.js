const { Client } = require('pg');
const {logger} = require('../logger.js')

require('dotenv').config();
const apiBaseUrl = process.env.API_BASE_URL;
const axios = require('axios');

const bcrypt = require('bcryptjs');
const { session } = require('electron');

// Handle Login 
// TODO: Switch to API instead of connect by client 
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

async function handleRegister(event, { username, password }) {
  logger.info(`Register Attempt Username: ${username} Password: ${password}`);
  
  const saltRounds = process.env.SALT_ROUNDS;

  try {
    const password_hash = await bcrypt.hash(password, saltRounds);
    try {
      const response = await axios.post(`${apiBaseUrl}RegisterCreate`, { username, password_hash});

      if (response.data.status === 'Username exists') {
        return { success: false, error: 'Username exists' };
      } 
      else if (response.data.status === 'Username is not a valid email') {
        return { success: false, error: 'Username is not a valid email' };
      } 
      else if (response.data.status === 'Success! Send verify OTP') {  
        // TODO: Is there any better way to send apiBaseUrl to front-end?
        sessionStorage.setItem('otpUsername', username);
        sessionStorage.setItem('apiBaseUrl', apiBaseUrl);

        window.location.href = '../templates/otp.html';
        return {success: true, message: 'OPT already sent to your email. Please verify and enter it.'};
      } 
      else {
        return { success: false, error: 'Unexpected response from server' };
      }
    } catch (apiError) {
      logger.error('API Request failed:', apiError.message);
      return { success: false, error: 'API request failed' };
    }

  } catch (err) {
    return { success: false, error: err.message };
  }
}


module.exports = { handleLogin, handleRegister};
