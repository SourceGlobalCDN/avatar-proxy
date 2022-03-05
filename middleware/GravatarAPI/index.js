const axios = require("axios");

const baseUrl = "https://www.gravatar.com";

const instance = axios.create({
    baseURL: baseUrl,
    withCredentials: false,
});

module.exports = instance;
