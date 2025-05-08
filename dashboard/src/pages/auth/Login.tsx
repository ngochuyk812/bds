import React, { useState, useEffect } from 'react';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Image, Input, message } from 'antd';
import { LoginCredentials } from '../../types/auth';
import { useAuthStore } from '../../store/auth';
import LayoutPublic from '../../layouts/public';
import { useNavigate } from 'react-router-dom';

const LoginPage: React.FC = () => {
  const { isAuthenticated, isLoading, error, login, logout, clearError } = useAuthStore();
  const navigate = useNavigate();

  const [credentials, setCredentials] = useState<LoginCredentials>({
    username: '',
    password: '',
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setCredentials(prev => ({ ...prev, [name]: value }));
  };

  const handleLogin = async () => {

    if (!credentials.username || !credentials.password) {
      message.error('Please enter both username and password');
      return;
    }

    await login(credentials);
  };

  useEffect(() => {
    if (isAuthenticated) {
      navigate('/');
      return
    }
    if (error) {
      message.error(error);
      clearError();
    }
  }, [error, clearError, isAuthenticated]);

  return (
    <LayoutPublic>
      <div className='h-[100%] flex items-center justify-center'>
        <div className='bg-white p-12 flex flex-col gap-[25px]'>
          <Image src="./logo.png" preview={false} className="max-w-[45%] m-auto" />

          <Input
            size="large"
            placeholder="Email/Username"
            prefix={<UserOutlined />}
            className="rounded-none min-w-[350px] text-[14px]"
            name="username"
            value={credentials.username}
            onChange={handleInputChange}
          />
          <Input
            size="large"
            placeholder="Password"
            prefix={<LockOutlined />}
            className="rounded-none min-w-[350px] text-[14px]"
            type="password"
            name="password"
            value={credentials.password}
            onChange={handleInputChange}
          />
          <Button
            type="primary"
            className='bg-primary mt-4'
            onClick={handleLogin}
            loading={isLoading}
          >
            Đăng nhập
          </Button>
        </div>
      </div>
    </LayoutPublic>
  );
};

export default LoginPage;