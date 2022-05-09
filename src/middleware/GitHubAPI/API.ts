import axios from "axios";

const baseUrl = "https://api.github.com";

const instance = axios.create({
    baseURL: baseUrl,
    withCredentials: false,
});

export default instance;
