import { useEffect, useState } from 'react'

export default function Nav({setOpen, open, options}) {
  const [minecraftPath, setMinecraftPath] = useState('')


  useEffect(() => {
    if (window.electronAPI) {
      window.electronAPI.getMinecraftPath()
        .then(result => {
          if (result.exists) {
            setMinecraftPath(result.path)
          } else {
            console.log('Minecraft not found at:', result.path)
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
      console.log('Minecraft path not found or electronAPI not available')
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

const Button = ({ children, onClick }) => {
  return (
    <button
      onClick={onClick}
      className="w-[400px] transition-all hover:scale-102 cursor-pointer hover:text-yellow-300 text-center p-2 shadow-black border-2 border-gray-800 text-white font-[Minecraft] bg-gray-500">
      {children}
    </button>
  )
}
