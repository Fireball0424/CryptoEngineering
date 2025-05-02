const { app, BrowserWindow, ipcMain } = require('electron');
const path = require('path');
const { handleLogin } = require('./backend/auth.js');

function createWindow() {
  const mainWindow = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      preload: path.join(__dirname, 'preload.js'), // 用來 expose ipcRenderer
      contextIsolation: true, // 隔離上下文（安全）
      nodeIntegration: false  // 不允許直接在 HTML 中 require Node.js 模組
    }
  });

  mainWindow.loadFile('templates/login.html');
}

app.whenReady().then(() => {
  createWindow();
});

app.on('window-all-closed', function() {
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

// Handle login event 
ipcMain.handle('login', handleLogin);

