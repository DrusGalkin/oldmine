import { Slide } from 'react-awesome-reveal'
import { useEffect, useState } from 'react'

export const ERROR = 'ОШИБКА'
export const SUCCESS = 'УСПЕХ'

const Message = ({ message, setMessage }) => {
  const [isVisible, setIsVisible] = useState(false)
  const { text, type } = message

  useEffect(() => {
    if (text !== '') {
      setIsVisible(true)

      setTimeout(() => {
        setIsVisible(false)
      }, 6000)

      setTimeout(()=>{
        setMessage(null)
      }, 7000)

    }
  }, [text])


  return (isVisible)
    ?
    (
      <Slide direction={'right'} triggerOnce>
        <div className="bg-white p-4">
          <div className="font-[Minecraft]">
            {
              type === ERROR
                ?
                <p className="text-red-600">Ахтунг!</p>
                :
                <p className="text-green-700">Сообщение</p>
            }
          </div>

          <hr />

          {text}
        </div>
      </Slide>
    )
    :
    <Slide direction={'right'} reverse={true}>
      <div className="bg-white p-4">
        <div className="font-[Minecraft]">
          {
            type === ERROR
              ?
              <p className="text-red-600">Ахтунг!</p>
              :
              <p className="text-green-700">Сообщение</p>
          }
        </div>

        <hr />

        {text}
      </div>
    </Slide>

}

export default Message