import {useEffect, useRef, useState} from 'react'
import {profile} from '../api/auth'
import Loader from './loader'
import {deleteSkin, SKIN_FORM_FILE_NAME, getSkin, getSkinPath, putSkin, saveSkin} from '../api/skins'
import skin from '../assets/images/default.png'
import {ReactSkinview3d} from 'react-skinview3d'
import {Button, Icon, Setting} from './nav'
import {ERROR, SUCCESS} from './message'
import {CLOAK_FORM_FILE_NAME, deleteCloak, getCloak, saveCloak} from "../api/cloaks";

export default function Profile({showMessage}) {
    const fileInputRef = useRef(null)
    const [file, setFile] = useState(null)
    const [loading, setLoading] = useState(true)
    const [isCloak, setCloak] = useState(false)
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
        const fileName  = isCloak ? CLOAK_FORM_FILE_NAME : SKIN_FORM_FILE_NAME
        console.log(fileName)
        const formData = new FormData()
        formData.append(fileName, files[0])
        setFile(formData)
    }

    const handleDeleteCloak = async (e) => {
        e.preventDefault()
        try {
            await deleteCloak(user.id)
            showMessage(
                "Плащ удален",
                SUCCESS
            )
        } catch (error) {
            showMessage(
                "Ошибка удаления плаща!",
                ERROR
            )
        } finally {
            setCloak(true)
        }
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
        } finally {
            setCloak(true)
        }
    }

    const handleDelete = (e) => {
        if (isCloak) return handleDeleteCloak(e)
        return handleDeleteSkin(e)
    }

    const uploadSkin = async () => {
        try {
            setLoading(true)
            const message = await saveSkin(file)
            showMessage(
                "Скин установлен! Он появится через 5-10 минут, или перезагрузите лаунчер",
                SUCCESS
            )
        } catch (error) {
            showMessage(
                'Ошибка обновления скина! Убедитесь что файл имеет разрешение 64x64 или 64x32',
                ERROR
            )
        } finally {
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

    const uploadFileManager = () => {
        if (!isCloak) {
            if (file === null) {
                return <Button
                    children={
                        user.skin_path !== '' || user.skin_path === skin
                            ? 'Установить скин'
                            : 'Обновить скин'
                    }
                    onClick={handleButtonClick}
                    width={400}/>
            } else {
                return <Button
                    children={'Сохранить скин'}
                    onClick={uploadSkin}
                    width={400}/>
            }
        } else {
            if (file === null) {
                return <Button
                    children={
                        user.skin_path !== '' || user.skin_path === skin
                            ? 'Установить плащ'
                            : 'Обновить плащ'
                    }
                    onClick={handleButtonClick}
                    width={400}/>
            } else {
                return <Button
                    children={'Сохранить плащ'}
                    onClick={uploadCloak}
                    width={400}/>
            }
        }
    }

    const uploadCloak = async () => {
        if (!file) {
            showMessage("Файл не выбран!", ERROR)
            return
        }

        try {
            setLoading(true)
            const message = await saveCloak(file)
            showMessage(
                message,
                SUCCESS
            )
        } catch (error) {
            console.log(error)
            showMessage(
                'Ошибка обновления плаща! Убедитесь что файл имеет разрешение 64x32 или 22x17',
                ERROR
            )
        } finally {
            setFile(null)
            setLoading(false)
            setCloak(false)
        }
    }


    const handlerChooseTypeFile = (e) => {
        setCloak(!isCloak)
        if (file !== null) {
            setFile(null)
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
            <h1 className="text-white text-[15px] gap-2 flex justify-end items-end text-outline-thin font-['Minecraft']">

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
                        capeUrl={getCloak(user.name)}
                        className="z-500"
                        height="350"
                        width="350"
                    />
                </div>

                <div className="flex w-full gap-2 mt-2">


                    {
                        uploadFileManager()
                    }


                    {
                        (user?.payment || user?.admin)
                            ? (<Setting onClick={handlerChooseTypeFile}/>)
                            : null
                    }


                    {
                        (user.skin_path !== '' & user.skin_path !== skin)
                            ? (<Icon onClick={handleDelete}/>)
                            : null
                    }
                </div>
            </div>

            <input
                type="file"
                ref={fileInputRef}
                onChange={handleFileChange}
                style={{display: 'none'}}
                accept=".png"
            />

        </div>
    )
}
