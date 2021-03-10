module.exports = {
    // 选项...
    publicPath: process.env.NODE_ENV === 'production'
        ? '/coupon-code/'
        : '/',
    css: {
        loaderOptions: {
            less: {
                modifyVars: {
                    'card-background-color': '#FFFFFF'
                }
            }
        }
    }
}