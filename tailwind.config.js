/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './component/**/*.templ',
  ],
  darkMode: 'media',
  theme: {
    container: {
      center: true,
    },
    fontFamily: {
      'sans': ['Poppins'],
      'display': ['Poppins'],
    },
    extend: {
      colors: {
        orange: {
          DEFAULT: "#d65d0e",
          500: "#d65d0e",
          400: "#fe8019",
        },
        red: {
          DEFAULT: "#cc241d",
          500: "#cc241d",
          400: "#fb4934",
        },
        green: {
          DEFAULT: "#98971a",
          500: "#98971a",
          400: "#b8bb26",
        },
        yellow: {
          DEFAULT: "#d79921",
          500: "#d79921",
          400: "#fabd2f",
        },
        blue: {
          DEFAULT: "#458588",
          500: "#458588",
          400: "#83a598",
        },
        purple: {
          DEFAULT: "#b16286",
          500: "#b16286",
          400: "#d3869b",
        },
        aqua: {
          DEFAULT: "#689d6a",
          500: "#689d6a",
          400: "#8ec07c",
        },
        gray: {
          DEFAULT: "#282828",
          800: "#282828",
          900: "#1d2021",
          700: "#3c3836",
          600: "#504945",
          500: "#665c54",
          400: "#7c6f64",
          300: "#928374",
        },
        light: {
          DEFAULT: "#fbf1c7",
          500: "#fbf1c7",
          600: "#ebdbb2",
          700: "#d5c4a1",
          800: "#bdae93",
          900: "#a89984",
        }
      }
    }
  },
  plugins: [
    '@tailwindcss/forms',
    '@tailwindcss/typography',
  ]
}
