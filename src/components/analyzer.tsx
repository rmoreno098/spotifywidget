import { useParams } from 'react-router-dom';
function LousPage() {

  const userid = useParams();
	console.log(userid)
	async function getAnalyzer(){
		const result = await fetch("http://localhost:8080/analyzer",
		{
        method: "POST",
        headers: { "Content-Type": "application/json"},
        body: JSON.stringify({ user_id: "minxical"}),
    })
		const items = await result.json();
    return items;
	}

  async function EventHandler(event: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
    event.preventDefault();
    // here is ur buttons functionality
		console.log("Event registered")
		return await getAnalyzer()
  }
    
  return (
    // here is where u add all the shit u wanna render on ur page
    <button onClick={(e)=>EventHandler(e)}>TEST</button>
  );
}

export default LousPage;
