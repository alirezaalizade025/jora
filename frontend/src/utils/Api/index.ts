import axios from 'axios';
import { axiosRequestConverter, axiosResponseConverter, axiosResponseError } from './utils';
import { apiBaseUrl } from 'utils/consts';

const API = axios.create({
  withCredentials: false,
  baseURL: apiBaseUrl,
  timeout: 300_000,
});

API.interceptors.request.use(axiosRequestConverter, (error) => {
  return Promise.reject(error);
});

API.interceptors.response.use(axiosResponseConverter, (err) => axiosResponseError(err, API));

export default API;
