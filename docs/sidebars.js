module.exports = {
    tutorialSidebar: [
        {
            type: 'category',
            label: '简介',
            collapsible: true,
            collapsed: false,
            items: [
                'introduction',
                'getting-started/5min-tutorial',
                'getting-started/installation',
            ],
        },
        {
            type: 'category',
            label: '概念',
            collapsible: true,
            collapsed: true,
            items: [
                'concepts/architecture',
                'concepts/scrapers',
            ],
        },
        {
            type: 'category',
            label: '指引',
            collapsible: true,
            collapsed: true,
            items: [
                'guides/developer',
            ],
        },
        {
            type: 'category',
            label: '参考',
            collapsible: true,
            collapsed: true,
            items: [
                'reference/config-file',
                'reference/http-api',
                'reference/cli',
            ]
        }
    ],
};
