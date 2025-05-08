import React, { useState } from 'react';
import {
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    UploadOutlined,
    UserOutlined,
    VideoCameraOutlined,
} from '@ant-design/icons';
import { Button, Layout, Menu, theme } from 'antd';
import SiderComponent from '../components/sider';
import HeaderComponent from '../components/header';

const { Header, Sider, Content } = Layout;

export default function LayoutDashboard({ children }: { children: React.ReactNode }) {
    const [collapsed, setCollapsed] = useState(false);

    const {
        token: { colorBgContainer, borderRadiusLG },
    } = theme.useToken();
    return (
        <Layout>
            <SiderComponent collapsed={collapsed} />

            <Layout>
                <HeaderComponent collapsed={collapsed} setCollapsed={setCollapsed} />
                <Content
                    style={{
                        margin: '24px 16px',
                        padding: 24,
                        minHeight: '100vh',
                        background: colorBgContainer,
                        borderRadius: borderRadiusLG,
                    }}
                >
                    {children}
                </Content>
            </Layout>
        </Layout>
    );
};
