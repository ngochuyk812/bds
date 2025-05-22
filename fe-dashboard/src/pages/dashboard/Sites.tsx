import React, { useState, useEffect } from 'react';
import { ColumnSearch, FilterItem, SearchAdvanceFilter } from '../../components/SearchAdvanceFilter';
import { Button, Popconfirm, Table, TableColumnsType } from 'antd';
import useWindowSize from '../../hooks/useWindowSize';
import TabelComponent from '../../components/TableComponent';
import { useTable } from '../../hooks/useTable';
import { grpcAuthClient } from '../../utils/connectrpc';
import { FetchSitesRequest, FetchSitesResponse, SiteModel } from '../../proto/genjs/auth/v1/site_pb';
import { PaginationRequest, PaginationRequestSchema } from '../../proto/genjs/utils/v1/utils_pb';
import { Message } from '@bufbuild/protobuf';
import { StatusCode } from '../../proto/genjs/statusmsg/v1/statusmsg_pb';
import CreateSiteModal, { FieldConfig } from '../../components/CreateGenericModal';
import CreateEntityModal from '../../components/CreateGenericModal';
import { useNotificationStore } from '../../store/notification';
import CreateGenericModal from '../../components/CreateGenericModal';
import UpdateGenericModal, { FieldUpdateConfig } from '../../components/UpdateGenericModal';






const searchs: ColumnSearch[] = [
  { dataIndex: 'name', title: 'Name', type: 'string' },
  // { dataIndex: 'dob', title: 'Ngày tạo', type: 'date' },
]
const fields: FieldConfig[] = [
  { dataIndex: 'siteId', label: 'Mã site', rules: [{ required: true, message: 'Nhập mã site' }], disabledUpdate: true },
  { dataIndex: 'name', label: 'Tên site', rules: [{ required: true, message: 'Nhập tên site' }] },
];
const SitesPage: React.FC = () => {
  const { width, height } = useWindowSize();

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
    {
      title: 'Action',
      dataIndex: '',
      key: 'x',
      width: 1,
      render: (col) => <div className='flex gap-2'>
        <Button onClick={() => showEditModal(col)} type="primary">Sửa</Button>
        <Popconfirm
          title="Xóa site này?"
          description="Bạn có chắc chắn muốn xóa site này không?"
          okText="Xác nhận"
          cancelText="Hủy"
          onConfirm={() => handleDelete(col)}
        >
          <Button type="primary" danger>Xóa</Button>
        </Popconfirm>

      </div>,
    },
  ];

  const handleDelete = (col: SiteModel) => {
    grpcAuthClient.deleteSite({
      guid: col.guid
    }).then((res) => {
      if (res.status?.code == StatusCode.SUCCESS) {
        refetch()
      } else {
        useNotificationStore.getState().errorExtras(StatusCode[res.status?.code ?? 0], res.status?.extras);
      }
    })
  }
  const handlerCreate = (payload: Record<string, any>) => {
    console.log(payload);

    grpcAuthClient.createSite({
      ...payload
    }).then((res) => {
      if (res.status?.code == StatusCode.SUCCESS) {
        refetch()
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

  const { data, fetchData, refetch } = useTable<FetchSitesRequest, FetchSitesResponse>(
    (): Promise<FetchSitesResponse> => grpcAuthClient.fetchSites({
      pagination: {
        pageNumber: BigInt(1),
        pageSize: BigInt(10)
      }
    }),
  );
  const [open, setOpen] = useState({
    create: false,
    update: false,
  });
  const [dataUpdate, setDataUpdate] = useState<{
    key: any,
    data: FieldUpdateConfig[]
  }>();

  const handlerUpdate = (payload: Record<string, any>) => {
    grpcAuthClient.updateSite({
      ...payload
    }).then((res) => {
      if (res.status?.code == StatusCode.SUCCESS) {
        refetch()
      } else {
        useNotificationStore.getState().errorExtras(StatusCode[res.status?.code ?? 0], res.status?.extras);
      }
    })
    console.log(payload);
  }

  const showEditModal = (col: SiteModel) => {
    const updatedFields: FieldUpdateConfig[] = fields.map(field => ({
      ...field,
      value: col[field.dataIndex as keyof SiteModel],
    }));

    setDataUpdate({
      key: col.guid,
      data: updatedFields
    });
    setOpen({ ...open, update: true });
  }


  return (
    <div>
      <div className='flex justify-between items-center mb-6'>
        <h1 className='text-2xl '>Quản lý Sites</h1>
        <Button type="primary" onClick={() => setOpen({ ...open, create: true })} >
          Thêm mới
        </Button>
      </div>
      <CreateGenericModal
        title="Tạo site mới"
        fields={fields}
        open={open.create}
        setOpen={(val: boolean) => setOpen({ ...open, create: val })}
        onCreate={(values) => handlerCreate(values)}
      />

      <UpdateGenericModal
        fields={dataUpdate?.data ?? []}
        id={dataUpdate?.key}
        open={open.update}
        setOpen={(val: boolean) => setOpen({ ...open, update: val })}
        onHandler={(values) => handlerUpdate(values)}
      />


      <TabelComponent<SiteModel> col={columns} searchs={searchs}
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