// import React from 'react';
import './login.css';

function generateRandomString(length: number): string {
  let text = '';
  let possible = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';

  for (let i = 0; i < length; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  }
  return text;
}

async function generateCodeChallenge(codeVerifier: string): Promise<string> {
  function base64encode(string: string): string {
    return btoa(string)
      .replace(/\+/g, '-')
      .replace(/\//g, '_')
      .replace(/=+$/, '');
  }

  const encoder = new TextEncoder();
  const data = encoder.encode(codeVerifier);
  const digest = await window.crypto.subtle.digest('SHA-256', data);

return base64encode(String.fromCharCode.apply(null, Array.from(new Uint8Array(digest))));
}

const handleLogin = () => {
  const clientId: string = '98fc1b94f1e445cebcfe067a505598ba';
  const redirectUri: string = 'http://localhost:8080/callback';

  let codeVerifier: string = generateRandomString(128);

  generateCodeChallenge(codeVerifier).then((codeChallenge: string) => {
    let state: string = generateRandomString(16);
    let scope: string = 'user-read-private user-read-email';

    localStorage.setItem('code_verifier', codeVerifier);

    let args = new URLSearchParams({
      response_type: 'code',
      client_id: clientId,
      scope: scope,
      redirect_uri: redirectUri,
      state: state,
      code_challenge_method: 'S256',
      code_challenge: codeChallenge,
    });

    // Data to send to the server
    const dataToSend = {
      codeVerifier: codeVerifier,
    };

    // Send a POST request to your Go server
    fetch('http://localhost:8080/incoming', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(dataToSend),
    })
      .then(response => {
        if (response.ok) {
          console.log('Data sent successfully');
        } else {
          console.error('Failed to send data to the server');
        }
      })
      .catch(error => {
        console.error('Error sending data:', error);
      });

    window.location.href = 'https://accounts.spotify.com/authorize?' + args;
  });
};

function LoginPage() {
  return (
    <div className="login-page">
      <div className="login-container">
        <h1>Spotify Login</h1>
        <button className="login-button" onClick={handleLogin}>
          Login
        </button>
      </div>
    </div>
  );
}

export default LoginPage;
