import { Form, Input, InputNumber, DatePicker, Space } from 'antd';
import { useEffect, useMemo, useCallback } from 'react';
import debounce from 'lodash.debounce';

const { RangePicker } = DatePicker;

export interface ColumnSearch {
    dataIndex: string;
    title: string;
    type: 'string' | 'number' | 'date';
}

export interface FilterItem {
    filterCol: string;
    filterType: 'equals' | 'contains' | 'greaterThan' | 'lessThan';
    value: any;
}

interface Props {
    columns: ColumnSearch[];
    onSearch: (filters: FilterItem[]) => void;
}

export const SearchAdvanceFilter: React.FC<Props> = ({ columns, onSearch }) => {
    const [form] = Form.useForm();



    const handleFinish = useCallback((values: any) => {
        const filters: FilterItem[] = [];
        for (const col of columns) {
            const value = values[col.dataIndex];

            if (value !== undefined && value !== null && value !== '') {
                if (col.type === 'number' && Array.isArray(value)) {
                    const [min, max] = value;
                    if (min) {
                        filters.push({
                            filterCol: col.dataIndex,
                            filterType: 'greaterThan',
                            value: min
                        });
                    }
                    if (max) {
                        filters.push({
                            filterCol: col.dataIndex,
                            filterType: 'lessThan',
                            value: max
                        });
                    }
                } else if (col.type === 'date' && Array.isArray(value)) {
                    const [start, end] = value;
                    if (start) {
                        filters.push({
                            filterCol: col.dataIndex,
                            filterType: 'greaterThan',
                            value: start.format('YYYY-MM-DD')
                        });
                    }
                    if (end) {
                        filters.push({
                            filterCol: col.dataIndex,
                            filterType: 'lessThan',
                            value: end.format('YYYY-MM-DD')
                        });
                    }
                } else {
                    filters.push({
                        filterCol: col.dataIndex,
                        filterType: 'contains',
                        value: value
                    });
                }
            }
        }

        onSearch(filters);
    }, [columns, onSearch]);

    const debouncedFinish = useMemo(() => debounce(handleFinish, 500), [handleFinish]);

    useEffect(() => {
        return () => {
            debouncedFinish.cancel();
        };
    }, [debouncedFinish]);
    const renderField = (col: ColumnSearch) => {

        switch (col.type) {
            case 'number':
                return (
                    <Space>
                        <Form.Item name={[col.dataIndex, 0]} >
                            <InputNumber placeholder="Min" />
                        </Form.Item>
                        <Form.Item name={[col.dataIndex, 1]} >
                            <InputNumber placeholder="Max" />
                        </Form.Item>
                    </Space>
                );
            case 'date':
                return <Form.Item className='max-w-[200px]' name={col.dataIndex}><RangePicker /></Form.Item>;
            default:
                return <Form.Item name={col.dataIndex}><Input placeholder="Search..." /></Form.Item>;
        }
    };

    const handleValuesChange = useCallback((_: any, allValues: any) => {
        debouncedFinish(allValues);
    }, [debouncedFinish]);

    return (
        <Form
            form={form}
            layout="inline"
            onFinish={handleFinish}
            className='mb-4'
            onValuesChange={handleValuesChange}
        >
            {columns.map(col => (
                <Form.Item key={col.dataIndex} label={col.title} className='mt-4 ' layout='vertical'>
                    {renderField(col)}
                </Form.Item>
            ))}
        </Form>
    );
};
