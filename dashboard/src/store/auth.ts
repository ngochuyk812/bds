import { create } from 'zustand';
import { persist, createJSONStorage, PersistOptions } from 'zustand/middleware';
import { axiosInstance } from '../utils/axios';
import { AuthState, LoginCredentials, User } from '../types/auth';

const initialState: AuthState = {
  user: null,
  accessToken: null,
  refreshToken: null,
  isAuthenticated: false,
  isLoading: false,
  error: null,
};

type AuthStore = AuthState & {
  login: (credentials: LoginCredentials) => Promise<void>;
  logout: () => void;
  setUser: (user: User) => void;
  clearError: () => void;
};

type AuthPersist = {
  user: User | null;
  accessToken: string | null;
  refreshToken: string | null;
  isAuthenticated: boolean;
};

const persistConfig: PersistOptions<AuthStore, AuthPersist> = {
  name: 'auth-storage',
  storage: createJSONStorage(() => localStorage),
  partialize: (state) => (state),
};

export const useAuthStore = create<AuthStore>()(
  persist(
    (set) => ({
      ...initialState,

      login: async (credentials: LoginCredentials) => {
        try {
          set({ isLoading: true, error: null });
          const response = await axiosInstance.post('/auth-service/auth.v1.AuthService/Login', credentials);
          const { accessToken, refreshToken } = response.data;

          localStorage.setItem('auth_token', accessToken);
          if (refreshToken) {
            localStorage.setItem('refresh_token', refreshToken);
          }
          set({
            accessToken,
            refreshToken: refreshToken || null,
            isAuthenticated: true,
            isLoading: false,
            error: null,
          });
        } catch (error: any) {
          console.log(error);
          const errorMessage = error.response?.data?.message || 'Login failed. Please try again.';
          set({
            ...initialState,
            error: errorMessage,
            isLoading: false,
          });
        }
      },


      logout: () => {
        localStorage.removeItem('auth_token');
        localStorage.removeItem('refresh_token');
        set(initialState);
      },


      setUser: (user: User) => {
        set({ user });
      },

      clearError: () => {
        set({ error: null });
      },
    }),
    persistConfig
  )
);
