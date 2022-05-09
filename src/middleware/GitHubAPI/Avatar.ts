import axios from "axios";

const baseUrl = "https://avatars.githubusercontent.com";

const instance = axios.create({
    baseURL: baseUrl,
    withCredentials: false,
    method: "GET",
});

export default instance;
