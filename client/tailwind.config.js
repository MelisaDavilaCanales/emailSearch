/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        white: '#ffffff',
        grayDark: '#737285',
        graySoft: '#d9d9d9',
        grayExtraSoft: '#f3f4f6',  //#f2f2f2 //f5f5f5
        primary: '#782be4',
        primaryDark: '#5e1fc7',
        primaryMiddle: '#8c52ff',
        primarySoft: '#c1a1ff',
        primaryExtraSoft: '#f0e9ff',
      },
      fontFamily: {
        roboto: ['Roboto', 'sans-serif'],
        opensans: ['Open Sans', 'sans-serif'],
      },
    },
  },
  plugins: [],
}
