import React, { useEffect, useState } from 'react';
import { Form, Select, DatePicker, Input, Button, Card, Row, Col, Table, message } from 'antd';
import { RoomInterface } from '../../interface/IRoom';
import { TimeInterface } from '../../interface/ITime';
import { AppointmentInterface } from '../../interface/IAppointment';
import { DepartmentInterface } from '../../interface/IDepartment';
import { CreateAppointment, GetAppointment, GetDepartment, GetRoom, GetTime } from '../../services/https';
import dayjs from 'dayjs'; // ใช้ dayjs แทน moment
import { ColumnsType } from 'antd/es/table';
const { TextArea } = Input;

const AppointmentForm1: React.FC = () => {
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

    useEffect(() => {
      console.log(departments);
      console.log(rooms);
    }, [departments, rooms]);
    
    const handleDepartmentChange = (value: React.SetStateAction<null>) => {
      setSelectedDept(value);
      // Reset room selection when department changes
      form.setFieldValue('RoomID', undefined);
    };
  // กำหนดวันไม่ให้เลือกวันในอดีต
  const disabledDate = (date: dayjs.Dayjs) => {
    // ป้องกันไม่ให้เลือกวันที่ในอดีต
    return date && date.isBefore(dayjs(), 'day'); // ใช้ isBefore แทน
  };
  console.log(data);

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
  
  // ข้อมูลสำหรับตาราง
  const columns : ColumnsType<AppointmentInterface> = [
    {
      title: 'ID',
      dataIndex: 'ID',
      key: 'id',
    },
    {
      title: "DepartmentID",
      key: "DepartmentID",
      render: (record) => <>{record?.department?.department}</>,
    },
    {
      title: "RoomID",
      key: "RoomID",
      render: (record) => <>{record?.room?.room}</>,
    }
      
,    
    {
      title: "Date",
      key: "date",
      render: (record) => <>{dayjs(record.date).format("DD/MM/YYYY")}</>,
    },
    {
      title: "TimeID",
      key: "TimeID",
      render: (record) => <>{record?.time?.time}</>,
    },
    {
      title: "Reason",
      dataIndex: "illness",
      key: "illness",
    },
  ];

  return (
    <div style={{ padding: '24px' }}>
      {contextHolder}
      <Card title="Medical Appointment Form" bordered={false}>
        <Row gutter={16}>
          <Col xs={24} sm={12}>
            {/* ฟอร์มสำหรับเลือกแผนกและห้อง */}
            <Form
              form={form}
              layout="vertical"
              onFinish={onFinish}
              autoComplete="off"
            >
              <Row gutter={16}>
              <Col xs={24} sm={12}>
          <Form.Item
            name="DepartmentID"
            label="Department"
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
            label="Examination Room"
            rules={[{ required: true, message: 'Please select a room' }]}
          >
            <Select
              placeholder="Select examination room"
              disabled={!selectedDept}
              loading={rooms.length === 0}
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
                    label="Date"
                    rules={[{ required: true, message: 'Please select a date' }]}
                  >
                    <DatePicker 
                      style={{ width: '100%' }}
                      format="DD/MM/YYYY"
                      disabledDate={disabledDate} // กำหนด disabledDate
                    />
                  </Form.Item>
                </Col>

                <Col xs={24} sm={12}>
                <Form.Item
            name="TimeID"
            label="time"
            rules={[{ required: true, message: 'Please select a time' }]}
          >
            <Select
              placeholder="Select time"
              onChange={handleDepartmentChange}
              loading={departments.length === 0}
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
                label="Initial Symptoms"
                rules={[{ required: true, message: 'Please describe your symptoms' }]}
              >
                <TextArea 
                  rows={4} 
                  placeholder="Please describe your initial symptoms"
                />
              </Form.Item>

              <Form.Item>
                <Button type="primary" htmlType="submit" block>
                  Confirm Appointment
                </Button>
              </Form.Item>
            </Form>
          </Col>

          {/* ตารางแสดงข้อมูลการนัดหมาย */}
          <Col xs={24} sm={12}>
            <Card title="Appointment Table" bordered={true}>
              <Table
                columns={columns}
                dataSource={data}
                rowKey="ID"
              />
            </Card>
          </Col>
        </Row>
      </Card>
    </div>
  );
};

export default AppointmentForm1;
