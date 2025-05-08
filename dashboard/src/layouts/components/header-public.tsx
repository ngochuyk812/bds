import { Button, Image, Menu, theme } from "antd";
import Sider from "antd/es/layout/Sider";
import type React from "react";
import {
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    UploadOutlined,
    UserOutlined,
    VideoCameraOutlined,
} from '@ant-design/icons';
import { Dispatch, SetStateAction, useState } from "react";
import { Header } from "antd/es/layout/layout";


export default function HeaderPublicComponent() {
    const {
        token: { colorBgContainer, borderRadiusLG },
    } = theme.useToken();
    return (
        <Header style={{ padding: 0, background: "#31313185", height: 80 }}>
            <Image height={80} src="./logo.png" preview={false} className="flex items-center ml-2" />
        </Header>
    );
}