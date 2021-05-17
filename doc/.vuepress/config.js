module.exports = {
    title: 'Gobatis',
    description: '适用于 golang 的数据库操作库',
    locales: {
        // 键名是该语言所属的子路径
        // 作为特例，默认语言可以使用 '/' 作为其路径。
        '/': {
            lang: 'zh-CN',
            title: 'VuePress',
            description: 'Vue 驱动的静态网站生成器'
        },
        '/en/': {
            lang: 'en-US', // 将会被设置为 <html> 的 lang 属性
            title: 'VuePress',
            description: 'Vue-powered Static Site Generator'
        },
    },
    themeConfig: {
        logo: '/assets/img/logo.png',
        search: false,
        searchMaxSuggestions: 10,
        smoothScroll: true,
        locales: {
            '/': {
                nav: [
                    { text: '文档', link: '/getting-started' },
                    { text: 'External', link: 'https://google.com', target: '_blank' },
                ],
                sidebar: [
                    {
                        title: '',   // 必要的
                        // path: '/foo/',      // 可选的, 标题的跳转链接，应为绝对路径且必须存在
                        collapsable: false, // 可选的, 默认值是 true,
                        sidebarDepth: 1,    // 可选的, 默认值是 1
                        children: [
                            ['/introduction', '介绍'],
                            ['/getting-started', '快速上手'],
                            ['/configuration', 'XML 配置'],
                            ['/sqlmap-xml', 'XML 映射文件'],
                            ['/dynamic-sql', '动态 SQL'],
                        ]
                    },

                ],
            },
            '/en/': {
                nav: [
                    { text: 'Home', link: '/' },
                    { text: 'External', link: 'https://google.com', target: '_blank' },
                ],
                sidebar: [
                    '/',
                    '/page-a',
                    ['/page-b', 'Explicit link text']
                ],
            }
        }
    }
}