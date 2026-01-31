import { contextBridge, ipcRenderer } from 'electron'
import { electronAPI } from '@electron-toolkit/preload'

const api = {}

if (process.contextIsolated) {
  try {
    contextBridge.exposeInMainWorld('electron', electronAPI)
    contextBridge.exposeInMainWorld('api', api)
    contextBridge.exposeInMainWorld('electronAPI', {
      getAppPaths: () => ipcRenderer.invoke('get-app-paths'),
      getPath: (pathName) => ipcRenderer.invoke('get-path', pathName),
      getMinecraftPath: () => ipcRenderer.invoke('get-minecraft-path'),
      openFolder: (path) => ipcRenderer.invoke('open-folder', path),

      launchMinecraft: (username) => ipcRenderer.invoke('launch-minecraft', username),
      launchMinecraftDirect: (username) => ipcRenderer.invoke('launch-minecraft-direct', username),

      ping: () => ipcRenderer.send('ping')
    })
  } catch (error) {
    console.error(error)
  }
} else {
  window.electron = electronAPI
  window.api = api
  window.electronAPI = {
    getAppPaths: () => ipcRenderer.invoke('get-app-paths'),
    getPath: (pathName) => ipcRenderer.invoke('get-path', pathName),
    getMinecraftPath: () => ipcRenderer.invoke('get-minecraft-path'),
    launchMinecraft: (username) => ipcRenderer.invoke('launch-minecraft', username),
    launchMinecraftDirect: (username) => ipcRenderer.invoke('launch-minecraft-direct', username),
    openPath: (path) => ipcRenderer.invoke('open-path', path),
    ping: () => ipcRenderer.send('ping')
  }
}
