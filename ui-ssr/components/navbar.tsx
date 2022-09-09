import { NextPage } from 'next'
import Link from 'next/link'

const NavBar: NextPage = () => {
  const navigations = [
    {
      name: 'Home',
      href: '/',
    },
    {
      name: 'Playground',
      href: '/plays',
    },
  ]

  return (
    <div className="fixed z-10 top-0 w-full bg-base-100 border-b border-slate-300/40 dark:border-slate-300/20">
      <div className="flex justify-between max-w-5xl mx-auto">
        <div>
          <a href="/" role="button" className="btn btn-ghost text-xl">
            Meetup
          </a>
        </div>
        <div>
          {navigations.map(({ name, href }) => (
            <Link href={href} key={name}>
              <a className="btn btn-link text-base-content">{name}</a>
            </Link>
          ))}
        </div>
        <div></div>
      </div>
    </div>
  )
}

export default NavBar
