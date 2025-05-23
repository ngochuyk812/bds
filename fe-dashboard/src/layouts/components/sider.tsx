import { Image, Menu } from "antd";
import Sider from "antd/es/layout/Sider";
import type React from "react";
import {
    BellOutlined,
    DashboardOutlined,
    DatabaseOutlined,
    GroupOutlined,
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
                navigate('/sites');
                break;
            case '3-1':
                navigate('/amenities');
                break;
            default:
                break;
        }
    };

    return (
        <Sider width={250} trigger={null} collapsible collapsed={collapsed}>
            <Image src="./logo.png" preview={false} className="max-w-[45%]  m-auto mt-4 mb-8" />

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
                        label: 'Thống kê',
                    },
                    {
                        key: '2',
                        icon: <GroupOutlined />,
                        label: 'Quản lý trang',
                    },
                    {
                        key: '3',
                        icon: <DatabaseOutlined />,
                        label: 'Danh mục hệ thống',
                        children: [
                            {
                                key: '3-1',
                                label: 'Tiện ích',
                            },

                        ],
                    },
                ]}
            />
        </Sider>
    );
}