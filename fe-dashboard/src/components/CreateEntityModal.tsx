import React from 'react';
import { Button, Modal, Form, Input, message } from 'antd';
import type { FormItemProps } from 'antd';

type FieldConfig = {
    name: string;
    label: string;
    rules?: FormItemProps['rules'];
    inputType?: 'text' | 'number' | 'password' | 'date';
};

type CreateEntityModalProps = {
    title?: string;
    open: boolean;
    setOpen: (val: boolean) => void;
    onCreate: (values: Record<string, any>) => void;
    fields: FieldConfig[];
};

const CreateEntityModal: React.FC<CreateEntityModalProps> = ({
    title = 'Tạo mới',
    open,
    setOpen,
    onCreate,
    fields,
}) => {
    const [form] = Form.useForm();

    const handleOk = async () => {
        try {
            const values = await form.validateFields();
            onCreate(values);
            setOpen(false);
            form.resetFields();
            message.success('Tạo thành công!');
        } catch (err) {
            // ignore
        }
    };

    const handleCancel = () => {
        setOpen(false);
        form.resetFields();
    };

    return (
        <Modal
            title={title}
            open={open}
            onOk={handleOk}
            onCancel={handleCancel}
            okText="Tạo"
            cancelText="Hủy"
        >
            <Form form={form} layout="vertical">
                {fields.map((field) => (
                    <Form.Item
                        key={field.name}
                        name={field.name}
                        label={field.label}
                        rules={field.rules}
                    >
                        <Input type={field.inputType ?? 'text'} />
                    </Form.Item>
                ))}
            </Form>
        </Modal>
    );
};

export default CreateEntityModal;
