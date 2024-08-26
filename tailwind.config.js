/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./resources/views/app.html",
    "./resources/**/*.{vue,js,ts,jsx,tsx}",
  ],
  darkMode: "class",
  theme: {
    extend: {
      fontFamily: {
        rubik: ["Rubik"],
      },
      colors: {
        primary: "#333333",
        secondary: "#b0b0b0",
        link: "#377fab",
        "link-hover": "#1f6793",
      },
    },
  },
  plugins: [require("@tailwindcss/typography"), require("@tailwindcss/forms")],
};
