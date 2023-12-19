import "../styles.css";
import { redirectToAuthCodeFlow } from "./auth";
import loginImage from "../images/welcome_page.jpeg";

function WelcomePage() {
    const clientId = "98fc1b94f1e445cebcfe067a505598ba";

    async function spotifyConnect(event: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
      event.preventDefault();
      redirectToAuthCodeFlow(clientId);
    }
    
    return (
    <div className="bg-gray-900 text-white h-screen">
      <section className="gap-5 flex max-md:flex-col max-md:items-stretch max-md:gap-0">
        <div className="flex flex-col items-stretch w-[55%] max-md:w-full max-md:ml-0 h-screen">
          <img
            loading="lazy"
            src={loginImage}
            alt="Spotify Welcome Page Image"
            className="object-cover w-full h-full"
          />
        </div>
        <div className="flex flex-col items-stretch w-[45%] ml-5 max-md:w-full max-md:ml-0">
          <div className="flex flex-col items-center my-auto px-5 max-md:mt-10 justify-center">
            <h1 className="text-4xl font-bold leading-10 tracking-tighter">
              Connect To Spotfiy Like You Never Have Before
            </h1>
            <button className="text-green-500 bg-green-200 px-4 py-2 rounded-md text-xl mt-6 max-w-xs" onClick={(event) => spotifyConnect(event)}>
              Get Started
            </button>
          </div>
        </div>
      </section>
    </div>
  );
}

export default WelcomePage;
