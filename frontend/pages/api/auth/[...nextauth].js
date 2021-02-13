import NextAuth from 'next-auth'
import Providers from 'next-auth/providers'
import axios from 'axios';

const providers = [
    Providers.GitHub({
        clientId: process.env.GITHUB_ID,
        clientSecret: process.env.GITHUB_SECRET
    }),
    Providers.Google({
        clientId: process.env.GOOGLE_ID,
        clientSecret: process.env.GOOGLE_SECRET
    }),
]

const callbacks = {}

callbacks.signIn = async function signIn(user, account, metadata) {
    console.log("callbacks:signIn")
    if (account.provider === 'github') {
        // grab email  from github because it does not reuturn by default
        let primaryEmail = ""
        try {
            const emailRes = await fetch('https://api.github.com/user/emails', {
                headers: {
                  Authorization: `token ${account.accessToken}`,
                },
              })
            const emails = await emailRes.json()
            primaryEmail = emails.find(e => e.primary).email; 
        } catch (error) {
            console.log(error)
        }
        // create a user obj
        const githubUser = {
            id: ''+metadata.id+'',
            login: metadata.login,
            name: metadata.name,
            image: user.image,
            email: primaryEmail
        }
    
        console.log("user created from provider data:", githubUser)

        // now that we have the user obj we validate and get a token from backend api
        // generates a token (maybe a user token id)
        const test = await axios.post("http://localhost:4061/auth/login", githubUser, {
            headers: {
              provider: 'github',
              Authorization: `Bearer ${account.accessToken}`,
        }})
        user.accessToken = test.data?.token
        console.log("token received from backend: ",test.data.token)
        console.log("sign in succefully")
        return true
    }
    console.log("sign in fail")
    return false;
}


callbacks.jwt = async function jwt(token, user) {
    console.log("callbacks:jwt ")
    if (user) {
        token = { accessToken: user.accessToken }
    }
    return token
}

callbacks.session = async function session(session, token) {
    console.log("callbacks:session ")
    console.log("callbacks:session call backend with token: ", token.accessToken)
    session.accessToken = token.accessToken
    // Gets user from backend
    const user = await axios.get("http://localhost:4061/user", {
        headers: {
          Authorization: `${session.accessToken}`,
    }})
    console.log("User received from backend:", user.data)
    session.user = user.data
    return session
}

const options = {
    providers,
    callbacks
}

export default (req, res) => NextAuth(req, res, options)