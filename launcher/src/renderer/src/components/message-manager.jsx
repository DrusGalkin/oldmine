import { useEffect, useState } from 'react'
import Message from './message'

const MessageManager = ({ message, setMessage }) => {
  const [content, setContent] = useState([])
  const [displayMessage, setDisplayMessage] = useState(null)

  useEffect(() => {
    if (message?.text !== ""){
      setContent(prev => [...prev, message])
      setMessage({
        text: '',
        type: ''
      })
    }
  }, [message, setMessage])


  useEffect(() => {
    if (!displayMessage && content.length > 0) {
      const [first, ...rest] = content
      setContent(rest)
      setDisplayMessage(first)

      const timer = setTimeout(() => {
        setDisplayMessage(null)
      }, 7000)
      return () => clearTimeout(timer)
    }
  }, [content, displayMessage])

  return (
    <div>
      {
        displayMessage && (
          <Message message={displayMessage} setMessage={setDisplayMessage} />
        )
      }
    </div>
  )
}

export default MessageManager