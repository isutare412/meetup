import type { NextPage } from 'next'
import Head from 'next/head'
import { useState } from 'react'

const Header: NextPage<{ title: string }> = ({ title }) => (
  <h1 className="text-2xl">{title ? title : 'Default title'}</h1>
)

const Home: NextPage = () => {
  const names = ['Ada Lovelace', 'Grace Hopper', 'Margaret Hamilton']
  const [likes, setLikes] = useState(0)

  function handleClick() {
    setLikes(likes + 1)
  }

  return (
    <>
      <Head>
        <title>Meetup Home</title>
      </Head>
      <Header title="Develop. Preview. Ship. 🚀" />
      <ul>
        {names.map((name) => (
          <li key={name}>{name}</li>
        ))}
      </ul>

      <button onClick={handleClick} className="btn btn-primary">
        Like ({likes})
      </button>
    </>
  )
}

export default Home
