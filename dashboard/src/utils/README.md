# Custom Axios and Authentication

This directory contains utilities for API communication and authentication.

## Custom Axios

The `axios.ts` file provides a custom Axios instance with the following features:

- Pre-configured base URL and timeout
- Authentication token handling via interceptors
- Error handling and token refresh logic
- Utility methods for token management

### Usage

```typescript
// Import the Axios instance
import { axiosInstance } from '../utils/axios';

// Make API requests
const fetchData = async () => {
  try {
    const response = await axiosInstance.get('/some-endpoint');
    return response.data;
  } catch (error) {
    console.error('Error fetching data:', error);
    throw error;
  }
};

// Set authentication token
import customAxios from '../utils/axios';
customAxios.setAuthToken('your-token-here');

// Clear tokens on logout
customAxios.clearTokens();
```

## Authentication Store

The authentication store (in `src/store/auth.ts`) provides:

- User authentication state management
- Login/logout functionality
- Persistent authentication across page refreshes
- Error handling

### Usage

```typescript
import { useAuthStore } from '../store/auth';

// In your component
const MyComponent = () => {
  // Access auth state and methods
  const { 
    user, 
    isAuthenticated, 
    isLoading, 
    error, 
    login, 
    logout 
  } = useAuthStore();

  // Login
  const handleLogin = async () => {
    await login({ username: 'user', password: 'pass' });
  };

  // Logout
  const handleLogout = () => {
    logout();
  };

  return (
    <div>
      {isAuthenticated ? (
        <div>
          <p>Welcome, {user?.name}!</p>
          <button onClick={handleLogout}>Logout</button>
        </div>
      ) : (
        <button onClick={handleLogin} disabled={isLoading}>
          {isLoading ? 'Logging in...' : 'Login'}
        </button>
      )}
      {error && <p>Error: {error}</p>}
    </div>
  );
};
```

## API Services

The `src/services/api.ts` file provides structured API service functions that use the custom Axios instance:

```typescript
import { authApi, userApi } from '../services/api';

// Authentication
const login = async () => {
  const authResponse = await authApi.login({ username: 'user', password: 'pass' });
  console.log(authResponse.user, authResponse.token);
};

// User data
const getUsers = async () => {
  const users = await userApi.getUsers();
  console.log(users);
};
```
