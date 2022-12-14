import type { AppProps } from 'next/app'
import '@styles/globals.css'
import Layout from '@components/layout'

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <div>
      <Layout>
        <Component {...pageProps} />
      </Layout>
    </div>
  )
}

export default MyApp
