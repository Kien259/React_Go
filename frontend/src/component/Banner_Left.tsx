import { useTypewriter, Cursor } from "react-simple-typewriter";
import { FadeIn } from "./FadeIn";
import SocialLink from "~/data/SocialLink";


const LeftBanner = () => {
  const [text] = useTypewriter({
    words: ["Professional Coder.", "Backend Developer."],
    loop: true,
    typeSpeed: 20,
    deleteSpeed: 10,
    delaySpeed: 2000,
  });
  return (
    <FadeIn className="w-full lgl:w-1/2 flex flex-col gap-20">
      <div className="flex flex-col gap-5">
        <h4 className=" text-lg font-normal">WELCOME TO MY PORTFOLIO</h4>
        <h1 className="text-6xl font-bold text-white">
          Hi, I'm <span className="text-designColor capitalize">Kyran Hoang</span>
        </h1>
        <h2 className="text-4xl font-bold text-white">
          a <span>{text}</span>
          <Cursor cursorStyle="|" cursorColor="#001eff" />
        </h2>
        <p className="text-base font-bodyFont leading-6 tracking-wider">
          I use animation as a third dimension by which to simplify experiences
          and kuiding thro each and every interaction. I'm not adding motion
          just to spruce things up, but doing it in ways that.
        </p>
      </div>
      <div className="flex flex-col xl:flex-row gap-6 lgl:gap-0 justify-between">
        <div>
          <h2 className="text-base uppercase font-titleFont mb-4">
            Find me in
          </h2>
          <div className="flex gap-4">
            <SocialLink />
          </div>
        </div>
      </div>
    </FadeIn>
  );
};

export default LeftBanner;