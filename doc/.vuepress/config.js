module.exports = {
    plugins: ['@vuepress/nprogress'],
    head: [
        ['script', {}, `
            var _hmt = _hmt || [];
            (function() {
              var hm = document.createElement("script");
              hm.src = "https://hm.baidu.com/hm.js?c307b48e1c6d0bf015b3d856d788356d";
              var s = document.getElementsByTagName("script")[0]; 
              s.parentNode.insertBefore(hm, s);
            })();
        `]
    ],
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
                selectText: '语言',
                nav: [
                    {text: '文档', link: '/introduction'},
                    {text: 'Github', link: 'https://github.com/gobatis/gobatis', target: '_blank'},
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
                            // ['/configuration', 'XML 配置'],
                            ['/sqlmap-xml', 'XML映射文件'],
                            ['/dynamic-sql', '动态 SQL'],
                            ['/expression', '表达式'],
                        ]
                    },
                    {
                        title: '高级主题',   // 必要的
                        collapsable: false, // 可选的, 默认值是 true,
                        sidebarDepth: 1,    // 可选的, 默认值是 1
                        children: [
                            ['/engine', 'Engine'],
                            ['/mapper', 'Mapper'],
                            ['/entity', 'Entity'],
                            ['/tx', '事务'],
                            ['/log', '日志'],
                            ['/pool', '连接池'],
                            ['/deploy', '部署'],
                            ['/performance', '性能'],
                        ]
                    },
                    {
                        title: '开发',   // 必要的
                        collapsable: false, // 可选的, 默认值是 true,
                        sidebarDepth: 1,    // 可选的, 默认值是 1
                        children: [
                            ['/dtd', 'XML语法提示'],
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