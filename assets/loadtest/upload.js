const formData = require('form-data')
const fs = require('fs')

function setupMultiFormData(requestParams, context, ee, next) {
	const form = new formData();
	form.append('file', fs.createReadStream('mock.png'));
	requestParams.body = form;
	return next()
}

module.exports = {
	setupMultiFormData,
}
