const { app, BrowserWindow, session, ipcMain } = require('electron');
const path = require('path');

function createWindow() {
    const mainWindow = new BrowserWindow({
        width: 1100,
        height: 600,
        webPreferences: {
            // Безопасная конфигурация для Electron v40+
            nodeIntegration: false, // отключаем в рендерере
            contextIsolation: true, // включаем изоляцию контекста
            enableRemoteModule: false,
            preload: path.join(__dirname, 'preload.js') // добавляем preload скрипт
        }
    });

    mainWindow.removeMenu();
    mainWindow.loadFile('./src/renderer/index.html');
}

// Обработчики IPC для работы с куками
ipcMain.handle('get-cookie', async (event, { name, url = 'http://localhost' }) => {
    try {
        const cookies = await session.defaultSession.cookies.get({ name, url });
        return cookies.length > 0 ? cookies[0] : null;
    } catch (error) {
        console.error('Error getting cookie:', error);
        return null;
    }
});

ipcMain.handle('set-cookie', async (event, cookieData) => {
    try {
        await session.defaultSession.cookies.set(cookieData);
        return { success: true };
    } catch (error) {
        console.error('Error setting cookie:', error);
        return { success: false, error: error.message };
    }
});

ipcMain.handle('remove-cookie', async (event, { name, url = 'http://localhost' }) => {
    try {
        await session.defaultSession.cookies.remove(url, name);
        return { success: true };
    } catch (error) {
        console.error('Error removing cookie:', error);
        return { success: false, error: error.message };
    }
});

// IPC для получения всех кук
ipcMain.handle('get-all-cookies', async (event, { url = 'http://localhost' }) => {
    try {
        return await session.defaultSession.cookies.get({ url });
    } catch (error) {
        console.error('Error getting all cookies:', error);
        return [];
    }
});

app.whenReady().then(() => {
    createWindow();

    app.on('activate', () => {
        if (BrowserWindow.getAllWindows().length === 0) {
            createWindow();
        }
    });
});

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit();
    }
});