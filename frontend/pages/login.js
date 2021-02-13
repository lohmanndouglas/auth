
import styles from '../styles/Home.module.css'
import axios from 'axios';

function handleLogin(provider) {
    console.log("You chose the login: "+provider);
    axios.get('http://localhost:4061/auth/'+provider);
}


function Login({ providers }) {
    console.log("providers:", providers)
    return (
        <div className={styles.container}>
            <h1> Available providers: </h1>
            <ul>
            {providers?.providers.map((provider) => (
                <button key={provider} onClick={() => handleLogin(provider)}>{provider}</button>
            ))}
            </ul>
        </div>
      )
}


Login.getInitialProps = async ctx => {
    try {
      const res = await axios.get('http://localhost:4061/auth/providers');
      const providers = res.data;
      return { providers };
    } catch (error) {
      return { error };
    }
  };

  
export default Login