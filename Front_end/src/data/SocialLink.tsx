import { FaFacebookF, FaGithub, FaYoutube } from "react-icons/fa"


const SocialLink = () => {
  return (
   <>
   <a href="https://www.youtube.com/" target="_blank">
      <span className="bannerIcon">
        <FaYoutube />
      </span>
    </a>
    <a
      href="https://www.linkedin.com/"
      target="_blank"
    >
      <span className="bannerIcon">
        <FaGithub />
      </span>
    </a>
    <a href="https://www.facebook.com/" target="_blank">
      <span className="bannerIcon">
        <FaFacebookF />
      </span>
    </a>
   </>
  )
}

export default SocialLink