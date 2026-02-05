import { useEffect, useRef, useState } from 'react'
import { profile } from '../api/auth'
import Loader from './loader'
import { deleteSkin, FORM_FILE_NAME, getSkin, getSkinPath, putSkin, saveSkin } from '../api/skins'
import skin from '../assets/images/default.png'
import { ReactSkinview3d } from 'react-skinview3d'
import { Button, Icon } from './nav'
import { ERROR, SUCCESS } from './message'

export default function Profile({ showMessage }) {
  const fileInputRef = useRef(null)
  const [file, setFile] = useState(null)
  const [loading, setLoading] = useState(true)
  const [user, setUser] = useState({
    name: '',
    created_at: '',
    id: 0,
    admin: false,
    payment: false,
    email: '',
    skin_path: ''
  })

  async function loadSkin(id) {
    try {
      setLoading(true)
      const response = await getSkinPath(id)
      setUser(prevUser => ({
        ...prevUser,
        skin_path: response?.data?.path
      }))
    } catch (error) {
      showMessage(
        "Ошибка загрузки скина, попробуйте перезагрузить лаунчер",
        ERROR
      )
    } finally {
      setLoading(false)
    }
  }

  const handleButtonClick = () => {
    fileInputRef.current.click()
  }

  const handleFileChange = async (e) => {
    const files = e.target.files
    const formData = new FormData()
    formData.append(FORM_FILE_NAME, files[0])
    setFile(formData)
  }

  const handleDeleteSkin = async (e) => {
    e.preventDefault()
    try {
      await deleteSkin(user.id)
      setUser(prevUser => ({
        ...prevUser,
        skin_path: skin
      }))

      showMessage(
        "Скин удален",
        SUCCESS
      )
    } catch (error) {
      showMessage(
        "Ошибка удаления скина!",
        ERROR
      )
    }
  }

  const uploadSkin = async () => {

    try {
      setLoading(true)
      const message = await saveSkin(file)
      showMessage(
        message,
        SUCCESS
      )
    } catch (error) {
      showMessage(
        'Ошибка обновления скина! Убедитесь что файл имеет разрешение 64x64 или 64x32',
        ERROR
      )
    }finally {
      setFile(null)
      setLoading(false)
    }
  }

  async function loadProfile() {
    try {
      setLoading(true)
      const response = await profile()
      setUser({
        name: response?.data?.name,
        created_at: response?.data?.created_at,
        id: response?.data?.id,
        admin: response?.data?.admin,
        payment: response?.data?.payment,
        email: response?.data?.email,
        skin_path: '',
      })
      await loadSkin(response?.data?.id)
    } catch (err) {
      showMessage(
        "Ошибка получения данных профиля!!!",
        ERROR
      )
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadProfile()
  }, [])

  if (loading) return (
    <div className="w-[468px] h-[236px] flex justify-center items-center">
      <Loader />
    </div>
  )
  else return (
    <div className="w-[400px] relative p-4">
      <h1 className="text-white text-[15px] gap-2 flex justify-end items-end text-outline-thin font-['Minecraft']">
        Пользователь
        <span
          className={
            user.admin
              ? 'text-[20px] text-red-500'
              : 'text-[20px] text-yellow-300'}>
            {user.name}
          </span>
      </h1>
      <div className="bg-gray-50 p-2 shadow-black">

        <div className="p-2 skin">
          <ReactSkinview3d
            skinUrl={
              user.skin_path === '' || user.skin_path === skin
                ?
                skin
                :
                getSkin(user.name)
            }
            className="z-500"
            height="350"
            width="350"
          />
        </div>

        <div className="flex w-full gap-2 mt-2">

          {
            file === null
              ?
              <Button
                children={
                  user.skin_path !== '' || user.skin_path === skin
                    ? 'Установить скин'
                    : 'Обновить скин'
                }
                onClick={handleButtonClick}
                width={400} />
              :
              <Button
                children={'Сохранить скин'}
                onClick={uploadSkin}
                width={400} />
          }


          {
            (user.skin_path !== '' & user.skin_path !== skin)
              ? (<Icon onClick={handleDeleteSkin} />)
              : null
          }
        </div>
      </div>

      <input
        type="file"
        ref={fileInputRef}
        onChange={handleFileChange}
        style={{ display: 'none' }}
        accept=".png"
      />

    </div>
  )
}
