import './assets/main.css'
import Login from './components/login'
import { checkAuth } from './api/auth'
import Profile from './components/profile'
import { useCallback, useEffect, useState } from 'react'
import Middle from './components/middle'
import MessageManager from './components/message-manager'
import Loader from './components/loader'

function App() {
  const [isLoading, setIsLoading] = useState(true)
  const [auth, setAuth] = useState(false)
  const [message, setMessage] = useState({
    text: '',
    type: '',
  })

  async function checkTotalSession(){
    try{
      const bool = await checkAuth()
      setAuth(bool)
    } catch(error){

    }finally {
      setIsLoading(false)
    }
  }

  const handlerMessages = useCallback( (text, type)=>{
     setMessage({
       text: text,
       type: type,
     })
  }, [setMessage])


  useEffect(() => {
    checkTotalSession()
  },[])

  if (isLoading) return <Loader />
  else
  return (
    <div className="p-2 w-full flex flex-col gap-6">
      <div className="p-2 w-full flex justify-between ">
        {
          auth
            ?
            <Profile showMessage={handlerMessages}/>
            :
            <Login showMessage={handlerMessages}/>
        }
        <Middle/>
      </div>

      <div className='absolute right-0 top-[450px]'>
        <MessageManager message={message} setMessage={setMessage}/>
      </div>
    </div>
  )
}

export default App