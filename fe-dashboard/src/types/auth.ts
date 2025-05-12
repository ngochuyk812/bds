export interface User {
  id: string;
  username: string;
  email: string;
  name?: string;
}

export interface LoginCredentials {
  username: string;
  password: string;
}
export interface SignUpCredentials {
  username: string;
  password: string;
  rePassword: string;
  name: string;
}

export interface VerifySignUpCredentials {
  username: string;
  otp: string;
}

export interface AuthResponse {
  user: User;
  accessToken: string;
  refreshToken?: string;
}

export interface AuthState {
  user: User | null;
  accessToken: string | null;
  refreshToken: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  error: string | null;
}
