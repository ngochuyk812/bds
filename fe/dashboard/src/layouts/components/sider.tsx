import { Menu } from "antd";
import Sider from "antd/es/layout/Sider";
import type React from "react";
import {
    BellOutlined,
    DashboardOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    UploadOutlined,
    UserOutlined,
    VideoCameraOutlined,
} from '@ant-design/icons';
import { Dispatch, SetStateAction, useState } from "react";
import { useNavigate } from "react-router-dom";

type Props = {
    collapsed: boolean
};
export default function SiderComponent({ collapsed }: Props) {
    const navigate = useNavigate();

    const handleMenuClick = (key: string) => {
        switch (key) {
            case '1':
                navigate('/');
                break;
            case '2':
                navigate('/notifications');
                break;
            default:
                break;
        }
    };

    return (
        <Sider trigger={null} collapsible collapsed={collapsed}>
            <div className="demo-logo-vertical" />
            <Menu
                theme="dark"
                mode="inline"
                defaultSelectedKeys={['1']}
                onClick={({ key }) => handleMenuClick(key)}
                items={[
                    {
                        key: '1',
                        icon: <DashboardOutlined />,
                        label: 'Dashboard',
                    },
                    {
                        key: '2',
                        icon: <BellOutlined />,
                        label: 'Notifications',
                    },
                    {
                        key: '3',
                        icon: <UserOutlined />,
                        label: 'Users',
                    },
                ]}
            />
        </Sider>
    );
}