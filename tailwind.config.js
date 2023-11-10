/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["**/*.{html,templ}"],
    theme: {
        extend: {
            fontFamily: {
                'sans': ['Inter', 'sans-serif'],
                'mono': ['"Fira Code"', 'monospace']
            },
        }
    },
    plugins: [
        require('@tailwindcss/typography'),
        require('@tailwindcss/forms'),
    ]
}
