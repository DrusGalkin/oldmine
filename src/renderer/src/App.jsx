import './assets/main.css'
import Login from './components/login'
import Logo from './components/logo'
import { checkAuth } from './api/auth'
import Profile from './components/profile'
import { useEffect, useState } from 'react'
import Middle from './components/middle'

function App() {
  const [auth, setAuth] = useState(false)

  async function checkTotalSession(){
    const bool = await checkAuth()
    setAuth(bool)
  }

  useEffect(() => {
    checkTotalSession()
  },[])

  return (
    <div className="p-2 w-full flex flex-col gap-6">
      <div className="p-2 w-full flex justify-between ">


        {
          auth
          ?
            <Profile/>
          :
            <Login />
        }

        <Middle/>
      </div>

    </div>

  )
}

export default App
