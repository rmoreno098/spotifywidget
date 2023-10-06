import React from 'react';
import { Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
// import {Container, Navbar, Nav, NavDropdown} from 'react-bootstrap';


function HomePage() {
  return (
    <div>
      <h1>Welcome to the Home Page</h1>
      <Link to="/">Login</Link>
    </div>
  );
}

export default HomePage;