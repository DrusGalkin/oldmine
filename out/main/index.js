"use strict";
const electron = require("electron");
const path = require("path");
const utils = require("@electron-toolkit/utils");
const fs = require("fs");
const child_process = require("child_process");
const os = require("os");
const icon = path.join(__dirname, "../../resources/icon.png");
function createWindow() {
  const mainWindow = new electron.BrowserWindow({
    width: 1100,
    height: 600,
    minWidth: 800,
    minHeight: 500,
    show: false,
    autoHideMenuBar: true,
    ...process.platform === "linux" ? { icon } : {},
    webPreferences: {
      preload: path.join(__dirname, "../preload/index.js"),
      sandbox: false,
      contextIsolation: true,
      nodeIntegration: false
    },
    icon
  });
  mainWindow.on("ready-to-show", () => {
    mainWindow.show();
    mainWindow.focus();
  });
  mainWindow.webContents.setWindowOpenHandler((details) => {
    electron.shell.openExternal(details.url);
    return { action: "deny" };
  });
  if (utils.is.dev && process.env["ELECTRON_RENDERER_URL"]) {
    mainWindow.loadURL(process.env["ELECTRON_RENDERER_URL"]);
  } else {
    mainWindow.loadFile(path.join(__dirname, "../renderer/index.html"));
  }
  return mainWindow;
}
electron.ipcMain.handle("open-path", (_, folderPath) => {
  try {
    if (fs.existsSync(folderPath)) {
      electron.shell.openPath(folderPath);
      return { success: true, path: folderPath };
    } else {
      return { success: false, error: "Path does not exist", path: folderPath };
    }
  } catch (error) {
    return { success: false, error: error.message, path: folderPath };
  }
});
const getOriginalMinecraftPath = () => {
  const platform = process.platform;
  const home = electron.app.getPath("home");
  switch (platform) {
    case "win32":
      return path.join(electron.app.getPath("appData"), ".minecraft");
    case "darwin":
      return path.join(home, "Library", "Application Support", "minecraft");
    default:
      return path.join(home, ".minecraft");
  }
};
const getAppDataPath = (appName = "oldmine") => {
  return path.join(electron.app.getPath("userData"), appName);
};
const initializeAppFolders = (appName = "oldmine") => {
  const appDataPath = getAppDataPath(appName);
  const folders = {
    base: appDataPath,
    config: path.join(appDataPath, "config"),
    cache: path.join(appDataPath, "cache"),
    logs: path.join(appDataPath, "logs"),
    instances: path.join(appDataPath, "instances"),
    downloads: path.join(appDataPath, "downloads"),
    temp: path.join(appDataPath, "temp"),
    screenshots: path.join(appDataPath, "screenshots"),
    backups: path.join(appDataPath, "backups")
  };
  Object.values(folders).forEach((folderPath) => {
    if (!fs.existsSync(folderPath)) {
      fs.mkdirSync(folderPath, { recursive: true });
    }
  });
  return folders;
};
const checkJavaInstallation = () => {
  return new Promise((resolve, reject) => {
    child_process.exec("java -version", (error, stdout, stderr) => {
      if (error) {
        resolve({ installed: false, error: error.message });
      } else {
        const versionMatch = stderr.match(/version "(.*?)"/) || stdout.match(/version "(.*?)"/);
        resolve({
          installed: true,
          version: versionMatch ? versionMatch[1] : "unknown",
          path: process.env.JAVA_HOME || ""
        });
      }
    });
  });
};
const getSystemInfo = () => {
  return {
    platform: process.platform,
    arch: process.arch,
    os: {
      type: os.type(),
      release: os.release(),
      version: os.version()
    },
    cpu: {
      cores: os.cpus().length,
      model: os.cpus()[0]?.model || "Unknown",
      speed: os.cpus()[0]?.speed || 0
    },
    memory: {
      total: os.totalmem(),
      free: os.freemem()
    },
    user: os.userInfo()
  };
};
const checkMinecraftFiles = (minecraftPath) => {
  const requiredFiles = [
    { path: path.join(minecraftPath, "bin", "minecraft.jar"), optional: false },
    { path: path.join(minecraftPath, "bin", "lwjgl.jar"), optional: false },
    { path: path.join(minecraftPath, "bin", "lwjgl_util.jar"), optional: false },
    { path: path.join(minecraftPath, "bin", "jinput.jar"), optional: false },
    { path: path.join(minecraftPath, "bin", "natives"), optional: false, isDir: true },
    { path: path.join(minecraftPath, "resources"), optional: true, isDir: true }
  ];
  const results = requiredFiles.map((file) => {
    const exists = file.isDir ? fs.existsSync(file.path) && fs.statSync(file.path).isDirectory() : fs.existsSync(file.path);
    return {
      ...file,
      exists,
      status: exists ? "found" : file.optional ? "optional_missing" : "required_missing"
    };
  });
  const missingRequired = results.filter((f) => !f.optional && !f.exists);
  const allRequiredExist = missingRequired.length === 0;
  return {
    files: results,
    allRequiredExist,
    missingRequired: missingRequired.map((f) => f.path),
    minecraftPath
  };
};
const launchMinecraft = async (options = {}) => {
  const {
    username = "Player",
    version = "1.0.0",
    maxMemory = "1024M",
    minMemory = "512M",
    windowWidth = 854,
    windowHeight = 480,
    useCustomPath = false,
    customPath = null
  } = options;
  const minecraftPath = useCustomPath && customPath ? customPath : getOriginalMinecraftPath();
  if (!fs.existsSync(minecraftPath)) {
    throw new Error(`Папка Minecraft не найдена: ${minecraftPath}`);
  }
  const javaInfo = await checkJavaInstallation();
  if (!javaInfo.installed) {
    throw new Error("Java не установлена. Пожалуйста, установите Java Runtime Environment.");
  }
  const minecraftFiles = checkMinecraftFiles(minecraftPath);
  if (!minecraftFiles.allRequiredExist) {
    console.warn("Отсутствуют некоторые обязательные файлы:", minecraftFiles.missingRequired);
    const result = await electron.dialog.showMessageBox({
      type: "warning",
      title: "Предупреждение",
      message: "Не найдены некоторые файлы Minecraft",
      detail: "Продолжить запуск? Некоторые функции могут не работать.",
      buttons: ["Продолжить", "Отмена"],
      defaultId: 0,
      cancelId: 1
    });
    if (result.response === 1) {
      throw new Error("Запуск отменен пользователем");
    }
  }
  const javaArgs = [
    `-Xmx${maxMemory}`,
    `-Xms${minMemory}`,
    `-Djava.library.path=${path.join(minecraftPath, "bin", "natives")}`,
    "-cp",
    [
      path.join(minecraftPath, "bin", "minecraft.jar"),
      path.join(minecraftPath, "bin", "lwjgl.jar"),
      path.join(minecraftPath, "bin", "lwjgl_util.jar"),
      path.join(minecraftPath, "bin", "jinput.jar")
    ].join(process.platform === "win32" ? ";" : ":"),
    "net.minecraft.client.Minecraft",
    username
  ];
  javaArgs.push(
    "--width",
    windowWidth.toString(),
    "--height",
    windowHeight.toString()
  );
  console.log("Запуск Minecraft с аргументами:", javaArgs);
  return new Promise((resolve, reject) => {
    const javaProcess = child_process.spawn("java", javaArgs, {
      cwd: minecraftPath,
      stdio: "pipe",
      detached: true,
      env: { ...process.env }
    });
    javaProcess.stdout.on("data", (data) => {
      console.log(`Minecraft stdout: ${data}`);
    });
    javaProcess.stderr.on("data", (data) => {
      console.error(`Minecraft stderr: ${data}`);
    });
    javaProcess.on("close", (code) => {
      console.log(`Minecraft завершился с кодом: ${code}`);
    });
    javaProcess.on("error", (error) => {
      console.error("Ошибка запуска Minecraft:", error);
      reject(error);
    });
    setTimeout(() => {
      javaProcess.unref();
    }, 1e3);
    resolve({
      success: true,
      pid: javaProcess.pid,
      minecraftPath,
      username,
      version,
      javaInfo,
      launchTime: (/* @__PURE__ */ new Date()).toISOString()
    });
  });
};
electron.app.whenReady().then(() => {
  utils.electronApp.setAppUserModelId("com.oldmine.launcher");
  const appFolders = initializeAppFolders("oldmine");
  console.log("Папки приложения инициализированы:", appFolders);
  electron.app.on("browser-window-created", (_, window) => {
    utils.optimizer.watchWindowShortcuts(window);
  });
  createWindow();
  electron.app.on("activate", function() {
    if (electron.BrowserWindow.getAllWindows().length === 0) {
      createWindow();
    }
  });
});
electron.app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    electron.app.quit();
  }
});
electron.ipcMain.handle("ping", () => {
  console.log("pong");
  return "pong";
});
electron.ipcMain.handle("get-minecraft-path", () => {
  const minecraftPath = getOriginalMinecraftPath();
  return {
    path: minecraftPath,
    exists: fs.existsSync(minecraftPath),
    platform: process.platform,
    files: checkMinecraftFiles(minecraftPath)
  };
});
electron.ipcMain.handle("get-app-paths", () => {
  const appFolders = initializeAppFolders("oldmine");
  return {
    ...appFolders,
    originalMinecraft: getOriginalMinecraftPath(),
    javaHome: process.env.JAVA_HOME || "",
    temp: electron.app.getPath("temp"),
    desktop: electron.app.getPath("desktop"),
    documents: electron.app.getPath("documents"),
    downloads: electron.app.getPath("downloads")
  };
});
electron.ipcMain.handle("get-system-info", () => {
  return getSystemInfo();
});
electron.ipcMain.handle("check-java", async () => {
  return await checkJavaInstallation();
});
electron.ipcMain.handle("launch-minecraft", async (_, options) => {
  try {
    return await launchMinecraft(options);
  } catch (error) {
    console.error("Ошибка запуска Minecraft:", error);
    throw error;
  }
});
electron.ipcMain.handle("open-folder", (_, folderPath) => {
  if (fs.existsSync(folderPath)) {
    electron.shell.openPath(folderPath);
    return { success: true, path: folderPath };
  } else {
    return { success: false, error: "Папка не существует" };
  }
});
electron.ipcMain.handle("select-folder", async () => {
  const result = await electron.dialog.showOpenDialog({
    properties: ["openDirectory", "createDirectory"],
    title: "Выберите папку Minecraft",
    defaultPath: getOriginalMinecraftPath()
  });
  if (!result.canceled && result.filePaths.length > 0) {
    return {
      success: true,
      path: result.filePaths[0],
      exists: fs.existsSync(result.filePaths[0])
    };
  }
  return { success: false, canceled: true };
});
electron.ipcMain.handle("save-config", (_, config) => {
  const configPath = path.join(getAppDataPath("oldmine"), "config", "settings.json");
  try {
    fs.writeFileSync(configPath, JSON.stringify(config, null, 2), "utf8");
    return { success: true, path: configPath };
  } catch (error) {
    return { success: false, error: error.message };
  }
});
electron.ipcMain.handle("load-config", () => {
  const configPath = path.join(getAppDataPath("oldmine"), "config", "settings.json");
  try {
    if (fs.existsSync(configPath)) {
      const data = fs.readFileSync(configPath, "utf8");
      return { success: true, config: JSON.parse(data) };
    } else {
      return {
        success: true,
        config: {
          username: "Player",
          memory: { max: "1024M", min: "512M" },
          window: { width: 854, height: 480 },
          launchOptions: {},
          lastPath: getOriginalMinecraftPath()
        }
      };
    }
  } catch (error) {
    return { success: false, error: error.message, config: null };
  }
});
electron.ipcMain.handle("download-file", async (_, { url, destination, filename }) => {
  return { success: false, message: "Not implemented" };
});
