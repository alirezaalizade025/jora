type axiosErrorInspectorType = {
  config: import('axios').AxiosRequestConfig & {
    __retryCount: number;
    maxRetryCount: number;
    startTime: Date;
    __spendTime: number;
    __toastId: string | number;
  };
  code?: string;
  request?: import('axios').AxiosRequestConfig;
  response?: import('axios').AxiosResponse;
  message?: string;
};

type PageGetParam = { page?: number; perPage?: number };

type Pagination = { total?: number; count?: number } & PageGetParam;

type PaginationParam = {
  pagination: Pagination;
};

type GetListResponse<T> = {
  pagination: Pagination;
  items: T[];
};
