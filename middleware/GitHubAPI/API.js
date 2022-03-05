const axios = require("axios");

const baseUrl = "https://api.github.com";

const instance = axios.create({
    baseURL: baseUrl,
    withCredentials: false,
});

module.exports = instance;
