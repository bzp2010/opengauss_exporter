const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

// With JSDoc @type annotations, IDEs can provide config autocompletion
/** @type {import('@docusaurus/types').DocusaurusConfig} */
(module.exports = {
  title: 'openGuass Exporter',
  tagline: 'Prometheus Exporter for openGauss',
  url: 'https://bzp2010.github.io/opengauss_exporter/',
  baseUrl: '/opengauss_exporter/',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  favicon: 'img/favicon.ico',
  organizationName: 'openGuass', // Usually your GitHub org/user name.
  projectName: 'opengauss_exporter', // Usually your repo name.

  i18n: {
    defaultLocale: 'zh-CN',
    locales: ['zh-CN', 'en'],
  },

  presets: [
    [
      '@docusaurus/preset-classic',
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          routeBasePath: '/',
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl: 'https://github.com/bzp2010/opengauss_exporter/edit/main/docs/',
        },
        blog: {
          showReadingTime: true,
          // Please change this to your repo.
          editUrl:
            'https://github.com/facebook/docusaurus/edit/main/website/blog/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        title: 'Exporter',
        logo: {
          alt: 'openGauss Logo',
          src: 'img/opengauss.svg',
        },
        items: [],
      },
      footer: {
        style: 'light',
        links: [
          {
            title: 'More',
            items: [
              {
                label: 'GitHub',
                href: 'https://github.com/bzp2010/opengauss_exporter',
              },
              {
                label: 'Gitee',
                href: 'https://gitee.com/bzp2010/opengauss_exporter',
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} openGauss Community`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
});
