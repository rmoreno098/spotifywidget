// import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginPage from './components/login';
import HomePage from './components/home';
import 'bootstrap/dist/css/bootstrap.min.css';
// import {Container, Navbar, Nav, NavDropdown} from 'react-bootstrap';

// const fetchData = async () => {
//   try {
//     const response = await fetch('http://localhost:8080/api/hello');
//     if (response.ok) {
//       const data = await response.json();
//       console.log(data);
//     } else {
//       throw new Error('API request failed');
//     }
//   } catch (error) {
//     console.error(error);
//   }
// };


function App() {
  return (
    <Router>
        <Routes>
          <Route
            path="/" 
            element={ <LoginPage />} 
          />
          <Route 
            path="/home" 
            element={ <HomePage />} 
          />
        </Routes>
    </Router>
  );
}

export default App;