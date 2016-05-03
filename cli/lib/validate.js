'use strict';

var validate = exports;

var validator = require('validator');

validate.name = function(input) {
  /**
   * TODO: Change js validation for json schema
   * https://github.com/arigatomachine/cli/issues/134
   */
  var error = 'Please provide your full name';
  return input.length < 3 || input.length > 64? error : true;
};

validate.email = function(input) {
  var error = 'Please enter a valid email address';
  return validator.isEmail(input)? true : error;
};

validate.passphrase = function(input) {
  var error = 'Passphrase must be at least 8 characters';
  return input.length < 8? error : true;
};

validate.code = function(input) {
  var error = 'Verification code must be exactly 9 characters';
  var trimmed = input.replace(/\s/g, '');
  return trimmed.length !== 9? error : true;
};
