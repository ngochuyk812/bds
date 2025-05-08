import React, { useState, useEffect } from 'react';
import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Button, Image, Input, message } from 'antd';
import LayoutDashboard from './layouts/private';
import LayoutPublic from './layouts/public';
import { useAuthStore } from './store/auth';
import { LoginCredentials } from './types/auth';
import routes from './router';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';

const App: React.FC = () => {
  const isAuthenticated = localStorage.getItem('auth_token') !== null;



  return (
    <BrowserRouter>
      <Routes>
        {routes.map((route, index) => {
          if (route.type === 'private') {
            return (
              <Route
                key={index}
                path={route.path}
                element={
                  isAuthenticated ? (
                    <LayoutDashboard>
                      <route.element />
                    </LayoutDashboard>
                  ) : (
                    <Navigate to="/login" replace />
                  )
                }
              />
            );
          }
          return (
            <Route
              key={index}
              path={route.path}
              element={<route.element />}
            />
          );
        })}
      </Routes>
    </BrowserRouter>
  )
};

export default App;