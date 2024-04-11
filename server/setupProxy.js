const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = function (app) {
  app.use(
    "/api",
    createProxyMiddleware({
      target: "http://localhost:3002", // Change the target to your backend URL
      changeOrigin: true,
    })
  );
};
