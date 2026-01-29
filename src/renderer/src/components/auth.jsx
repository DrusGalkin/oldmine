import { useState } from 'react'
import { login } from '../api/auth'

export default function Auth() {
  const [data, setData] = useState({ email: '', password: '' })
  const [loading, setLoading] = useState(false)
  const [str, setStr] = useState('')
  const inputStyle =
    'p-3 cursor-pointer hover:scale-101 transition-all rounded border-b-2 rounded-b-none focus:border-b-blue-600 border-b-gray-400 focus:outline-none focus:ring-0 focus:ring-offset-0 focus:border-transparent focus:text-blue-600 hover:outline-none active:outline-none'

  const loginHandler = async (e) => {
    e.preventDefault()
    try{
      setLoading(true)
      console.log(data)
      let str = await login(data)
      setStr(str)
    } catch (error) {
      console.error(error)
    }finally {
      setLoading(false)
    }

  }

  return (
    <div className="w-[500px] relative p-4">
      <h1 className="text-yellow-300 text-outline-thin font-['Minecraft']">Вход в аккаунт</h1>

      <p>
        {str}
      </p>
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
          <p className="text-gray-400 transition-all hover:scale-105 hover:text-blue-600 cursor-pointer">
            Регистрация
          </p>
        </div>
      </form>
    </div>
  )
}
