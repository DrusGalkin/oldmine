import './assets/main.css'
import Auth from './components/auth'
import Logo from './components/logo'

function App() {
  return (
    <div className="p-2 w-full flex justify-between ">
      <Auth />
      <Logo />
    </div>
  )
}

export default App
