const defaultTheme = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Inter", ...defaultTheme.fontFamily.sans],
        block: ["Noway", ...defaultTheme.fontFamily.sans],
        title: ["n27", ...defaultTheme.fontFamily.sans],
        icons: ["'Noway Icons'"],
      }
    },
  },
  plugins: [],
}
