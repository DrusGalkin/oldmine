import { useState } from 'react'

export default function Settings({setOpen, open, setOptions, options}) {
  const [memory, setMemory] = useState({
    maxMemory: options.maxMemory.replace(/[a-zA-Z]/g, ''),
    minMemory: options.minMemory.replace(/[a-zA-Z]/g, ''),
  })

  const savedMemory = (e) => {
    e.preventDefault()
    setOptions(
      {
        ...options,
        maxMemory: memory.maxMemory + "M",
        minMemory: memory.minMemory + "M",
      })
    localStorage.setItem('maxMemory', memory.maxMemory + "M")
    localStorage.setItem('minMemory', memory.minMemory + "M")
  }

  return (
    <div className='w-[400px] p-4 bg-emerald-50'>

      <p
        onClick={()=>setOpen(!open)}
        className='font-[Minecraft] text-2xl transition hover:scale-102 hover:text-outline hover:text-yellow-300 cursor-pointer'>
        {"<"}
      </p>

      <form onSubmit={savedMemory}>
        <input
          type="number"
          value={memory.minMemory}
          onChange={(e) => setMemory({ ...memory, minMemory: e.target.value })}
        />
        <input
          type="number"
          value={memory.maxMemory}
          onChange={(e) => setMemory({ ...memory, maxMemory: e.target.value })}
        />
        <button>
          Сохранить
        </button>
      </form>
    </div>
  )
}
