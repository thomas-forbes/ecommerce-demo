import Image from 'next/image'
import Link from 'next/link'

export default function Home() {
  return (
    <div className="flex flex-col w-full items-center *:w-full *:max-w-screen-xl p-10">
      <div className="grid grid-cols-2 gap-10">
        <ProductCard />
        <ProductCard />
        <ProductCard />
        <ProductCard />
      </div>
    </div>
  )
}

function ProductCard() {
  return (
    <Link href="/">
      <div className="border border-zinc-900 rounded-md hover:border-zinc-400 hover:border-dashed duration-300 flex flex-col gap-4 group relative">
        <Image
          src="/shirt.avif"
          alt="shirt"
          width={600}
          height={600}
          className="scale-90 group-hover:scale-100 duration-300"
        />
        <div className="absolute bottom-5 left-5 font-mono group-hover:bottom-8 group-hover:left-8 duration-300 flex gap-2 group-hover:gap-3 items-center">
          <span className="text-zinc-300 group-hover:text-zinc-100 duration-200">
            shirt
          </span>
          <span className="text-zinc-700">–</span>
          <span className="text-zinc-500 group-hover:text-zinc-400 duration-200">
            $100
          </span>
        </div>
      </div>
    </Link>
  )
}
