import Banner from "./component/Banner";
import Contact from "./component/Contact";
import Feature from "./component/Feature";
import Footer from "./component/Footer";
import Navbar from "./component/Navbar";
import Projects from "./component/Projects";
import Resume from "./component/Resume";
import Testimonial from "./component/Testimonial";


function App() {
  return (
    <main className="font-bodyFont w-full h-auto bg-bodyColor text-lightText">
      <Navbar />
      <div className="px-4">
        <div className="max-w-screen-xl mx-auto">
          <Banner />
          <Feature />
          <Projects />
          <Resume />
          <Testimonial />
          <Contact />
          <Footer />
        </div>
      </div>
    </main>
  );
}

export default App;
