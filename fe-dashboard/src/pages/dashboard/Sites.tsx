import React, { useState, useEffect } from 'react';
import { SearchAdvanceFilter } from '../../components/SearchAdvanceFilter';
import { Button, Table, TableColumnsType } from 'antd';
import useWindowSize from '../../hooks/useWindowSize';


interface DataType {
  key: React.Key;
  name: string;
  siteId: string;
  domain: string;
}

const columns: TableColumnsType<DataType> = [
  {
    title: 'Id',
    dataIndex: 'key',
    rowScope: 'row',
    fixed: 'left',
    width: 50
  },

  { title: 'Site Id', dataIndex: 'siteId', key: 'siteId', width: 100 },
  { title: 'Name', dataIndex: 'name', key: 'name', width: 100 },
  { title: 'Domain', dataIndex: 'domain', key: 'domain', width: 100 },

  {
    title: 'Action',
    dataIndex: '',
    key: 'x',
    width: 100,
    render: () => <div >
      <button className='text-blue-500 mr-4'>Edit</button>
      <button className='text-red-500'>Delete</button>
    </div>,
  },
];

const data: DataType[] = [
  {
    key: 1,
    name: 'John Brown',
    siteId: '1100',
    domain: 'https://google.com'
  },
  {
    key: 2,
    name: 'Jim Green',
    siteId: '25',
    domain: 'https://google.com'
  },
  {
    key: 3,
    name: 'Not Expandable',
    siteId: '28',
    domain: 'https://google.com'
  },
  {
    key: 4,
    name: 'Joe Black',
    siteId: '12',
    domain: 'https://google.com'
  },
];


const SitesPage: React.FC = () => {
  const { width, height } = useWindowSize();

  return (
    <div>
      <div className='flex justify-between items-center mb-6'>
        <h1 className='text-2xl '>Quản lý Sites</h1>
        <Button type="primary" >
          Thêm mới
        </Button>

      </div>
      <SearchAdvanceFilter
        columns={[
          { dataIndex: 'name', title: 'Name', type: 'string' },
          { dataIndex: 'siteId', title: 'Site Id', type: 'string' },
          { dataIndex: 'domain', title: 'Domain', type: 'string' },
          // { dataIndex: 'age', title: 'Age', type: 'number' },
          { dataIndex: 'dob', title: 'Ngày tạo', type: 'date' },
        ]}
        onSearch={(payload) => {
          console.log('Sending to server:', payload);
        }}
      />

      <Table<DataType>
        columns={columns}
        expandable={{
          rowExpandable: (record) => record.name !== 'Not Expandable',
        }}
        virtual
        scroll={{ x: width - 200 }}
        dataSource={data}
      />

    </div>
  );
};

export default SitesPage;