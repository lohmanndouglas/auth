import Header from '../components/Header'
import styles from '../styles/Home.module.css'
import { useSession } from 'next-auth/client'

export default function Home() {

  const [session, loading] = useSession();

  if (session) {
    console.log('User:', session.user, 'Token:',session.accessToken)
  }

  return (
    <div className={styles.container}>
      <Header>
        <title>POC authentication Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Header>

      <main className={styles.main}>
        <p className={styles.description}>
          Home page
        </p>
        <div className={styles.user}>
          {loading && <div className={styles.title}>Loading...</div>}
          {session && <> <p style={{ marginBottom: '10px' }}> Welcome, {session.user.name ?? session.user.email}</p> <br />
            <img src={session.user.image} alt="" className={styles.avatar} />
          </>}
          {!session &&
            <>
              <p className={styles.title}>Please Sign in</p>
              <img src="http://s2.glbimg.com/-p1otaWHnNYONNJu4xkKVDfUdtY=/123x60:976x679/695x504/s.glbimg.com/po/tt2/f/original/2016/02/18/giphy.gif" alt=""  width="400" height="450" className={styles.avatar} />          
            </>
          }
        </div>
      </main>
    </div>
  )
}
