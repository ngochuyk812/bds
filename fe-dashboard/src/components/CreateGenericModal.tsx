import React from 'react';
import { Button, Modal, Form, Input, message } from 'antd';
import type { FormItemProps } from 'antd';

export interface FieldConfig {
    dataIndex: string;
    label: string;
    rules?: FormItemProps['rules'];
    inputType?: 'text' | 'number' | 'password' | 'date';
    disabledUpdate?: boolean;
};

type CreateGenericModalProps = {
    title?: string;
    open: boolean;
    setOpen: (val: boolean) => void;
    onCreate: (values: Record<string, any>) => void;
    fields: FieldConfig[];
};

const CreateGenericModal: React.FC<CreateGenericModalProps> = ({
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
                        key={field.dataIndex}
                        name={field.dataIndex}
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

export default CreateGenericModal;
