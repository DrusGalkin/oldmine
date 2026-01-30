
export default function Nav() {
  return (
    <div className="p-3 flex flex-col items-center gap-3 shadow-2xl ">
      <Button children="Играть"/>
      <Button children="Настройки"/>
      <Button children="Выйти"/>
    </div>
  )
}

const Button = ({children, onClick}) => {
  return (
    <button
      onClick={onClick}
      className="w-[400px] transition-all hover:scale-102 cursor-pointer hover:text-yellow-300 text-center p-2 shadow-black border-2 border-gray-800 text-white font-[Minecraft] bg-gray-500">
      {children}
    </button>
  )
}
