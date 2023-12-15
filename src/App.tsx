// import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginPage from './components/login';
import DashboardPage from './components/dashboard';
import 'bootstrap/dist/css/bootstrap.min.css';

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
            path="/dashboard" 
            element={ <DashboardPage />} 
          />
        </Routes>
    </Router>
  );
}

export default App;