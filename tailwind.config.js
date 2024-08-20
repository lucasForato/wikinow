/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./component/**/*.templ"
  ],
  theme: {
    extend: {
      extend: {
        colors: {
          'primary': '#3490dc',
          'secondary': '#ffed4a',
          'danger': '#e3342f',
        }
      }
    },
  },
  plugins: [
    '@tailwindcss/typography',
  ],
}

