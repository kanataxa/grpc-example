const path = require('path');

module.exports = ( env, argv ) => ({
	entry: './src/main.js',
	output: {
		filename: 'main.js',
		path: path.resolve(__dirname, 'dist')
	},
	watch: true,
	devServer: {
		open: true,
		openPage: "index.html",
		contentBase: __dirname,
		watchContentBase: true,
		port: 3000,
	},

});
