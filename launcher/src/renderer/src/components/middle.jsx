import Logo from './logo'
import Nav from './nav'
import { useState } from 'react'
import Settings from './settings'

export default function Middle({ setMessage }) {
  const [open, setOpen] = useState(false)
  const [options, setOptions] = useState({
    username: localStorage.getItem('name'),
    maxMemory: localStorage.getItem('maxMemory') === null ? '1024M' : localStorage.getItem('maxMemory'),
    minMemory: localStorage.getItem('minMemory') === null ? '512M' : localStorage.getItem('minMemory') ,
  })

  return (
    <div className="flex flex-col items-center">
      <Logo/>

      {
        open
          ?
            <Settings setOpen={setOpen} setOptions={setOptions} open={open} options={options}/>
          :
            <Nav setOpen={setOpen} options={options} open={open} setMessage={setMessage}/>
      }


    </div>
  )
}
