import React, { useState, useEffect } from 'react';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Image, Input, Modal } from 'antd';
import { SignUpCredentials } from '../../types/auth';
import { useAuthStore } from '../../store/auth';
import LayoutPublic from '../../layouts/public';
import { Link, useNavigate } from 'react-router-dom';
import { useNotificationStore } from '../../store/notification';
import { OTPProps } from 'antd/es/input/OTP';

const SignUpPage: React.FC = () => {
  const { isAuthenticated, isLoading, error, signUp, logout, clearError } = useAuthStore();
  const { addNotification, success, error: showError } = useNotificationStore();
  const [isModalOpen, setIsModalOpen] = useState(false);

  const showModal = () => {
    setIsModalOpen(true);
  };

  const handleOk = () => {
    setIsModalOpen(false);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
  };
  const [credentials, setCredentials] = useState<SignUpCredentials>({
    username: '',
    password: '',
    rePassword: '',
    name: ''
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setCredentials(prev => ({ ...prev, [name]: value }));
  };

  const handleLogin = async () => {

    const { username, password, rePassword, name } = credentials;

    if ([username, password, rePassword, name].some(field => !field)) {
      showError('Vui lòng nhập đủ thông tin');
      return;
    }
    if (!username.includes('@')) {
      showError('Email không hợp lệ');
      return;
    }
    if (credentials.password !== credentials.rePassword) {
      showError('Mật khẩu không khớp');
      return;
    }


    var rsSignup = await signUp(credentials);
    if (rsSignup) {
      success('Đăng ký thành công', 'Vui lòng kiểm tra mail để  lấy OTP kích hoạt tài khoản');
      showModal()
    }
  };

  return (
    <LayoutPublic>
      <InputOTP isModalOpen={isModalOpen} setIsModalOpen={setIsModalOpen} ussername={credentials.username} />
      <div className='h-[100%] flex items-center justify-center'>
        <div className='bg-white p-12 flex flex-col gap-[25px] mx-5 md:mx-0'>
          <Image src="./logo.png" preview={false} className="max-w-[45%]  m-auto" />

          <Input
            size="large"
            placeholder="Email/Username"
            prefix={<UserOutlined />}
            className="rounded-none md:min-w-[350px] text-[14px]"
            name="username"
            value={credentials.username}
            onChange={handleInputChange}
          />
          <Input
            size="large"
            placeholder="Họ tên"
            prefix={<LockOutlined />}
            className="rounded-none md:min-w-[350px] text-[14px]"
            type="name"
            name="name"
            value={credentials.name}
            onChange={handleInputChange}
          />
          <Input
            size="large"
            placeholder="Mạt khẩu"
            prefix={<LockOutlined />}
            className="rounded-none md:min-w-[350px] text-[14px]"
            type="password"
            name="password"
            value={credentials.password}
            onChange={handleInputChange}
          />
          <Input
            size="large"
            placeholder="Nập lại mật khẩu"
            prefix={<LockOutlined />}
            className="rounded-none md:min-w-[350px] text-[14px]"
            type="password"
            name="rePassword"
            value={credentials.rePassword}
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
          <div className='flex justify-end w-full'>
            <Link to="/login" className='text-primary '>Đã có tài khoản? Đăng nhập ngay</Link>
          </div>

        </div>
      </div>
    </LayoutPublic>
  );
};

type InputOTPProps = {
  isModalOpen: boolean;
  setIsModalOpen: (value: boolean) => void;
  ussername: string;
};
function InputOTP({ isModalOpen, setIsModalOpen, ussername }: InputOTPProps) {
  const verifySignUp = useAuthStore(state => state.verifySignUp);
  const showError = useNotificationStore(state => state.error);
  const sussess = useNotificationStore(state => state.success);
  const navigate = useNavigate();

  const [otp, setOtp] = useState("");

  const onChange: OTPProps['onChange'] = (value) => {
    setOtp(value);
  };

  const onFinish = async () => {
    if (otp.length != 7) {
      showError('Vui lòng nhập đủ 7 số');
      return;
    }
    var rsVerify = await verifySignUp({
      otp: otp,
      username: ussername
    });
    if (rsVerify) {
      sussess('Thành công', 'Kích hoạt tài khoản thành công');
      setIsModalOpen(false);
      navigate('/login');
    }

  };
  const sharedProps: OTPProps = {
    onChange,
  };
  return (
    <Modal
      title="Nhập mã OTP"
      closable={{ 'aria-label': 'Custom Close Button' }}
      okText="Xác nhận"
      cancelText="Hủy"
      open={isModalOpen}
      onOk={onFinish}
      centered
      onCancel={() => {
        setIsModalOpen(false)
        setOtp("")
      }}
    >
      <div className='flex justify-center'>
        <Input.OTP
          size='large'
          length={7}
          formatter={(str) => str.toUpperCase()} {...sharedProps} />
      </div>
    </Modal>
  );
}

export default SignUpPage;