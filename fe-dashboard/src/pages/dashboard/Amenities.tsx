import React, { useState, useEffect } from 'react';
import { ColumnSearch, FilterItem, SearchAdvanceFilter } from '../../components/SearchAdvanceFilter';
import { Button, Popconfirm, Table, TableColumnsType } from 'antd';
import useWindowSize from '../../hooks/useWindowSize';
import TabelComponent from '../../components/TableComponent';
import { useTable } from '../../hooks/useTable';
import { grpcAuthClient, grpcPropertyClient } from '../../utils/connectrpc';
import { FetchSitesRequest, FetchSitesResponse, SiteModel } from '../../proto/genjs/auth/v1/site_pb';
import { PaginationRequest, PaginationRequestSchema, SearchAdvanceRequest } from '../../proto/genjs/utils/v1/utils_pb';
import { Message } from '@bufbuild/protobuf';
import { StatusCode } from '../../proto/genjs/statusmsg/v1/statusmsg_pb';
import CreateSiteModal, { FieldConfig } from '../../components/CreateGenericModal';
import CreateEntityModal from '../../components/CreateGenericModal';
import { useNotificationStore } from '../../store/notification';
import CreateGenericModal from '../../components/CreateGenericModal';
import UpdateGenericModal, { FieldUpdateConfig } from '../../components/UpdateGenericModal';
import { AmenityModel, FetchAmenitiesRequest, FetchAmenitiesResponse, SearchAdvanceAmenitiesResponse } from '../../proto/genjs/property/v1/amenity_pb';
import { convertFilterToSearchAdvance } from '../../helpers/convertFiterToSearchAdvance';






const searchs: ColumnSearch[] = [
  { dataIndex: 'name', title: 'Name', type: 'string' },
  // { dataIndex: 'dob', title: 'Ngày tạo', type: 'date' },
]
const fields: FieldConfig[] = [
  { dataIndex: 'name', label: 'Tên tiện ích', rules: [{ required: true, message: 'Nhập tên tiện ích' }] },
  { dataIndex: 'description', label: 'Mô tả', rules: [{ required: false, message: 'Mô tả tiện ích' }] },
  { dataIndex: 'icon', label: 'Icon', rules: [{ required: true, message: 'Chọn icon' }] },
];
const AmenitiesPage: React.FC = () => {
  const [currentPage, setCurrentPage] = useState(1);

  const grpcFetchData = (params: Record<string, any>, page: number = 1, pageSize: number = 10) => {
    return grpcPropertyClient.searchAdvanceAmenities({
      startRow: (page - 1) * pageSize,
      endRow: page * pageSize,
      ...params
    });
  }
  const columns: TableColumnsType<AmenityModel> = [
    {
      title: 'Id',
      dataIndex: 'id',
      rowScope: 'row',
      fixed: 'left',
      width: 0.3
    },

    { title: 'Name', dataIndex: 'name', key: 'name', width: 1 },
    { title: 'Description', dataIndex: 'description', key: 'description', width: 1 },
    { title: 'Icon', dataIndex: 'icon', key: 'icon', width: 1 },
    {
      title: 'Action',
      dataIndex: '',
      key: 'x',
      width: 1,
      render: (col) => <div className='flex gap-2'>
        <Button onClick={() => showEditModal(col)} type="primary">Sửa</Button>
        <Popconfirm
          title="Xóa tiện ích này?"
          description="Bạn có chắc chắn muốn xóa tiện ích này không?"
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
    grpcPropertyClient.deleteAmenity({
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

    grpcPropertyClient.createAmenity({
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
    const filters = {
      filters: convertFilterToSearchAdvance(payload)
    };

    fetchData(() => {
      return grpcFetchData(filters);
    });
  }
  const handlerUpdate = (payload: Record<string, any>) => {
    grpcPropertyClient.updateAmenity({
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
  const { data, fetchData, refetch } = useTable<SearchAdvanceRequest, SearchAdvanceAmenitiesResponse>(
    (): Promise<SearchAdvanceAmenitiesResponse> => grpcFetchData({}),
  );
  const [open, setOpen] = useState({
    create: false,
    update: false,
  });
  const [dataUpdate, setDataUpdate] = useState<{
    key: any,
    data: FieldUpdateConfig[]
  }>();



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
        <h1 className='text-2xl '>Danh sách tiện ích</h1>
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


      <TabelComponent<AmenityModel> col={columns} searchs={searchs}
        data={
          data?.status?.code == StatusCode.SUCCESS ?
            {
              data: data?.rows ?? [],
              total: Number(data?.total),
              pageSize: 10,
              current: currentPage
            } : undefined
        }
        onChangePage={(page, pageSize) => {
          setCurrentPage(page);
          fetchData(() => {
            return grpcFetchData({}, page, pageSize);
          })
        }}
        onSearch={handlerSearch}
      />
    </div>
  );
};

export default AmenitiesPage;