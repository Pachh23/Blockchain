import React, { useEffect, useState } from 'react';
import { Table, Card, message } from 'antd';
import { ColumnsType } from 'antd/es/table';
import dayjs from 'dayjs';
import { BlockInterface } from '../../interface/IBlock';
import { GetAllBlock } from '../../services/https';

const BlockchainDataPage = () => {
  const [block, setBlock] = useState<BlockInterface[]>([]);
  const [messageApi, contextHolder] = message.useMessage();



  // Fetch Initial Data
  const getฺBlocks = async () => {
    let res = await GetAllBlock();
    if (res.status === 200) {
      setBlock(res.data);
    } else {
      messageApi.error("Block not found");
    }
  };

  useEffect(() => {
    getฺBlocks();
  }, []); // useEffect จะทำงานเมื่อ component ถูก mount

  // กำหนดคอลัมน์ของตาราง
  const columns: ColumnsType<any> = [
    {
      title: 'Block Index',
      dataIndex: 'index',
      key: 'index',
    },
    {
      title: 'Timestamp',
      dataIndex: 'timestamp',
      key: 'timestamp',
      render: (text: string) => <>{dayjs(text).format("YYYY-MM-DD HH:mm:ss")}</>,
    },
    {
      title: 'Previous Hash',
      dataIndex: 'previous_hash',
      key: 'previous_hash',
    },
    {
      title: 'Block Hash',
      dataIndex: 'hash',
      key: 'hash',
    },
    {
      title: 'Data',
      dataIndex: 'data',
      key: 'data',
    },
  ];

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-blue-100 p-6">
      <Card
        title={<div className="text-xl font-semibold">Blockchain Transactions</div>}
        bordered={false}
        className="shadow-lg"
      >
        <Table
          columns={columns}
          dataSource={block}
          rowKey="index"
          pagination={{ pageSize: 10 }}
        />
      </Card>
    </div>
  );
};

export default BlockchainDataPage;
