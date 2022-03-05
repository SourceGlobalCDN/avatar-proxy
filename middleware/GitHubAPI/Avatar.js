const axios = require("axios");

const baseUrl = "https://avatars.githubusercontent.com";

const instance = axios.create({
    baseURL: baseUrl,
    withCredentials: false,
    method: "GET",
});

module.exports = instance;
