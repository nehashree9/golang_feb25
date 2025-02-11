import Greeting from "./Greeting";
import Navbar from "./header/Navbar";

export default function App() {
  return (
    <>
    <Navbar/>
      <h1>
        <marquee>Welcome to CLASS</marquee>
      </h1>
      <Greeting/>
    </>
  );
}