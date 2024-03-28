/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        bodyFont: ["Poppins", "sans-serif"],
        titleFont: ["Montserrat", "sans-serif"],
      },
      screens: {
        xs: "320px",
        sm: "375px",
        sml: "500px",
        md: "667px",
        mdl: "768px",
        lg: "960px",
        lgl: "1024px",
        xl: "1280px",
      },
      colors: {
        navColor: "#ff008d",
        bodyColor: "#1c1c1c",
        lightText: "#c4cfde",
        designColor: "#8900ff",
        titleColor: "#ffffff",
      },
      boxShadow: {
        shadowOne: "10px 10px 19px #030712, -10px -10px 19px #0f172a",
      },
    },
  },
  plugins: [],
};