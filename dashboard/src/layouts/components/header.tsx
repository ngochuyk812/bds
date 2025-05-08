import { Button, Dropdown, Menu, MenuProps, theme } from "antd";
import Sider from "antd/es/layout/Sider";
import type React from "react";
import {
    CaretDownOutlined,
    LogoutOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    SmileOutlined,
    UploadOutlined,
    UserOutlined,
    VideoCameraOutlined,
} from '@ant-design/icons';
import { Dispatch, SetStateAction, useState } from "react";
import { Header } from "antd/es/layout/layout";
import { useAuthStore } from "../../store/auth";
import { useNavigate } from "react-router-dom";

type Props = {
    setCollapsed: Dispatch<SetStateAction<boolean>>;
    collapsed: boolean
};




export default function HeaderComponent({ collapsed, setCollapsed }: Props) {
    const { isAuthenticated, isLoading, error, login, logout, clearError } = useAuthStore();
    let navigate = useNavigate()
    const {
        token: { colorBgContainer, borderRadiusLG },
    } = theme.useToken();
    const logoutHandler = () => {
        logout();
        navigate('/login');
    };

    const items: MenuProps['items'] = [
        {
            key: '1',
            label: (
                <button>Thông tin</button>
            ),
            icon: <UserOutlined />,
        },
        {
            key: '2',
            label: (
                <button onClick={logoutHandler}>Đăng xuất</button>
            ),
            icon: <LogoutOutlined />,

        },

    ];

    return (
        <Header className="p-0 flex justify-between items-center" style={{ background: colorBgContainer }}>
            <Button
                type="text"
                icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
                onClick={() => setCollapsed(v => !v)}
                style={{
                    fontSize: '16px',
                    width: 64,
                    height: 64,
                }}
            />
            <Dropdown menu={{ items }} className="mr-5">
                <a onClick={(e) => e.preventDefault()}>
                    <Button type="primary" >
                        Thông tin <CaretDownOutlined />
                    </Button>
                </a>
            </Dropdown>


        </Header>

    );
}