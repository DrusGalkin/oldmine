import { useEffect, useState } from 'react'
import { profile } from '../api/auth'
import Loader from './loader'
import { getSkin } from '../api/skins'

export default function Profile() {
  const [user, setUser] = useState({})
  const [loading, setLoading] = useState(true);

  async function loadSkin(id) {
    try {
      setLoading(true);
      const response = await getSkin(id)
      setUser({ ...response, skin_path: response?.data?.path })
    } catch (error) {
      console.log(error)
    }finally {
      setLoading(false);
    }
  }

  async function loadProfile() {
    try{
      setLoading(true)
      const response = await profile()
      setUser(response?.data)
      // console.log(response?.data)
      // await loadSkin(response?.data?.id)
    }catch(err){
      console.log(err)
    }finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadProfile()
  }, [])

  if (loading) return (
    <div className="w-[468px] h-[236px] flex justify-center items-center">
      <Loader/>
    </div>
  )
  else return (
      <div className="w-[400px] relative p-4">
        <h1 className="text-white text-center text-outline-thin font-['Minecraft']">
          Профиль пользователя <span className="text-[20px] p-2 text-yellow-300">{user.name}</span>
        </h1>
        <div className="bg-gray-50 p-2 h-[450px] shadow-black">

        </div>
      </div>
    )



}
