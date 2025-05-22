import React, { useEffect } from 'react';
import { Button, Modal, Form, Input, message } from 'antd';
import type { FormItemProps } from 'antd';
import { FieldConfig } from './CreateGenericModal';

export interface FieldUpdateConfig extends FieldConfig {
    value: any;
}

type UpdateGenericModalProps = {
    title?: string;
    open: boolean;
    setOpen: (val: boolean) => void;
    onHandler: (values: Record<string, any>) => void;
    fields: FieldUpdateConfig[];
    id: string;
};

const UpdateGenericModal: React.FC<UpdateGenericModalProps> = ({
    title = 'Cập nhật',
    open,
    setOpen,
    onHandler,
    fields,
    id
}) => {
    const [form] = Form.useForm();

    const handleOk = async () => {
        try {
            const values = await form.validateFields();
            onHandler({ ...values, guid: id });
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

    useEffect(() => {
        if (open) {
            const initialValues = fields.reduce((acc, field) => {
                acc[field.dataIndex] = field.value;
                return acc;
            }, {} as Record<string, any>);
            form.setFieldsValue(initialValues);
        } else {
            form.resetFields();
        }
    }, [open, fields, form]);

    return (
        <Modal
            title={title}
            open={open}
            onOk={handleOk}
            onCancel={handleCancel}
            okText="Cập nhật"
            cancelText="Hủy"
        >
            <Form form={form} layout="vertical">
                {fields.map((field) => (
                    <Form.Item
                        key={field.dataIndex}
                        name={field.dataIndex}
                        label={field.label}
                        rules={field.rules}
                        initialValue={field.value}
                    >
                        <Input type={field.inputType ?? 'text'} disabled={field.disabledUpdate} />
                    </Form.Item>
                ))}
            </Form>
        </Modal>
    );
};

export default UpdateGenericModal;
