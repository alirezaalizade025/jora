import { AxiosInstance, AxiosRequestConfig, AxiosRequestHeaders, AxiosResponse } from 'axios';
import changeStringCaseRecursive, { changeBooleanToNumber } from 'utils/changeStringCaseRecursive';
import { cookieGetter } from 'utils/cookieUtils';

type axiosErrorInspectorType = {
  config: AxiosRequestConfig & {
    __retryCount: number;
    maxRetryCount: number;
  };
  message?: string;
};

export const shouldAxiosRequestRetry = (error: axiosErrorInspectorType) => {
  const ALLOWED_RETRY_METHODS = ['get', 'put', 'delete', 'head', 'options', 'post'];
  const isAllowMethod = ALLOWED_RETRY_METHODS.includes(error.config?.method);
  const isAllowStatus = false; //ALLOWED_RETRY_CODE.includes(error.response?.status); disable retry server errors
  const isNetworkErr = error.message === 'Network Error';
  const canRetryMore = (error.config?.__retryCount || 0) < error.config?.maxRetryCount;

  return canRetryMore && (isNetworkErr || (isAllowMethod && isAllowStatus));
};

export const axiosRequestConverter = (config: AxiosRequestConfig) => {
  const jwt = getJwtToken();
  const MAX_ATTEMPTS = 5;

  return {
    ...config,
    params: changeBooleanToNumber(changeStringCaseRecursive(config.params, 'toSnakeCase')),
    data: changeStringCaseRecursive(config.data, 'toSnakeCase'),
    headers: {
      ...config.headers,
      ...(jwt && { Authorization: `Bearer ${jwt}` }),
    } as AxiosRequestHeaders,
    maxRetryCount: MAX_ATTEMPTS,
  };
};

export const axiosResponseConverter = (res): AxiosResponse => {
  if (res.config?.responseType == 'arraybuffer') {
    return res;
  }

  return {
    ...res,
    data: changeStringCaseRecursive(res.data, 'toCamelCase'),
  };
};

const isWarningApi = (e) => e.response?.data?.type === 'warning';

export const axiosResponseError = async (
  error: axiosErrorInspectorType,
  axiosInstance: AxiosInstance,
) => {
  if (shouldAxiosRequestRetry(error)) {
    const { __retryCount: retryCount = 0, url, data } = error.config;
    error.config.__retryCount = retryCount + 1;
    return new Promise((resolve) => {
      setTimeout(
        () =>
          resolve(axiosInstance({ ...error.config, ...(data ? { data: parseJson(data) } : {}) })),
        5000,
      );
    });
  }
  throw error;
};

export const getJwtToken = () => {
  const jwt = cookieGetter({ name: 'jwt' }) || null;
  return jwt;
};

export const removeNulls = (obj) =>
  Object.keys(obj).reduce((acc, key) => (obj[key] == null ? acc : { ...acc, [key]: obj[key] }), {});

export const removeNullables = (obj) =>
  Object.keys(obj).reduce(
    (acc, key) =>
      obj[key] === null || obj[key] === '' || obj[key] === 0 ? acc : { ...acc, [key]: obj[key] },
    {},
  );

export const convertBooleanToNullable = (obj) => (obj == true ? true : (obj == false ? false : null));

const parseJson = (stringify) => {
  if (typeof stringify !== 'string') {
    return stringify;
  }
  try {
    return JSON.parse(stringify);
  } catch {
    return stringify;
  }
};

export default {};
