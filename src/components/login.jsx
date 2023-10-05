import React from 'react';
import './login.css';

function generateRandomString(length) {
  let text = '';
  let possible = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';

  for (let i = 0; i < length; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  }
  return text;
}


async function generateCodeChallenge(codeVerifier) {
  function base64encode(string) {
    return btoa(String.fromCharCode.apply(null, new Uint8Array(string)))
      .replace(/\+/g, '-')
      .replace(/\//g, '_')
      .replace(/=+$/, '');
  }

  const encoder = new TextEncoder();
  const data = encoder.encode(codeVerifier);
  const digest = await window.crypto.subtle.digest('SHA-256', data);

  return base64encode(digest);
}


const handleLogin = () => {
  const clientId = '98fc1b94f1e445cebcfe067a505598ba';
  const redirectUri = 'http://localhost:8080/callback';

  let codeVerifier = generateRandomString(128);

  generateCodeChallenge(codeVerifier).then(codeChallenge => {
    let state = generateRandomString(16);
    let scope = 'user-read-private user-read-email';

    localStorage.setItem('code_verifier', codeVerifier);

    let args = new URLSearchParams({
      response_type: 'code',
      client_id: clientId,
      scope: scope,
      redirect_uri: redirectUri,
      state: state,
      code_challenge_method: 'S256',
      code_challenge: codeChallenge
    });

    window.location = 'https://accounts.spotify.com/authorize?' + args;
  });

};

const urlParams = new URLSearchParams(window.location.search);
let code = urlParams.get('code');

console.log("key from login page \n" + code);


function LoginPage() {
  return (
    <div className="login-page">
      <div className="login-container">
        <h1>Spotify Login</h1>
        <button className="login-button" onClick={handleLogin}>Login</button>
      </div>
    </div>
  );
}

export default LoginPage;
