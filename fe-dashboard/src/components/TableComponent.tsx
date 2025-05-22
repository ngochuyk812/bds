import React, { useState, useEffect } from 'react';
import { Button, Table, TableColumnsType } from 'antd';
import useWindowSize from '../hooks/useWindowSize';
import { ColumnSearch, FilterItem, SearchAdvanceFilter } from './SearchAdvanceFilter';




interface TabelComponentProps<T> {
    col: TableColumnsType<T>;
    searchs: ColumnSearch[]
    isEdit?: boolean;
    isDelete?: boolean;
    data?: PaginationData<T>;
    onChangePage: (page: number, pageSize: number) => void;
    onSearch?: (payload: FilterItem[]) => void;
}

interface PaginationData<T> {
    current: number;
    pageSize: number;
    total: number;
    data?: T[];
}

const TabelComponent = <T,>(props: TabelComponentProps<T>) => {
    const { width, height } = useWindowSize();
    var col = props.col;

    if (props.isEdit || props.isDelete)
        col = [
            ...props.col,
            {
                title: 'Action',
                dataIndex: '',
                key: 'x',
                width: 1,
                render: () => <div >
                    {props.isEdit && <button className='text-blue-500 mr-4'>Edit</button>}
                    {props.isDelete && <button className='text-red-500'>Delete</button>}
                </div>,
            },
        ];



    return (
        <div>

            <SearchAdvanceFilter
                columns={props.searchs}
                onSearch={(payload) => {
                    props.onSearch && props.onSearch(payload);
                }}
            />

            <Table<T>
                columns={col}
                virtual
                scroll={{ x: width - 300 }}
                dataSource={props.data?.data}
                pagination={{
                    current: props.data?.current,
                    pageSize: props.data?.pageSize,
                    total: props.data?.total,
                }}
                onChange={(page) => {
                    props.onChangePage(page.current ?? 1, page.pageSize ?? 10);
                }}
            />
        </div>
    );
};

export default TabelComponent;