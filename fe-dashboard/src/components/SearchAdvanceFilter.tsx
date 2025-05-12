import { Form, Input, InputNumber, DatePicker, Button, Space, Select } from 'antd';

const { RangePicker } = DatePicker;

interface Column {
    dataIndex: string;
    title: string;
    type: 'string' | 'number' | 'date';
}

interface FilterItem {
    filterCol: string;
    filterType: 'equals' | 'contains' | 'greaterThan' | 'lessThan';
    value: any;
}

interface Props {
    columns: Column[];
    onSearch: (filters: FilterItem[]) => void;
}

export const SearchAdvanceFilter: React.FC<Props> = ({ columns, onSearch }) => {
    const [form] = Form.useForm();

    const handleFinish = (values: any) => {
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
    };

    const renderField = (col: Column) => {
        switch (col.type) {
            case 'number':
                return (
                    <Space>
                        <Form.Item name={[col.dataIndex, 0]} noStyle>
                            <InputNumber placeholder="Min" />
                        </Form.Item>
                        <Form.Item name={[col.dataIndex, 1]} noStyle>
                            <InputNumber placeholder="Max" />
                        </Form.Item>
                    </Space>
                );
            case 'date':
                return <Form.Item name={col.dataIndex} noStyle><RangePicker /></Form.Item>;
            default:
                return <Input placeholder="Search..." />;
        }
    };

    return (
        <Form form={form} layout="inline" onFinish={handleFinish} className='mb-4'>
            {columns.map(col => (
                <Form.Item key={col.dataIndex} label={col.title} className='mt-4'>
                    {renderField(col)}
                </Form.Item>
            ))}


            {/* <Form.Item>
                <Button htmlType="submit" type="primary" className='mt-4'>Tìm kiếm</Button>
            </Form.Item> */}
        </Form>
    );
};
