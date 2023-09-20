const formData = require("form-data");
const fs = require("fs");

function upload(requestParams, _context, _ee, next) {
  const form = new formData();
  form.append("image", fs.createReadStream("mock.png"));
  requestParams.body = form;
  return next();
}

module.exports = {
  upload,
};
