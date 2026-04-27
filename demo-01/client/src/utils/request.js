import axios from "axios";

const request = axios.create({
  baseURL: "http://localhost:8080",
  timeout: 10000,
});

// 请求拦截器
request.interceptors.request.use((config) => {
  // 可加 token
  // const token = localStorage.getItem('token')
  // if (token) config.headers.Authorization = `Bearer ${token}`
  return config;
});

// 响应拦截器
request.interceptors.response.use(
  (res) => {
    return res.data;
  },
  (err) => {
    console.error("请求异常", err);
    return Promise.reject(err);
  },
);

export default request;
