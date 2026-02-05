import { useState } from 'react'
import { Button } from './nav'

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

    setOpen(!open)
  }

  return (
    <div className=' p-4 bg-emerald-50'>

      <p
        onClick={()=>setOpen(!open)}
        className='font-[Minecraft] text-2xl transition hover:scale-102 hover:text-outline hover:text-yellow-300 cursor-pointer'>
        {"<"}
      </p>

      <form
        className='p-2 gap-4 bg-gray-100 border border-gray-200 flex flex-col'
        onSubmit={savedMemory}>
        <div className='flex gap-4'>
          <label htmlFor="min">Мин</label>

          <input
            type="number"
            id='min'
            className='w-full border-b-gray-400 border-b-2 '
            value={memory.minMemory}
            onChange={(e) => setMemory({ ...memory, minMemory: e.target.value })}
          />
        </div>


        <div className='flex gap-4'>
          <label htmlFor="max">Макс</label>

          <input
            id='max'
            className='w-full border-b-gray-400 border-b-2 '
            type="number"
            value={memory.maxMemory}
            onChange={(e) => setMemory({ ...memory, maxMemory: e.target.value })}
          />
        </div>

        <Button children={"Сохранить"}/>
      </form>
    </div>
  )
}
