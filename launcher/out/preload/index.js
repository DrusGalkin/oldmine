"use strict";
const electron = require("electron");
const preload = require("@electron-toolkit/preload");
const api = {};
if (process.contextIsolated) {
  try {
    electron.contextBridge.exposeInMainWorld("electron", preload.electronAPI);
    electron.contextBridge.exposeInMainWorld("api", api);
    electron.contextBridge.exposeInMainWorld("electronAPI", {
      getAppPaths: () => electron.ipcRenderer.invoke("get-app-paths"),
      getPath: (pathName) => electron.ipcRenderer.invoke("get-path", pathName),
      getMinecraftPath: () => electron.ipcRenderer.invoke("get-minecraft-path"),
      openFolder: (path) => electron.ipcRenderer.invoke("open-folder", path),
      launchMinecraft: (username) => electron.ipcRenderer.invoke("launch-minecraft", username),
      launchMinecraftDirect: (username) => electron.ipcRenderer.invoke("launch-minecraft-direct", username),
      ping: () => electron.ipcRenderer.send("ping")
    });
  } catch (error) {
    console.error(error);
  }
} else {
  window.electron = preload.electronAPI;
  window.api = api;
  window.electronAPI = {
    getAppPaths: () => electron.ipcRenderer.invoke("get-app-paths"),
    getPath: (pathName) => electron.ipcRenderer.invoke("get-path", pathName),
    getMinecraftPath: () => electron.ipcRenderer.invoke("get-minecraft-path"),
    launchMinecraft: (username) => electron.ipcRenderer.invoke("launch-minecraft", username),
    launchMinecraftDirect: (username) => electron.ipcRenderer.invoke("launch-minecraft-direct", username),
    openPath: (path) => electron.ipcRenderer.invoke("open-path", path),
    ping: () => electron.ipcRenderer.send("ping")
  };
}
