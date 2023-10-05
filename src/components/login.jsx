import React from 'react';
import './login.css'; // You can create a separate CSS file for styling

function LoginPage() {
  return (
    <div className="login-page">
      <div className="login-container">
        <h1>Spotify Login</h1>
        <button className="login-button">Login</button>
      </div>
    </div>
  );
}

export default LoginPage;
