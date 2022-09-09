import { NextPage } from 'next'
import Head from 'next/head'
import { ReactElement } from 'react'
import NavBar from '@components/navbar'

const Layout: NextPage<{ children: ReactElement }> = ({ children }) => (
  <>
    <Head>
      <title>Meetup</title>
    </Head>
    <NavBar>{children}</NavBar>
  </>
)

export default Layout
