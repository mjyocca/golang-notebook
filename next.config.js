const withNextra = require('nextra')({
  theme: 'nextra-theme-docs',
  themeConfig: './theme.config.tsx',
  defaultShowCopyCode: true,
  // mdxOptions: {
  //   rehypePrettyCodeOptions: {
  //     theme: 'github-dark'
  //   }
  // }
})

let nextProps = {}

if (process.env.NODE_ENV !== 'development') {
  nextProps = {
    output: "export",
    basePath: "/golang-notebook",
    images: {
      unoptimized: true,
    },
  }
}

module.exports = withNextra(nextProps)
