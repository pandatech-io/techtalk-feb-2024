import axios, { AxiosError, AxiosRequestConfig } from 'axios';

const createClient = (
    baseURL = process.env.REACT_APP_BASE_URL,
    config = {}
  ) => {
    const axiosInstance = axios.create({
      baseURL,
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      ...config,
    });
  
    return axiosInstance;
  }

export const request = async <T>(
  config: AxiosRequestConfig,
  options?: AxiosRequestConfig
): Promise<T> => {
    const { url, ...restOfConfig } = config;
    const client = createClient("http://localhost:8080");
    const response = await client({
        url: url,
        ...restOfConfig,
        ...options,
    });
    return response?.data;
};

export type ErrorType<Error> = AxiosError<Error>;

export type BodyType<BodyData> = BodyData;