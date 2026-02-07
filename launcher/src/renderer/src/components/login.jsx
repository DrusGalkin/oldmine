import { useState } from 'react'
import { login } from '../api/auth'
import Loader from './loader'
import Profile from './profile'
import { ERROR, SUCCESS } from './message'

export default function Login({ showMessage }) {
  const [data, setData] = useState({ email: 'andrew.galkin2018@gmail.com', password: 'andrew.galkin2018@gmail.com' })
  const [loading, setLoading] = useState(false)
  const [auth, setAuth] = useState("")
  const inputStyle =
    'p-3 cursor-pointer hover:scale-101 transition-all rounded border-b-2 rounded-b-none focus:border-b-blue-600 border-b-gray-400 focus:outline-none focus:ring-0 focus:ring-offset-0 focus:border-transparent focus:text-blue-600 hover:outline-none active:outline-none'

  const loginHandler = async (e) => {
    e.preventDefault()
    try{
      setLoading(true)
      let bool = await login(data)
      setAuth(bool)

      if (bool && showMessage) showMessage(
        'Привет ' + localStorage.getItem('name') + "!",
        SUCCESS
      )

    } catch (error) {
      console.error(error)
      showMessage(
        'Ошибка авторизации! Неверный логин или пароль',
        ERROR
      )
    }finally {
      setLoading(false)
    }
  }

  if (auth === true) return <Profile showMessage={showMessage} />

  if(loading){
    return (
      <div className="w-[468px] h-[236px] flex justify-center items-center">
        <Loader/>
      </div>
    )
  }
   return (
    <div className="w-[500px] relative p-4">
      <h1 className="text-yellow-300 text-outline-thin font-['Minecraft']">Вход в аккаунт</h1>
      <form
        onSubmit={loginHandler}
        className="flex bg-gray-50 p-6 flex-col gap-6 shadow-lg shadow-black"
      >

        <input
          id="email"
          value={data.email}
          onChange={(e) => setData({ ...data, email: e.target.value })}
          className={inputStyle}
          type="email"
          placeholder="Почта" />


        <input
          id="password"
          value={data.password}
          onChange={(e) => setData({ ...data, password: e.target.value })}
          className={inputStyle}
          type="password"
          placeholder="Пароль" />


        <div className="flex justify-between items-center">
          <button
            className="p-2 border-[1px] cursor-pointer transition-all hover:scale-105 text-white rounded bg-gray-400 border-gray-500 w-[150px]"
            type="submit">
            Войти
          </button>
          <p onClick={()=> window.location.href = ""} className="text-gray-400 transition-all hover:scale-105 hover:text-blue-600 cursor-pointer">
            Регистрация
          </p>
        </div>
      </form>
    </div>
  )
}
