const { app, BrowserWindow, ipcMain } = require('electron');
const path = require('path');
const {handleLogin} = require('./backend/auth.js');

app.whenReady().then(() => {
  const mainWindow = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      nodeIntegration: true
    }
  });

  mainWindow.loadFile('templates/login.html');
});

// Handle login event 
ipcMain.handle('login', handleLogin);