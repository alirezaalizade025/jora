import axios from 'axios';
import { axiosRequestConverter, axiosResponseConverter, axiosResponseError } from './utils';

const API = axios.create({
  withCredentials: false,
  baseURL: 'jora.com',
  timeout: 300_000,
});

API.interceptors.request.use(axiosRequestConverter, (error) => {
  return Promise.reject(error);
});

API.interceptors.response.use(axiosResponseConverter, (err) => axiosResponseError(err, API));

export default API;
