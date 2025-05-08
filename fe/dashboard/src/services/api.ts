import { axiosInstance } from '../utils/axios';
import { AuthResponse, LoginCredentials, User } from '../types/auth';

/**
 * Authentication API service
 */
export const authApi = {
  /**
   * Login with username and password
   */
  login: async (credentials: LoginCredentials): Promise<AuthResponse> => {
    const response = await axiosInstance.post<AuthResponse>('/auth/login', credentials);
    return response.data;
  },

  /**
   * Logout the current user
   */
  logout: async (): Promise<void> => {
    await axiosInstance.post('/auth/logout');
  },

  /**
   * Get the current user profile
   */
  getProfile: async (): Promise<User> => {
    const response = await axiosInstance.get<User>('/auth/profile');
    return response.data;
  },

  /**
   * Refresh the authentication token
   */
  refreshToken: async (refreshToken: string): Promise<{ token: string }> => {
    const response = await axiosInstance.post<{ token: string }>('/auth/refresh', { refreshToken });
    return response.data;
  },
};

/**
 * Example of another API service using the same axios instance
 */
export const userApi = {
  /**
   * Get all users
   */
  getUsers: async (): Promise<User[]> => {
    const response = await axiosInstance.get<User[]>('/users');
    return response.data;
  },

  /**
   * Get user by ID
   */
  getUserById: async (id: string): Promise<User> => {
    const response = await axiosInstance.get<User>(`/users/${id}`);
    return response.data;
  },
};
