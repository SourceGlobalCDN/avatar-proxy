import axios from "axios";

const baseUrl = "https://www.gravatar.com";

const instance = axios.create({
    baseURL: baseUrl,
    withCredentials: false,
});

export default instance;
