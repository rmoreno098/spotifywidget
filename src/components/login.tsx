import "./login.css";
import { Link } from 'react-router-dom';

function LoginPage() {
    
  return (
    <div className="login-page">
      <div className="login-container">
        <h1>Spotify Widget</h1>
        <Link to="/dashboard">Go to Dashboard!</Link>
      </div>
    </div>
  );
}

export default LoginPage;
