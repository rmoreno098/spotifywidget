import React from 'react';
import { BrowserRouter, Routes, Route} from 'react-router-dom';
import LoginPage from './components/login'
import HomePage from './components/home';
import 'bootstrap/dist/css/bootstrap.min.css';
// import {Container, Navbar, Nav, NavDropdown} from 'react-bootstrap';

function App() {
  return (
    <BrowserRouter>
      <div>
        <Routes>
          <Route path="/" element={ <LoginPage />} />
          <Route path="/home" element={ <HomePage />} />
          {/* more routes go here */}
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
