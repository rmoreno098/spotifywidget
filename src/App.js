import './App.css';

const fetchData = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/hello');
    if (response.ok) {
      const data = await response.json();
      console.log(data);
    } else {
      throw new Error('API request failed');
    }
  } catch (error) {
    console.error(error);
  }
};


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <button onClick={fetchData}>Click me</button>
      </header>
    </div>
  );
}

export default App;
