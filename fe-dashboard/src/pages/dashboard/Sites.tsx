import React, { useState, useEffect } from 'react';
import { ColumnSearch, FilterItem, SearchAdvanceFilter } from '../../components/SearchAdvanceFilter';
import { Button, Table, TableColumnsType } from 'antd';
import useWindowSize from '../../hooks/useWindowSize';
import TabelComponent from '../../components/TableComponent';
import { useTable } from '../../hooks/useTable';
import { grpcAuthClient } from '../../utils/connectrpc';
import { FetchSitesRequest, FetchSitesResponse, SiteModel } from '../../proto/genjs/auth/v1/site_pb';
import { PaginationRequest, PaginationRequestSchema } from '../../proto/genjs/utils/v1/utils_pb';
import { Message } from '@bufbuild/protobuf';
import { StatusCode } from '../../proto/genjs/statusmsg/v1/statusmsg_pb';
import CreateSiteModal from '../../components/CreateEntityModal';
import CreateEntityModal from '../../components/CreateEntityModal';
import { useNotificationStore } from '../../store/notification';






const searchs: ColumnSearch[] = [
  { dataIndex: 'name', title: 'Name', type: 'string' },
  { dataIndex: 'siteId', title: 'Site Id', type: 'string' },
  // { dataIndex: 'dob', title: 'Ngày tạo', type: 'date' },
]

const SitesPage: React.FC = () => {
  const { width, height } = useWindowSize();

  const handlerCreate = (payload: Record<string, any>) => {
    console.log(payload);

    grpcAuthClient.createSite({
      ...payload
    }).then((res) => {
      if (res.status?.code == StatusCode.SUCCESS) {
        fetchData(() => {
          return grpcAuthClient.fetchSites({
            pagination: {
              pageNumber: BigInt(1),
              pageSize: BigInt(10)
            }
          })
        })
      } else {
        useNotificationStore.getState().errorExtras(StatusCode[res.status?.code ?? 0], res.status?.extras);
      }
    })
  }
  const handlerSearch = (payload: FilterItem[]) => {
    const filters = payload.reduce((acc, filter) => {
      if (filter.filterCol === 'name' || filter.filterCol === 'siteId') {
        acc[filter.filterCol] = filter.value;
      }
      return acc;
    }, {} as Record<string, string>);

    fetchData(() => {
      return grpcAuthClient.fetchSites({
        pagination: {
          pageNumber: BigInt(1),
          pageSize: BigInt(10),
        },
        ...filters
      });
    });
  }

  const { data, fetchData } = useTable<FetchSitesRequest, FetchSitesResponse>(
    (): Promise<FetchSitesResponse> => grpcAuthClient.fetchSites({
      pagination: {
        pageNumber: BigInt(1),
        pageSize: BigInt(10)
      }
    }),
  );
  const [open, setOpen] = useState(false);

  const columns: TableColumnsType<SiteModel> = [
    {
      title: 'Id',
      dataIndex: 'id',
      rowScope: 'row',
      fixed: 'left',
      width: 0.3
    },

    { title: 'Site', dataIndex: 'siteId', key: 'siteId', width: 1 },
    { title: 'Name', dataIndex: 'name', key: 'name', width: 1 },
  ];
  return (
    <div>
      <div className='flex justify-between items-center mb-6'>
        <h1 className='text-2xl '>Quản lý Sites</h1>
        <Button type="primary" onClick={() => setOpen(true)} >
          Thêm mới
        </Button>
      </div>
      <CreateEntityModal
        title="Tạo site mới"
        fields={[
          { name: 'siteId', label: 'Mã site', rules: [{ required: true, message: 'Nhập mã site' }] },
          { name: 'name', label: 'Tên site', rules: [{ required: true, message: 'Nhập tên site' }] },
        ]}
        open={open}
        setOpen={setOpen}
        onCreate={(values) => handlerCreate(values)}
      />

      <TabelComponent<SiteModel> col={columns} searchs={searchs} isEdit={true}
        data={
          data?.status?.code == StatusCode.SUCCESS ?
            {
              data: data?.items ?? [],
              total: Number(data?.pagination?.total),
              pageSize: Number(data?.pagination?.pageSize),
              current: Number(data?.pagination?.currentPage)
            } : undefined
        }
        onChangePage={(page, pageSize) => {
          fetchData(() => {
            return grpcAuthClient.fetchSites({
              pagination: {
                pageNumber: BigInt(page),
                pageSize: BigInt(pageSize)
              }
            })
          })
        }}
        onSearch={handlerSearch}
      />
    </div>
  );
};

export default SitesPage;