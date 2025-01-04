import React, { useEffect, useState } from 'react';
import { Form, Select, DatePicker, Input, Button, Card, Row, Col, Table, message, ConfigProvider } from 'antd';
import { RoomInterface } from '../../interface/IRoom';
import { TimeInterface } from '../../interface/ITime';
import { AppointmentInterface } from '../../interface/IAppointment';
import { DepartmentInterface } from '../../interface/IDepartment';
import { CreateAppointment, GetAppointment, GetDepartment, GetRoom, GetTime } from '../../services/https';
import dayjs from 'dayjs'; // ใช้ dayjs แทน moment
import { ColumnsType } from 'antd/es/table';
const { TextArea } = Input;

function AppointmentForm1() {
  const [form] = Form.useForm();
  const [messageApi, contextHolder] = message.useMessage();
  const [selectedDept, setSelectedDept] = useState(null);
  const [rooms, setRooms] = useState<RoomInterface[]>([]);
  const [departments, setDepartments] = useState<DepartmentInterface[]>([]);
  const [times, setTimes] = useState<TimeInterface[]>([]);
  const [data, setDatas] = useState<AppointmentInterface[]>([]);

  const getDepartments = async () => {
    let res = await GetDepartment();
    if (res.status === 200) {
      setDepartments(res.data);
    } else {
      messageApi.error("Departments not found");
    }
  };

  // Fetch Initial Data
  const getRooms = async () => {
    let res = await GetRoom();
    if (res.status === 200) {
      setRooms(res.data);
    } else {
      messageApi.error("Rooms not found");
    }
  };

  // Fetch Initial Data
  const getTimes = async () => {
    let res = await GetTime();
    if (res.status === 200) {
      setTimes(res.data);
    } else {
      messageApi.error("Times not found");
    }
  };

  const getDatas = async () => {
    let res = await GetAppointment();

    if (res.status == 200) {
      setDatas(res.data);
    } else {
      setDatas([]);
      messageApi.open({
        type: "error",
        content: res.data.error,
      });
    }
  };

  useEffect(() => {
    getDepartments();
    getRooms();
    getTimes();
    getDatas();
  }, []);

  const handleDepartmentChange = (value: React.SetStateAction<null>) => {
    setSelectedDept(value);
    form.setFieldValue('RoomID', undefined);
  };

  const disabledDate = (date: dayjs.Dayjs) => {
    return date && date.isBefore(dayjs(), 'day');
  };

  const handleTimeChange = (value: any) => {
    form.setFieldValue('TimeID', value);
  };

  const handleRoomChange = (value: any) => {
    form.setFieldValue('RoomID', value);
  };

  const onFinish = async (values: AppointmentInterface) => {
    let res = await CreateAppointment(values);

    if (res.status == 201) {
      messageApi.open({
        type: "success",
        content: res.data.message,
      });
      setTimeout(function () {
        window.location.reload();
      }, 2000);
    } else {
      messageApi.open({
        type: "error",
        content: res.data.error,
      });
    }
  };

  // Logout function
  const handleLogout = () => {
    // Clear any authentication data (e.g., localStorage)
    localStorage.removeItem('NationalID');
    // Redirect to login page (assuming login page is at '/login')
    window.location.replace('/');
  };

  const columns: ColumnsType<AppointmentInterface> = [
    {
      title: 'ID',
      dataIndex: 'ID',
      key: 'id',
    },
    {
      title: "แผนก",
      key: "DepartmentID",
      render: (record) => <>{record?.department?.department}</>,
    },
    {
      title: "ห้องตรวจ",
      key: "RoomID",
      render: (record) => <>{record?.room?.room}</>,
    },
    {
      title: "วันที่",
      key: "date",
      render: (record) => <>{dayjs(record.date).format("DD/MM/YYYY")}</>,
    },
    {
      title: "เวลา",
      key: "TimeID",
      render: (record) => <>{record?.time?.time}</>,
    },
    {
      title: "อาการเบื้องต้น",
      dataIndex: "illness",
      key: "illness",
    },
  ];

  return (
    <ConfigProvider
      theme={{
        token: {
          colorPrimary: '#faad14',
          colorBgContainer: '#fff',
          borderRadius: 8,
        },
        components: {
          Card: {
            headerBg: '#faad14',
            headerFontSize: 16,
            headerFontSizeSM: 14,
            headerHeight: 50,
          }
        }
      }}
    >
      <div className="min-h-screen bg-gradient-to-br from-orange-50 to-orange-100 p-6">
        {contextHolder}
        <Card
          title={<div className="flex items-center text-white font-semibold">
            Medical Appointment Form
          </div>}
          bordered={false}
          className="shadow-lg"
        >


          <Row gutter={24}>
  <Col xs={24} lg={8}> {/* Reduce form column width */}
    <Form
      form={form}
      layout="vertical"
      requiredMark="optional"
      onFinish={onFinish}
      autoComplete="off"
      size="large"
      className="p-4"
    >

                <Row gutter={16}>
                  <Col xs={24} sm={12}>
                    <Form.Item
                      name="DepartmentID"
                      label={<span className="text-gray-700 font-medium">แผนก</span>}
                      rules={[{ required: true, message: 'Please select a department' }]}
                    >
                      <Select
                        placeholder="Select department"
                        onChange={handleDepartmentChange}
                        loading={departments.length === 0}
                      >
                        {departments.map((item) => (
                          <Select.Option value={item.ID} key={item.ID}>
                            {item.department}
                          </Select.Option>
                        ))}
                      </Select>
                    </Form.Item>
                  </Col>
                  <Col xs={24} sm={12}>
                    <Form.Item
                      name="RoomID"
                      label={<span className="text-gray-700 font-medium">ห้องตรวจ</span>}
                      rules={[{ required: true, message: 'Please select a room' }]}
                    >
                      <Select
                        placeholder="Select examination room"
                        disabled={!selectedDept}
                        loading={rooms.length === 0}
                        onChange={handleRoomChange}
                      >
                        {rooms
                          .filter((room) => room.department_id === selectedDept)
                          .map((room) => (
                            <Select.Option key={room.ID} value={room.ID}>
                              {room.room}
                            </Select.Option>
                          ))}
                      </Select>
                    </Form.Item>
                  </Col>
                </Row>

                <Row gutter={16}>
                  <Col xs={24} sm={12}>
                    <Form.Item
                      name="date"
                      label={<span className="text-gray-700 font-medium">วันที่</span>}
                      rules={[{ required: true, message: 'Please select a date' }]}
                    >
                      <DatePicker
                        className="w-full"
                        format="DD/MM/YYYY"
                        disabledDate={disabledDate}
                        style={{ width: '100%' }}
                      />
                    </Form.Item>
                  </Col>
                  <Col xs={24} sm={12}>
                    <Form.Item
                      name="TimeID"
                      label={<span className="text-gray-700 font-medium">เวลา</span>}
                      rules={[{ required: true, message: 'Please select a time' }]}
                    >
                      <Select
                        placeholder="Select time"
                        onChange={handleTimeChange}
                        loading={times.length === 0}
                      >
                        {times.map((item) => (
                          <Select.Option value={item.ID} key={item.ID}>
                            {item.time}
                          </Select.Option>
                        ))}
                      </Select>
                    </Form.Item>
                  </Col>
                </Row>

                <Form.Item
                  name="illness"
                  label={<span className="text-gray-700 font-medium">อาการเบื้องต้น</span>}
                  rules={[{ required: true, message: 'Please describe your symptoms' }]}
                >
                  <TextArea
                    rows={4}
                    placeholder="Please describe your initial symptoms"
                    className="text-base" />
                </Form.Item>

                <Form.Item>
  <Button
    type="primary"
    htmlType="submit"
    block
    size="large"
    className="h-12 text-base font-medium shadow-md hover:shadow-lg transition-all"
  >
    Confirm Appointment
  </Button>
</Form.Item>

<Form.Item>
  <Button
    type="primary"
    danger
    onClick={handleLogout}
    block
    size="large"
    className="h-12 text-base font-medium bg-red-900 text-white shadow-md hover:bg-red-900 hover:shadow-lg transition-all"
  
  >
    Logout
  </Button>
</Form.Item>

              </Form>
            </Col>

            <Col xs={24} lg={16}> {/* Increase table column width */}
              <Card
                title={<div className="flex items-center text-white font-semibold">
                  Appointment Table
                </div>}
                bordered={true}
                className="shadow-md h-full"
              >
                <Table
                  columns={columns}
                  dataSource={data}
                  rowKey="ID"
                  className="border border-gray-100 rounded-lg overflow-hidden"
                  pagination={{
                    pageSize: 5,
                    className: "pt-4"
                  }} />
              </Card>
            </Col>
          </Row>
        </Card>
      </div>
    </ConfigProvider>
  );
};

export default AppointmentForm1;
