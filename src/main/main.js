const { app, BrowserWindow } = require('electron');
const path = require('path');



function createWindow() {
    const mainWindow = new BrowserWindow({
        width: 1100,
        height: 600,
        webPreferences: {
            nodeIntegration: true,
            contextIsolation: false,
        }
    });

    mainWindow.removeMenu();

    mainWindow.loadFile('./src/renderer/index.html');
}

app.whenReady().then(() => {
    createWindow();

    app.on('activate', function () {
        if (BrowserWindow.getAllWindows().length === 0) createWindow();
    });
});

app.on('window-all-closed', function () {
    if (process.platform !== 'darwin') app.quit();
});