import { useEffect, useState } from 'react'
import { ERROR } from './message'

export default function Nav({ setOpen, open, options, setMessage}) {
  const [minecraftPath, setMinecraftPath] = useState('')


  useEffect(() => {
    if (window.electronAPI) {
      window.electronAPI.getMinecraftPath()
        .then(result => {
          if (result.exists) {
            setMinecraftPath(result.path)
          } else {
            setMessage({
              text: 'Путь до Minecraft не найден! Проверите путь',
              type: ERROR
            })
          }
        })
        .catch(console.error)
    }
  }, [])

  const openPholderClient = async () => {
    await window.electronAPI.openFolder(minecraftPath)
  }

  const handlePlayClick = async () => {
    if (window.electronAPI && minecraftPath) {
      await window.electronAPI.launchMinecraft(options)
    } else {
      setMessage({
        text: 'Ошибка клиента! Обратитесь в поддержку или обновите лаунчер',
        type: ERROR
      })
    }
  }

  const handleSettingsClick = () => {
    setOpen(!open)
  }

  const handleExitClick = () => {
    window.close()
  }

  return (
    <div className="p-3 flex flex-col items-center gap-3">
      <Button onClick={handlePlayClick}>Играть</Button>
      <Button onClick={handleSettingsClick}>Настройки</Button>
      <Button onClick={openPholderClient}>Папка с клиентом</Button>
      <Button onClick={handleExitClick}>Выйти</Button>
    </div>
  )
}

export const Button = ({ children, onClick, width = 400 }) => {
  return (
    <button
      onClick={onClick}
      className={`w-[${width}px] transition-all hover:scale-102 cursor-pointer hover:text-yellow-300 text-center p-2 shadow-black border-2 border-gray-800 text-white font-[Minecraft] bg-gray-500`}>
      {children}
    </button>
  )
}

export const Icon = ({ width, onClick }) => {
  return (
    <button
      onClick={onClick}
      className={`w-[${width}px] transition-all hover:scale-102 cursor-pointer hover:text-red-500 text-center p-2 shadow-black border-2 border-gray-800 text-white font-[Minecraft] bg-gray-500`}>

      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M3 6h18"/>
        <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6"/>
        <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"/>
        <line x1="10" y1="11" x2="10" y2="17"/>
        <line x1="14" y1="11" x2="14" y2="17"/>
      </svg>
    </button>
  )
}
