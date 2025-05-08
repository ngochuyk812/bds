import axios, { AxiosError, AxiosInstance, AxiosRequestConfig } from 'axios';

const API_URL = process.env.REACT_APP_API_URL || 'https://api-dev.nnh.io.vn';

class CustomAxios {
  private instance: AxiosInstance;

  constructor() {
    console.log(API_URL)
    this.instance = axios.create({
      baseURL: API_URL,
      timeout: 30000,
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
    });

    this.setupInterceptors();
  }

  private setupInterceptors(): void {
    this.instance.interceptors.request.use(
      (config) => {
        const token = localStorage.getItem('auth_token');

        if (token && config.headers) {
          config.headers.Authorization = `Bearer ${token}`;
        }

        return config;
      },
      (error) => Promise.reject(error)
    );

    this.instance.interceptors.response.use(
      (response) => response,
      async (error: AxiosError) => {
        const originalRequest = error.config as AxiosRequestConfig & { _retry?: boolean };

        if (error.response?.status === 401 && !originalRequest._retry) {
          originalRequest._retry = true;

          try {
            const refreshToken = localStorage.getItem('refresh_token');

            if (refreshToken) {
              // Example refresh token request
              // const response = await axios.post(`${API_URL}/auth/refresh`, { refreshToken });
              // const { token } = response.data;
              // localStorage.setItem('auth_token', token);

              // Update the failed request with the new token and retry
              // if (originalRequest.headers) {
              //   originalRequest.headers.Authorization = `Bearer ${token}`;
              // }
              // return this.instance(originalRequest);
            }
          } catch (refreshError) {
            localStorage.removeItem('auth_token');
            localStorage.removeItem('refresh_token');
            window.location.href = '/login';
            return Promise.reject(refreshError);
          }
        }

        if (error.response?.status === 403) {
          console.error('Access forbidden');
        } else if (error.response?.status === 404) {
          console.error('Resource not found');
        } else if (error.response?.status === 500) {
          console.error('Server error');
        }
        return Promise.reject(error);
      }
    );
  }


  public getAxiosInstance(): AxiosInstance {
    return this.instance;
  }


  public setAuthToken(token: string | null): void {
    if (token) {
      localStorage.setItem('auth_token', token);
    } else {
      localStorage.removeItem('auth_token');
    }
  }


  public setRefreshToken(token: string | null): void {
    if (token) {
      localStorage.setItem('refresh_token', token);
    } else {
      localStorage.removeItem('refresh_token');
    }
  }


  public clearTokens(): void {
    localStorage.removeItem('auth_token');
    localStorage.removeItem('refresh_token');
  }
}

const customAxios = new CustomAxios();
export default customAxios;

export const axiosInstance = customAxios.getAxiosInstance();
