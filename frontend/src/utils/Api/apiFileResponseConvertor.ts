import { AxiosResponse } from 'axios';

export const apiResponseToBlob = (res: AxiosResponse<any, any>) => {
  const blob = new Blob([res.data], { type: res.headers['content-type'] });

  return blob;
};

export const apiResponseToDataUrl = (res: AxiosResponse<any, any>) => {
  const blob = new Blob([res.data], { type: res.headers['content-type'] });
  const dataUrl = URL.createObjectURL(blob);
  return dataUrl;
};

export default {};
