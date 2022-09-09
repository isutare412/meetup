import { NextPage } from 'next'
import Link from 'next/link'
import Octocat from '@components/icon/octocat'

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
          <Link href="/">
            <a role="button" className="btn btn-ghost text-xl">
              Meetup
            </a>
          </Link>
        </div>
        <div>
          {navigations.map(({ name, href }) => (
            <Link href={href} key={name}>
              <a className="btn btn-link text-base-content">{name}</a>
            </Link>
          ))}
        </div>
        <div className="flex content-center">
          <a
            href="https://github.com/isutare412"
            role="button"
            className="btn btn-link btn-sm px-0 my-auto mr-2 text-base-content hover:text-gray-500"
          >
            <Octocat width={30} height={30} className="fill-current" />
          </a>
        </div>
      </div>
    </div>
  )
}

export default NavBar
