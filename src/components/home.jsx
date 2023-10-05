import React from 'react';
import { Link } from 'react-router-dom';

const urlParams = new URLSearchParams(window.location.search);
let code = urlParams.get('code');

console.log("key from home page \n" + code);

function HomePage() {
  return (
    <div>
      <h1>Welcome to the Home Page</h1>
      <Link to="/">Login</Link>
    </div>
  );
}

export default HomePage;