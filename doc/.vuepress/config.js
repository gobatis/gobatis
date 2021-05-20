module.exports = {
    locales: {
        // 键名是该语言所属的子路径
        // 作为特例，默认语言可以使用 '/' 作为其路径。
        '/': {
            lang: 'zh-CN',
            title: 'Gobatis',
            description: 'Golang 持久层框架'
        },
        '/en/': {
            lang: 'en-US', // 将会被设置为 <html> 的 lang 属性
            title: 'Gobatis',
            description: 'Golang database ORM'
        },
    },
    themeConfig: {
        // logo: '/assets/img/logo.png',
        search: false,
        searchMaxSuggestions: 10,
        smoothScroll: true,
        locales: {
            '/': {
                nav: [
                    {text: '文档', link: '/getting-started'},
                    {text: 'External', link: 'https://google.com', target: '_blank'},
                ],
                sidebar: [
                    {
                        title: '入门指南',   // 必要的
                        collapsable: false, // 可选的, 默认值是 true,
                        sidebarDepth: 1,    // 可选的, 默认值是 1
                        children: [
                            ['/introduction', '介绍'],
                            ['/getting-started', '快速上手'],
                        ]
                    },
                    {
                        title: 'SQL',   // 必要的
                        collapsable: false, // 可选的, 默认值是 true,
                        sidebarDepth: 1,    // 可选的, 默认值是 1
                        children: [
                            ['/configuration', 'XML 配置'],
                            ['/sqlmap-xml', 'XML 映射文件'],
                            ['/dynamic-sql', '动态 SQL'],
                            ['/expression', '表达式'],
                            ['/curd', 'CURD'],
                        ]
                    },
                    {
                        title: '高级主题',   // 必要的
                        collapsable: false, // 可选的, 默认值是 true,
                        sidebarDepth: 1,    // 可选的, 默认值是 1
                        children: [
                            ['/engine', 'Engine'],
                            ['/mapper', 'Mapper'],
                            ['/mapper', '事务处理'],
                            ['/mapper', '日志'],
                            ['/mapper', '连接池'],
                            ['/mapper', '性能'],
                            ['/deploy', '部署'],
                        ]
                    },

                ],
            },
            // '/en/': {
            //     nav: [
            //         {text: 'Home', link: '/'},
            //         {text: 'External', link: 'https://google.com', target: '_blank'},
            //     ],
            //     sidebar: [
            //         '/',
            //         '/page-a',
            //         ['/page-b', 'Explicit link text']
            //     ],
            // }
        }
    }
}