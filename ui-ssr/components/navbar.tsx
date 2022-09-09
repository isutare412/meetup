import { NextPage } from 'next'
import Link from 'next/link'
import { ReactElement } from 'react'
import Octocat from '@components/icon/octocat'

type NavTarget = {
  name: string
  href: string
}

const NavBar: NextPage<{ children: ReactElement }> = ({ children }) => {
  const navTargets: NavTarget[] = [
    {
      name: 'Dev',
      href: '/dev',
    },
    {
      name: 'Playground',
      href: '/plays',
    },
  ]

  const sidebarNavTargets: NavTarget[] = [
    {
      name: 'Home',
      href: '/',
    },
    ...navTargets,
  ]

  function toggleSidebar() {
    document.getElementById('sidebar-button')?.click()
  }

  return (
    <div className="drawer">
      <input id="sidebar-button" type="checkbox" className="drawer-toggle" />
      <div className="drawer-content">
        {/* <!-- Global Navigation Bar --> */}
        <div className="fixed z-10 top-0 w-full bg-base-100 border-b border-slate-300/40 dark:border-slate-300/20">
          <div className="flex justify-between max-w-5xl mx-auto">
            {/* <!-- Left Chunk --> */}
            <div className="flex">
              <div className="sm:hidden">
                <label
                  htmlFor="sidebar-button"
                  className="btn btn-square btn-ghost"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    className="inline-block w-6 h-6 stroke-current"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M4 6h16M4 12h16M4 18h16"
                    ></path>
                  </svg>
                </label>
              </div>
              <div>
                <Link href="/">
                  <a
                    role="button"
                    className="btn btn-ghost text-xl px-2 sm:px-4"
                  >
                    Meetup
                  </a>
                </Link>
              </div>
            </div>
            {/* <!-- Center Chunk --> */}
            <ul className="hidden sm:block">
              {navTargets.map(({ name, href }) => (
                <li key={name} className="inline">
                  <Link href={href}>
                    <a className="btn btn-link text-base-content">{name}</a>
                  </Link>
                </li>
              ))}
            </ul>
            {/* <!-- Right Chunk --> */}
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
        {/* <!-- Page Contents --> */}
        <div className="sm:border-x border-slate-300/40 dark:border-slate-300/20 max-w-5xl p-6 pt-16 mx-auto min-h-screen">
          <main>{children}</main>
        </div>
      </div>
      {/* <!-- Sidebar --> */}
      <div className="drawer-side">
        <label htmlFor="sidebar-button" className="drawer-overlay"></label>
        <ul className="menu p-4 overflow-y-auto w-64 bg-base-100 border-r dark:border-slate-300/20">
          {sidebarNavTargets.map(({ name, href }) => (
            <li key={name} onClick={toggleSidebar}>
              <Link href={href}>{name}</Link>
            </li>
          ))}
        </ul>
      </div>
    </div>
  )
}

export default NavBar
