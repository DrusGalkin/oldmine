@ECHO OFF
SET BINDIR=%~dp0
CD /D "%BINDIR%"
"%ProgramFiles%\Java\jre-1.8\bin\java.exe" -Xmx1024M -Xms1024M -jar mcpc-plus-191.jar
PAUSE