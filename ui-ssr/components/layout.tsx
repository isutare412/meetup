import { NextPage } from 'next'
import Head from 'next/head'
import { ReactElement } from 'react'
import NavBar from '@components/navbar'

const Layout: NextPage<{ children: ReactElement }> = ({ children }) => {
  return (
    <>
      <Head>
        <title>Meetup</title>
      </Head>
      <NavBar />
      <div className="border-x border-slate-300/40 dark:border-slate-300/20 max-w-5xl p-4 mx-auto mt-12">
        <main>{children}</main>
      </div>
    </>
  )
}

export default Layout
