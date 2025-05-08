import React, { useState } from 'react';
import {
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    UploadOutlined,
    UserOutlined,
    VideoCameraOutlined,
} from '@ant-design/icons';
import { Button, Layout, Menu, theme } from 'antd';
import HeaderPublicComponent from '../components/header-public';

const { Header, Sider, Content } = Layout;

export default function LayoutPublic({ children }: { children: React.ReactNode }) {
    const [collapsed, setCollapsed] = useState(false);

    const {
        token: { colorBgContainer, borderRadiusLG },
    } = theme.useToken();
    return (
        <Layout className="bg-cover bg-[position-y:80%]" style={{
            background: 'url("./bg-auth.jpg")',
        }}>
            {/* <HeaderPublicComponent /> */}
            <Content
                className='h-[calc(100vh)]'
                style={{
                    borderRadius: borderRadiusLG,
                }}
            >
                {children}
            </Content>
        </Layout>
    );
};
