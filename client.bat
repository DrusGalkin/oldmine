@echo off
cd /d "C:\Users\AndrewGalkin\AppData\Roaming\.minecraft\"
java -Xmx1024M -Xms512M ^
-Djava.library.path="bin/natives" ^
-cp "bin/minecraft.jar;bin/lwjgl.jar;bin/lwjgl_util.jar;bin/jinput.jar" ^
net.minecraft.client.Minecraft ^
aaa
pause