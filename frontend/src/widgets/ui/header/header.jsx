import {Nav} from "../../../shared/tools/nav.js";

export default function Header(){
    return (
        <header className="bg-red-600">
            <img
                className='h-[400px] z-0 absolute w-full object-cover border-yellow-300 border-b-[7px]'
                src="https://ru-minecraft.ru/uploads/posts/2020-01/1577873178_kpm2yyq.jpg"
                alt=""/>

            <div className='z-10 relative text-minecraft'>

                <section className='flex justify-around items-end h-auto'>
                    <h1 className='text-[120px] text-white text-outline'>
                        OldMine
                    </h1>

                    <nav className='flex text-white text-outline gap-2'>
                        {
                            Nav.map((item, index) => (
                                <p key={index}>
                                    {item.title}
                                </p>
                            ))
                        }
                    </nav>
                </section>

            </div>
        </header>
    )
}