import { create } from 'zustand';
import { persist, createJSONStorage, PersistOptions } from 'zustand/middleware';
import { AuthState, LoginCredentials, User } from '../types/auth';
import { grpcAuthClient } from '../utils/connectrpc';
import { AuthService } from '../proto/genjs/auth/v1/auth_service_pb';
import { useNotificationStore } from './notification';
import { StatusCode } from '../proto/genjs/statusmsg/v1/statusmsg_pb';


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
          const response = await grpcAuthClient.login({
            email: credentials.username,
            password: credentials.password,
          });
          const { accessToken, refreshToken, status } = response;
          if (status?.code != StatusCode.SUCCESS) {
            useNotificationStore.getState().error(StatusCode[status?.code ?? 0], status?.extras?.join('\n'));
            set({ isLoading: false, error: status?.extras?.join('\n') });
            return;
          }

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
          useNotificationStore.getState().error(error.response?.data?.message || 'Login failed. Please try again.');
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
