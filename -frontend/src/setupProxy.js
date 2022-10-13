const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function (app) {
    app.use(
        '/api',
        createProxyMiddleware({
            target: 'http://localhost:8080',
            changeOrigin: true,
            onProxyReq: function (proxyReq, req, res) {
                Object.keys(req.headers).forEach(function (key) {
                    proxyReq.setHeader(key, req.headers[key])
                })
            },
            onProxyRes: function (proxyRes, req, res) {
                Object.keys(proxyRes.headers).forEach(function (key) {
                    res.append(key, proxyRes.headers[key])
                })
            }
        })
    );
};