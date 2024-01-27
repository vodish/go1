; инструкция
; https://webtous.ru/poleznye-sovety/kak-pomenyat-raskladku-klaviatury-nazhatiem-vsego-odnoj-klavishi.html
; скрипт под прогу версии 1
; для автозагрузки положить файл в папку
; C:\Пользователи\ИмяПользователя\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup

SendMode Input
SetWorkingDir %A_ScriptDir%

CapsLock::Send, {Alt Down}{Shift Down}{Shift Up}{Alt Up}